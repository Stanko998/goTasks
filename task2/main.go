/*
As a security researcher, I need a simple subdomain bruteforcer tool to identify potential subdomains of a target domain for authorized security testing.
The tool should take a target domain (e.g., example.com) and a wordlist of potential subdomain names (e.g., subdomains.txt), then attempt to resolve each
subdomain (e.g., sub1.example.com, sub2.example.com) using DNS requests. If a subdomain resolves successfully, it should be logged as a discovered subdomain.
This implementation should be as simple as possible, focusing on core functionality without advanced features like concurrency, rate
limiting, or wildcard detection.

Requirements
Input:
A target domain (e.g., example.com) provided as a command-line argument.

A wordlist file (e.g., subdomains.txt) containing potential subdomain names, one per line (e.g., dev, api, staging).

Process:
For each subdomain in the wordlist, construct a full domain (e.g., dev.example.com).

Attempt to resolve the full domain using a DNS request.

If the DNS request succeeds (i.e., the subdomain exists), log it to the console.

If the DNS request fails (e.g., NXDOMAIN error), skip it.

Output:
Print each discovered subdomain to the console (e.g., Found: dev.example.com).

Error Handling:
Handle basic errors like invalid domain input or missing wordlist file.

Dependencies:
Use standard libraries only (e.g., Python’s socket or dns.resolver for DNS lookups).

No external dependencies beyond what’s available in the standard library.
*/
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	domain := flag.String("domain", "example.com", "Domain to bruteforce")
	wordlist := flag.String("wordlist", "task2/wordlist.txt", "Wordlist file")
	flag.Parse()

	if len(*domain) == 0 || len(*wordlist) == 0 {
		fmt.Println("error you need 3 arguments! \nExample: go run main.go --domain doman.com --wordlist wordlist.txt")
		os.Exit(1)
	}

	file, err := os.ReadFile(*wordlist)
	if err != nil {
		fmt.Println("error opening file: ", err)
		return
	}

	for _, line := range strings.Split(string(file), "\r\n") {
		subdomain := line + "." + *domain
		_, err := net.LookupHost(subdomain)
		if err == nil {
			fmt.Println("Found: ", subdomain)
		} else {
			fmt.Println("Not Found: ", subdomain)
		}
	}

}
