// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: model/adset.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId                   int64                  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AssetFeedId                 int64                  `protobuf:"varint,3,opt,name=asset_feed_id,json=assetFeedId,proto3" json:"asset_feed_id,omitempty"`
	BidAmount                   int64                  `protobuf:"varint,4,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount,omitempty"`
	BidStrategy                 string                 `protobuf:"bytes,5,opt,name=bid_strategy,json=bidStrategy,proto3" json:"bid_strategy,omitempty"`
	BillingEvent                string                 `protobuf:"bytes,6,opt,name=billing_event,json=billingEvent,proto3" json:"billing_event,omitempty"`
	BudgetRemaining             int64                  `protobuf:"varint,7,opt,name=budget_remaining,json=budgetRemaining,proto3" json:"budget_remaining,omitempty"`
	CampaignActiveTime          int64                  `protobuf:"varint,8,opt,name=campaign_active_time,json=campaignActiveTime,proto3" json:"campaign_active_time,omitempty"`
	CampaignAttribution         string                 `protobuf:"bytes,9,opt,name=campaign_attribution,json=campaignAttribution,proto3" json:"campaign_attribution,omitempty"`
	CampaignId                  int64                  `protobuf:"varint,10,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	ConfiguredStatus            string                 `protobuf:"bytes,11,opt,name=configured_status,json=configuredStatus,proto3" json:"configured_status,omitempty"`
	CreatedTime                 *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	DailyMinSpendTarget         int64                  `protobuf:"varint,14,opt,name=daily_min_spend_target,json=dailyMinSpendTarget,proto3" json:"daily_min_spend_target,omitempty"`
	DailySpendCap               int64                  `protobuf:"varint,15,opt,name=daily_spend_cap,json=dailySpendCap,proto3" json:"daily_spend_cap,omitempty"`
	DestinationType             string                 `protobuf:"bytes,16,opt,name=destination_type,json=destinationType,proto3" json:"destination_type,omitempty"`
	DsaBeneficiary              string                 `protobuf:"bytes,17,opt,name=dsa_beneficiary,json=dsaBeneficiary,proto3" json:"dsa_beneficiary,omitempty"`
	DsaPayor                    string                 `protobuf:"bytes,18,opt,name=dsa_payor,json=dsaPayor,proto3" json:"dsa_payor,omitempty"`
	EffectiveStatus             string                 `protobuf:"bytes,19,opt,name=effective_status,json=effectiveStatus,proto3" json:"effective_status,omitempty"`
	EndTime                     *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	InstagramActorId            int64                  `protobuf:"varint,21,opt,name=instagram_actor_id,json=instagramActorId,proto3" json:"instagram_actor_id,omitempty"`
	IsBudgetScheduleEnabled     bool                   `protobuf:"varint,22,opt,name=is_budget_schedule_enabled,json=isBudgetScheduleEnabled,proto3" json:"is_budget_schedule_enabled,omitempty"`
	IsDynamicCreative           bool                   `protobuf:"varint,23,opt,name=is_dynamic_creative,json=isDynamicCreative,proto3" json:"is_dynamic_creative,omitempty"`
	LifetimeBudget              int64                  `protobuf:"varint,24,opt,name=lifetime_budget,json=lifetimeBudget,proto3" json:"lifetime_budget,omitempty"`
	LifetimeImps                int64                  `protobuf:"varint,25,opt,name=lifetime_imps,json=lifetimeImps,proto3" json:"lifetime_imps,omitempty"`
	LifetimeMinSpendTarget      int64                  `protobuf:"varint,26,opt,name=lifetime_min_spend_target,json=lifetimeMinSpendTarget,proto3" json:"lifetime_min_spend_target,omitempty"`
	LifetimeSpendCap            int64                  `protobuf:"varint,27,opt,name=lifetime_spend_cap,json=lifetimeSpendCap,proto3" json:"lifetime_spend_cap,omitempty"`
	MultiOptimizationGoalWeight string                 `protobuf:"bytes,28,opt,name=multi_optimization_goal_weight,json=multiOptimizationGoalWeight,proto3" json:"multi_optimization_goal_weight,omitempty"`
	Name                        string                 `protobuf:"bytes,29,opt,name=name,proto3" json:"name,omitempty"`
	OptimizationGoal            string                 `protobuf:"bytes,30,opt,name=optimization_goal,json=optimizationGoal,proto3" json:"optimization_goal,omitempty"`
	OptimizationSubEvent        string                 `protobuf:"bytes,31,opt,name=optimization_sub_event,json=optimizationSubEvent,proto3" json:"optimization_sub_event,omitempty"`
	RecurringBudgetSemantics    bool                   `protobuf:"varint,32,opt,name=recurring_budget_semantics,json=recurringBudgetSemantics,proto3" json:"recurring_budget_semantics,omitempty"`
	ReviewFeedback              string                 `protobuf:"bytes,33,opt,name=review_feedback,json=reviewFeedback,proto3" json:"review_feedback,omitempty"`
	RfPredictionId              int64                  `protobuf:"varint,34,opt,name=rf_prediction_id,json=rfPredictionId,proto3" json:"rf_prediction_id,omitempty"`
	SourceAdsetId               int64                  `protobuf:"varint,35,opt,name=source_adset_id,json=sourceAdsetId,proto3" json:"source_adset_id,omitempty"`
	StartTime                   *timestamppb.Timestamp `protobuf:"bytes,36,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Status                      string                 `protobuf:"bytes,37,opt,name=status,proto3" json:"status,omitempty"`
	UpdatedTime                 *timestamppb.Timestamp `protobuf:"bytes,38,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	UseNewAppClick              bool                   `protobuf:"varint,39,opt,name=use_new_app_click,json=useNewAppClick,proto3" json:"use_new_app_click,omitempty"`
}

func (x *AdSet) Reset() {
	*x = AdSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_adset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdSet) ProtoMessage() {}

func (x *AdSet) ProtoReflect() protoreflect.Message {
	mi := &file_model_adset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdSet.ProtoReflect.Descriptor instead.
func (*AdSet) Descriptor() ([]byte, []int) {
	return file_model_adset_proto_rawDescGZIP(), []int{0}
}

func (x *AdSet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdSet) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *AdSet) GetAssetFeedId() int64 {
	if x != nil {
		return x.AssetFeedId
	}
	return 0
}

func (x *AdSet) GetBidAmount() int64 {
	if x != nil {
		return x.BidAmount
	}
	return 0
}

func (x *AdSet) GetBidStrategy() string {
	if x != nil {
		return x.BidStrategy
	}
	return ""
}

func (x *AdSet) GetBillingEvent() string {
	if x != nil {
		return x.BillingEvent
	}
	return ""
}

func (x *AdSet) GetBudgetRemaining() int64 {
	if x != nil {
		return x.BudgetRemaining
	}
	return 0
}

func (x *AdSet) GetCampaignActiveTime() int64 {
	if x != nil {
		return x.CampaignActiveTime
	}
	return 0
}

func (x *AdSet) GetCampaignAttribution() string {
	if x != nil {
		return x.CampaignAttribution
	}
	return ""
}

func (x *AdSet) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *AdSet) GetConfiguredStatus() string {
	if x != nil {
		return x.ConfiguredStatus
	}
	return ""
}

func (x *AdSet) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *AdSet) GetDailyMinSpendTarget() int64 {
	if x != nil {
		return x.DailyMinSpendTarget
	}
	return 0
}

func (x *AdSet) GetDailySpendCap() int64 {
	if x != nil {
		return x.DailySpendCap
	}
	return 0
}

func (x *AdSet) GetDestinationType() string {
	if x != nil {
		return x.DestinationType
	}
	return ""
}

func (x *AdSet) GetDsaBeneficiary() string {
	if x != nil {
		return x.DsaBeneficiary
	}
	return ""
}

func (x *AdSet) GetDsaPayor() string {
	if x != nil {
		return x.DsaPayor
	}
	return ""
}

func (x *AdSet) GetEffectiveStatus() string {
	if x != nil {
		return x.EffectiveStatus
	}
	return ""
}

func (x *AdSet) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *AdSet) GetInstagramActorId() int64 {
	if x != nil {
		return x.InstagramActorId
	}
	return 0
}

func (x *AdSet) GetIsBudgetScheduleEnabled() bool {
	if x != nil {
		return x.IsBudgetScheduleEnabled
	}
	return false
}

func (x *AdSet) GetIsDynamicCreative() bool {
	if x != nil {
		return x.IsDynamicCreative
	}
	return false
}

func (x *AdSet) GetLifetimeBudget() int64 {
	if x != nil {
		return x.LifetimeBudget
	}
	return 0
}

func (x *AdSet) GetLifetimeImps() int64 {
	if x != nil {
		return x.LifetimeImps
	}
	return 0
}

func (x *AdSet) GetLifetimeMinSpendTarget() int64 {
	if x != nil {
		return x.LifetimeMinSpendTarget
	}
	return 0
}

func (x *AdSet) GetLifetimeSpendCap() int64 {
	if x != nil {
		return x.LifetimeSpendCap
	}
	return 0
}

func (x *AdSet) GetMultiOptimizationGoalWeight() string {
	if x != nil {
		return x.MultiOptimizationGoalWeight
	}
	return ""
}

func (x *AdSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdSet) GetOptimizationGoal() string {
	if x != nil {
		return x.OptimizationGoal
	}
	return ""
}

func (x *AdSet) GetOptimizationSubEvent() string {
	if x != nil {
		return x.OptimizationSubEvent
	}
	return ""
}

func (x *AdSet) GetRecurringBudgetSemantics() bool {
	if x != nil {
		return x.RecurringBudgetSemantics
	}
	return false
}

func (x *AdSet) GetReviewFeedback() string {
	if x != nil {
		return x.ReviewFeedback
	}
	return ""
}

func (x *AdSet) GetRfPredictionId() int64 {
	if x != nil {
		return x.RfPredictionId
	}
	return 0
}

func (x *AdSet) GetSourceAdsetId() int64 {
	if x != nil {
		return x.SourceAdsetId
	}
	return 0
}

func (x *AdSet) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *AdSet) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AdSet) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *AdSet) GetUseNewAppClick() bool {
	if x != nil {
		return x.UseNewAppClick
	}
	return false
}

var File_model_adset_proto protoreflect.FileDescriptor

var file_model_adset_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x64, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x0d, 0x0a, 0x05, 0x41, 0x64, 0x53, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x69, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x6d,
	0x69, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x4d, 0x69, 0x6e, 0x53,
	0x70, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x43,
	0x61, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x64, 0x73, 0x61, 0x5f, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x73, 0x61, 0x42, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x73, 0x61, 0x5f, 0x70, 0x61,
	0x79, 0x6f, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x73, 0x61, 0x50, 0x61,
	0x79, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x5f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69, 0x73, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69,
	0x73, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x69, 0x66, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6d, 0x70, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x70, 0x73, 0x12, 0x39,
	0x0a, 0x19, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x16, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6e, 0x53, 0x70,
	0x65, 0x6e, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x61, 0x70, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x53,
	0x70, 0x65, 0x6e, 0x64, 0x43, 0x61, 0x70, 0x12, 0x43, 0x0a, 0x1e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x6f,
	0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x34, 0x0a,
	0x16, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75,
	0x62, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x1a, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x66,
	0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x22,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x66, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61,
	0x64, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x23, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x24, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x26, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x18, 0x27, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x4e, 0x65,
	0x77, 0x41, 0x70, 0x70, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x42, 0x1e, 0x5a, 0x1c, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_model_adset_proto_rawDescOnce sync.Once
	file_model_adset_proto_rawDescData = file_model_adset_proto_rawDesc
)

func file_model_adset_proto_rawDescGZIP() []byte {
	file_model_adset_proto_rawDescOnce.Do(func() {
		file_model_adset_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_adset_proto_rawDescData)
	})
	return file_model_adset_proto_rawDescData
}

var file_model_adset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_adset_proto_goTypes = []interface{}{
	(*AdSet)(nil),                 // 0: AdSet
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_model_adset_proto_depIdxs = []int32{
	1, // 0: AdSet.created_time:type_name -> google.protobuf.Timestamp
	1, // 1: AdSet.end_time:type_name -> google.protobuf.Timestamp
	1, // 2: AdSet.start_time:type_name -> google.protobuf.Timestamp
	1, // 3: AdSet.updated_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_model_adset_proto_init() }
func file_model_adset_proto_init() {
	if File_model_adset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_adset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_adset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_adset_proto_goTypes,
		DependencyIndexes: file_model_adset_proto_depIdxs,
		MessageInfos:      file_model_adset_proto_msgTypes,
	}.Build()
	File_model_adset_proto = out.File
	file_model_adset_proto_rawDesc = nil
	file_model_adset_proto_goTypes = nil
	file_model_adset_proto_depIdxs = nil
}
