# GoDig CLI Tool

## Introduction

GoDig is a modern, fast, and user-friendly command-line tool for querying DNS records. Developed in Go and utilizing the Cobra CLI package, GoDig is inspired by the classic functionality of the `dig` command but reinvented with a contemporary twist and the efficiency of Go.

## Why GoDig?

In an era where network responsiveness is key, traditional DNS lookup tools often fall short. GoDig steps in with:

- **Speed**: Built in Go for rapid query responses.
- **Ease of Use**: Intuitive commands and clean output formats.
- **Modern Touch**: Enhances traditional DNS query functionalities to meet today's networking demands.

## Features

- Support for multiple DNS record types (A, AAAA, CNAME, MX, etc.).
- High-speed query execution.
- Readable and clear output for ease of interpretation.
- Advanced query options for in-depth DNS analysis.

## Installation

[Instructions on how to install GoDig]

## Example Usage

### Basic DNS Query

```bash
godig query a drod.dev
Domain                         Class  TTL      Type     IP Address
-----------------------------------------------------------------------
drod.dev                       IN     300      A        185.199.108.153
drod.dev                       IN     300      A        185.199.109.153
drod.dev                       IN     300      A        185.199.110.153
drod.dev                       IN     300      A        185.199.111.153
```
This retrieves the A record for www.example.com.

## Advanced Query Options
::TODO

## Contributing
::TODO
## License
::TODO