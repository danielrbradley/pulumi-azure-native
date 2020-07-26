// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The iSCSI disk.
type ManagerDeviceIscsiserverDisk struct {
	pulumi.CustomResourceState

	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties.
	Properties ISCSIDiskPropertiesResponseOutput `pulumi:"properties"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagerDeviceIscsiserverDisk registers a new resource with the given unique name, arguments, and options.
func NewManagerDeviceIscsiserverDisk(ctx *pulumi.Context,
	name string, args *ManagerDeviceIscsiserverDiskArgs, opts ...pulumi.ResourceOption) (*ManagerDeviceIscsiserverDisk, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.IscsiServerName == nil {
		return nil, errors.New("missing required argument 'IscsiServerName'")
	}
	if args == nil || args.ManagerName == nil {
		return nil, errors.New("missing required argument 'ManagerName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagerDeviceIscsiserverDiskArgs{}
	}
	var resource ManagerDeviceIscsiserverDisk
	err := ctx.RegisterResource("azurerm:storsimple:ManagerDeviceIscsiserverDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagerDeviceIscsiserverDisk gets an existing ManagerDeviceIscsiserverDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagerDeviceIscsiserverDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerDeviceIscsiserverDiskState, opts ...pulumi.ResourceOption) (*ManagerDeviceIscsiserverDisk, error) {
	var resource ManagerDeviceIscsiserverDisk
	err := ctx.ReadResource("azurerm:storsimple:ManagerDeviceIscsiserverDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagerDeviceIscsiserverDisk resources.
type managerDeviceIscsiserverDiskState struct {
	// The name.
	Name *string `pulumi:"name"`
	// The properties.
	Properties *ISCSIDiskPropertiesResponse `pulumi:"properties"`
	// The type.
	Type *string `pulumi:"type"`
}

type ManagerDeviceIscsiserverDiskState struct {
	// The name.
	Name pulumi.StringPtrInput
	// The properties.
	Properties ISCSIDiskPropertiesResponsePtrInput
	// The type.
	Type pulumi.StringPtrInput
}

func (ManagerDeviceIscsiserverDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceIscsiserverDiskState)(nil)).Elem()
}

type managerDeviceIscsiserverDiskArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The iSCSI server name.
	IscsiServerName string `pulumi:"iscsiServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The disk name.
	Name string `pulumi:"name"`
	// The properties.
	Properties ISCSIDiskProperties `pulumi:"properties"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagerDeviceIscsiserverDisk resource.
type ManagerDeviceIscsiserverDiskArgs struct {
	// The device name.
	DeviceName pulumi.StringInput
	// The iSCSI server name.
	IscsiServerName pulumi.StringInput
	// The manager name
	ManagerName pulumi.StringInput
	// The disk name.
	Name pulumi.StringInput
	// The properties.
	Properties ISCSIDiskPropertiesInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
}

func (ManagerDeviceIscsiserverDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceIscsiserverDiskArgs)(nil)).Elem()
}