// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about appliance definition.
//
// Deprecated: Version v20160901preview will be removed in the next major version of the provider. Upgrade to version v20190701 or later.
func LookupApplianceDefinition(ctx *pulumi.Context, args *LookupApplianceDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplianceDefinitionResult, error) {
	var rv LookupApplianceDefinitionResult
	err := ctx.Invoke("azure-native:solutions/v20160901preview:getApplianceDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplianceDefinitionArgs struct {
	// The name of the appliance definition.
	ApplianceDefinitionName string `pulumi:"applianceDefinitionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about appliance definition.
type LookupApplianceDefinitionResult struct {
	// The collection of appliance artifacts. The portal will use the files specified as artifacts to construct the user experience of creating an appliance from an appliance definition.
	Artifacts []ApplianceArtifactResponse `pulumi:"artifacts"`
	// The appliance provider authorizations.
	Authorizations []ApplianceProviderAuthorizationResponse `pulumi:"authorizations"`
	// The appliance definition description.
	Description *string `pulumi:"description"`
	// The appliance definition display name.
	DisplayName *string `pulumi:"displayName"`
	// Resource ID
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Resource location
	Location *string `pulumi:"location"`
	// The appliance lock level.
	LockLevel string `pulumi:"lockLevel"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// Resource name
	Name string `pulumi:"name"`
	// The appliance definition package file Uri.
	PackageFileUri string `pulumi:"packageFileUri"`
	// The SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupApplianceDefinitionOutput(ctx *pulumi.Context, args LookupApplianceDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupApplianceDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplianceDefinitionResult, error) {
			args := v.(LookupApplianceDefinitionArgs)
			r, err := LookupApplianceDefinition(ctx, &args, opts...)
			var s LookupApplianceDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplianceDefinitionResultOutput)
}

type LookupApplianceDefinitionOutputArgs struct {
	// The name of the appliance definition.
	ApplianceDefinitionName pulumi.StringInput `pulumi:"applianceDefinitionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplianceDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceDefinitionArgs)(nil)).Elem()
}

// Information about appliance definition.
type LookupApplianceDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupApplianceDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceDefinitionResult)(nil)).Elem()
}

func (o LookupApplianceDefinitionResultOutput) ToLookupApplianceDefinitionResultOutput() LookupApplianceDefinitionResultOutput {
	return o
}

func (o LookupApplianceDefinitionResultOutput) ToLookupApplianceDefinitionResultOutputWithContext(ctx context.Context) LookupApplianceDefinitionResultOutput {
	return o
}

// The collection of appliance artifacts. The portal will use the files specified as artifacts to construct the user experience of creating an appliance from an appliance definition.
func (o LookupApplianceDefinitionResultOutput) Artifacts() ApplianceArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) []ApplianceArtifactResponse { return v.Artifacts }).(ApplianceArtifactResponseArrayOutput)
}

// The appliance provider authorizations.
func (o LookupApplianceDefinitionResultOutput) Authorizations() ApplianceProviderAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) []ApplianceProviderAuthorizationResponse {
		return v.Authorizations
	}).(ApplianceProviderAuthorizationResponseArrayOutput)
}

// The appliance definition description.
func (o LookupApplianceDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The appliance definition display name.
func (o LookupApplianceDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource ID
func (o LookupApplianceDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupApplianceDefinitionResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Resource location
func (o LookupApplianceDefinitionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The appliance lock level.
func (o LookupApplianceDefinitionResultOutput) LockLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.LockLevel }).(pulumi.StringOutput)
}

// ID of the resource that manages this resource.
func (o LookupApplianceDefinitionResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupApplianceDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The appliance definition package file Uri.
func (o LookupApplianceDefinitionResultOutput) PackageFileUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.PackageFileUri }).(pulumi.StringOutput)
}

// The SKU of the resource.
func (o LookupApplianceDefinitionResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o LookupApplianceDefinitionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupApplianceDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplianceDefinitionResultOutput{})
}
