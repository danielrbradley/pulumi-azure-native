// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRemediationAtResourceGroup(ctx *pulumi.Context, args *LookupRemediationAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtResourceGroupResult, error) {
	var rv LookupRemediationAtResourceGroupResult
	err := ctx.Invoke("azurerm:policyinsights/v20190701:getRemediationAtResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtResourceGroupArgs struct {
	// The name of the remediation.
	Name string `pulumi:"name"`
	// Resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The remediation definition.
type LookupRemediationAtResourceGroupResult struct {
	// The name of the remediation.
	Name string `pulumi:"name"`
	// Properties for the remediation.
	Properties RemediationPropertiesResponse `pulumi:"properties"`
	// The type of the remediation.
	Type string `pulumi:"type"`
}