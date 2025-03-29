package apps

import (
	"crypto/tls"
	"fmt"
	"strings"
	"net/http"
	"io"
	"log"
	"time"
  )

func CheckAwake(up string, ip string, i int) bool {

	url := fmt.Sprintf("https://%s@%s:10000/xmlrpc.cgi", up, ip)
	method := "GET"
  
    payload := strings.NewReader(`<?xml version="1.0"?>
    <methodCall>
        <methodName>servers::get_my_address</methodName>
        <params>
            <param>
                <value>apiUser</value>
            </param>
            <param>
                <value>apiKey</value>
            </param>
        </params>
    </methodCall>`)
   
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }
    // Retry the request up to "i" times
    for attempt := 1; attempt <= i; attempt++ {
        req, err := http.NewRequest(method, url, payload)
        if err != nil {
            log.Printf("checkAwake: attempt %d: Error creating request: %v\n", attempt, err)
            continue
        }
        req.Header.Add("Content-Type", "application/xml")

        res, err := client.Do(req)
        if err != nil {
            log.Printf("checkAwake: attempt %d: Error making request: %v\n", attempt, err)
            continue
        }
        defer res.Body.Close()

        // Check if the response status code is 200
        if res.StatusCode == http.StatusOK {
            body, err := io.ReadAll(res.Body)
            if err != nil {
                log.Printf("checkAwake: attempt %d: Error reading response body: %v\n", attempt, err)
                continue
            }
            log.Printf("checkAwake: attempt %d: checkAwake: command executed successfully!\n\n%s\n", attempt, string(body))
            return true
        } else {
            log.Printf("checkAwake: attempt %d: Request failed with status code: %d\n", attempt, res.StatusCode)
        }

		// Add a delay before the next attempt
		if attempt < i {
			log.Printf("checkAwake: attempt %d: Retrying in 2 seconds...\n", attempt)
			time.Sleep(2 * time.Second)
		}
    }

    // If all attempts fail, return false
    log.Println("checkAwake: all attempts failed. The server is not awake.")
    return false

}