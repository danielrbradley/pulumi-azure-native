// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDedicatedHost(ctx *pulumi.Context, args *LookupDedicatedHostArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostResult, error) {
	var rv LookupDedicatedHostResult
	err := ctx.Invoke("azurerm:compute/v20191201:getDedicatedHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostArgs struct {
	// The name of the dedicated host group.
	HostGroupName string `pulumi:"hostGroupName"`
	// The name of the dedicated host.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the Dedicated host.
type LookupDedicatedHostResult struct {
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Properties of the dedicated host.
	Properties DedicatedHostPropertiesResponse `pulumi:"properties"`
	// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
	Sku SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}