// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
//
// Deprecated: Version v20210111 will be removed in the next major version of the provider. Upgrade to version v20210308 or later.
func LookupPrivateEndpointConnectionsAdtAPI(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsAdtAPIArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsAdtAPIResult, error) {
	var rv LookupPrivateEndpointConnectionsAdtAPIResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210111:getPrivateEndpointConnectionsAdtAPI", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsAdtAPIArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionsAdtAPIResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Required property for system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsAdtAPIOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsAdtAPIOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsAdtAPIResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsAdtAPIResult, error) {
			args := v.(LookupPrivateEndpointConnectionsAdtAPIArgs)
			r, err := LookupPrivateEndpointConnectionsAdtAPI(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsAdtAPIResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsAdtAPIResultOutput)
}

type LookupPrivateEndpointConnectionsAdtAPIOutputArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsAdtAPIOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsAdtAPIArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionsAdtAPIResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsAdtAPIResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsAdtAPIResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) ToLookupPrivateEndpointConnectionsAdtAPIResultOutput() LookupPrivateEndpointConnectionsAdtAPIResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) ToLookupPrivateEndpointConnectionsAdtAPIResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsAdtAPIResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Required property for system data
func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsAdtAPIResultOutput{})
}
