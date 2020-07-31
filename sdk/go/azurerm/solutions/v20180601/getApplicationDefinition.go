// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupApplicationDefinition(ctx *pulumi.Context, args *LookupApplicationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationDefinitionResult, error) {
	var rv LookupApplicationDefinitionResult
	err := ctx.Invoke("azurerm:solutions/v20180601:getApplicationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationDefinitionArgs struct {
	// The name of the managed application definition.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about managed application definition.
type LookupApplicationDefinitionResult struct {
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Resource location
	Location *string `pulumi:"location"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// Resource name
	Name string `pulumi:"name"`
	// The managed application definition properties.
	Properties ApplicationDefinitionPropertiesResponse `pulumi:"properties"`
	// The SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}