// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupLabVirtualnetwork(ctx *pulumi.Context, args *LookupLabVirtualnetworkArgs, opts ...pulumi.InvokeOption) (*LookupLabVirtualnetworkResult, error) {
	var rv LookupLabVirtualnetworkResult
	err := ctx.Invoke("azurerm:devtestlab:getLabVirtualnetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabVirtualnetworkArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the virtual network.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A virtual network.
type LookupLabVirtualnetworkResult struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties of the resource.
	Properties VirtualNetworkPropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}