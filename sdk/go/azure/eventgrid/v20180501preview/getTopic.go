// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventGrid Topic
//
// Deprecated: Version v20180501preview will be removed in the next major version of the provider. Upgrade to version v20200401preview or later.
func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20180501preview:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTopicArgs struct {
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the topic
	TopicName string `pulumi:"topicName"`
}

// EventGrid Topic
type LookupTopicResult struct {
	// Endpoint for the topic.
	Endpoint string `pulumi:"endpoint"`
	// Fully qualified identifier of the resource
	Id string `pulumi:"id"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location string `pulumi:"location"`
	// Name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the topic.
	ProvisioningState string `pulumi:"provisioningState"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupTopicResult
func (val *LookupTopicResult) Defaults() *LookupTopicResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InputSchema) {
		inputSchema_ := "EventGridSchema"
		tmp.InputSchema = &inputSchema_
	}
	return &tmp
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
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the topic
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicArgs)(nil)).Elem()
}

// EventGrid Topic
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

// Endpoint for the topic.
func (o LookupTopicResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Fully qualified identifier of the resource
func (o LookupTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// This determines the format that Event Grid should expect for incoming events published to the topic.
func (o LookupTopicResultOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.InputSchema }).(pulumi.StringPtrOutput)
}

// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
func (o LookupTopicResultOutput) InputSchemaMapping() JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *JsonInputSchemaMappingResponse { return v.InputSchemaMapping }).(JsonInputSchemaMappingResponsePtrOutput)
}

// Location of the resource
func (o LookupTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource
func (o LookupTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the topic.
func (o LookupTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Tags of the resource
func (o LookupTopicResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTopicResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource
func (o LookupTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicResultOutput{})
}
