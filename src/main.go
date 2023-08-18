package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io"
)

// Global variables
func main() {
	fmt.Println("Lets do this!!!")
	account_id := os.Getenv("ACCOUNT_ID")
	
	fmt.Println(account_id)
	access_token := os.Getenv("ACCESS_TOKEN")
	fmt.Println(access_token)

	// Extract all campaigns
	base_url := "https://graph.facebook.com/v17.0/" + account_id
	campaigns_path := "/campaigns"

	// Build the request
	req, err := http.NewRequest("GET", base_url + campaigns_path, nil)
	if err != nil {
		// TODO: Handle
	}
	req.URL.RawQuery = url.Values{
		"access_token": { access_token },
	}.Encode()

	// Execute it
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("There was an error making the request!")
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("There was an error parsing the response body!")
		fmt.Println(err)
	}
	fmt.Printf("%s\n", body)
	fmt.Println("We got liftoff!")
}
