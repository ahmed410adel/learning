package main

import (
	"flag"
)

func main() {
domainName := flag.String("u", "", "Domain to scan")
customHeader := flag.String("h", "", "Custom header name")
handle := flag.String("v", "", "Your handle")
flag.Parse()
SQLpayload := "' OR 1=1--"
requester(*domainName + SQLpayload, *customHeader, *handle)
}
