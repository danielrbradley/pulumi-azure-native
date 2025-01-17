// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A storage system being cached by a Cache.
//
// Deprecated: Version v20190801preview will be removed in the next major version of the provider. Upgrade to version v20210301 or later.
func LookupStorageTarget(ctx *pulumi.Context, args *LookupStorageTargetArgs, opts ...pulumi.InvokeOption) (*LookupStorageTargetResult, error) {
	var rv LookupStorageTargetResult
	err := ctx.Invoke("azure-native:storagecache/v20190801preview:getStorageTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageTargetArgs struct {
	// Name of cache.
	CacheName string `pulumi:"cacheName"`
	// Target resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of storage target.
	StorageTargetName string `pulumi:"storageTargetName"`
}

// A storage system being cached by a Cache.
type LookupStorageTargetResult struct {
	// Properties when clfs target.
	Clfs *ClfsTargetResponse `pulumi:"clfs"`
	// Resource Id
	Id string `pulumi:"id"`
	// List of cache namespace to target namespace associations.
	Junctions []NamespaceJunctionResponse `pulumi:"junctions"`
	// A fully qualified URL.
	Name string `pulumi:"name"`
	// Properties when nfs3 target.
	Nfs3 *Nfs3TargetResponse `pulumi:"nfs3"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *string `pulumi:"provisioningState"`
	// Type for storage target.
	TargetType *string `pulumi:"targetType"`
	// Type for the storage target; Microsoft.StorageCache/Cache/StorageTarget
	Type string `pulumi:"type"`
	// Properties when unknown target.
	Unknown *UnknownTargetResponse `pulumi:"unknown"`
}

func LookupStorageTargetOutput(ctx *pulumi.Context, args LookupStorageTargetOutputArgs, opts ...pulumi.InvokeOption) LookupStorageTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageTargetResult, error) {
			args := v.(LookupStorageTargetArgs)
			r, err := LookupStorageTarget(ctx, &args, opts...)
			var s LookupStorageTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageTargetResultOutput)
}

type LookupStorageTargetOutputArgs struct {
	// Name of cache.
	CacheName pulumi.StringInput `pulumi:"cacheName"`
	// Target resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of storage target.
	StorageTargetName pulumi.StringInput `pulumi:"storageTargetName"`
}

func (LookupStorageTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageTargetArgs)(nil)).Elem()
}

// A storage system being cached by a Cache.
type LookupStorageTargetResultOutput struct{ *pulumi.OutputState }

func (LookupStorageTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageTargetResult)(nil)).Elem()
}

func (o LookupStorageTargetResultOutput) ToLookupStorageTargetResultOutput() LookupStorageTargetResultOutput {
	return o
}

func (o LookupStorageTargetResultOutput) ToLookupStorageTargetResultOutputWithContext(ctx context.Context) LookupStorageTargetResultOutput {
	return o
}

// Properties when clfs target.
func (o LookupStorageTargetResultOutput) Clfs() ClfsTargetResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *ClfsTargetResponse { return v.Clfs }).(ClfsTargetResponsePtrOutput)
}

// Resource Id
func (o LookupStorageTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of cache namespace to target namespace associations.
func (o LookupStorageTargetResultOutput) Junctions() NamespaceJunctionResponseArrayOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) []NamespaceJunctionResponse { return v.Junctions }).(NamespaceJunctionResponseArrayOutput)
}

// A fully qualified URL.
func (o LookupStorageTargetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties when nfs3 target.
func (o LookupStorageTargetResultOutput) Nfs3() Nfs3TargetResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *Nfs3TargetResponse { return v.Nfs3 }).(Nfs3TargetResponsePtrOutput)
}

// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
func (o LookupStorageTargetResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Type for storage target.
func (o LookupStorageTargetResultOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *string { return v.TargetType }).(pulumi.StringPtrOutput)
}

// Type for the storage target; Microsoft.StorageCache/Cache/StorageTarget
func (o LookupStorageTargetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) string { return v.Type }).(pulumi.StringOutput)
}

// Properties when unknown target.
func (o LookupStorageTargetResultOutput) Unknown() UnknownTargetResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *UnknownTargetResponse { return v.Unknown }).(UnknownTargetResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageTargetResultOutput{})
}
