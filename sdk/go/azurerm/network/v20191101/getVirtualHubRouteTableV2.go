// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVirtualHubRouteTableV2(ctx *pulumi.Context, args *LookupVirtualHubRouteTableV2Args, opts ...pulumi.InvokeOption) (*LookupVirtualHubRouteTableV2Result, error) {
	var rv LookupVirtualHubRouteTableV2Result
	err := ctx.Invoke("azurerm:network/v20191101:getVirtualHubRouteTableV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubRouteTableV2Args struct {
	// The name of the VirtualHubRouteTableV2.
	Name string `pulumi:"name"`
	// The resource group name of the VirtualHubRouteTableV2.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// VirtualHubRouteTableV2 Resource.
type LookupVirtualHubRouteTableV2Result struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the virtual hub route table v2.
	Properties VirtualHubRouteTableV2PropertiesResponse `pulumi:"properties"`
}