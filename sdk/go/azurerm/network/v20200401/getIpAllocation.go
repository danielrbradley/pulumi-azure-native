// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIpAllocation(ctx *pulumi.Context, args *LookupIpAllocationArgs, opts ...pulumi.InvokeOption) (*LookupIpAllocationResult, error) {
	var rv LookupIpAllocationResult
	err := ctx.Invoke("azurerm:network/v20200401:getIpAllocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpAllocationArgs struct {
	// The name of the IpAllocation.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// IpAllocation resource.
type LookupIpAllocationResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Properties of the IpAllocation.
	Properties IpAllocationPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}