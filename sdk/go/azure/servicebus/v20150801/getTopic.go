// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of topic resource.
//
// Deprecated: Version v20150801 will be removed in the next major version of the provider. Upgrade to version v20170401 or later.
func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:servicebus/v20150801:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// Description of topic resource.
type LookupTopicResult struct {
	// Last time the message was sent, or a request was received, for this topic.
	AccessedAt string `pulumi:"accessedAt"`
	// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Message Count Details.
	CountDetails MessageCountDetailsResponse `pulumi:"countDetails"`
	// Exact time the message was created.
	CreatedAt string `pulumi:"createdAt"`
	// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive *string `pulumi:"defaultMessageTimeToLive"`
	// TimeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Entity availability status for the topic.
	EntityAvailabilityStatus *string `pulumi:"entityAvailabilityStatus"`
	// Whether messages should be filtered before publishing.
	FilteringMessagesBeforePublishing *bool `pulumi:"filteringMessagesBeforePublishing"`
	// Resource Id
	Id string `pulumi:"id"`
	// Value that indicates whether the message is accessible anonymously.
	IsAnonymousAccessible *bool `pulumi:"isAnonymousAccessible"`
	IsExpress             *bool `pulumi:"isExpress"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
	MaxSizeInMegabytes *float64 `pulumi:"maxSizeInMegabytes"`
	// Resource name
	Name string `pulumi:"name"`
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Size of the topic, in bytes.
	SizeInBytes float64 `pulumi:"sizeInBytes"`
	// Enumerates the possible values for the status of a messaging entity.
	Status *string `pulumi:"status"`
	// Number of subscriptions.
	SubscriptionCount int `pulumi:"subscriptionCount"`
	// Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `pulumi:"supportOrdering"`
	// Resource type
	Type string `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupTopicOutput(ctx *pulumi.Context, args LookupTopicOutputArgs, opts ...pulumi.InvokeOption) LookupTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicResult, error) {
			args := v.(LookupTopicArgs)
			r, err := LookupTopic(ctx, &args, opts...)
			var s LookupTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicResultOutput)
}

type LookupTopicOutputArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The topic name.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicArgs)(nil)).Elem()
}

// Description of topic resource.
type LookupTopicResultOutput struct{ *pulumi.OutputState }

func (LookupTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicResult)(nil)).Elem()
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutput() LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutputWithContext(ctx context.Context) LookupTopicResultOutput {
	return o
}

// Last time the message was sent, or a request was received, for this topic.
func (o LookupTopicResultOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.AccessedAt }).(pulumi.StringOutput)
}

// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
func (o LookupTopicResultOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

// Message Count Details.
func (o LookupTopicResultOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v LookupTopicResult) MessageCountDetailsResponse { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

// Exact time the message was created.
func (o LookupTopicResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
func (o LookupTopicResultOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

// TimeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
func (o LookupTopicResultOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

// Value that indicates whether server-side batched operations are enabled.
func (o LookupTopicResultOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
func (o LookupTopicResultOutput) EnableExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnableExpress }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
func (o LookupTopicResultOutput) EnablePartitioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnablePartitioning }).(pulumi.BoolPtrOutput)
}

// Entity availability status for the topic.
func (o LookupTopicResultOutput) EntityAvailabilityStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.EntityAvailabilityStatus }).(pulumi.StringPtrOutput)
}

// Whether messages should be filtered before publishing.
func (o LookupTopicResultOutput) FilteringMessagesBeforePublishing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.FilteringMessagesBeforePublishing }).(pulumi.BoolPtrOutput)
}

// Resource Id
func (o LookupTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// Value that indicates whether the message is accessible anonymously.
func (o LookupTopicResultOutput) IsAnonymousAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.IsAnonymousAccessible }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) IsExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.IsExpress }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o LookupTopicResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
func (o LookupTopicResultOutput) MaxSizeInMegabytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *float64 { return v.MaxSizeInMegabytes }).(pulumi.Float64PtrOutput)
}

// Resource name
func (o LookupTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// Value indicating if this topic requires duplicate detection.
func (o LookupTopicResultOutput) RequiresDuplicateDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.RequiresDuplicateDetection }).(pulumi.BoolPtrOutput)
}

// Size of the topic, in bytes.
func (o LookupTopicResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupTopicResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

// Enumerates the possible values for the status of a messaging entity.
func (o LookupTopicResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Number of subscriptions.
func (o LookupTopicResultOutput) SubscriptionCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTopicResult) int { return v.SubscriptionCount }).(pulumi.IntOutput)
}

// Value that indicates whether the topic supports ordering.
func (o LookupTopicResultOutput) SupportOrdering() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.SupportOrdering }).(pulumi.BoolPtrOutput)
}

// Resource type
func (o LookupTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

// The exact time the message was updated.
func (o LookupTopicResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicResultOutput{})
}
