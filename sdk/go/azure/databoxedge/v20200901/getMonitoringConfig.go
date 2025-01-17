// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The metric setting details for the role
//
// Deprecated: Version v20200901 will be removed in the next major version of the provider. Upgrade to version v20201201 or later.
func LookupMonitoringConfig(ctx *pulumi.Context, args *LookupMonitoringConfigArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringConfigResult, error) {
	var rv LookupMonitoringConfigResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901:getMonitoringConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringConfigArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The role name.
	RoleName string `pulumi:"roleName"`
}

// The metric setting details for the role
type LookupMonitoringConfigResult struct {
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The metrics configuration details
	MetricConfigurations []MetricConfigurationResponse `pulumi:"metricConfigurations"`
	// The object name.
	Name string `pulumi:"name"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupMonitoringConfigOutput(ctx *pulumi.Context, args LookupMonitoringConfigOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoringConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitoringConfigResult, error) {
			args := v.(LookupMonitoringConfigArgs)
			r, err := LookupMonitoringConfig(ctx, &args, opts...)
			var s LookupMonitoringConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitoringConfigResultOutput)
}

type LookupMonitoringConfigOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The role name.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (LookupMonitoringConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringConfigArgs)(nil)).Elem()
}

// The metric setting details for the role
type LookupMonitoringConfigResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoringConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringConfigResult)(nil)).Elem()
}

func (o LookupMonitoringConfigResultOutput) ToLookupMonitoringConfigResultOutput() LookupMonitoringConfigResultOutput {
	return o
}

func (o LookupMonitoringConfigResultOutput) ToLookupMonitoringConfigResultOutputWithContext(ctx context.Context) LookupMonitoringConfigResultOutput {
	return o
}

// The path ID that uniquely identifies the object.
func (o LookupMonitoringConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// The metrics configuration details
func (o LookupMonitoringConfigResultOutput) MetricConfigurations() MetricConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) []MetricConfigurationResponse { return v.MetricConfigurations }).(MetricConfigurationResponseArrayOutput)
}

// The object name.
func (o LookupMonitoringConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The hierarchical type of the object.
func (o LookupMonitoringConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoringConfigResultOutput{})
}
