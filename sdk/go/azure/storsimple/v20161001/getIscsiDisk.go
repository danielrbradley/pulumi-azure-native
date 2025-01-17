// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The iSCSI disk.
//
// Deprecated: Version v20161001 will be removed in the next major version of the provider. Upgrade to version v20170601 or later.
func LookupIscsiDisk(ctx *pulumi.Context, args *LookupIscsiDiskArgs, opts ...pulumi.InvokeOption) (*LookupIscsiDiskResult, error) {
	var rv LookupIscsiDiskResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getIscsiDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiDiskArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The disk name.
	DiskName string `pulumi:"diskName"`
	// The iSCSI server name.
	IscsiServerName string `pulumi:"iscsiServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The iSCSI disk.
type LookupIscsiDiskResult struct {
	// The access control records.
	AccessControlRecords []string `pulumi:"accessControlRecords"`
	// The data policy.
	DataPolicy string `pulumi:"dataPolicy"`
	// The description.
	Description *string `pulumi:"description"`
	// The disk status.
	DiskStatus string `pulumi:"diskStatus"`
	// The identifier.
	Id string `pulumi:"id"`
	// The local used capacity in bytes.
	LocalUsedCapacityInBytes float64 `pulumi:"localUsedCapacityInBytes"`
	// The monitoring.
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// The name.
	Name string `pulumi:"name"`
	// The provisioned capacity in bytes.
	ProvisionedCapacityInBytes float64 `pulumi:"provisionedCapacityInBytes"`
	// The type.
	Type string `pulumi:"type"`
	// The used capacity in bytes.
	UsedCapacityInBytes float64 `pulumi:"usedCapacityInBytes"`
}

func LookupIscsiDiskOutput(ctx *pulumi.Context, args LookupIscsiDiskOutputArgs, opts ...pulumi.InvokeOption) LookupIscsiDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIscsiDiskResult, error) {
			args := v.(LookupIscsiDiskArgs)
			r, err := LookupIscsiDisk(ctx, &args, opts...)
			var s LookupIscsiDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIscsiDiskResultOutput)
}

type LookupIscsiDiskOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The disk name.
	DiskName pulumi.StringInput `pulumi:"diskName"`
	// The iSCSI server name.
	IscsiServerName pulumi.StringInput `pulumi:"iscsiServerName"`
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIscsiDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiDiskArgs)(nil)).Elem()
}

// The iSCSI disk.
type LookupIscsiDiskResultOutput struct{ *pulumi.OutputState }

func (LookupIscsiDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiDiskResult)(nil)).Elem()
}

func (o LookupIscsiDiskResultOutput) ToLookupIscsiDiskResultOutput() LookupIscsiDiskResultOutput {
	return o
}

func (o LookupIscsiDiskResultOutput) ToLookupIscsiDiskResultOutputWithContext(ctx context.Context) LookupIscsiDiskResultOutput {
	return o
}

// The access control records.
func (o LookupIscsiDiskResultOutput) AccessControlRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) []string { return v.AccessControlRecords }).(pulumi.StringArrayOutput)
}

// The data policy.
func (o LookupIscsiDiskResultOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.DataPolicy }).(pulumi.StringOutput)
}

// The description.
func (o LookupIscsiDiskResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The disk status.
func (o LookupIscsiDiskResultOutput) DiskStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.DiskStatus }).(pulumi.StringOutput)
}

// The identifier.
func (o LookupIscsiDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

// The local used capacity in bytes.
func (o LookupIscsiDiskResultOutput) LocalUsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupIscsiDiskResult) float64 { return v.LocalUsedCapacityInBytes }).(pulumi.Float64Output)
}

// The monitoring.
func (o LookupIscsiDiskResultOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.MonitoringStatus }).(pulumi.StringOutput)
}

// The name.
func (o LookupIscsiDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioned capacity in bytes.
func (o LookupIscsiDiskResultOutput) ProvisionedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupIscsiDiskResult) float64 { return v.ProvisionedCapacityInBytes }).(pulumi.Float64Output)
}

// The type.
func (o LookupIscsiDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

// The used capacity in bytes.
func (o LookupIscsiDiskResultOutput) UsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupIscsiDiskResult) float64 { return v.UsedCapacityInBytes }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupIscsiDiskResultOutput{})
}
