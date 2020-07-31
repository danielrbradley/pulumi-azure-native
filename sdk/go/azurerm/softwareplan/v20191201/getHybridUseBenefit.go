// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupHybridUseBenefit(ctx *pulumi.Context, args *LookupHybridUseBenefitArgs, opts ...pulumi.InvokeOption) (*LookupHybridUseBenefitResult, error) {
	var rv LookupHybridUseBenefitResult
	err := ctx.Invoke("azurerm:softwareplan/v20191201:getHybridUseBenefit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridUseBenefitArgs struct {
	// This is a unique identifier for a plan. Should be a guid.
	Name string `pulumi:"name"`
	// The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
	Scope string `pulumi:"scope"`
}

// Response on GET of a hybrid use benefit
type LookupHybridUseBenefitResult struct {
	// Indicates the revision of the hybrid use benefit
	Etag int `pulumi:"etag"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Property bag for a hybrid use benefit response
	Properties HybridUseBenefitPropertiesResponse `pulumi:"properties"`
	// Hybrid use benefit SKU
	Sku SkuResponse `pulumi:"sku"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}