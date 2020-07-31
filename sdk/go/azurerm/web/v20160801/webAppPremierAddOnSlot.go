// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Premier add-on.
type WebAppPremierAddOnSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// PremierAddOn resource specific properties
	Properties PremierAddOnResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppPremierAddOnSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppPremierAddOnSlot(ctx *pulumi.Context,
	name string, args *WebAppPremierAddOnSlotArgs, opts ...pulumi.ResourceOption) (*WebAppPremierAddOnSlot, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &WebAppPremierAddOnSlotArgs{}
	}
	var resource WebAppPremierAddOnSlot
	err := ctx.RegisterResource("azurerm:web/v20160801:WebAppPremierAddOnSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppPremierAddOnSlot gets an existing WebAppPremierAddOnSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppPremierAddOnSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPremierAddOnSlotState, opts ...pulumi.ResourceOption) (*WebAppPremierAddOnSlot, error) {
	var resource WebAppPremierAddOnSlot
	err := ctx.ReadResource("azurerm:web/v20160801:WebAppPremierAddOnSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppPremierAddOnSlot resources.
type webAppPremierAddOnSlotState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// PremierAddOn resource specific properties
	Properties *PremierAddOnResponseProperties `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppPremierAddOnSlotState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// PremierAddOn resource specific properties
	Properties PremierAddOnResponsePropertiesPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppPremierAddOnSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnSlotState)(nil)).Elem()
}

type webAppPremierAddOnSlotArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Add-on name.
	Name string `pulumi:"name"`
	// PremierAddOn resource specific properties
	Properties *PremierAddOnProperties `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the named add-on for the production slot.
	Slot string `pulumi:"slot"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebAppPremierAddOnSlot resource.
type WebAppPremierAddOnSlotArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringInput
	// Add-on name.
	Name pulumi.StringInput
	// PremierAddOn resource specific properties
	Properties PremierAddOnPropertiesPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will update the named add-on for the production slot.
	Slot pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (WebAppPremierAddOnSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnSlotArgs)(nil)).Elem()
}