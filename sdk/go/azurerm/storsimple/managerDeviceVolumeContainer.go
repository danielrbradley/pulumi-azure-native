// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The volume container.
type ManagerDeviceVolumeContainer struct {
	pulumi.CustomResourceState

	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The volume container properties.
	Properties VolumeContainerPropertiesResponseOutput `pulumi:"properties"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagerDeviceVolumeContainer registers a new resource with the given unique name, arguments, and options.
func NewManagerDeviceVolumeContainer(ctx *pulumi.Context,
	name string, args *ManagerDeviceVolumeContainerArgs, opts ...pulumi.ResourceOption) (*ManagerDeviceVolumeContainer, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
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
		args = &ManagerDeviceVolumeContainerArgs{}
	}
	var resource ManagerDeviceVolumeContainer
	err := ctx.RegisterResource("azurerm:storsimple:ManagerDeviceVolumeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagerDeviceVolumeContainer gets an existing ManagerDeviceVolumeContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagerDeviceVolumeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerDeviceVolumeContainerState, opts ...pulumi.ResourceOption) (*ManagerDeviceVolumeContainer, error) {
	var resource ManagerDeviceVolumeContainer
	err := ctx.ReadResource("azurerm:storsimple:ManagerDeviceVolumeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagerDeviceVolumeContainer resources.
type managerDeviceVolumeContainerState struct {
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name *string `pulumi:"name"`
	// The volume container properties.
	Properties *VolumeContainerPropertiesResponse `pulumi:"properties"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type ManagerDeviceVolumeContainerState struct {
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The name of the object.
	Name pulumi.StringPtrInput
	// The volume container properties.
	Properties VolumeContainerPropertiesResponsePtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (ManagerDeviceVolumeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceVolumeContainerState)(nil)).Elem()
}

type managerDeviceVolumeContainerArgs struct {
	// The device name
	DeviceName string `pulumi:"deviceName"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The name of the volume container.
	Name string `pulumi:"name"`
	// The volume container properties.
	Properties VolumeContainerProperties `pulumi:"properties"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagerDeviceVolumeContainer resource.
type ManagerDeviceVolumeContainerArgs struct {
	// The device name
	DeviceName pulumi.StringInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The name of the volume container.
	Name pulumi.StringInput
	// The volume container properties.
	Properties VolumeContainerPropertiesInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
}

func (ManagerDeviceVolumeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceVolumeContainerArgs)(nil)).Elem()
}