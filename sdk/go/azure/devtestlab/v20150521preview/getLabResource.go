// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A lab.
//
// Deprecated: Version v20150521preview will be removed in the next major version of the provider. Upgrade to version v20180915 or later.
func LookupLabResource(ctx *pulumi.Context, args *LookupLabResourceArgs, opts ...pulumi.InvokeOption) (*LookupLabResourceResult, error) {
	var rv LookupLabResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getLabResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabResourceArgs struct {
	// The name of the lab.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A lab.
type LookupLabResourceResult struct {
	// The artifact storage account of the lab.
	ArtifactsStorageAccount *string `pulumi:"artifactsStorageAccount"`
	// The creation date of the lab.
	CreatedDate *string `pulumi:"createdDate"`
	// The lab's default storage account.
	DefaultStorageAccount *string `pulumi:"defaultStorageAccount"`
	// The default virtual network identifier of the lab.
	DefaultVirtualNetworkId *string `pulumi:"defaultVirtualNetworkId"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// The type of the lab storage.
	LabStorageType *string `pulumi:"labStorageType"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The storage accounts of the lab.
	StorageAccounts []string `pulumi:"storageAccounts"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The name of the key vault of the lab.
	VaultName *string `pulumi:"vaultName"`
}

func LookupLabResourceOutput(ctx *pulumi.Context, args LookupLabResourceOutputArgs, opts ...pulumi.InvokeOption) LookupLabResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResourceResult, error) {
			args := v.(LookupLabResourceArgs)
			r, err := LookupLabResource(ctx, &args, opts...)
			var s LookupLabResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResourceResultOutput)
}

type LookupLabResourceOutputArgs struct {
	// The name of the lab.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLabResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResourceArgs)(nil)).Elem()
}

// A lab.
type LookupLabResourceResultOutput struct{ *pulumi.OutputState }

func (LookupLabResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResourceResult)(nil)).Elem()
}

func (o LookupLabResourceResultOutput) ToLookupLabResourceResultOutput() LookupLabResourceResultOutput {
	return o
}

func (o LookupLabResourceResultOutput) ToLookupLabResourceResultOutputWithContext(ctx context.Context) LookupLabResourceResultOutput {
	return o
}

// The artifact storage account of the lab.
func (o LookupLabResourceResultOutput) ArtifactsStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.ArtifactsStorageAccount }).(pulumi.StringPtrOutput)
}

// The creation date of the lab.
func (o LookupLabResourceResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

// The lab's default storage account.
func (o LookupLabResourceResultOutput) DefaultStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.DefaultStorageAccount }).(pulumi.StringPtrOutput)
}

// The default virtual network identifier of the lab.
func (o LookupLabResourceResultOutput) DefaultVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.DefaultVirtualNetworkId }).(pulumi.StringPtrOutput)
}

// The identifier of the resource.
func (o LookupLabResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The type of the lab storage.
func (o LookupLabResourceResultOutput) LabStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.LabStorageType }).(pulumi.StringPtrOutput)
}

// The location of the resource.
func (o LookupLabResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupLabResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o LookupLabResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The storage accounts of the lab.
func (o LookupLabResourceResultOutput) StorageAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLabResourceResult) []string { return v.StorageAccounts }).(pulumi.StringArrayOutput)
}

// The tags of the resource.
func (o LookupLabResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupLabResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The name of the key vault of the lab.
func (o LookupLabResourceResultOutput) VaultName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.VaultName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResourceResultOutput{})
}
