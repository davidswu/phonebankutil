package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	To   = ""
	From = ""
)

func main() {
	// TODO: Set account keys & information from env var or cli
	accountSid := ""
	authToken := ""
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Create possible message bodies
	quotes := []string{
		"",
	}

	// quotes index
	i := 0

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", To)
	msgData.Set("From", From)

	for {
		msgData.Set("Body", quotes[i])
		msgDataReader := *strings.NewReader(msgData.Encode())

		// Create HTTP request client
		client := &http.Client{}
		req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// Make HTTP POST request and return message SID
		resp, _ := client.Do(req)
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var data map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			err := decoder.Decode(&data)
			if err == nil {
				fmt.Println(data["sid"])
			}
		}
		time.Sleep(time.Hour * 24)
		i++
		if i >= len(quotes) {
			i = 0
		}
	}
}
