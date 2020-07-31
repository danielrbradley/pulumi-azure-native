// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azurerm:notificationhubs/v20160301:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceArgs struct {
	// The namespace name.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a Namespace resource.
type LookupNamespaceResult struct {
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Properties of the Namespace.
	Properties NamespacePropertiesResponse `pulumi:"properties"`
	// The sku of the created namespace
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}