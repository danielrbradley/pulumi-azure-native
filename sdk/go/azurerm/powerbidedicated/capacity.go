// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package powerbidedicated

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an instance of a Dedicated Capacity resource.
type Capacity struct {
	pulumi.CustomResourceState

	// Location of the PowerBI Dedicated resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the PowerBI Dedicated resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the provision operation request.
	Properties DedicatedCapacityPropertiesResponseOutput `pulumi:"properties"`
	// The SKU of the PowerBI Dedicated resource.
	Sku ResourceSkuResponseOutput `pulumi:"sku"`
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the PowerBI Dedicated resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCapacity registers a new resource with the given unique name, arguments, and options.
func NewCapacity(ctx *pulumi.Context,
	name string, args *CapacityArgs, opts ...pulumi.ResourceOption) (*Capacity, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &CapacityArgs{}
	}
	var resource Capacity
	err := ctx.RegisterResource("azurerm:powerbidedicated:Capacity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacity gets an existing Capacity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityState, opts ...pulumi.ResourceOption) (*Capacity, error) {
	var resource Capacity
	err := ctx.ReadResource("azurerm:powerbidedicated:Capacity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Capacity resources.
type capacityState struct {
	// Location of the PowerBI Dedicated resource.
	Location *string `pulumi:"location"`
	// The name of the PowerBI Dedicated resource.
	Name *string `pulumi:"name"`
	// Properties of the provision operation request.
	Properties *DedicatedCapacityPropertiesResponse `pulumi:"properties"`
	// The SKU of the PowerBI Dedicated resource.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
	// The type of the PowerBI Dedicated resource.
	Type *string `pulumi:"type"`
}

type CapacityState struct {
	// Location of the PowerBI Dedicated resource.
	Location pulumi.StringPtrInput
	// The name of the PowerBI Dedicated resource.
	Name pulumi.StringPtrInput
	// Properties of the provision operation request.
	Properties DedicatedCapacityPropertiesResponsePtrInput
	// The SKU of the PowerBI Dedicated resource.
	Sku ResourceSkuResponsePtrInput
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapInput
	// The type of the PowerBI Dedicated resource.
	Type pulumi.StringPtrInput
}

func (CapacityState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityState)(nil)).Elem()
}

type capacityArgs struct {
	// Location of the PowerBI Dedicated resource.
	Location string `pulumi:"location"`
	// The name of the Dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
	Name string `pulumi:"name"`
	// Properties of the provision operation request.
	Properties *DedicatedCapacityProperties `pulumi:"properties"`
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the PowerBI Dedicated resource.
	Sku ResourceSku `pulumi:"sku"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Capacity resource.
type CapacityArgs struct {
	// Location of the PowerBI Dedicated resource.
	Location pulumi.StringInput
	// The name of the Dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
	Name pulumi.StringInput
	// Properties of the provision operation request.
	Properties DedicatedCapacityPropertiesPtrInput
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput
	// The SKU of the PowerBI Dedicated resource.
	Sku ResourceSkuInput
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapInput
}

func (CapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityArgs)(nil)).Elem()
}