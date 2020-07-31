// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVpnGateway(ctx *pulumi.Context, args *LookupVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpnGatewayResult, error) {
	var rv LookupVpnGatewayResult
	err := ctx.Invoke("azurerm:network/v20180601:getVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnGatewayArgs struct {
	// The name of the gateway.
	Name string `pulumi:"name"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// VpnGateway Resource.
type LookupVpnGatewayResult struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Parameters for VpnGateway
	Properties VpnGatewayPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}