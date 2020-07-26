// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAppServicePlanVirtualNetworkConnectionRoute(ctx *pulumi.Context, args *LookupAppServicePlanVirtualNetworkConnectionRouteArgs, opts ...pulumi.InvokeOption) (*LookupAppServicePlanVirtualNetworkConnectionRouteResult, error) {
	var rv LookupAppServicePlanVirtualNetworkConnectionRouteResult
	err := ctx.Invoke("azurerm:web:getAppServicePlanVirtualNetworkConnectionRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServicePlanVirtualNetworkConnectionRouteArgs struct {
	// Name of the App Service plan.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Virtual Network route.
	RouteName string `pulumi:"routeName"`
	// Name of the Virtual Network.
	VnetName string `pulumi:"vnetName"`
}

type LookupAppServicePlanVirtualNetworkConnectionRouteResult struct {
}