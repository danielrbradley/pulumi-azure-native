// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190615

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azurerm:cdn/v20190615:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	Name string `pulumi:"name"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
type LookupEndpointResult struct {
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The JSON object that contains the properties required to create an endpoint.
	Properties EndpointPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}