// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSiteDeploymentSlot(ctx *pulumi.Context, args *LookupSiteDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteDeploymentSlotResult, error) {
	var rv LookupSiteDeploymentSlotResult
	err := ctx.Invoke("azurerm:web/v20150801:getSiteDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteDeploymentSlotArgs struct {
	// Id of the deployment
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
}

// Represents user credentials used for publishing activity
type LookupSiteDeploymentSlotResult struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name       *string                      `pulumi:"name"`
	Properties DeploymentResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}