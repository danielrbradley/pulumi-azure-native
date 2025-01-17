// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The remediation definition.
//
// Deprecated: Version v20180701preview will be removed in the next major version of the provider. Upgrade to version v20190701 or later.
func LookupRemediationAtSubscription(ctx *pulumi.Context, args *LookupRemediationAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtSubscriptionResult, error) {
	var rv LookupRemediationAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights/v20180701preview:getRemediationAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtSubscriptionArgs struct {
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
}

// The remediation definition.
type LookupRemediationAtSubscriptionResult struct {
	// The time at which the remediation was created.
	CreatedOn string `pulumi:"createdOn"`
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus *RemediationDeploymentSummaryResponse `pulumi:"deploymentStatus"`
	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFiltersResponse `pulumi:"filters"`
	// The ID of the remediation.
	Id string `pulumi:"id"`
	// The time at which the remediation was last updated.
	LastUpdatedOn string `pulumi:"lastUpdatedOn"`
	// The name of the remediation.
	Name string `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The status of the remediation.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the remediation.
	Type string `pulumi:"type"`
}

func LookupRemediationAtSubscriptionOutput(ctx *pulumi.Context, args LookupRemediationAtSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationAtSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationAtSubscriptionResult, error) {
			args := v.(LookupRemediationAtSubscriptionArgs)
			r, err := LookupRemediationAtSubscription(ctx, &args, opts...)
			var s LookupRemediationAtSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationAtSubscriptionResultOutput)
}

type LookupRemediationAtSubscriptionOutputArgs struct {
	// The name of the remediation.
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
}

func (LookupRemediationAtSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtSubscriptionArgs)(nil)).Elem()
}

// The remediation definition.
type LookupRemediationAtSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationAtSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtSubscriptionResult)(nil)).Elem()
}

func (o LookupRemediationAtSubscriptionResultOutput) ToLookupRemediationAtSubscriptionResultOutput() LookupRemediationAtSubscriptionResultOutput {
	return o
}

func (o LookupRemediationAtSubscriptionResultOutput) ToLookupRemediationAtSubscriptionResultOutputWithContext(ctx context.Context) LookupRemediationAtSubscriptionResultOutput {
	return o
}

// The time at which the remediation was created.
func (o LookupRemediationAtSubscriptionResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// The deployment status summary for all deployments created by the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) DeploymentStatus() RemediationDeploymentSummaryResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *RemediationDeploymentSummaryResponse {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponsePtrOutput)
}

// The filters that will be applied to determine which resources to remediate.
func (o LookupRemediationAtSubscriptionResultOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *RemediationFiltersResponse { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

// The ID of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The time at which the remediation was last updated.
func (o LookupRemediationAtSubscriptionResultOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

// The name of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the policy assignment that should be remediated.
func (o LookupRemediationAtSubscriptionResultOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
func (o LookupRemediationAtSubscriptionResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

// The status of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationAtSubscriptionResultOutput{})
}
