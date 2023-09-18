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

// Types
// Nodes
type Node interface{}

type AdAccount struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
}

type AdSet struct {
	ID                          string    `json:"id"`
	AccountID                   string    `json:"account_id"`
	AdCampaignID                string    `json:"ad_campaign_id"`
	AssetFeedID                 string    `json:"asset_feed_id"`
	BidAmount                   int       `json:"bid_amount"`
	BidStrategy                 string    `json:"bid_strategy"`
	BillingEvent                string    `json:"billing_event"`
	BudgetRemaining             float64   `json:"budget_remaining"`
	CampaignActiveTime          float64   `json:"campaign_active_time"`
	CampaignAttribution         string    `json:"campaign_attribution"`
	CampaignID                  float64   `json:"campaign_id"`
	ConfiguredStatus            string    `json:"configured_status"`
	CreatedTime                 time.Time `json:"created_time"`
	DailyBudget                 float64   `json:"daily_budget"`
	DailyMinSpendTarget         float64   `json:"daily_min_spend_target"`
	DailySpendCap               float64   `json:"daily_spend_cap"`
	DestinationType             string    `json:"destination_type"`
	DSABeneficiary              string    `json:"dsa_beneficiary"`
	DSAPayor                    string    `json:"dsa_payor"`
	EffectiveStatus             string    `json:"effective_status"`
	EndTime                     time.Time `json:"end_time"`
	InstagramActorID            string    `json:"instagram_actor_id"`
	IsBudgetScheduleEnabled     bool      `json:"is_budget_schedule_enabled"`
	IsDynamicCreative           bool      `json:"is_dynamic_creative"`
	LifetimeBudget              string    `json:"lifetime_budget"`
	LifetimeImpressions         int       `json:"lifetime_imps"`
	LifetimeMinSpendTarget      float64   `json:"lifetime_min_spend_target"`
	LifetimeSpendCap            string    `json:"lifetime_spend_cap"`
	MultiOptimizationGoalWeight string    `json:"multi_optimization_goal_weight"`
	Name                        string    `json:"name"`
	OptimizationGoal            string    `json:"optimization_goal"`
	OptimizationSubEvent        string    `json:"optimization_sub_event"`
	RecurringBudgetSemantics    bool      `json:"recurring_budget_semantics"`
	ReviewFeedback              string    `json:"review_feedback"`
	RFPredictionID              string    `json:"rf_prediction_id"`
	SourceAdSetID               string    `json:"source_adset_id"`
	StartTime                   time.Time `json:"start_time"`
	Status                      string    `json:"status"`
	UpdatedTime                 time.Time `json:"updated_time"`
	UseNewAppClick              bool      `json:"use_new_app_click"`
}

type Ad struct {
	ID                   string    `json:"id"`
	AccountID            string    `json:"account_id"`
	AdActiveTime         float64   `json:"ad_active_time"`
	AdScheduleEndTime    time.Time `json:"ad_schedule_end_time"`
	AdScheduleStartTime  time.Time `json:"ad_schedule_start_time"`
	AdSetID              string    `json:"adset_id"`
	BidAmount            int       `json:"bid_amount"`
	CampaignID           string    `json:"campaign_id"`
	ConfiguredStatus     string    `json:"configured_status"`
	ConversionDomain     string    `json:"conversion_domain"`
	CreatedTime          time.Time `json:"created_time"`
	EffectiveStatus      string    `json:"effective_status"`
	MetaRewardAdgroupStatus string `json:"meta_reward_adgroup_status"`
	Name                 string    `json:"name"`
	PreviewShareableLink string    `json:"preview_shareable_link"`
	SourceAdID           string    `json:"source_ad_id"`
	Status               string    `json:"status"`
	UpdatedTime          time.Time `json:"updated_time"`
}

type AdCreative struct {
	ID                           string `json:"id"`
	AccountID                    string `json:"account_id"`
	ActorID                      string `json:"actor_id"`
	ApplinkTreatment             string `json:"applink_treatment"`
	AuthorizationCategory        string `json:"authorization_category"`
	Body                         string `json:"body"`
	BundleFolderID               string `json:"bundle_folder_id"`
	CallToActionType             string `json:"call_to_action_type"`
	CategorizationCriteria       string `json:"categorization_criteria"`
	CategoryMediaSource          string `json:"category_media_source"`
	CollaborativeAdsLSBImageBankID string `json:"collaborative_ads_lsb_image_bank_id"`
	DestinationSetID             string `json:"destination_set_id"`
	DynamicAdVoice               string `json:"dynamic_ad_voice"`
	EffectiveAuthorizationCategory string `json:"effective_authorization_category"`
	EffectiveInstagramMediaID    string `json:"effective_instagram_media_id"`
	EffectiveInstagramStoryID    string `json:"effective_instagram_story_id"`
	EffectiveObjectStoryID       string `json:"effective_object_story_id"`
	EnableDirectInstall          bool   `json:"enable_direct_install"`
	EnableLaunchInstantApp       bool   `json:"enable_launch_instant_app"`
	ImageHash                    string `json:"image_hash"`
	ImageURL                     string `json:"image_url"`
	InstagramActorID             string `json:"instagram_actor_id"`
	InstagramPermalinkURL        string `json:"instagram_permalink_url"`
	InstagramStoryID             string `json:"instagram_story_id"`
	InstagramUserID              string `json:"instagram_user_id"`
	LinkDestinationDisplayURL    string `json:"link_destination_display_url"`
	LinkOGID                     string `json:"link_og_id"`
	LinkURL                      string `json:"link_url"`
	MessengerSponsoredMessage    string `json:"messenger_sponsored_message"`
	Name                         string `json:"name"`
	ObjectID                     string `json:"object_id"`
	ObjectStoreURL               string `json:"object_store_url"`
	ObjectType                   string `json:"object_type"`
	ObjectURL                    string `json:"object_url"`
	PlacePageSetID               string `json:"place_page_set_id"`
	PlayableAssetID              string `json:"playable_asset_id"`
	ProductSetID                 string `json:"product_set_id"`
	ReferralID                   string `json:"referral_id"`
	SourceInstagramMediaID       string `json:"source_instagram_media_id"`
	Status                       string `json:"status"`
	TemplateURL                  string `json:"template_url"`
	ThumbnailURL                 string `json:"thumbnail_url"`
	Title                        string `json:"title"`
	URLTags                       string `json:"url_tags"`
	UsePageActorOverride         bool   `json:"use_page_actor_override"`
	VideoID                      string `json:"video_id"`
}

type AdInsight struct {
	ID              string    `json:"id"`
	AccountID       string    `json:"account_id"`
	AccountName     string    `json:"account_name"`
	AdID            string    `json:"ad_id"`
	AdName          string    `json:"ad_name"`
	AdSetID         string    `json:"adset_id"`
	AdSetName       string    `json:"adset_name"`
	CampaignID      string    `json:"campaign_id"`
	CampaignName    string    `json:"campaign_name"`
	Clicks          int64     `json:"clicks"`
	CPC             float64   `json:"cpc"`
	CPM             float64   `json:"cpm"`
	CPP             float64   `json:"cpp"`
	CreatedTime     time.Time `json:"created_time"`
	CTR             float64   `json:"ctr"`
	DateStart       time.Time `json:"date_start"`
	DateStop        time.Time `json:"date_stop"`
	Frequency       float64   `json:"frequency"`
	FullViewImpressions int64 `json:"full_view_impressions"`
	FullViewReach   int64     `json:"full_view_reach"`
	Impressions     int64     `json:"impressions"`
	Reach           int64     `json:"reach"`
	SocialSpend     float64   `json:"social_spend"`
	Spend           float64   `json:"spend"`
	UpdatedTime     time.Time `json:"updated_time"`
}

type UserLeadGenInfo struct {
	ID          string          `json:"id"`
	AdID        string          `json:"ad_id"`
	AdName      string          `json:"ad_name"`
	AdSetID     string          `json:"adset_id"`
	AdSetName   string          `json:"adset_name"`
	CampaignID  string          `json:"campaign_id"`
	CampaignName string         `json:"campaign_name"`
	CreatedTime time.Time       `json:"created_time"`
	FieldData   json.RawMessage `json:"field_data"`
	FormID      string          `json:"form_id"`
	IsOrganic   bool            `json:"is_organic"`
	PartnerName string          `json:"partner_name"`
	Platform    string          `json:"platform"`
	Post        string          `json:"post"`
}
// Paging
type Paging struct {
	Cursors struct {
		After string `json:"after"`
		Before string `json:"before"`
	}
	Previous string `json:"previous"`
	Next string `json:"next"`
}

// Error
type MetaGraphAPIError struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Code int `json:"code"`
}

// Response object
type MetaGraphAPIResponse struct {
	Data []Node
	Paging Paging
	Error MetaGraphAPIError
}

func exec_request(req http.Request) (MetaGraphAPIResponse, error) {

	// Create objet to write results into
	var data MetaGraphAPIResponse
	
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
		switch data.Error.Code {

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

func make_jsonl(json_data []Node) ([]byte, error){
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

// This is the thing that is wrong
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
		data := response.Data
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
		paging := response.Paging
		next := paging.Next
		if next == "" {
			break
		}
		
		// Build next request
		req, err = http.NewRequest("GET", next, nil)
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
		"last_budget_toggling_time",
		"lifetime_budget",
		"name",
		"objective",
		"primary_attribution",
		"smart_promotion_type",
		"source_campaign_id",
		"special_ad_category",
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
		"account_id",
		"ad_campaign_id",
		"asset_feed_id",
		"bid_amount",
		"bid_strategy",
		"billing_event",
		"budget_remaining",
		"campaign_active_time",
		"campaign_attribution",
		"campaign_id",
		"configured_status",
		"created_time",
		"daily_budget",
		"daily_min_spend_target",
		"daily_spend_cap",
		"destination_type",
		"dsa_beneficiary",
		"dsa_payor",
		"effective_status",
		"end_time",
		"instagram_actor_id",
		"is_budget_schedule_enabled",
		"is_dynamic_creative",
		"lifetime_budget",
		"lifetime_imps",
		"lifetime_min_spend_target",
		"lifetime_spend_cap",
		"multi_optimization_goal_weight",
		"name",
		"optimization_goal",
		"optimization_sub_event",
		"recurring_budget_semantics",
		"review_feedback",
		"rf_prediction_id",
		"source_adset_id",
		"start_time",
		"status",
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
		"ad_schedule_end_time",
		"ad_schedule_start_time",
		"adset_id",
		"bid_amount",
		"campaign_id",
		"configured_status",
		"conversion_domain",
		"created_time",
		"effective_status",
		"meta_reward_adgroup_status",
		"name",
		"preview_shareable_link",
		"source_ad_id",
		"status",
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
