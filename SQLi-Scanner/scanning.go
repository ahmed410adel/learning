package main

import ("net/http"
	"fmt"
	"io"
)
func normalSQLiScan(req *http.Request){
	for key, value := range req.Header {
		if key == "Cookie" {
			
		}
		newValues := []string{}
		for _, val := range value {
		newVal := val + "'\"`"
		newValues = append(newValues, newVal)
		}
		req.Header[key] = newValues
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()


		body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:")
	fmt.Println(string(body))
}