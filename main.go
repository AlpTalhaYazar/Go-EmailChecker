package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please write the domain you want to check\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: could not lookup MX records for %s: %v\n", domain, err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	spfRecords, err := net.LookupTXT(fmt.Sprintf("_spf.%s", domain))

	if err != nil {
		log.Printf("Error: could not lookup TXT records for %s: %v\n", domain, err)
	}
	if len(spfRecords) > 0 {
		hasSPF = true
		spfRecord = spfRecords[0]
	}

	dmarcRecords, err := net.LookupTXT(fmt.Sprintf("_dmarc.%s", domain))

	if err != nil {
		log.Printf("Error: could not lookup TXT records for %s: %v\n", domain, err)
	}
	if len(dmarcRecords) > 0 {
		hasDMARC = true
		dmarcRecord = dmarcRecords[0]
	}

	fmt.Printf("\nhasMX : %t\n\nhasSPF : %t\n\nspfRecord : %s\n\nhasDMARC : %t\n\ndmarcRecord : %s\n", hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
