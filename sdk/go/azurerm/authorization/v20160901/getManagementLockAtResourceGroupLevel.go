// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160901

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagementLockAtResourceGroupLevel(ctx *pulumi.Context, args *LookupManagementLockAtResourceGroupLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtResourceGroupLevelResult, error) {
	var rv LookupManagementLockAtResourceGroupLevelResult
	err := ctx.Invoke("azurerm:authorization/v20160901:getManagementLockAtResourceGroupLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtResourceGroupLevelArgs struct {
	// The name of the lock to get.
	Name string `pulumi:"name"`
	// The name of the locked resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The lock information.
type LookupManagementLockAtResourceGroupLevelResult struct {
	// The name of the lock.
	Name string `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponse `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type string `pulumi:"type"`
}