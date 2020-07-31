// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupBackupScheduleGroup(ctx *pulumi.Context, args *LookupBackupScheduleGroupArgs, opts ...pulumi.InvokeOption) (*LookupBackupScheduleGroupResult, error) {
	var rv LookupBackupScheduleGroupResult
	err := ctx.Invoke("azurerm:storsimple/v20161001:getBackupScheduleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupScheduleGroupArgs struct {
	// The name of the device.
	DeviceName string `pulumi:"deviceName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The name of the schedule group.
	Name string `pulumi:"name"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Backup Schedule Group
type LookupBackupScheduleGroupResult struct {
	// The name.
	Name string `pulumi:"name"`
	// Properties of BackupScheduleGroup
	Properties BackupScheduleGroupPropertiesResponse `pulumi:"properties"`
	// The type.
	Type string `pulumi:"type"`
}