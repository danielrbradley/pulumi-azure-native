// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupExpressRouteCrossConnectionPeering(ctx *pulumi.Context, args *LookupExpressRouteCrossConnectionPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCrossConnectionPeeringResult, error) {
	var rv LookupExpressRouteCrossConnectionPeeringResult
	err := ctx.Invoke("azurerm:network/v20181201:getExpressRouteCrossConnectionPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCrossConnectionPeeringArgs struct {
	// The name of the ExpressRouteCrossConnection.
	CrossConnectionName string `pulumi:"crossConnectionName"`
	// The name of the peering.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Peering in an ExpressRoute Cross Connection resource.
type LookupExpressRouteCrossConnectionPeeringResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       *string                                              `pulumi:"name"`
	Properties ExpressRouteCrossConnectionPeeringPropertiesResponse `pulumi:"properties"`
}