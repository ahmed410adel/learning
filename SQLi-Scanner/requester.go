package main

import (
	"fmt"
	//"io"
	"net/http"
	// "net/url"
	// "crypto/tls"
)



func requester(domain string, headerName string, handle string) *http.Request {
	if domain == "" {
		fmt.Println("You did not enter a domain")
		return nil
	}
	req, _ := http.NewRequest("GET", "https://"+domain, nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Te", "trailers")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Ahmed-Adel/Scanner")
	if headerName != "" && handle != "" {
		req.Header.Set(headerName, handle)
	}
	// new part
	// proxyURL, _ := url.Parse("http://127.0.0.1:8080") // Burp default

	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		Proxy: http.ProxyURL(proxyURL),
	// 		TLSClientConfig: &tls.Config{
	// 			InsecureSkipVerify: true, // Skip TLS verification for testing
	// 		},
	// 	},
	// }
	// resp, err := client.Do(req)

	// end of new part
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching the domain:", err)
		return nil 
	}
	// } else {
	// 	defer resp.Body.Close()
	// 	fmt.Println("Successfully fetched the domain:", domain)
	// 	fmt.Println("Response Status:", resp.Status)

	// 	fmt.Println("Response Headers:")
	// 	for key, value := range resp.Header {
	// 		fmt.Printf("%s: %s\n", key, value)
	// 	}
	// 	fmt.Printf("\n\n")
	// 	fmt.Println("Response Body:")

	// 	body, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		fmt.Println("Error reading response body:", err)
	// 	}
	// 	fmt.Println(string(body))
	// 	fmt.Printf("\n\n")
	// 	fmt.Println("Domain graphed successfully")
	// }

	fmt.Println("Successfully fetched the domain:", domain)
	fmt.Println("Response Status:", resp.Status)

	if resp.Header.Get("Set-Cookie") != "" {
	req.Header.Set("Cookie", resp.Header.Get("Set-Cookie"))
	}
	// //again


	return req
}
