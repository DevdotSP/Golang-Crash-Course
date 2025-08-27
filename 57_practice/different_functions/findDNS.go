package diffuc

import (
	"fmt"
	"net"
)

// ExampleFindDNS demonstrates different DNS lookup operations
func ExampleFindDNS() {
	// Uncomment the functions to test specific DNS queries:
	// DNSLookUpIP()      // Lookup A and AAAA records (IPv4 and IPv6 addresses)
	// DNSCname()         // Lookup canonical name (CNAME) records
	// DNSPTR()           // Lookup PTR (reverse DNS) records
	// DNSNameServer()    // Lookup NS (Name Server) records
	DNSMaxLookUp() // Lookup MX (Mail Exchange) records
}

// DNSLookUpIP retrieves the IP addresses (A and AAAA records) of a host
func DNSLookUpIP() {
	iprecords, err := net.LookupIP("facebook.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, ip := range iprecords {
		fmt.Println("IP:", ip)
	}
}

// DNSCname retrieves the canonical name (CNAME) for a host
func DNSCname() {
	cname, err := net.LookupCNAME("m.facebook.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("CNAME:", cname)
}

// DNSPTR performs a reverse DNS lookup for a given IP address
func DNSPTR() {
	ptr, err := net.LookupAddr("8.8.8.8") // Google's public DNS
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, ptrvalue := range ptr {
		fmt.Println("PTR:", ptrvalue)
	}
}

// DNSNameServer retrieves the name server (NS) records for a domain
func DNSNameServer() {
	nameservers, err := net.LookupNS("facebook.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, ns := range nameservers {
		fmt.Println("NS:", ns.Host)
	}
}

// DNSMaxLookUp retrieves the mail exchange (MX) records for a domain
func DNSMaxLookUp() {
	mxrecords, err := net.LookupMX("facebook.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, mx := range mxrecords {
		fmt.Printf("MX Host: %s, Preference: %d\n", mx.Host, mx.Pref)
	}
}
