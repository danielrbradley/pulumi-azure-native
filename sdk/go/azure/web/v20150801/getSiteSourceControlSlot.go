// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes the source control configuration for web app
//
// Deprecated: Version v20150801 will be removed in the next major version of the provider. Upgrade to version v20150801preview or later.
func LookupSiteSourceControlSlot(ctx *pulumi.Context, args *LookupSiteSourceControlSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteSourceControlSlotResult, error) {
	var rv LookupSiteSourceControlSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSourceControlSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSourceControlSlotArgs struct {
	// Name of web app
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
}

// Describes the source control configuration for web app
type LookupSiteSourceControlSlotResult struct {
	// Name of branch to use for deployment
	Branch *string `pulumi:"branch"`
	// Whether to manual or continuous integration
	DeploymentRollbackEnabled *bool `pulumi:"deploymentRollbackEnabled"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Whether to manual or continuous integration
	IsManualIntegration *bool `pulumi:"isManualIntegration"`
	// Mercurial or Git repository type
	IsMercurial *bool `pulumi:"isMercurial"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Repository or source control url
	RepoUrl *string `pulumi:"repoUrl"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

func LookupSiteSourceControlSlotOutput(ctx *pulumi.Context, args LookupSiteSourceControlSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteSourceControlSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteSourceControlSlotResult, error) {
			args := v.(LookupSiteSourceControlSlotArgs)
			r, err := LookupSiteSourceControlSlot(ctx, &args, opts...)
			var s LookupSiteSourceControlSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteSourceControlSlotResultOutput)
}

type LookupSiteSourceControlSlotOutputArgs struct {
	// Name of web app
	Name pulumi.StringInput `pulumi:"name"`
	// Name of resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupSiteSourceControlSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSourceControlSlotArgs)(nil)).Elem()
}

// Describes the source control configuration for web app
type LookupSiteSourceControlSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteSourceControlSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSourceControlSlotResult)(nil)).Elem()
}

func (o LookupSiteSourceControlSlotResultOutput) ToLookupSiteSourceControlSlotResultOutput() LookupSiteSourceControlSlotResultOutput {
	return o
}

func (o LookupSiteSourceControlSlotResultOutput) ToLookupSiteSourceControlSlotResultOutputWithContext(ctx context.Context) LookupSiteSourceControlSlotResultOutput {
	return o
}

// Name of branch to use for deployment
func (o LookupSiteSourceControlSlotResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

// Whether to manual or continuous integration
func (o LookupSiteSourceControlSlotResultOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *bool { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

// Resource Id
func (o LookupSiteSourceControlSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Whether to manual or continuous integration
func (o LookupSiteSourceControlSlotResultOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *bool { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

// Mercurial or Git repository type
func (o LookupSiteSourceControlSlotResultOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *bool { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

// Kind of resource
func (o LookupSiteSourceControlSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location
func (o LookupSiteSourceControlSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o LookupSiteSourceControlSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Repository or source control url
func (o LookupSiteSourceControlSlotResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o LookupSiteSourceControlSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupSiteSourceControlSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteSourceControlSlotResultOutput{})
}
