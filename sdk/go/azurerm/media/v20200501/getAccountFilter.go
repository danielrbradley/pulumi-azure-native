// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAccountFilter(ctx *pulumi.Context, args *LookupAccountFilterArgs, opts ...pulumi.InvokeOption) (*LookupAccountFilterResult, error) {
	var rv LookupAccountFilterResult
	err := ctx.Invoke("azurerm:media/v20200501:getAccountFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Account Filter name
	Name string `pulumi:"name"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Account Filter.
type LookupAccountFilterResult struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// The Media Filter properties.
	Properties MediaFilterPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}