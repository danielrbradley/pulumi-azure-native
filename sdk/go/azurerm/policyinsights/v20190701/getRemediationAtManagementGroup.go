// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRemediationAtManagementGroup(ctx *pulumi.Context, args *LookupRemediationAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtManagementGroupResult, error) {
	var rv LookupRemediationAtManagementGroupResult
	err := ctx.Invoke("azurerm:policyinsights/v20190701:getRemediationAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtManagementGroupArgs struct {
	// Management group ID.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	// The name of the remediation.
	Name string `pulumi:"name"`
}

// The remediation definition.
type LookupRemediationAtManagementGroupResult struct {
	// The name of the remediation.
	Name string `pulumi:"name"`
	// Properties for the remediation.
	Properties RemediationPropertiesResponse `pulumi:"properties"`
	// The type of the remediation.
	Type string `pulumi:"type"`
}