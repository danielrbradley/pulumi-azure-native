// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupNamespaceEventhubAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceEventhubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceEventhubAuthorizationRuleResult, error) {
	var rv LookupNamespaceEventhubAuthorizationRuleResult
	err := ctx.Invoke("azurerm:eventhub:getNamespaceEventhubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceEventhubAuthorizationRuleArgs struct {
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// The authorization rule name.
	Name string `pulumi:"name"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in a List or Get AuthorizationRule operation
type LookupNamespaceEventhubAuthorizationRuleResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// Properties supplied to create or update AuthorizationRule
	Properties AuthorizationRuleResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}