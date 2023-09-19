package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"meta_marketing_extract/model"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/Valgard/godotenv"

	storage "cloud.google.com/go/bigquery/storage/apiv1beta2"
	storagepb "cloud.google.com/go/bigquery/storage/apiv1beta2/storagepb"
	"cloud.google.com/go/bigquery/storage/managedwriter/adapt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Program Arguments
type stringSliceFlag []string
func (f *stringSliceFlag) String() string {
	return fmt.Sprintf("%v", *f)
}
func (f *stringSliceFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
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

// node
type Node map[string]interface{}

// Response object
type MetaGraphAPIResponse struct {
	Data []Node
	Paging Paging
	Error MetaGraphAPIError
}


// Conversions to proto
// Conversions to proto

func nodeToCampaign(node Node) (*model.Campaign, error) {
	const layout string = "2006-01-02T15:04:05-0700"
	campaign := &model.Campaign{}

	if val, ok := node["id"].(string); ok {
		id, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		campaign.Id = int64(id)
	}

	if val, ok := node["account_id"].(string); ok {
		// Convert to into
		id, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		campaign.AccountId = int64(id)
	}

	if val, ok := node["bid_strategy"].(string); ok {
		campaign.BidStrategy = val
	}
	
	if val, ok := node["boosted_object_id"].(string); ok {
		id, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		campaign.BoostedObjectId = int64(id)
	}

	if val, ok := node["budget_rebalance_flag"].(bool); ok {
		campaign.BudgetRebalanceFlag = val
	}

	if val, ok := node["budget_remaining"].(string); ok {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		campaign.BudgetRemaining = int64(intVal) 
	}

	if val, ok := node["buying_type"].(string); ok {
		campaign.BuyingType = val
	}

	if val, ok := node["can_create_brand_lift_study"].(bool); ok {
		campaign.CanCreateBrandLiftStudy = val
	}

	if val, ok := node["can_use_spend_cap"].(bool); ok {
		campaign.CanUseSpendCap = val
	}

	if val, ok := node["configured_status"].(string); ok {
		campaign.ConfiguredStatus = val
	}

	if val, ok := node["created_time"].(string); ok {
		// Convert to Timestamp
		time, err := time.Parse(layout, val)
		if err != nil {
			return nil, err
		}
		campaign.CreatedTime = timestamppb.New(time)
	}

	if val, ok := node["daily_budget"].(string); ok {
		// Convert to int
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		campaign.DailyBudget = int64(intVal)
	}

	if val, ok := node["effective_status"].(string); ok {
		campaign.EffectiveStatus = val
	}

	if val, ok := node["has_secondary_skadnetwork_reporting"].(bool); ok {
		campaign.HasSecondarySkadnetworkReporting = val
	}

	if val, ok := node["is_budget_schedule_enabled"].(bool); ok {
		campaign.IsBudgetScheduleEnabled = val
	}

	if val, ok := node["is_skadnetwork_attribution"].(bool); ok {
		campaign.IsSkadnetworkAttribution = val
	}

	if val, ok := node["is_skadnetwork_attribution"].(bool); ok {
		campaign.IsSkadnetworkAttribution = val
	}

	if val, ok := node["last_budget_toggling_time"].(string); ok {
		// Convert to Timestamp
		time, err := time.Parse(layout, val)
		if err != nil { return nil, err
		}
		campaign.LastBudgetTogglingTime = timestamppb.New(time)
	}

	if val, ok := node["lifetime_budget"].(string); ok {
		// Convert to int
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		campaign.LifetimeBudget = int64(intVal)
	}

	if val, ok := node["name"].(string); ok {
		campaign.Name = val
	}
	if val, ok := node["objective"].(string); ok {
		campaign.Objective = val
	}
	if val, ok := node["primary_attribution"].(string); ok {
		campaign.PrimaryAttribution = val
	}
	if val, ok := node["smart_promotion_type"].(string); ok {
		campaign.SmartPromotionType = val
	}
	if val, ok := node["source_campaign_id"].(string); ok {
		// Convert to int
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		campaign.SourceCampaignId = int64(intVal)
	}
	if val, ok := node["special_ad_category"].(string); ok {
		campaign.SpecialAdCategory = val
	}
	if val, ok := node["spend_cap"].(string); ok {
		campaign.SpendCap = val
	}
	if val, ok := node["start_time"].(string); ok {
		// Convert to Timestamp
		time, err := time.Parse(layout, val)
		if err != nil {
			return nil, err
		}
		campaign.StartTime = timestamppb.New(time)
	}
	if val, ok := node["status"].(string); ok {
		campaign.Status = val
	}
	if val, ok := node["stop_time"].(string); ok {
		// Convert to Timestamp
		time, err := time.Parse(layout, val)
		if err != nil {
			return nil, err
		}
		campaign.StopTime = timestamppb.New(time)
	}
	if val, ok := node["topline_id"].(string); ok {
		// Convert to int
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		campaign.ToplineId = int64(intVal)
	}
	if val, ok := node["updated_time"].(string); ok {
		// Convert to Timestamp
		time, err := time.Parse(layout, val)
		if err != nil {
			return nil, err
		}
		campaign.UpdatedTime = timestamppb.New(time)
	}

	return campaign, nil
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

// Convert 


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

func extract(req *http.Request) ([]Node, error) {

	// Data
	var data []Node

	for page := 1; true; page++{

		// Execute the request
		fmt.Printf("Extracting page %d\n", page)
		response, err := exec_request(*req)
		if err != nil {
			log.Fatal(err)
		}

		// Append the data
		data = append(data, response.Data...)

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
	return data, nil
}

// Create bigquery client
func createBQClient() *storage.BigQueryWriteClient {

	ctx := context.Background()

	// create the bigquery client
	log.Println("creating the bigquery client...")
	client, err := storage.NewBigQueryWriteClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// descriptor
func getDescriptor(message protoreflect.ProtoMessage) *descriptorpb.DescriptorProto {
	descriptor, err := adapt.NormalizeDescriptor(message.ProtoReflect().Descriptor())
	if err != nil {
		log.Fatal("NormalizeDescriptor: ", err)
	}
	return descriptor
}

// Write Data
func writeRows(
	client *storage.BigQueryWriteClient,
	descriptor *descriptorpb.DescriptorProto,
	rows []protoreflect.ProtoMessage,
	project string, dataset string, table string,
	trace string,
) {

	ctx := context.Background()

	// create the write stream
	// a COMMITTED write stream inserts data immediately into bigquery
	log.Println("creating the write stream...")
	resp, err := client.CreateWriteStream(ctx, &storagepb.CreateWriteStreamRequest{
		Parent: fmt.Sprintf("projects/%s/datasets/%s/tables/%s", project, dataset, table),
		WriteStream: &storagepb.WriteStream{
			Type: storagepb.WriteStream_COMMITTED,
		},
	})
	if err != nil {
		log.Fatal("CreateWriteStream: ", err)
	}

	// get the stream by calling AppendRows
	log.Println("calling AppendRows...")
	stream, err := client.AppendRows(ctx)
	if err != nil {
		log.Fatal("AppendRows: ", err)
	}

	// serialize the rows
	log.Println("marshalling the rows...")
	var opts proto.MarshalOptions
	var data [][]byte
	for _, row := range rows {
		buf, err := opts.Marshal(row)
		if err != nil {
			log.Fatal("protobuf.Marshal: ", err)
		}
		data = append(data, buf)
	}

	// send the rows to bigquery
	log.Println("sending the data...")
	err = stream.Send(&storagepb.AppendRowsRequest{
		WriteStream: resp.Name,
		TraceId:     trace, // identifies this client
		Rows: &storagepb.AppendRowsRequest_ProtoRows{
			ProtoRows: &storagepb.AppendRowsRequest_ProtoData{
				// protocol buffer schema
				WriterSchema: &storagepb.ProtoSchema{
					ProtoDescriptor: descriptor,
				},
				// protocol buffer data
				Rows: &storagepb.ProtoRows{
					SerializedRows: data, // serialized protocol buffer data
				},
			},
		},
	})
	if err != nil {
		log.Fatal("AppendRows.Send: ", err)
	}

	// get the response, which will tell us whether it worked
	log.Println("waiting for response...")
	r, err := stream.Recv()
	if err != nil {
		log.Fatal("AppendRows.Recv: ", err)
	}

	if rErr := r.GetError(); rErr != nil {
		log.Printf("result was error: %v", rErr)
	} else if rResult := r.GetAppendResult(); rResult != nil {
		log.Printf("now stream offset is %d", rResult.Offset.Value)
	}

	log.Println("done")
}


// Extraction functions
func extract_campaings(AccountId string, bq *storage.BigQueryWriteClient) {
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
	
	edge := fmt.Sprintf("/%s/campaigns", AccountId)
	req, err := build_request(edge, params, campaign_fields)
	if err != nil {
		fmt.Println("Error building request")
	}
	fmt.Println("Extracting campaigns...")
	data, err := extract(req)
	if err != nil {
		//TODO: Handle
		fmt.Println(err)
	}
	fmt.Printf("Total rows extracted: %d\n", len(data))

	// Serialize the Data
	log.Println("Serializing json data into proto messages")
	var campaingsData []protoreflect.ProtoMessage
	for _, node := range(data){ 
		// Convert the Node to the campaings objective
		campaignProto, err := nodeToCampaign(node)
		if err != nil {
			log.Fatal(err)
		}
		messageProto := proto.Message(campaignProto)
		campaingsData = append(campaingsData, messageProto)
	}

	// Write data
	log.Println("Writing rows")
	var campaign model.Campaign
	desc := getDescriptor(&campaign)
	project := os.Getenv("PROJECT_ID")
	dataset := os.Getenv("DATASET_ID")
	table := "campaigns"
	trace := "historical-extraction"
	writeRows(bq, desc, campaingsData, project, dataset, table, trace)
}


func main() {

	// Initialize the program
	environmentPtr := flag.String("environment", "", "Specify the environment (e.g., dev or prod)")
	adAccountIdPtr := flag.String("ad-account-id", "", "Account Id to extract the data from")
	var entities stringSliceFlag
	flag.Var(&entities, "entities", "List of entities to extract")
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

	// BQ client
	bq := createBQClient()
	
	for _, entity := range entities {
		fmt.Printf("Extracting entity: %s\n", entity)
		switch entity {
		case "campaigns":
			extract_campaings(*adAccountIdPtr, bq)
		}
	}
	
	/*
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
	*/
}
