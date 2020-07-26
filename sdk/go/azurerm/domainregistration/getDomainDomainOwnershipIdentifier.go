// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package domainregistration

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDomainDomainOwnershipIdentifier(ctx *pulumi.Context, args *LookupDomainDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*LookupDomainDomainOwnershipIdentifierResult, error) {
	var rv LookupDomainDomainOwnershipIdentifierResult
	err := ctx.Invoke("azurerm:domainregistration:getDomainDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainDomainOwnershipIdentifierArgs struct {
	// Name of domain.
	DomainName string `pulumi:"domainName"`
	// Name of identifier.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Domain ownership Identifier.
type LookupDomainDomainOwnershipIdentifierResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// DomainOwnershipIdentifier resource specific properties
	Properties DomainOwnershipIdentifierResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}