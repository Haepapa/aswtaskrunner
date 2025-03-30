package apps

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func CyclePlugOn(ip string, password string, device string, token string) {
    if password == "" {
        log.Fatal("login: Password is not set. Please set LOGIN_PASSWORD as an environment variable or provide it at build time.")
    }

    // Power off
    offURL := fmt.Sprintf("http://%s:8000/actions/p100/off?device=%s", ip, device)
    log.Printf("power off output: %s\n", makeRequest(offURL, token))

    // Power on
    onURL := fmt.Sprintf("http://%s:8000/actions/p100/on?device=%s", ip, device)
    log.Printf("power on output: %s\n", makeRequest(onURL, token))
}

func makeRequest(url string, token string) string {
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    return string(body)
}