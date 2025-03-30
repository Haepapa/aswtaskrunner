package apps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(ip string, password string) string{

    if password == "" {
        log.Fatal("login: Password is not set. Please set LOGIN_PASSWORD as an environment variable or provide it at build time.")
    }

    // Login URL and payload
    loginURL := fmt.Sprintf("http://%s:8000/login", ip)
    loginPayload := map[string]string{"password": password}

    // Convert payload to JSON
    payloadBytes, err := json.Marshal(loginPayload)
    if err != nil {
        log.Fatalf("login: Error marshalling login payload: %v", err)
    }

    // Make the login request
    resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        log.Fatalf("login: Error making login request: %v", err)
    }
    defer resp.Body.Close()

    // Check if login was successful
    if resp.StatusCode != http.StatusOK {
        log.Fatalf("login: Login failed with status code: %d", resp.StatusCode)
    }

    // Read the response body to get the token
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("login: Error reading login response body: %v", err)
    }

	token := string(body)
	if token == "" {
		log.Fatalf("login: Token not found in login response")
	}

	return token
}