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

func save_data(data []interface{}, name string, prefix string) error {

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

func build_request(edge string, params url.Values, fields []string) (*http.Request, error) {
	
	baseUrl := "https://graph.facebook.com/v17.0/" + os.Getenv("ACCOUNT_ID") + edge
	// Add access Token to params
	params.Set("access_token", os.Getenv("ACCESS_TOKEN"))
	// Add fields to params
	for i := range(fields) {
		params.Add("fields", fields[i])
	}
	// Build the request
	var req *http.Request
	var err error
	req, err = http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return req, err
	}

	req.URL.RawQuery = params.Encode()

	return req, nil
}

func extract(req *http.Request, prefix string) error {
	
	// Variables
	h := md5.New()
	var filename string

	for page := 1; true; page++{

		// Execute the request
		fmt.Printf("Extractint page %d\n", page)
		data, err := exec_request(*req)
		if err != nil {
			return err
		}

		// Save the results
		io.WriteString(h, req.URL.String())
		filename = fmt.Sprintf("%x.json", h.Sum(nil))
		err = save_data(data["data"].([]interface{}), filename, prefix)
		if err != nil {
			return err
		}

		// Check if there is a nex page
		next := data["paging"].(map[string]interface{})["next"]
		if next == nil {
			// End extraction if not
			break
		}
		
		// Build next request
		req, err = http.NewRequest("GET", next.(string), nil)
		if err != nil {
			fmt.Println("Error building request: ", err)
		}
	}
	fmt.Println("Pagination ended")
	return nil
}

func main() {

	// Load environment
	fmt.Println("Loading environment")
	godotenv.Load("../.env")
	
	// CAMPAIGNS
	params := url.Values {
		"date_preset": { "maximum" },
		"limit": { "500" },
	}
	campaign_fields := []string{"id", "account_id", "name"}
	req, err := build_request("/campaigns", params, campaign_fields)
	if err != nil {
		fmt.Println("Error building request")
	}
	fmt.Println("Extracting campaigns...")
	err = extract(req, "campaigns/")
	if err != nil {
		//TODO: Handle
		fmt.Println(err)
	}

	// AD SETS
	params = url.Values {
		"date_preset": { "maximum" },
		"limit": { "500" },
	}
	ad_sets_fields := []string{
		"id",
		"account_id",
		"adlabels",
	}
	req, err = build_request("/ad_sets", params, ad_sets_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ad Sets")
	err = extract(req, "ad_sets/")
	if err != nil {
		fmt.Println("Error extraction ad sets: ", err)
	}

	// ADS
	params = url.Values {
		"date_preset": { "maximum" },
		"limit": { "500" },
	}
	ads_fields := []string{
		"id",
		"account_id",
		"ad_active_time",
	}
	req, err = build_request("/ads", params, ads_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ads")
	// Params are the same for campaigns
	err = extract(req, "ads/")
	if err != nil {
		fmt.Println("Error extracting ads: ", err)
	}

	// Extract ads insights
	params = url.Values {
		"date_preset": { "maximum" },
		"level": { "ad" },
		"limit": { "500" },
	}
	ads_insights_fields := []string {
		"id",
		"account_currency",
		"account_name",
	}
	req, err = build_request("/insights", params, ads_insights_fields)
	fmt.Println("Extracting ads insights")
	err = extract(req, "insights/")
	if err != nil {
		fmt.Println("Error extracting insights: ", err)
	}
}
