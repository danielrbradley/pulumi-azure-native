// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Source control configuration for an app.
type AppServiceSlotSourcecontrol struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SiteSourceControl resource specific properties
	Properties SiteSourceControlResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServiceSlotSourcecontrol registers a new resource with the given unique name, arguments, and options.
func NewAppServiceSlotSourcecontrol(ctx *pulumi.Context,
	name string, args *AppServiceSlotSourcecontrolArgs, opts ...pulumi.ResourceOption) (*AppServiceSlotSourcecontrol, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AppServiceSlotSourcecontrolArgs{}
	}
	var resource AppServiceSlotSourcecontrol
	err := ctx.RegisterResource("azurerm:web:AppServiceSlotSourcecontrol", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServiceSlotSourcecontrol gets an existing AppServiceSlotSourcecontrol resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServiceSlotSourcecontrol(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceSlotSourcecontrolState, opts ...pulumi.ResourceOption) (*AppServiceSlotSourcecontrol, error) {
	var resource AppServiceSlotSourcecontrol
	err := ctx.ReadResource("azurerm:web:AppServiceSlotSourcecontrol", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServiceSlotSourcecontrol resources.
type appServiceSlotSourcecontrolState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// SiteSourceControl resource specific properties
	Properties *SiteSourceControlResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AppServiceSlotSourcecontrolState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// SiteSourceControl resource specific properties
	Properties SiteSourceControlResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AppServiceSlotSourcecontrolState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceSlotSourcecontrolState)(nil)).Elem()
}

type appServiceSlotSourcecontrolArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the deployment slot. If a slot is not specified, the API will update the source control configuration for the production slot.
	Name string `pulumi:"name"`
	// SiteSourceControl resource specific properties
	Properties *SiteSourceControlProperties `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AppServiceSlotSourcecontrol resource.
type AppServiceSlotSourcecontrolArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the deployment slot. If a slot is not specified, the API will update the source control configuration for the production slot.
	Name pulumi.StringInput
	// SiteSourceControl resource specific properties
	Properties SiteSourceControlPropertiesPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (AppServiceSlotSourcecontrolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceSlotSourcecontrolArgs)(nil)).Elem()
}