/*
Copyright © 2023 Daniel Rodriguez dantech2000@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var dnsServer string // DNS server to use for queries

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godig",
	Short: "GoDig is a fast and flexible DNS query tool built in Go.",
	Long: `
   ______      ____  _
  / ____/___  / __ \(_)___ _
 / / __/ __ \/ / / / / __  /
/ /_/ / /_/ / /_/ / / /_/ /
\____/\____/_____/_/\__, /
                   /____/

GoDig is a command-line DNS query tool that provides quick and easy access to DNS records.

Built using Go, it offers performance and flexibility for modern networking needs.

It's inspired by traditional tools like 'dig', but with additional features and a user-friendly interface.

Examples of using GoDig:

Query A record: godig query a example.com
Query MX record: godig query mx example.com`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dnsServer, "server", "s", "1.1.1.1:53", "DNS server to use (IP address or hostname)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
