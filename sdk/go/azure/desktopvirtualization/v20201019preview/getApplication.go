// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201019preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for Application properties.
//
// Deprecated: Version v20201019preview will be removed in the next major version of the provider. Upgrade to version v20210201preview or later.
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20201019preview:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	// The name of the application group
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	// The name of the application within the specified application group
	ApplicationName string `pulumi:"applicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Schema for Application properties.
type LookupApplicationResult struct {
	// Resource Type of Application.
	ApplicationType *string `pulumi:"applicationType"`
	// Command Line Arguments for Application.
	CommandLineArguments *string `pulumi:"commandLineArguments"`
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting string `pulumi:"commandLineSetting"`
	// Description of Application.
	Description *string `pulumi:"description"`
	// Specifies a path for the executable file for the application.
	FilePath *string `pulumi:"filePath"`
	// Friendly name of Application.
	FriendlyName *string `pulumi:"friendlyName"`
	// the icon a 64 bit string as a byte array.
	IconContent string `pulumi:"iconContent"`
	// Hash of the icon.
	IconHash string `pulumi:"iconHash"`
	// Index of the icon.
	IconIndex *int `pulumi:"iconIndex"`
	// Path to icon.
	IconPath *string `pulumi:"iconPath"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Specifies the package application Id for MSIX applications
	MsixPackageApplicationId *string `pulumi:"msixPackageApplicationId"`
	// Specifies the package family name for MSIX applications
	MsixPackageFamilyName *string `pulumi:"msixPackageFamilyName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal *bool `pulumi:"showInPortal"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	// The name of the application group
	ApplicationGroupName pulumi.StringInput `pulumi:"applicationGroupName"`
	// The name of the application within the specified application group
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

// Schema for Application properties.
type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

// Resource Type of Application.
func (o LookupApplicationResultOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ApplicationType }).(pulumi.StringPtrOutput)
}

// Command Line Arguments for Application.
func (o LookupApplicationResultOutput) CommandLineArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.CommandLineArguments }).(pulumi.StringPtrOutput)
}

// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
func (o LookupApplicationResultOutput) CommandLineSetting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.CommandLineSetting }).(pulumi.StringOutput)
}

// Description of Application.
func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies a path for the executable file for the application.
func (o LookupApplicationResultOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

// Friendly name of Application.
func (o LookupApplicationResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// the icon a 64 bit string as a byte array.
func (o LookupApplicationResultOutput) IconContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.IconContent }).(pulumi.StringOutput)
}

// Hash of the icon.
func (o LookupApplicationResultOutput) IconHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.IconHash }).(pulumi.StringOutput)
}

// Index of the icon.
func (o LookupApplicationResultOutput) IconIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *int { return v.IconIndex }).(pulumi.IntPtrOutput)
}

// Path to icon.
func (o LookupApplicationResultOutput) IconPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.IconPath }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the package application Id for MSIX applications
func (o LookupApplicationResultOutput) MsixPackageApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.MsixPackageApplicationId }).(pulumi.StringPtrOutput)
}

// Specifies the package family name for MSIX applications
func (o LookupApplicationResultOutput) MsixPackageFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.MsixPackageFamilyName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether to show the RemoteApp program in the RD Web Access server.
func (o LookupApplicationResultOutput) ShowInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *bool { return v.ShowInPortal }).(pulumi.BoolPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
