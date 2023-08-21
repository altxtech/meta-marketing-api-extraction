package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"io"
	"github.com/Valgard/godotenv"
)

type PagingInfo struct {
	Cursors struct {
		Before string `json:"before"`
		After string `json:"after"`
	}
	Next string `json:"next"`
}

type Campaign struct {
	ID string `json:"id"`
}
type CampaignEdge []Campaign

type GetCampaignsResponse struct {
	Data CampaignEdge
	Paging PagingInfo
}

// Global variables
func main() {

	// Load environgment
	godotenv.Load("../.env")
	account_id := os.Getenv("ACCOUNT_ID")
	fmt.Println("Account id: ", account_id)
	access_token := os.Getenv("ACCESS_TOKEN")
	fmt.Println("Access token: ", access_token)

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
	page_counter := 1
	fmt.Println("Extracting page: ", page_counter)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("There was an error making the request!")
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var result GetCampaignsResponse
	fmt.Println(resp.StatusCode)
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))
	json.NewDecoder(resp.Body).Decode(&result)
	// print the next token for debugging purposes
	fmt.Println(result.Paging.Next)
	page_counter += 1
	
	// Save results so I can look at it
	file, _ := json.MarshalIndent(result, "", "    ")
	_ = os.WriteFile("first_page.json", file, 0644)

	for result.Paging.Next != "" {
		// Create request for next page
		fmt.Println("Extracting page: ", page_counter)
		req, err = http.NewRequest("GET", result.Paging.Next, nil) // The next URL already inclues all necessary parameters
		if err != nil {
			fmt.Println("Error building request")
		}
		// Execute request
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			//TODO: Handle
		}

		// Execute it
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("There was an error making the request!")
			fmt.Println(err)
		}
		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Println(result.Paging.Next)
		page_counter += 1
	}
}
