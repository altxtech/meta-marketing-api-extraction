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
func parseAsString(node Node, key string) (string, bool) {
	val, ok := node[key].(string)
	return val, ok
}

func parseAsNumericString(node Node, key string) (int64, bool) {
	valStr, ok := node[key].(string)
	if !ok {
		return 0, false
	}
	intVal, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing %s:%s as int64: %v", key, valStr, err)
	}
	return intVal, true
}

func parseAsFloatString(node Node, key string) (float64, bool) {
	valStr, ok := node[key].(string)
	if !ok {
		return 0, false
	}
	floatVal, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		log.Fatalf("Error parsing %s:%s as int64: %v", key, valStr, err)
	}
	return floatVal, true
}

func parseAsInt(node Node, key string) (int64, bool) {
	val, ok := node[key].(int)
	if !ok {
		return 0, false
	}
	return int64(val), true
}

func parseAsTimestamp(node Node, key string, layout string) (*timestamppb.Timestamp, bool) {
	valStr, ok := node[key].(string)
	if !ok {
		return nil, false
	}
	timeVal, err := time.Parse(layout, valStr)
	if err != nil {
		log.Fatalf("Error parsing %s:%s as timestamp: %v", key, valStr, err)
	}
	return timestamppb.New(timeVal), true
}

func parseAsBool(node Node, key string) (bool, bool) {
	val, ok := node[key].(bool)
	return val, ok
}

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

func nodeToAdset( node Node ) ( *model.AdSet, error ){
	adset := &model.AdSet{}

	if val, ok := parseAsNumericString(node, "id"); ok {
		adset.Id = val
	}
	if val, ok := parseAsNumericString(node, "account_id"); ok {
		adset.AccountId = val
	}
	if val, ok := parseAsNumericString(node, "asset_feed_id"); ok {
		adset.AssetFeedId = val
	}
	if val, ok := parseAsInt(node, "bid_amount"); ok {
		adset.BidAmount = val
	}
	if val, ok := parseAsString(node, "bid_strategy"); ok {
		adset.BidStrategy = val
	}
	if val, ok := parseAsString(node, "billing_event"); ok {
		adset.BillingEvent = val
	}
	if val, ok := parseAsNumericString(node, "budget_remaining"); ok {
		adset.BudgetRemaining = val
	}
	if val, ok := parseAsNumericString(node, "campaign_active_time"); ok {
		adset.CampaignActiveTime = val
	}
	if val, ok := parseAsString(node, "campaign_attribution"); ok {
		adset.CampaignAttribution = val
	}
	if val, ok := parseAsNumericString(node, "campaign_id"); ok {
		adset.CampaignId = val
	}
	if val, ok := parseAsString(node, "configured_status"); ok {
		adset.ConfiguredStatus = val
	}
	if val, ok := parseAsTimestamp(node, "created_time", "2006-01-02T15:04:05-0700"); ok {
		adset.CreatedTime = val
	}
	if val, ok := parseAsNumericString(node, "daily_min_spend_target"); ok {
		adset.DailyMinSpendTarget = val
	}
	if val, ok := parseAsNumericString(node, "daily_spend_cap"); ok {
		adset.DailySpendCap = val
	}
	if val, ok := parseAsString(node, "destination_type"); ok {
		adset.DestinationType = val
	}
	if val, ok := parseAsString(node, "dsa_beneficiary"); ok {
		adset.DsaBeneficiary = val
	}
	if val, ok := parseAsString(node, "dsa_payor"); ok {
		adset.DsaPayor = val
	}
	if val, ok := parseAsString(node, "effective_status"); ok {
		adset.EffectiveStatus = val
	}
	if val, ok := parseAsTimestamp(node, "end_time", "2006-01-02T15:04:05-0700"); ok {
		adset.EndTime = val
	}
	if val, ok := parseAsNumericString(node, "instagram_actor_id"); ok {
		adset.InstagramActorId = val
	 }
	if val, ok := parseAsBool(node, "is_budget_schedule_enabled"); ok {
		adset.IsBudgetScheduleEnabled = val
	}
	if val, ok := parseAsBool(node, "is_dynamic_creative"); ok {
		adset.IsDynamicCreative = val
	}
	if val, ok := parseAsNumericString(node, "lifetime_budget"); ok {
		adset.LifetimeBudget = val
	}
	if val, ok := parseAsInt(node, "lifetime_imps"); ok {
		adset.LifetimeImps = val
	}
	if val, ok := parseAsNumericString(node, "lifetime_min_spend_target"); ok {
		adset.LifetimeMinSpendTarget = val
	}
	if val, ok := parseAsNumericString(node, "lifetime_spend_cap"); ok {
		adset.LifetimeSpendCap = val
	}
	if val, ok := parseAsString(node, "multi_optimization_goal_weight"); ok {
		adset.MultiOptimizationGoalWeight = val
	}
	if val, ok := parseAsString(node, "name"); ok {
		adset.Name = val
	}
	if val, ok := parseAsString(node, "optimization_goal"); ok {
		adset.OptimizationGoal = val
	}
	if val, ok := parseAsString(node, "optimization_sub_event"); ok {
		adset.OptimizationSubEvent = val
	}
	if val, ok := parseAsBool(node, "recurring_budget_semantics"); ok {
		adset.RecurringBudgetSemantics = val
	}
	if val, ok := parseAsString(node, "review_feedback"); ok {
		adset.ReviewFeedback = val
	}
	if val, ok := parseAsInt(node, "rf_prediction_id"); ok {
		adset.RfPredictionId = val
	}
	if val, ok := parseAsInt(node, "source_adset_id"); ok {
		adset.SourceAdsetId = val
	}
	if val, ok := parseAsTimestamp(node, "start_time", "2006-01-02T15:04:05-0700"); ok {
		adset.StartTime = val
	}
	if val, ok := parseAsString(node, "status"); ok {
		adset.Status = val
	}
	if val, ok := parseAsTimestamp(node, "updated_time", "2006-01-02T15:04:05-0700"); ok {
		adset.UpdatedTime = val
	}
	if val, ok := parseAsBool(node, "use_new_app_click"); ok {
		adset.UseNewAppClick = val
	}
	return adset, nil
}

func nodeToAdInsight(node Node) (*model.AdInsight, error) {
	adInsight := &model.AdInsight{}

	var ok bool
	var intVal int64
	var strVal string
	var floatVal float64

	if intVal, ok = parseAsNumericString(node, "account_id"); ok {
		adInsight.AccountId = intVal
	}
	if strVal, ok = parseAsString(node, "account_name"); ok {
		adInsight.AccountName = strVal
	}
	if intVal, ok = parseAsNumericString(node, "ad_id"); ok {
		adInsight.AdId = intVal
	}
	if strVal, ok = parseAsString(node, "ad_name"); ok {
		adInsight.AdName = strVal
	}
	if intVal, ok = parseAsNumericString(node, "adset_id"); ok {
		adInsight.AdsetId = intVal
	}
	if strVal, ok = parseAsString(node, "adset_name"); ok {
		adInsight.AdsetName = strVal
	}
	if intVal, ok = parseAsNumericString(node, "cammpaign_id"); ok {
		adInsight.CampaignId = intVal
	}
	if strVal, ok = parseAsString(node, "campaign_name"); ok {
		adInsight.CampaignName = strVal
	}
	if intVal, ok = parseAsNumericString(node, "clicks"); ok {
		adInsight.Clicks = intVal
	}
	if floatVal, ok = parseAsFloatString(node, "cpc"); ok {
		adInsight.Cpc = floatVal
	}
	if floatVal, ok = parseAsFloatString(node, "cpm"); ok {
		adInsight.Cpm = floatVal
	}
	if floatVal, ok = parseAsFloatString(node, "cpp"); ok {
		adInsight.Cpp = floatVal
	}
	if floatVal, ok = parseAsFloatString(node, "ctr"); ok {
		adInsight.Ctr = floatVal
	}
	if timestampVal, ok := parseAsTimestamp(node, "date_start", "2006-01-02"); ok {
		adInsight.DateStart = timestampVal
	}	
	if timestampVal, ok := parseAsTimestamp(node, "date_stop", "2006-01-02"); ok {
		adInsight.DateStop = timestampVal
	}	
	if floatVal, ok = parseAsFloatString(node, "frequency"); ok {
		adInsight.Ctr = floatVal
	}
	if intVal, ok = parseAsNumericString(node, "full_view_impressions"); ok {
		adInsight.FullViewImpressions = intVal
	}
	if intVal, ok = parseAsNumericString(node, "full_view_reach"); ok {
		adInsight.FullViewReach = intVal
	}
	if intVal, ok = parseAsNumericString(node, "impressions"); ok {
		adInsight.Impressions = intVal
	}
	if intVal, ok = parseAsNumericString(node, "reach"); ok {
		adInsight.Impressions = intVal
	}
	if strVal, ok = parseAsString(node, "social_spend"); ok {
		adInsight.SocialSpend = strVal
	}
	if strVal, ok = parseAsString(node, "spend"); ok {
		adInsight.Spend = strVal
	}
	if timestampVal, ok := parseAsTimestamp(node, "updated_time", "2006-01-02"); ok {
		adInsight.UpdatedTime = timestampVal
	}	
	return adInsight, nil
}
func nodeToAd(node Node) (*model.Ad, error) {
	ad := &model.Ad{}

	// Parse and set fields
	var ok bool
	var val int64

	// id
	if val, ok = parseAsNumericString(node, "id"); ok {
		ad.Id = val
	}

	// account_id
	if val, ok = parseAsNumericString(node, "account_id"); ok {
		ad.AccountId = val
	}

	// ad_active_time
	if val, ok = parseAsNumericString(node, "ad_active_time"); ok {
		ad.AdActiveTime = val
	}

	// adset_id
	if val, ok = parseAsNumericString(node, "adset_id"); ok {
		ad.AdsetId = val
	}

	// bid_amount
	if val, ok = parseAsNumericString(node, "bid_amount"); ok {
		ad.BidAmount = val
	}

	// campaign_id
	if val, ok = parseAsNumericString(node, "campaign_id"); ok {
		ad.CampaignId = val
	}

	// configured_status
	if valStr, ok := parseAsString(node, "configured_status"); ok {
		ad.ConfiguredStatus = valStr
	}

	// conversion_domain
	if valStr, ok := parseAsString(node, "conversion_domain"); ok {
		ad.ConversionDomain = valStr
	}

	// created_time
	if ts, ok := parseAsTimestamp(node, "created_time", "2006-01-02T15:04:05-0700"); ok {
		ad.CreatedTime = ts
	}

	// effective_status
	if valStr, ok := parseAsString(node, "effective_status"); ok {
		ad.EffectiveStatus = valStr
	}

	// meta_reward_adgroup_status
	if valStr, ok := parseAsString(node, "meta_reward_adgroup_status"); ok {
		ad.MetaRewardAdgroupStatus = valStr
	}

	// name
	if valStr, ok := parseAsString(node, "name"); ok {
		ad.Name = valStr
	}

	// preview_shareable_link
	if valStr, ok := parseAsString(node, "preview_shareable_link"); ok {
		ad.PreviewShareableLink = valStr
	}

	// source_ad_id
	if val, ok = parseAsNumericString(node, "source_ad_id"); ok {
		ad.SourceAdId = val
	}

	// status
	if valStr, ok := parseAsString(node, "status"); ok {
		ad.Status = valStr
	}

	// updated_time
	if ts, ok := parseAsTimestamp(node, "updated_time", "2006-01-02T15:04:05-0700"); ok {
		ad.UpdatedTime = ts
	}

	return ad, nil
}

func nodeToUserLeadGenInfo(node Node) (*model.UserLeadGenInfo, error){

	lead := &model.UserLeadGenInfo{}
	var ok bool
	var intVal int64
	var strVal string

	if intVal, ok = parseAsNumericString(node, "id"); ok {
		lead.Id = intVal
	}
	if intVal, ok = parseAsNumericString(node, "ad_id"); ok {
		lead.AdId = intVal
	}
	if strVal, ok = parseAsString(node, "ad_name"); ok {
		lead.AdName = strVal
	}
	if intVal, ok = parseAsNumericString(node, "adset_id"); ok {
		lead.AdsetId = intVal
	}
	if strVal, ok = parseAsString(node, "adset_name"); ok {
		lead.AdsetName = strVal
	}
	if intVal, ok = parseAsNumericString(node, "campaign_id"); ok {
		lead.CampaignId = intVal
	}
	if strVal, ok = parseAsString(node, "campaign_name"); ok {
		lead.CampaignName = strVal
	}
	if ts, ok := parseAsTimestamp(node, "created_time", "2006-01-02T15:04:05-0700"); ok {
		lead.CreatedTime = ts
	}

	if mapVal, ok := node["field_data"].([]map[string]interface{}); ok {
		var fieldData []*model.FieldDataEntry	
		// Iterate through field data entry
		for _, jsonEntry := range(mapVal) {
			entry := &model.FieldDataEntry{}
			if entryName, ok := jsonEntry["name"].(string); ok {
				entry.Name = entryName
			}
			if entryValues, ok := jsonEntry["values"].([]string); ok {
				entry.Values = entryValues
			}

			fieldData = append(fieldData, entry)
		}
		lead.FieldData = fieldData
	}

	if intVal, ok = parseAsNumericString(node, "form_id"); ok {
		lead.FormId = intVal
	}
	if boolVal, ok := parseAsBool(node, "is_organic"); ok {
		lead.IsOrganic = boolVal
	}
	if strVal, ok = parseAsString(node, "partner_name"); ok {
		lead.PartnerName = strVal
	}
	if strVal, ok = parseAsString(node, "platform"); ok {
		lead.Platform = strVal
	}
	if strVal, ok = parseAsString(node, "post"); ok {
		lead.Post = strVal
	}

  	return lead, nil 
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


func StdQueryParams() url.Values {
	return url.Values {
		"date_preset": { "maximum" },
		"limit": { "100" },
	}
}

// Extraction functions
func extractCampaigns(AccountId string, bq *storage.BigQueryWriteClient) {
	// CAMPAIGNS
	params := StdQueryParams()
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

func extractAdSets(AccountId string, bq *storage.BigQueryWriteClient){
	
	// AD SETS
	params := StdQueryParams()
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
	edge := fmt.Sprintf("/%s/adsets", AccountId)
	req, err := build_request(edge, params, adsets_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ad Sets")
	data, err := extract(req)
	if err != nil {
		fmt.Println("Error extraction ad sets: ", err)
	}
	fmt.Printf("Total rows extracted: %d\n", len(data))

	// Serialize the Data
	log.Println("Serializing json data into proto messages")
	var adsetsData []protoreflect.ProtoMessage
	for _, node := range(data){ 
		// Convert the Node to the campaings objective
		adsetProto, err := nodeToAdset(node)
		if err != nil {
			log.Fatal(err)
		}
		messageProto := proto.Message(adsetProto)
		adsetsData = append(adsetsData, messageProto)
	}

	// Write data
	log.Println("Writing rows")
	var adset model.AdSet
	desc := getDescriptor(&adset)
	project := os.Getenv("PROJECT_ID")
	dataset := os.Getenv("DATASET_ID")
	table := "adsets"
	trace := "historical-extraction"
	writeRows(bq, desc, adsetsData, project, dataset, table, trace)
}

func extractAds(AccountId string, bq *storage.BigQueryWriteClient){
	
	// ADS
	params := StdQueryParams()
	ads_fields := []string{
		"id",
		"account_id",
		"ad_active_time",
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
	edge := fmt.Sprintf("/%s/ads", AccountId)
	req, err := build_request(edge, params, ads_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ad Sets")
	data, err := extract(req)
	if err != nil {
		fmt.Println("Error extraction ad sets: ", err)
	}
	fmt.Printf("Total rows extracted: %d\n", len(data))

	// Serialize the Data
	log.Println("Serializing json data into proto messages")
	var adsData []protoreflect.ProtoMessage
	for _, node := range(data){ 
		// Convert the Node to the campaings objective
		adProto, err := nodeToAd(node)
		if err != nil {
			log.Fatal(err)
		}
		messageProto := proto.Message(adProto)
		adsData = append(adsData, messageProto)
	}

	// Write data
	log.Println("Writing rows")
	var ad model.Ad
	desc := getDescriptor(&ad)
	project := os.Getenv("PROJECT_ID")
	dataset := os.Getenv("DATASET_ID")
	table := "ads"
	trace := "historical-extraction"
	writeRows(bq, desc, adsData, project, dataset, table, trace)
}
func extractAdInsights(AccountId string, bq *storage.BigQueryWriteClient){
	
	// ADS 
	params := url.Values {
		"date_preset": { "maximum" },
		"level": { "ad" },
		"limit": { "100" },
	}

	ad_insights_fields := []string{
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
	edge := fmt.Sprintf("/%s/insights", AccountId)
	req, err := build_request(edge, params, ad_insights_fields)
	if err != nil {
		fmt.Println("Error building request: ", err)
	}
	fmt.Println("Extracting Ad Insights")
	data, err := extract(req)
	if err != nil {
		fmt.Println("Error extraction ad sets: ", err)
	}
	fmt.Printf("Total rows extracted: %d\n", len(data))

	// Serialize the Data
	log.Println("Serializing json data into proto messages")
	var adInsightsData []protoreflect.ProtoMessage
	for _, node := range(data){ 
		// Convert the Node to the campaings objective
		adInsightProto, err := nodeToAdInsight(node)
		if err != nil {
			log.Fatal(err)
		}
		messageProto := proto.Message(adInsightProto)
		adInsightsData = append(adInsightsData, messageProto)
	}

	// Write data
	log.Println("Writing rows")
	var adInsight model.AdInsight
	desc := getDescriptor(&adInsight)
	project := os.Getenv("PROJECT_ID")
	dataset := os.Getenv("DATASET_ID")
	table := "ad_insights"
	trace := "historical-extraction"
	writeRows(bq, desc, adInsightsData, project, dataset, table, trace)
}

func extractUserLeadGenInfo(AccountId string, bq *storage.BigQueryWriteClient){

	// AD LEADS
	// List existing ads
	params := StdQueryParams()
	adFields := []string{ "id" }
	edge := fmt.Sprintf("%s/ads", AccountId)
	req, err := build_request(edge, params, adFields)
	if err != nil {
		log.Fatal("Error build_request: ", err)
	}
	adNodes, err := extract(req)
	if err != nil {
		log.Fatal("Error extracting ads: ", err)
	}

	// adsIds	
	adsIds := []string{}
	for _, node := range(adNodes){
		if val, ok := node["id"].(string); ok {
			adsIds = append(adsIds, val) 
		}
	}

	// For each ad
	for _, adId := range(adsIds){

		log.Println("Extracting leads for ad ", adId)

		// This edge takes no query parameters and no fields
		params = url.Values{}
		leadsFields := []string{}
		edge := fmt.Sprintf("%s/leads", adId)
		req, err := build_request(edge, params, leadsFields)
		if err != nil {
			log.Println("Error building request: ", err)
			continue
		}

		// Execute the request
		// Lead files are going to be small enough that
		// I won't worry about partitioning by ad_id
		leadNodes, err := extract(req)
		if err != nil {
			fmt.Println("Error extracting data: ", err)
		}

		fmt.Printf("Total rows extracted: %d\n", len(leadNodes))

		// Serialize the Data
		log.Println("Serializing json data into proto messages")
		var userLeadGenInfoData []protoreflect.ProtoMessage
		for _, node := range(leadNodes){ 
			// fmt.Println(node) // Debug
			// Convert the Node to the campaings objective
			userLeadGenInfoProto, err := nodeToUserLeadGenInfo(node)
			if err != nil {
				log.Fatal(err)
			}
			messageProto := proto.Message(userLeadGenInfoProto)
			userLeadGenInfoData = append(userLeadGenInfoData, messageProto)
		}

		// Write data
		log.Println("Writing rows")
		var userLeadGenInfo model.UserLeadGenInfo
		desc := getDescriptor(&userLeadGenInfo)
		project := os.Getenv("PROJECT_ID")
		dataset := os.Getenv("DATASET_ID")
		table := "user_lead_gen_info"
		trace := "historical-extraction"
		writeRows(bq, desc, userLeadGenInfoData, project, dataset, table, trace)
	}
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
			extractCampaigns(*adAccountIdPtr, bq)
		
		case "adsets":
			extractAdSets(*adAccountIdPtr, bq)

		case "ads":
			extractAds(*adAccountIdPtr, bq)
		
		case "ad_insights":
			extractAdInsights(*adAccountIdPtr, bq)

		case "ad_leads":
			extractUserLeadGenInfo(*adAccountIdPtr, bq)
		}
	}
	
}
