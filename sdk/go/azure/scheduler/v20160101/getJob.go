// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: Version v20160101 will be removed in the next major version of the provider. Upgrade to version v20160301 or later.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:scheduler/v20160101:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The job collection name.
	JobCollectionName string `pulumi:"jobCollectionName"`
	// The job name.
	JobName string `pulumi:"jobName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobResult struct {
	// Gets the job resource identifier.
	Id string `pulumi:"id"`
	// Gets the job resource name.
	Name string `pulumi:"name"`
	// Gets or sets the job properties.
	Properties JobPropertiesResponse `pulumi:"properties"`
	// Gets the job resource type.
	Type string `pulumi:"type"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	// The job collection name.
	JobCollectionName pulumi.StringInput `pulumi:"jobCollectionName"`
	// The job name.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

// Gets the job resource identifier.
func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets the job resource name.
func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the job properties.
func (o LookupJobResultOutput) Properties() JobPropertiesResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobPropertiesResponse { return v.Properties }).(JobPropertiesResponseOutput)
}

// Gets the job resource type.
func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
