// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupClusterDatabaseDataConnection(ctx *pulumi.Context, args *LookupClusterDatabaseDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupClusterDatabaseDataConnectionResult, error) {
	var rv LookupClusterDatabaseDataConnectionResult
	err := ctx.Invoke("azurerm:kusto:getClusterDatabaseDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterDatabaseDataConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the data connection.
	Name string `pulumi:"name"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing an data connection.
type LookupClusterDatabaseDataConnectionResult struct {
	// Kind of the endpoint for the data connection
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}