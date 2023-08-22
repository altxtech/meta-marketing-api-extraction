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
	"time"
)

func exec_request(req http.Request) (map[string]interface{}, error) {

	// Create objet to write results into
	var data map[string]interface{}
	
	backoff_time := 100
	for attempt := 1; attempt <= 9; attempt++ {

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
		
		// Evaluate status code
		if resp.StatusCode != 200 {

			// Rate Limiting error -> Backoff
			if data["error"].(map[string]interface{})["code"].(float64) == 17 {
				// backoff
				fmt.Printf("Rate limit exceeded. Backing off by %dms", backoff_time)
				time.Sleep(time.Duration(backoff_time) * time.Millisecond)
				backoff_time *= 2
				continue
			}

			// Other errors
			fmt.Printf(string(content))
			return data, fmt.Errorf("HTTP Error - Status code %d", resp.StatusCode)
		}
		return data, nil
	}
	return data, fmt.Errorf("HTTP Error - Retry limit exceeded")
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

func extract(req *http.Request, prefix string) ([]interface{}, error) {
	
	// Variables
	h := md5.New()
	var filename string
	var ids []interface{}

	for page := 1; true; page++{

		// Execute the request
		fmt.Printf("Extractint page %d\n", page)
		response, err := exec_request(*req)
		if err != nil {
			return ids, err
		}

		// Add the ids to the list of data to return
		data := response["data"].([]interface{})
		for i := range(data){
			ids = append(ids, data[i].(map[string]interface{})["id"])
		}

		// Save the results
		io.WriteString(h, req.URL.String())
		filename = fmt.Sprintf("%x.json", h.Sum(nil))
		err = save_data(data, filename, "data/" + prefix)
		if err != nil {
			return ids, err
		}

		// Check if there is a nex page
		next := response["paging"].(map[string]interface{})["next"]
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

	// Load environment
	fmt.Println("Loading environment")
	godotenv.Load("../.env")
	
	// CAMPAIGNS
	params := url.Values {
		"date_preset": { "maximum" },
		"limit": { "500" },
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
	req, err := build_request("/campaigns", params, campaign_fields)
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
		"limit": { "500" },
	}
	adsets_fields := []string{
		"id",
		"account_id",
		"adlabelds",
		"adset_schedule",
		"asset_feed_id",
		"attribution_spec",
		"bid_adjustments",
		"bid_amount",
		"bid_constraints",
		"bid_info",
		"bid_strategy",
		"billing_event",
		"budget_remaining",
		"campaign",
		"campaign_active_time",
		"campaign_attribution",
		"campaign_id",
		"configured_status",
		"contextual_bundling_spec",
		"created_time",
		"creative_sequence",
		"daily_budget",
		"daily_min_spend_target",
		"daily_spend_cap",
		"destination_type",
		"dsa_beneficiary",
		"dsa_payor",
		"effective_status",
		"end_time",
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
	req, err = build_request("/adsets", params, adsets_fields)
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
		"limit": { "500" },
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
	req, err = build_request("/ads", params, ads_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ads")
	// Params are the same for campaigns
	_, err = extract(req, "ads/")
	if err != nil {
		fmt.Println("Error extracting ads: ", err)
	}

	// ADS INSIGHTS
	params = url.Values { 
		"date_preset": { "maximum" },
		"level": { "ad" }, "limit": { "500" }, 
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
	req, err = build_request("/insights", params, ads_insights_fields)
	fmt.Println("Extracting ads insights")
	_, err = extract(req, "insights/")
	if err != nil { 
		fmt.Println("Error extracting insights: ", err)
	}

	// AD LEADS
	// We won't use the standard functions, because it is a little different
	// TODO: Implement
}
