// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupFactory(ctx *pulumi.Context, args *LookupFactoryArgs, opts ...pulumi.InvokeOption) (*LookupFactoryResult, error) {
	var rv LookupFactoryResult
	err := ctx.Invoke("azurerm:datafactory/v20180601:getFactory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFactoryArgs struct {
	// The factory name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Factory resource type.
type LookupFactoryResult struct {
	// Etag identifies change in the resource.
	ETag string `pulumi:"eTag"`
	// Managed service identity of the factory.
	Identity *FactoryIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// Properties of the factory.
	Properties FactoryPropertiesResponse `pulumi:"properties"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}