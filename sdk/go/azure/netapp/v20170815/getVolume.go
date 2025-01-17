// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Volume resource
//
// Deprecated: Version v20170815 will be removed in the next major version of the provider. Upgrade to version v20201201 or later.
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20170815:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVolumeArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// Volume resource
type LookupVolumeResult struct {
	// A unique file path for the volume. Used when creating mount targets
	CreationToken string `pulumi:"creationToken"`
	// Export policy rule
	ExportPolicy *VolumePropertiesResponseExportPolicy `pulumi:"exportPolicy"`
	// Unique FileSystem Identifier.
	FileSystemId string `pulumi:"fileSystemId"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// The service level of the file system
	ServiceLevel string `pulumi:"serviceLevel"`
	// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
	SubnetId *string `pulumi:"subnetId"`
	// Resource tags
	Tags interface{} `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB.
	UsageThreshold *float64 `pulumi:"usageThreshold"`
}

// Defaults sets the appropriate defaults for LookupVolumeResult
func (val *LookupVolumeResult) Defaults() *LookupVolumeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceLevel) {
		tmp.ServiceLevel = "Premium"
	}
	if isZero(tmp.UsageThreshold) {
		usageThreshold_ := 107374182400.0
		tmp.UsageThreshold = &usageThreshold_
	}
	return &tmp
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}

// Volume resource
type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

// A unique file path for the volume. Used when creating mount targets
func (o LookupVolumeResultOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.CreationToken }).(pulumi.StringOutput)
}

// Export policy rule
func (o LookupVolumeResultOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumePropertiesResponseExportPolicy { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

// Unique FileSystem Identifier.
func (o LookupVolumeResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The service level of the file system
func (o LookupVolumeResultOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ServiceLevel }).(pulumi.StringOutput)
}

// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
func (o LookupVolumeResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o LookupVolumeResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVolumeResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

// Resource type
func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB.
func (o LookupVolumeResultOutput) UsageThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.UsageThreshold }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
