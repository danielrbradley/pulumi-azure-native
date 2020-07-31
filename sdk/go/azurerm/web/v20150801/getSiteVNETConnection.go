// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSiteVNETConnection(ctx *pulumi.Context, args *LookupSiteVNETConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSiteVNETConnectionResult, error) {
	var rv LookupSiteVNETConnectionResult
	err := ctx.Invoke("azurerm:web/v20150801:getSiteVNETConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteVNETConnectionArgs struct {
	// The name of the Virtual Network
	Name string `pulumi:"name"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// VNETInfo contract. This contract is public and is a stripped down version of VNETInfoInternal
type LookupSiteVNETConnectionResult struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name       *string                    `pulumi:"name"`
	Properties VnetInfoResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}