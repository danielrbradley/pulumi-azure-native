// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListAppServicePlanHybridConnectionNamespaceRelayKeys(ctx *pulumi.Context, args *ListAppServicePlanHybridConnectionNamespaceRelayKeysArgs, opts ...pulumi.InvokeOption) (*ListAppServicePlanHybridConnectionNamespaceRelayKeysResult, error) {
	var rv ListAppServicePlanHybridConnectionNamespaceRelayKeysResult
	err := ctx.Invoke("azurerm:web:listAppServicePlanHybridConnectionNamespaceRelayKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAppServicePlanHybridConnectionNamespaceRelayKeysArgs struct {
	// The name of the Service Bus relay.
	Name string `pulumi:"name"`
	// The name of the Service Bus namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Hybrid Connection key contract. This has the send key name and value for a Hybrid Connection.
type ListAppServicePlanHybridConnectionNamespaceRelayKeysResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// HybridConnectionKey resource specific properties
	Properties HybridConnectionKeyResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}