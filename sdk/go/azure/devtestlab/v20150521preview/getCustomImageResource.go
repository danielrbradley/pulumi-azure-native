// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A custom image.
//
// Deprecated: Version v20150521preview will be removed in the next major version of the provider. Upgrade to version v20180915 or later.
func LookupCustomImageResource(ctx *pulumi.Context, args *LookupCustomImageResourceArgs, opts ...pulumi.InvokeOption) (*LookupCustomImageResourceResult, error) {
	var rv LookupCustomImageResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getCustomImageResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomImageResourceArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the custom image.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A custom image.
type LookupCustomImageResourceResult struct {
	// The author of the custom image.
	Author *string `pulumi:"author"`
	// The creation date of the custom image.
	CreationDate *string `pulumi:"creationDate"`
	// The description of the custom image.
	Description *string `pulumi:"description"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The OS type of the custom image.
	OsType *string `pulumi:"osType"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The VHD from which the image is to be created.
	Vhd *CustomImagePropertiesCustomResponse `pulumi:"vhd"`
	// Properties for creating a custom image from a virtual machine.
	Vm *CustomImagePropertiesFromVmResponse `pulumi:"vm"`
}

func LookupCustomImageResourceOutput(ctx *pulumi.Context, args LookupCustomImageResourceOutputArgs, opts ...pulumi.InvokeOption) LookupCustomImageResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomImageResourceResult, error) {
			args := v.(LookupCustomImageResourceArgs)
			r, err := LookupCustomImageResource(ctx, &args, opts...)
			var s LookupCustomImageResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomImageResourceResultOutput)
}

type LookupCustomImageResourceOutputArgs struct {
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the custom image.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomImageResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomImageResourceArgs)(nil)).Elem()
}

// A custom image.
type LookupCustomImageResourceResultOutput struct{ *pulumi.OutputState }

func (LookupCustomImageResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomImageResourceResult)(nil)).Elem()
}

func (o LookupCustomImageResourceResultOutput) ToLookupCustomImageResourceResultOutput() LookupCustomImageResourceResultOutput {
	return o
}

func (o LookupCustomImageResourceResultOutput) ToLookupCustomImageResourceResultOutputWithContext(ctx context.Context) LookupCustomImageResourceResultOutput {
	return o
}

// The author of the custom image.
func (o LookupCustomImageResourceResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

// The creation date of the custom image.
func (o LookupCustomImageResourceResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

// The description of the custom image.
func (o LookupCustomImageResourceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the resource.
func (o LookupCustomImageResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The location of the resource.
func (o LookupCustomImageResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupCustomImageResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The OS type of the custom image.
func (o LookupCustomImageResourceResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o LookupCustomImageResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o LookupCustomImageResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupCustomImageResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The VHD from which the image is to be created.
func (o LookupCustomImageResourceResultOutput) Vhd() CustomImagePropertiesCustomResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *CustomImagePropertiesCustomResponse { return v.Vhd }).(CustomImagePropertiesCustomResponsePtrOutput)
}

// Properties for creating a custom image from a virtual machine.
func (o LookupCustomImageResourceResultOutput) Vm() CustomImagePropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *CustomImagePropertiesFromVmResponse { return v.Vm }).(CustomImagePropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomImageResourceResultOutput{})
}
