// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Capacity pool resource
//
// Deprecated: Version v20200601 will be removed in the next major version of the provider. Upgrade to version v20201201 or later.
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:netapp/v20200601:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPoolArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Capacity pool resource
type LookupPoolResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// UUID v4 used to identify the Pool
	PoolId string `pulumi:"poolId"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// The qos type of the pool
	QosType *string `pulumi:"qosType"`
	// The service level of the file system
	ServiceLevel string `pulumi:"serviceLevel"`
	// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
	Size float64 `pulumi:"size"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Total throughput of pool in Mibps
	TotalThroughputMibps float64 `pulumi:"totalThroughputMibps"`
	// Resource type
	Type string `pulumi:"type"`
	// Utilized throughput of pool in Mibps
	UtilizedThroughputMibps float64 `pulumi:"utilizedThroughputMibps"`
}

// Defaults sets the appropriate defaults for LookupPoolResult
func (val *LookupPoolResult) Defaults() *LookupPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.QosType) {
		qosType_ := "Auto"
		tmp.QosType = &qosType_
	}
	if isZero(tmp.ServiceLevel) {
		tmp.ServiceLevel = "Premium"
	}
	return &tmp
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPoolResult, error) {
			args := v.(LookupPoolArgs)
			r, err := LookupPool(ctx, &args, opts...)
			var s LookupPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPoolResultOutput)
}

type LookupPoolOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}

// Capacity pool resource
type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

// Resource Id
func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// UUID v4 used to identify the Pool
func (o LookupPoolResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.PoolId }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The qos type of the pool
func (o LookupPoolResultOutput) QosType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.QosType }).(pulumi.StringPtrOutput)
}

// The service level of the file system
func (o LookupPoolResultOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ServiceLevel }).(pulumi.StringOutput)
}

// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
func (o LookupPoolResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// Resource tags
func (o LookupPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Total throughput of pool in Mibps
func (o LookupPoolResultOutput) TotalThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.TotalThroughputMibps }).(pulumi.Float64Output)
}

// Resource type
func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// Utilized throughput of pool in Mibps
func (o LookupPoolResultOutput) UtilizedThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.UtilizedThroughputMibps }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
