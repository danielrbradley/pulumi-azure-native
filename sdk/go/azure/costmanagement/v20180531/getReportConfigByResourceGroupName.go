// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180531

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A report config resource.
//
// Deprecated: Version v20180531 will be removed in the next major version of the provider. Upgrade to version v20180801preview or later.
func LookupReportConfigByResourceGroupName(ctx *pulumi.Context, args *LookupReportConfigByResourceGroupNameArgs, opts ...pulumi.InvokeOption) (*LookupReportConfigByResourceGroupNameResult, error) {
	var rv LookupReportConfigByResourceGroupNameResult
	err := ctx.Invoke("azure-native:costmanagement/v20180531:getReportConfigByResourceGroupName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportConfigByResourceGroupNameArgs struct {
	// Report Config Name.
	ReportConfigName string `pulumi:"reportConfigName"`
	// Azure Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A report config resource.
type LookupReportConfigByResourceGroupNameResult struct {
	// Has definition for the report config.
	Definition ReportConfigDefinitionResponse `pulumi:"definition"`
	// Has delivery information for the report config.
	DeliveryInfo ReportConfigDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Has schedule information for the report config.
	Schedule *ReportConfigScheduleResponse `pulumi:"schedule"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupReportConfigByResourceGroupNameOutput(ctx *pulumi.Context, args LookupReportConfigByResourceGroupNameOutputArgs, opts ...pulumi.InvokeOption) LookupReportConfigByResourceGroupNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportConfigByResourceGroupNameResult, error) {
			args := v.(LookupReportConfigByResourceGroupNameArgs)
			r, err := LookupReportConfigByResourceGroupName(ctx, &args, opts...)
			var s LookupReportConfigByResourceGroupNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportConfigByResourceGroupNameResultOutput)
}

type LookupReportConfigByResourceGroupNameOutputArgs struct {
	// Report Config Name.
	ReportConfigName pulumi.StringInput `pulumi:"reportConfigName"`
	// Azure Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReportConfigByResourceGroupNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportConfigByResourceGroupNameArgs)(nil)).Elem()
}

// A report config resource.
type LookupReportConfigByResourceGroupNameResultOutput struct{ *pulumi.OutputState }

func (LookupReportConfigByResourceGroupNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportConfigByResourceGroupNameResult)(nil)).Elem()
}

func (o LookupReportConfigByResourceGroupNameResultOutput) ToLookupReportConfigByResourceGroupNameResultOutput() LookupReportConfigByResourceGroupNameResultOutput {
	return o
}

func (o LookupReportConfigByResourceGroupNameResultOutput) ToLookupReportConfigByResourceGroupNameResultOutputWithContext(ctx context.Context) LookupReportConfigByResourceGroupNameResultOutput {
	return o
}

// Has definition for the report config.
func (o LookupReportConfigByResourceGroupNameResultOutput) Definition() ReportConfigDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) ReportConfigDefinitionResponse {
		return v.Definition
	}).(ReportConfigDefinitionResponseOutput)
}

// Has delivery information for the report config.
func (o LookupReportConfigByResourceGroupNameResultOutput) DeliveryInfo() ReportConfigDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) ReportConfigDeliveryInfoResponse {
		return v.DeliveryInfo
	}).(ReportConfigDeliveryInfoResponseOutput)
}

// The format of the report being delivered.
func (o LookupReportConfigByResourceGroupNameResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupReportConfigByResourceGroupNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupReportConfigByResourceGroupNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) string { return v.Name }).(pulumi.StringOutput)
}

// Has schedule information for the report config.
func (o LookupReportConfigByResourceGroupNameResultOutput) Schedule() ReportConfigScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) *ReportConfigScheduleResponse { return v.Schedule }).(ReportConfigScheduleResponsePtrOutput)
}

// Resource tags.
func (o LookupReportConfigByResourceGroupNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupReportConfigByResourceGroupNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportConfigByResourceGroupNameResultOutput{})
}
