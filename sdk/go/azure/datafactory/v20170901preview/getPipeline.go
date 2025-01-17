// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pipeline resource type.
//
// Deprecated: Version v20170901preview will be removed in the next major version of the provider. Upgrade to version v20180601 or later.
func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("azure-native:datafactory/v20170901preview:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The pipeline name.
	PipelineName string `pulumi:"pipelineName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Pipeline resource type.
type LookupPipelineResult struct {
	// List of activities in pipeline.
	Activities []interface{} `pulumi:"activities"`
	// List of tags that can be used for describing the Pipeline.
	Annotations []interface{} `pulumi:"annotations"`
	// The max number of concurrent runs for the pipeline.
	Concurrency *int `pulumi:"concurrency"`
	// The description of the pipeline.
	Description *string `pulumi:"description"`
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// List of parameters for pipeline.
	Parameters map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineResult, error) {
			args := v.(LookupPipelineArgs)
			r, err := LookupPipeline(ctx, &args, opts...)
			var s LookupPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineResultOutput)
}

type LookupPipelineOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The pipeline name.
	PipelineName pulumi.StringInput `pulumi:"pipelineName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}

// Pipeline resource type.
type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

// List of activities in pipeline.
func (o LookupPipelineResultOutput) Activities() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []interface{} { return v.Activities }).(pulumi.ArrayOutput)
}

// List of tags that can be used for describing the Pipeline.
func (o LookupPipelineResultOutput) Annotations() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []interface{} { return v.Annotations }).(pulumi.ArrayOutput)
}

// The max number of concurrent runs for the pipeline.
func (o LookupPipelineResultOutput) Concurrency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *int { return v.Concurrency }).(pulumi.IntPtrOutput)
}

// The description of the pipeline.
func (o LookupPipelineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Etag identifies change in the resource.
func (o LookupPipelineResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of parameters for pipeline.
func (o LookupPipelineResultOutput) Parameters() ParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v LookupPipelineResult) map[string]ParameterSpecificationResponse { return v.Parameters }).(ParameterSpecificationResponseMapOutput)
}

// The resource type.
func (o LookupPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
