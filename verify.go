package main

import (
	"fmt"
	"net"
	"strings"
)

func checkDomain(domain string) (bool, bool, string, bool, string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	textRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	for _, textRec := range textRecords {
		if strings.HasPrefix(textRec, "v=spf1") {
			hasSPF = true
			spfRecord = textRec
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	for _, dRec := range dmarcRecords {
		if strings.HasPrefix(dRec, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = dRec
			break
		}
	}
	return hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
}
