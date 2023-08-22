package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"io"
	"github.com/Valgard/godotenv"
	"path/filepath"
	"crypto/md5"
)

func exec_request(req http.Request) (map[string]interface{}, error) {

	// Create objet to write results into
	var data map[string]interface{}

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	// Parse the json data and return
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	
	// Evaluate status code
	if resp.StatusCode != 200 {
		fmt.Println(string(content))
		return data, fmt.Errorf("Status code %d", resp.StatusCode)
	}
	
	json.Unmarshal(content, &data)
	return data, nil
}

func save_data(data map[string]interface{}, name string, prefix string) error {

	// Create directories if then don't exist
	dirPath := filepath.Dir(prefix + name)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directories: ", err)
	}

	// Encode object into json
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Write it file
	err = os.WriteFile(prefix + name, json_data, 0666)
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
	fmt.Println("Building request")
	
	req_url := "https://graph.facebook.com/v17.0/" + account_id + "/campaigns"

	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		return err
	}
	
	// Prepare query params
	fmt.Println("Preparing query params")
	
	params := url.Values{}
	params.Add("access_token", access_token)
	params.Add("limit", "100")
	params.Add("date_preset", "maximum")
	for i := range(fields){
		params.Add("fields", fields[i])
	}
	req.URL.RawQuery = params.Encode()

	// Execute it
	page_counter := 1
	fmt.Printf("Extractint page %d\n", page_counter)
	
	data, err := exec_request(*req)
	if err != nil {
		return err
	}

	// Save the results
	h := md5.New()
	io.WriteString(h, req.URL.String())
	filename := fmt.Sprintf("%x.json", h.Sum(nil))
	prefix := fmt.Sprintf("data/%s/", edge)
	err = save_data(data, filename, prefix)
	if err != nil {
		return err
	}
	next := data["paging"].(map[string]interface{})["next"]

	// Now paginate :D
	for next != nil {
		page_counter += 1
		// Update the query paramenters
		fmt.Printf("Extracting page %d - %s\n", page_counter, next)
		
		req, err = http.NewRequest("GET", next.(string), nil)
		if err != nil {
			fmt.Println("Error preparing request: ", err)
		}

		// Execute it
		data, err = exec_request(*req)
		if err != nil {
			return err
		}
		// Save the results
		filename = fmt.Sprintf("%s.json", next)
		err = save_data(data, filename, path)
		if err != nil {
			return err
		}
		next = data["paging"].(map[string]interface{})["next"]
	}
	fmt.Println("Pagination ended")
	
	return nil
}

func main() {

	// Load environment
	fmt.Println("Loading environment")
	godotenv.Load("../.env")

	// Extract campaigns
	campaign_fields := []string{"id", "account_id", "name"}
	fmt.Println("Extracting campaigns...")
	err := extract("campaigns", campaign_fields)
	if err != nil {
		//TODO: Handle
		fmt.Println(err)
	}

}
