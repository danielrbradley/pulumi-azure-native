// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource information.
//
// Deprecated: Version v20151101 will be removed in the next major version of the provider. Upgrade to version v20190501 or later.
func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	var rv LookupResourceResult
	err := ctx.Invoke("azure-native:resources/v20151101:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceArgs struct {
	// Resource identity.
	ParentResourcePath string `pulumi:"parentResourcePath"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource identity.
	ResourceName string `pulumi:"resourceName"`
	// Resource identity.
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	// Resource identity.
	ResourceType string `pulumi:"resourceType"`
}

// Resource information.
type LookupResourceResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets or sets the plan of the resource.
	Plan *PlanResponse `pulumi:"plan"`
	// Gets or sets the resource properties.
	Properties interface{} `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupResourceOutput(ctx *pulumi.Context, args LookupResourceOutputArgs, opts ...pulumi.InvokeOption) LookupResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceResult, error) {
			args := v.(LookupResourceArgs)
			r, err := LookupResource(ctx, &args, opts...)
			var s LookupResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceResultOutput)
}

type LookupResourceOutputArgs struct {
	// Resource identity.
	ParentResourcePath pulumi.StringInput `pulumi:"parentResourcePath"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Resource identity.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// Resource identity.
	ResourceProviderNamespace pulumi.StringInput `pulumi:"resourceProviderNamespace"`
	// Resource identity.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceArgs)(nil)).Elem()
}

// Resource information.
type LookupResourceResultOutput struct{ *pulumi.OutputState }

func (LookupResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceResult)(nil)).Elem()
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutput() LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutputWithContext(ctx context.Context) LookupResourceResultOutput {
	return o
}

// Resource Id
func (o LookupResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupResourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the plan of the resource.
func (o LookupResourceResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

// Gets or sets the resource properties.
func (o LookupResourceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupResourceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags
func (o LookupResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceResultOutput{})
}
