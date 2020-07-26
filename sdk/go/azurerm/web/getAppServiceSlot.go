// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAppServiceSlot(ctx *pulumi.Context, args *LookupAppServiceSlotArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceSlotResult, error) {
	var rv LookupAppServiceSlotResult
	err := ctx.Invoke("azurerm:web:getAppServiceSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceSlotArgs struct {
	// Name of the deployment slot. By default, this API returns the production slot.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A web app, a mobile app backend, or an API app.
type LookupAppServiceSlotResult struct {
	// Managed service identity.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Site resource specific properties
	Properties SiteResponseProperties `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}