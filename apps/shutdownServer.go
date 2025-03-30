package apps

import(
	"crypto/tls"
	"fmt"
	"strings"
	"net/http"
	"io"
)

func ShutdownServer(up string, ip string){

	url := fmt.Sprintf("https://%s@%s:10000/xmlrpc.cgi", up, ip)
	method := "GET"
  
    payload := strings.NewReader(`<?xml version="1.0"?>
    <methodCall>
        <methodName>init::shutdown_system</methodName>
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
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return
	}
	req.Header.Add("Content-Type", "application/xml")
  
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := io.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(string(body))
}