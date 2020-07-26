// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package search

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSearchServiceSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupSearchServiceSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupSearchServiceSharedPrivateLinkResourceResult, error) {
	var rv LookupSearchServiceSharedPrivateLinkResourceResult
	err := ctx.Invoke("azurerm:search:getSearchServiceSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSearchServiceSharedPrivateLinkResourceArgs struct {
	// The name of the shared private link resource managed by the Azure Cognitive Search service within the specified resource group.
	Name string `pulumi:"name"`
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	SearchServiceName string `pulumi:"searchServiceName"`
}

// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
type LookupSearchServiceSharedPrivateLinkResourceResult struct {
	// The name of the shared private link resource.
	Name string `pulumi:"name"`
	// Describes the properties of a Shared Private Link Resource managed by the Azure Cognitive Search service.
	Properties SharedPrivateLinkResourcePropertiesResponse `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}