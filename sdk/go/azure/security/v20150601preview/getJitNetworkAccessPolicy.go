// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: Version v20150601preview will be removed in the next major version of the provider. Upgrade to version v20170801preview or later.
func LookupJitNetworkAccessPolicy(ctx *pulumi.Context, args *LookupJitNetworkAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupJitNetworkAccessPolicyResult, error) {
	var rv LookupJitNetworkAccessPolicyResult
	err := ctx.Invoke("azure-native:security/v20150601preview:getJitNetworkAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName string `pulumi:"jitNetworkAccessPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJitNetworkAccessPolicyResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState string                            `pulumi:"provisioningState"`
	Requests          []JitNetworkAccessRequestResponse `pulumi:"requests"`
	// Resource type
	Type string `pulumi:"type"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines []JitNetworkAccessPolicyVirtualMachineResponse `pulumi:"virtualMachines"`
}

func LookupJitNetworkAccessPolicyOutput(ctx *pulumi.Context, args LookupJitNetworkAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupJitNetworkAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJitNetworkAccessPolicyResult, error) {
			args := v.(LookupJitNetworkAccessPolicyArgs)
			r, err := LookupJitNetworkAccessPolicy(ctx, &args, opts...)
			var s LookupJitNetworkAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJitNetworkAccessPolicyResultOutput)
}

type LookupJitNetworkAccessPolicyOutputArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation pulumi.StringInput `pulumi:"ascLocation"`
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName pulumi.StringInput `pulumi:"jitNetworkAccessPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJitNetworkAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJitNetworkAccessPolicyArgs)(nil)).Elem()
}

type LookupJitNetworkAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupJitNetworkAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJitNetworkAccessPolicyResult)(nil)).Elem()
}

func (o LookupJitNetworkAccessPolicyResultOutput) ToLookupJitNetworkAccessPolicyResultOutput() LookupJitNetworkAccessPolicyResultOutput {
	return o
}

func (o LookupJitNetworkAccessPolicyResultOutput) ToLookupJitNetworkAccessPolicyResultOutputWithContext(ctx context.Context) LookupJitNetworkAccessPolicyResultOutput {
	return o
}

// Resource Id
func (o LookupJitNetworkAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the resource
func (o LookupJitNetworkAccessPolicyResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o LookupJitNetworkAccessPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupJitNetworkAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the Just-in-Time policy.
func (o LookupJitNetworkAccessPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) Requests() JitNetworkAccessRequestResponseArrayOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) []JitNetworkAccessRequestResponse { return v.Requests }).(JitNetworkAccessRequestResponseArrayOutput)
}

// Resource type
func (o LookupJitNetworkAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Configurations for Microsoft.Compute/virtualMachines resource type.
func (o LookupJitNetworkAccessPolicyResultOutput) VirtualMachines() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) []JitNetworkAccessPolicyVirtualMachineResponse {
		return v.VirtualMachines
	}).(JitNetworkAccessPolicyVirtualMachineResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJitNetworkAccessPolicyResultOutput{})
}
