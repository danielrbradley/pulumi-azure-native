// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServiceSubscription(ctx *pulumi.Context, args *LookupServiceSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupServiceSubscriptionResult, error) {
	var rv LookupServiceSubscriptionResult
	err := ctx.Invoke("azurerm:apimanagement:getServiceSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceSubscriptionArgs struct {
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Subscription details.
type LookupServiceSubscriptionResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// Subscription contract properties.
	Properties SubscriptionContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}