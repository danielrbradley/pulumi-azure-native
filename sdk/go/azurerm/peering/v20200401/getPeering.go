// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPeering(ctx *pulumi.Context, args *LookupPeeringArgs, opts ...pulumi.InvokeOption) (*LookupPeeringResult, error) {
	var rv LookupPeeringResult
	err := ctx.Invoke("azurerm:peering/v20200401:getPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringArgs struct {
	// The name of the peering.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
type LookupPeeringResult struct {
	// The kind of the peering.
	Kind string `pulumi:"kind"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties that define a peering.
	Properties PeeringPropertiesResponse `pulumi:"properties"`
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSkuResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}