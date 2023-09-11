package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"cloud.google.com/go/storage"
	"github.com/Valgard/godotenv"
	"flag"
)

func exec_request(req http.Request) (map[string]interface{}, error) {

	// Create objet to write results into
	var data map[string]interface{}
	
	backoff_time := 100
	// The Meta API is really strict when it comes to Rate-Limiting, so yeah... 20 attempts
	for attempt := 1; attempt <= 20; attempt++ {

		// Trye to execute the request
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
		json.Unmarshal(content, &data)

		if resp.StatusCode == 200 {
			return data, nil
		}
		
		// Rate Limiting error -> Backoff 80004
		switch data["error"].(map[string]interface{})["code"].(float64) {

			case 17: 
				// backoff
				fmt.Printf("Rate limit exceeded. Backing off by %dms\n", backoff_time)
				time.Sleep(time.Duration(backoff_time) * time.Millisecond)
				backoff_time *= 2
			
			case 80004:
				// backoff
				fmt.Printf("Rate limit exceeded. Backing off by %dms\n", backoff_time)
				time.Sleep(time.Duration(backoff_time) * time.Millisecond)
				backoff_time *= 2

			default:
				fmt.Printf(string(content))
				return data, fmt.Errorf("HTTP Error - Status code %d", resp.StatusCode)
		}
	}

	// Retry limit exceede
	return data, fmt.Errorf("HTTP Error - Retry limit exceeded")
}

func make_jsonl(json_data []interface{}) ([]byte, error){
	var jsonl_data []byte
	var buf []byte
	var err error
	for i := range(json_data){
		// Encode object into json
		buf, err = json.Marshal(json_data[i])
		if err != nil {
			return jsonl_data, err
		}
		// Append the new json_data
		jsonl_data = append(jsonl_data, buf...)
		// Append a newline
		jsonl_data = append(jsonl_data, byte('\n'))
	}
	return jsonl_data, nil
}

func build_request(edge string, params url.Values, fields []string) (*http.Request, error) {
	
	baseUrl := "https://graph.facebook.com/v17.0/" + edge
	// Add access Token to params
	params.Set("access_token", os.Getenv("ACCESS_TOKEN"))
	// Add fields to params, if necessar
	if len(fields) > 0 {
		field_str := ""
		for i := range(fields) {
			field_str += fields[i] + ","
		}
		// Remove trailing comma
		field_str = field_str[:len(field_str)-1]
		params.Set("fields", field_str)
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

func extract(req *http.Request, prefix string) ([]interface{}, error) {

	// Initialize Google Cloud storage client
	ctx := context.Background()
	gcs, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	bucket := gcs.Bucket(os.Getenv("BUCKET_NAME"))
	
	// Variables
	h := md5.New()
	var key string
	var ids []interface{}

	for page := 1; true; page++{

		// Execute the request
		fmt.Printf("Extracting page %d\n", page)
		response, err := exec_request(*req)
		if err != nil {
			return ids, err
		}

		// Add the ids to the list of data to return
		data := response["data"].([]interface{})
		for i := range(data){
			ids = append(ids, data[i].(map[string]interface{})["id"])
		}

		// Convert the data to jsonl
		jsonl_data, err := make_jsonl(data)
		if err != nil {
			return ids, err
		}

		// Save the results
		io.WriteString(h, req.URL.String())
		key = fmt.Sprintf("%x.json", h.Sum(nil))
		ctx := context.Background()
		w := bucket.Object(prefix + key).NewWriter(ctx)
		_, err = w.Write(jsonl_data)
		if err != nil {
			return ids, err
		}
		err = w.Close()
		if err != nil {
			log.Fatal(err)
		}

		// Check if there is a next page
		paging := response["paging"]
		if paging == nil {
			break
		}
		next := paging.(map[string]interface{})["next"]
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
	return ids, nil
}

func main() {

	// Define a command-line flag for specifying the environment
	environmentPtr := flag.String("environment", "", "Specify the environment (e.g., dev or prod)")
	flag.Parse()

	// Load base environment variables from the .env file
	fmt.Println("Loading base environment")
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading base environment:", err)
		os.Exit(1)
	}

	// Load specific environment variables if an environment is specified
	if *environmentPtr != "" {
		envFileName := fmt.Sprintf("%s.env", *environmentPtr)
		fmt.Printf("Loading %s environment\n", *environmentPtr)
		if err := godotenv.Load(envFileName); err != nil {
			fmt.Printf("Error loading %s environment: %v\n", *environmentPtr, err)
			os.Exit(1)
		}
	}
	
	// CAMPAIGNS
	params := url.Values {
		"date_preset": { "maximum" },
		"limit": { "100" },
	}
	campaign_fields := []string{
		"id",
		"account_id",
		"adlabels",
		"bid_strategy",
		"boosted_object_id",
		"budget_rebalance_flag",
		"budget_remaining",
		"buying_type",
		"can_create_brand_lift_study",
		"can_use_spend_cap",
		"configured_status",
		"created_time",
		"daily_budget",
		"effective_status",
		"has_secondary_skadnetwork_reporting",
		"is_budget_schedule_enabled",
		"is_skadnetwork_attribution",
		"issues_info",
		"last_budget_toggling_time",
		"lifetime_budget",
		"name",
		"objective",
		"pacing_type",
		"primary_attribution",
		"promoted_object",
		"recommendations",
		"smart_promotion_type",
		"source_campaign",
		"source_campaign_id",
		"special_ad_categories",
		"special_ad_category",
		"special_ad_category_country",
		"spend_cap",
		"start_time",
		"status",
		"stop_time",
		"topline_id",
		"updated_time",
	}
	edge := fmt.Sprintf("/%s/campaigns", os.Getenv("ACCOUNT_ID"))
	req, err := build_request(edge, params, campaign_fields)
	if err != nil {
		fmt.Println("Error building request")
	}
	
	fmt.Println("Extracting campaigns...")
	_, err = extract(req, "campaigns/")
	if err != nil {
		//TODO: Handle
		fmt.Println(err)
	}

	// AD SETS
	params = url.Values {
		"date_preset": { "maximum" },
		"limit": { "100" },
	}
	adsets_fields := []string{
		"id",
		"frequency_control_specs",
		"instagram_actor_id",
		"is_budget_schedule_enabled",
		"is_dynamic_creative",
		"issues_info",
		"learning_stage_info",
		"lifetime_budget",
		"lifetime_imps",
		"lifetime_min_spend_target",
		"lifetime_spend_cap",
		"multi_optimization_goal_weight",
		"name",
		"optimization_goal",
		"optimization_sub_event",
		"pacing_type",
		"promoted_object",
		"recommendations",
		"recurring_budget_semantics",
		"review_feedback",
		"rf_prediction_id",
		"source_adset",
		"source_adset_id",
		"start_time",
		"status",
		"targeting",
		"targeting_optimization_types",
		"time_based_ad_rotation_id_blocks",
		"time_based_ad_rotation_intervals",
		"updated_time",
		"use_new_app_click",
	}
	edge = fmt.Sprintf("/%s/adsets", os.Getenv("ACCOUNT_ID"))
	req, err = build_request(edge, params, adsets_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ad Sets")
	_, err = extract(req, "adsets/")
	if err != nil {
		fmt.Println("Error extraction ad sets: ", err)
	}

	// ADS
	params = url.Values {
		"date_preset": { "maximum" },
		"limit": { "100" },
	}
	ads_fields := []string{
		"id",
		"account_id",
		"ad_active_time",
		"ad_review_feedback",
		"adlabels",
		"adset",
		"adset_id",
		"bid_amount",
		"campaign",
		"campaign_id",
		"configured_status",
		"conversion_domain",
		"created_time",
		"creative",
		"effective_status",
		"issues_info",
		"last_updated_by_app_id",
		"meta_reward_adgroup_status",
		"name",
		"preview_shareable_link",
		"recommendation",
		"source_ad",
		"source_ad_id",
		"status",
		"tracking_specs",
		"updated_time",
	}
	edge = fmt.Sprintf("/%s/ads", os.Getenv("ACCOUNT_ID"))
	req, err = build_request(edge, params, ads_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ads")
	// Params are the same for campaigns
	ads_ids, err := extract(req, "ads/")
	if err != nil {
		fmt.Println("Error extracting ads: ", err)
	}

	// ADS INSIGHTS
	params = url.Values { 
		"date_preset": { "maximum" },
		"level": { "ad" }, "limit": { "100" }, 
	} 
	ads_insights_fields := []string {
        "account_id",
        "account_name",
        "ad_id",
        "ad_name",
        "adset_id",
        "adset_name",
        "campaign_id",
        "campaign_name",
        "clicks",
        "conversions",
        "cost_per_ad_click",
        "cpc",
        "cpm",
        "cpp",
        "created_time",
        "ctr",
        "date_start",
        "date_stop",
        "frequency",
        "full_view_impressions",
        "full_view_reach",
        "impressions",
        "reach",
        "social_spend",
        "spend",
        "updated_time",
	}
	edge = fmt.Sprintf("/%s/insights", os.Getenv("ACCOUNT_ID"))
	req, err = build_request(edge, params, ads_insights_fields)
	fmt.Println("Extracting ads insights")
	_, err = extract(req, "insights/")
	if err != nil { 
		fmt.Println("Error extracting insights: ", err)
	}

	// AD LEADS
	// For each ad
	for i := range(ads_ids){

		fmt.Println("Extracting leads for ad ", ads_ids[i])
		// This edge takes no query parameters and no fields
		params = url.Values{}
		leads_fields := []string{}
		edge := fmt.Sprintf("%s/leads", ads_ids[i])
		req, err = build_request(edge, params, leads_fields)
		if err != nil {
			fmt.Println("Error building request: ", err)
		}
		// Execute the request
		// Lead files are going to be small enough that
		// I won't worry about partitioning by ad_id
		_, err := extract(req, fmt.Sprintf("leads/%d/", ads_ids[i]))
		if err != nil {
			fmt.Println("Error extracting data: ", err)
		}
	}
}
