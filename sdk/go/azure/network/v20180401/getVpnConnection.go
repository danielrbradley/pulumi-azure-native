// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VpnConnection Resource.
//
// Deprecated: Version v20180401 will be removed in the next major version of the provider. Upgrade to version v20180501 or later.
func LookupVpnConnection(ctx *pulumi.Context, args *LookupVpnConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpnConnectionResult, error) {
	var rv LookupVpnConnectionResult
	err := ctx.Invoke("azure-native:network/v20180401:getVpnConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnConnectionArgs struct {
	// The name of the vpn connection.
	ConnectionName string `pulumi:"connectionName"`
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// VpnConnection Resource.
type LookupVpnConnectionResult struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth int `pulumi:"connectionBandwidth"`
	// The connection status.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// Egress bytes transferred.
	EgressBytesTransferred float64 `pulumi:"egressBytesTransferred"`
	// EnableBgp flag
	EnableBgp *bool `pulumi:"enableBgp"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Ingress bytes transferred.
	IngressBytesTransferred float64 `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite *SubResourceResponse `pulumi:"remoteVpnSite"`
	// routing weight for vpn connection.
	RoutingWeight *int `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey *string `pulumi:"sharedKey"`
}

func LookupVpnConnectionOutput(ctx *pulumi.Context, args LookupVpnConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVpnConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnConnectionResult, error) {
			args := v.(LookupVpnConnectionArgs)
			r, err := LookupVpnConnection(ctx, &args, opts...)
			var s LookupVpnConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnConnectionResultOutput)
}

type LookupVpnConnectionOutputArgs struct {
	// The name of the vpn connection.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVpnConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionArgs)(nil)).Elem()
}

// VpnConnection Resource.
type LookupVpnConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVpnConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionResult)(nil)).Elem()
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutput() LookupVpnConnectionResultOutput {
	return o
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutputWithContext(ctx context.Context) LookupVpnConnectionResultOutput {
	return o
}

// Expected bandwidth in MBPS.
func (o LookupVpnConnectionResultOutput) ConnectionBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) int { return v.ConnectionBandwidth }).(pulumi.IntOutput)
}

// The connection status.
func (o LookupVpnConnectionResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// Egress bytes transferred.
func (o LookupVpnConnectionResultOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVpnConnectionResult) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

// EnableBgp flag
func (o LookupVpnConnectionResultOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o LookupVpnConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVpnConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Ingress bytes transferred.
func (o LookupVpnConnectionResultOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVpnConnectionResult) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

// The IPSec Policies to be considered by this connection.
func (o LookupVpnConnectionResultOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupVpnConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o LookupVpnConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Id of the connected vpn site.
func (o LookupVpnConnectionResultOutput) RemoteVpnSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *SubResourceResponse { return v.RemoteVpnSite }).(SubResourceResponsePtrOutput)
}

// routing weight for vpn connection.
func (o LookupVpnConnectionResultOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

// SharedKey for the vpn connection.
func (o LookupVpnConnectionResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnConnectionResultOutput{})
}
