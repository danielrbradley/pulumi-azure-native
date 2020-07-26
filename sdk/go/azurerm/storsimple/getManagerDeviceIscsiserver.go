// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagerDeviceIscsiserver(ctx *pulumi.Context, args *LookupManagerDeviceIscsiserverArgs, opts ...pulumi.InvokeOption) (*LookupManagerDeviceIscsiserverResult, error) {
	var rv LookupManagerDeviceIscsiserverResult
	err := ctx.Invoke("azurerm:storsimple:getManagerDeviceIscsiserver", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerDeviceIscsiserverArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The iSCSI server name.
	Name string `pulumi:"name"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The iSCSI server.
type LookupManagerDeviceIscsiserverResult struct {
	// The name.
	Name string `pulumi:"name"`
	// The properties.
	Properties ISCSIServerPropertiesResponse `pulumi:"properties"`
	// The type.
	Type string `pulumi:"type"`
}