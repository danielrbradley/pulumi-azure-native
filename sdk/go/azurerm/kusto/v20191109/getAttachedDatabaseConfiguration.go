// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191109

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAttachedDatabaseConfiguration(ctx *pulumi.Context, args *LookupAttachedDatabaseConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAttachedDatabaseConfigurationResult, error) {
	var rv LookupAttachedDatabaseConfigurationResult
	err := ctx.Invoke("azurerm:kusto/v20191109:getAttachedDatabaseConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttachedDatabaseConfigurationArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the attached database configuration.
	Name string `pulumi:"name"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing an attached database configuration.
type LookupAttachedDatabaseConfigurationResult struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The properties of the attached database configuration.
	Properties AttachedDatabaseConfigurationPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}