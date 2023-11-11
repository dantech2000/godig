/*
Copyright © 2023 Daniel Rodriguez dantech2000@gmail.com

*/
package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query [record type] [domain]",
	Short: "Query DNS records for a given domain",
	Long: `Use godig to query DNS records for a specified domain.
For example:
godig query a example.com`,
	Args: cobra.MinimumNArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Split the server address into host and port
		host, port, err := net.SplitHostPort(dnsServer)
		if err != nil { // If splitting fails, assume it's a hostname without a port
			host = dnsServer
			port = "53" // Default port
		}

		// Resolve the DNS server address if it's not an IP
		if net.ParseIP(host) == nil {
			ips, err := net.LookupIP(host)
			if err != nil || len(ips) == 0 {
				return fmt.Errorf("error resolving DNS server address: %v", err)
			}
			host = ips[0].String()
		}

		dnsServer = net.JoinHostPort(host, port)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		recordType := strings.ToLower(args[0])
		domain := args[1]
		switch recordType {
		case "a":
			c := new(dns.Client)
			m := new(dns.Msg)

			m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
			r, _, err := c.Exchange(m, dnsServer) // Using Cloudflare's DNS server as default
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if len(r.Answer) == 0 {
				fmt.Printf("No A records found for %s\n", domain)
				return
			}

			fmt.Printf("%-30s %-6s %-8s %-8s %-16s\n", "Domain", "Class", "TTL", "Type", "IP Address")
			fmt.Println(strings.Repeat("-", 71)) // Separator line

			for _, a := range r.Answer {
				record, ok := a.(*dns.A)
				if ok {
					fmt.Printf("%-30s %-6s %-8d %-8s %-16s\n", domain, dns.ClassToString[record.Hdr.Class], record.Hdr.Ttl, "A", record.A.String())
				}
			}

		case "mx":
			c := new(dns.Client)
			m := new(dns.Msg)

			m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
			r, _, err := c.Exchange(m, dnsServer) // Using Cloudflare's DNS server as default
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if len(r.Answer) == 0 {
				fmt.Printf("No MX records found for %s\n", domain)
				return
			}

			fmt.Printf("%-20s %-8s %-6s %-8s %-30s\n", "Domain", "TTL", "Class", "Type", "Mail Exchange")
			fmt.Println(strings.Repeat("-", 72)) // Separator line

			for _, a := range r.Answer {
				mx, ok := a.(*dns.MX)
				if ok {
					fmt.Printf("%-20s %-8d %-6s %-8s %-30s\n", domain, mx.Hdr.Ttl, dns.ClassToString[mx.Hdr.Class], "MX", mx.Mx)
				}
			}

		default:
			fmt.Println("Unsupported record type")
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	// Here you can define flags and configuration settings.
}
