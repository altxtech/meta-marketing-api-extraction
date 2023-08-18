package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url"
	"encoding/json"
	"io"
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

	account_id := os.Getenv("ACCOUNT_ID")
	access_token := os.Getenv("ACCESS_TOKEN")

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

	// Print body content
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		//TODO: handle
	}
	
	var result GetCampaignsResponse
	if err := json.Unmarshal(content, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result.Data)
	fmt.Println(result.Paging.Next)
	fmt.Printf("%T\n", result.Data)
	for i := range result.Data {
		fmt.Printf("%T\n", i)
	}
}
