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

func exec_request(req http.Request) (map[string]interface{}, error) {
	
	// Create objet to write results into
	var data map[string]interface{}

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	// Evaluate status code
	if resp.StatusCode != 200 {
		return data, fmt.Errorf("Status code %d", resp.StatusCode)
	}

	// Parse the json data and return
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(content, &data)
	return data, nil
}

func save_data(data map[string]interface{}, name string) error {

	// Encode object into json
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Write it file
	err = os.WriteFile(name, json_data, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Global variables
func extract(edge string, fields []string) error {

	account_id := os.Getenv("ACCOUNT_ID")
	access_token := os.Getenv("ACCESS_TOKEN")

	// Build the requests
	base_url := "https://graph.facebook.com/v17.0/" + account_id + edge

	req, err := http.NewRequest("GET", base_url + edge, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = url.Values{
		"access_token": { access_token },
		"limit": { "100" },
		"date_preset": { "maximum" },
		"fields": fields,
	}.Encode()

	// Execute it
	page_counter := 1
	data, err := exec_request(*req)
	if err != nil {
		return err
	}
	// Save the results
	filename := fmt.Sprintf("page_%d.json", page_counter)
	err = save_data(data, filename)
	if err != nil {
		return err
	}
	after_cursor := data["paging"].(map[string]interface{})["cursors"].(map[string]interface{})["after"]

	// Now paginate :D:

	for after_cursor != nil {
		page_counter += 1
		// Update the query paramenters
		req.URL.RawQuery = url.Values{
			"access_token": { access_token },
			"limit": { "100" },
			"date_preset": { "maximum" },
			"fields": fields,
			"after": { after_cursor.(string) },
		}.Encode()

		// Execute it
		data, err = exec_request(*req)
		if err != nil {
			return err
		}
		// Save the results
		filename = fmt.Sprintf("page_%d.json", page_counter)
		err = save_data(data, filename)
		if err != nil {
			return err
		}
		after_cursor = data["paging"].(map[string]interface{})["cursors"].(map[string]interface{})["after"]
	}
	return nil
}

func main() {
	// Load environment
	godotenv.Load("../.env")

	// Extract campaigns
	fmt.Println("Extracting campaigns...")
	campaign_fields := []string{"id", "account_id", "name"}
	err := extract("/campaigns", campaign_fields)
	if err != nil {
		//TODO: Handle
		fmt.Println(err)
	}
}