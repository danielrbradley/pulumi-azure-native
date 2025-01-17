// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the PUT rollout request body.
//
// Deprecated: Version v20180901preview will be removed in the next major version of the provider. Upgrade to version v20191101preview or later.
func LookupRollout(ctx *pulumi.Context, args *LookupRolloutArgs, opts ...pulumi.InvokeOption) (*LookupRolloutResult, error) {
	var rv LookupRolloutResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20180901preview:getRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRolloutArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Rollout retry attempt ordinal to get the result of. If not specified, result of the latest attempt will be returned.
	RetryAttempt *int `pulumi:"retryAttempt"`
	// The rollout name.
	RolloutName string `pulumi:"rolloutName"`
}

// Defines the rollout.
type LookupRolloutResult struct {
	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The version of the build being deployed.
	BuildVersion string `pulumi:"buildVersion"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Operational information of the rollout.
	OperationInfo RolloutOperationInfoResponse `pulumi:"operationInfo"`
	// The detailed information on the services being deployed.
	Services []ServiceResponse `pulumi:"services"`
	// The current status of the rollout.
	Status string `pulumi:"status"`
	// The list of step groups that define the orchestration.
	StepGroups []StepResponse `pulumi:"stepGroups"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyId string `pulumi:"targetServiceTopologyId"`
	// The cardinal count of total number of retries performed on the rollout at a given time.
	TotalRetryAttempts int `pulumi:"totalRetryAttempts"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupRolloutOutput(ctx *pulumi.Context, args LookupRolloutOutputArgs, opts ...pulumi.InvokeOption) LookupRolloutResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRolloutResult, error) {
			args := v.(LookupRolloutArgs)
			r, err := LookupRollout(ctx, &args, opts...)
			var s LookupRolloutResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRolloutResultOutput)
}

type LookupRolloutOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Rollout retry attempt ordinal to get the result of. If not specified, result of the latest attempt will be returned.
	RetryAttempt pulumi.IntPtrInput `pulumi:"retryAttempt"`
	// The rollout name.
	RolloutName pulumi.StringInput `pulumi:"rolloutName"`
}

func (LookupRolloutOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolloutArgs)(nil)).Elem()
}

// Defines the rollout.
type LookupRolloutResultOutput struct{ *pulumi.OutputState }

func (LookupRolloutResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolloutResult)(nil)).Elem()
}

func (o LookupRolloutResultOutput) ToLookupRolloutResultOutput() LookupRolloutResultOutput {
	return o
}

func (o LookupRolloutResultOutput) ToLookupRolloutResultOutputWithContext(ctx context.Context) LookupRolloutResultOutput {
	return o
}

// The reference to the artifact source resource Id where the payload is located.
func (o LookupRolloutResultOutput) ArtifactSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRolloutResult) *string { return v.ArtifactSourceId }).(pulumi.StringPtrOutput)
}

// The version of the build being deployed.
func (o LookupRolloutResultOutput) BuildVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.BuildVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRolloutResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupRolloutResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupRolloutResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupRolloutResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRolloutResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operational information of the rollout.
func (o LookupRolloutResultOutput) OperationInfo() RolloutOperationInfoResponseOutput {
	return o.ApplyT(func(v LookupRolloutResult) RolloutOperationInfoResponse { return v.OperationInfo }).(RolloutOperationInfoResponseOutput)
}

// The detailed information on the services being deployed.
func (o LookupRolloutResultOutput) Services() ServiceResponseArrayOutput {
	return o.ApplyT(func(v LookupRolloutResult) []ServiceResponse { return v.Services }).(ServiceResponseArrayOutput)
}

// The current status of the rollout.
func (o LookupRolloutResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Status }).(pulumi.StringOutput)
}

// The list of step groups that define the orchestration.
func (o LookupRolloutResultOutput) StepGroups() StepResponseArrayOutput {
	return o.ApplyT(func(v LookupRolloutResult) []StepResponse { return v.StepGroups }).(StepResponseArrayOutput)
}

// Resource tags.
func (o LookupRolloutResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRolloutResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
func (o LookupRolloutResultOutput) TargetServiceTopologyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.TargetServiceTopologyId }).(pulumi.StringOutput)
}

// The cardinal count of total number of retries performed on the rollout at a given time.
func (o LookupRolloutResultOutput) TotalRetryAttempts() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRolloutResult) int { return v.TotalRetryAttempts }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRolloutResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRolloutResultOutput{})
}
