// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServiceApiRelease(ctx *pulumi.Context, args *LookupServiceApiReleaseArgs, opts ...pulumi.InvokeOption) (*LookupServiceApiReleaseResult, error) {
	var rv LookupServiceApiReleaseResult
	err := ctx.Invoke("azurerm:apimanagement:getServiceApiRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceApiReleaseArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Release identifier within an API. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// ApiRelease details.
type LookupServiceApiReleaseResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// ApiRelease entity contract properties.
	Properties ApiReleaseContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}