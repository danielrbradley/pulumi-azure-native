// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Data Box Edge/Gateway device.
type DataBoxEdgeDevice struct {
	pulumi.CustomResourceState

	// The etag for the devices.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringOutput `pulumi:"location"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the Data Box Edge/Gateway device.
	Properties DataBoxEdgeDevicePropertiesResponseOutput `pulumi:"properties"`
	// The SKU type.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataBoxEdgeDevice registers a new resource with the given unique name, arguments, and options.
func NewDataBoxEdgeDevice(ctx *pulumi.Context,
	name string, args *DataBoxEdgeDeviceArgs, opts ...pulumi.ResourceOption) (*DataBoxEdgeDevice, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DataBoxEdgeDeviceArgs{}
	}
	var resource DataBoxEdgeDevice
	err := ctx.RegisterResource("azurerm:databoxedge:DataBoxEdgeDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataBoxEdgeDevice gets an existing DataBoxEdgeDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataBoxEdgeDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataBoxEdgeDeviceState, opts ...pulumi.ResourceOption) (*DataBoxEdgeDevice, error) {
	var resource DataBoxEdgeDevice
	err := ctx.ReadResource("azurerm:databoxedge:DataBoxEdgeDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataBoxEdgeDevice resources.
type dataBoxEdgeDeviceState struct {
	// The etag for the devices.
	Etag *string `pulumi:"etag"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location *string `pulumi:"location"`
	// The object name.
	Name *string `pulumi:"name"`
	// The properties of the Data Box Edge/Gateway device.
	Properties *DataBoxEdgeDevicePropertiesResponse `pulumi:"properties"`
	// The SKU type.
	Sku *SkuResponse `pulumi:"sku"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags map[string]string `pulumi:"tags"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type DataBoxEdgeDeviceState struct {
	// The etag for the devices.
	Etag pulumi.StringPtrInput
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringPtrInput
	// The object name.
	Name pulumi.StringPtrInput
	// The properties of the Data Box Edge/Gateway device.
	Properties DataBoxEdgeDevicePropertiesResponsePtrInput
	// The SKU type.
	Sku SkuResponsePtrInput
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (DataBoxEdgeDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataBoxEdgeDeviceState)(nil)).Elem()
}

type dataBoxEdgeDeviceArgs struct {
	// The etag for the devices.
	Etag *string `pulumi:"etag"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location string `pulumi:"location"`
	// The device name.
	Name string `pulumi:"name"`
	// The properties of the Data Box Edge/Gateway device.
	Properties *DataBoxEdgeDeviceProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU type.
	Sku *Sku `pulumi:"sku"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataBoxEdgeDevice resource.
type DataBoxEdgeDeviceArgs struct {
	// The etag for the devices.
	Etag pulumi.StringPtrInput
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringInput
	// The device name.
	Name pulumi.StringInput
	// The properties of the Data Box Edge/Gateway device.
	Properties DataBoxEdgeDevicePropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The SKU type.
	Sku SkuPtrInput
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapInput
}

func (DataBoxEdgeDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataBoxEdgeDeviceArgs)(nil)).Elem()
}