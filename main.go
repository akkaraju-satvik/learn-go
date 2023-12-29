package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	// utils.Insert()
	auth0_url := "https://satvik.uk.auth0.com/oauth/device/code"
	params := url.Values{}
	params.Add("client_id", "28LESuEgj4TnYD68bgTfq3nEPMevJHS5")
	params.Add("scope", "openid profile email")
	payload := strings.NewReader(params.Encode())
	req, _ := http.NewRequest("POST", auth0_url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	fmt.Printf("Please open the following link in the browser: %s\n\n", result["verification_uri_complete"])
	tokenChan := make(chan string)
	fmt.Printf("Waiting for authorization...\n\n")
	go pollToken(result["device_code"].(string), &tokenChan)
	var token string = <-tokenChan
	fmt.Printf("Your token is: %s\n", token)
}

func pollToken(device_code string, tokenChan *chan string) {
	auth0_token_url := "https://satvik.uk.auth0.com/oauth/token"
	token_params := url.Values{}
	token_params.Add("client_id", "28LESuEgj4TnYD68bgTfq3nEPMevJHS5")
	token_params.Add("device_code", device_code)
	token_params.Add("grant_type", "urn:ietf:params:oauth:grant-type:device_code")
	token_payload := strings.NewReader(token_params.Encode())
	token_req, _ := http.NewRequest("POST", auth0_token_url, token_payload)
	token_req.Header.Add("content-type", "application/x-www-form-urlencoded")
	token_res, err := http.DefaultClient.Do(token_req)
	if err != nil {
		fmt.Println(err)
		if token_res.StatusCode == 403 {
			fmt.Println("Please check your email and approve the login request")
		}
	}
	defer token_res.Body.Close()
	token_body, _ := io.ReadAll(token_res.Body)
	var token_result map[string]interface{}
	json.Unmarshal([]byte(token_body), &token_result)
	if token_result["error"] == "authorization_pending" {
		time.Sleep(10 * time.Second)
		pollToken(device_code, tokenChan)
		return
	}
	*tokenChan <- token_result["id_token"].(string)
}
