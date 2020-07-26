// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.
type FrontDoorRulesEngine struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Rules Engine Configuration.
	Properties RulesEnginePropertiesResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFrontDoorRulesEngine registers a new resource with the given unique name, arguments, and options.
func NewFrontDoorRulesEngine(ctx *pulumi.Context,
	name string, args *FrontDoorRulesEngineArgs, opts ...pulumi.ResourceOption) (*FrontDoorRulesEngine, error) {
	if args == nil || args.FrontDoorName == nil {
		return nil, errors.New("missing required argument 'FrontDoorName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FrontDoorRulesEngineArgs{}
	}
	var resource FrontDoorRulesEngine
	err := ctx.RegisterResource("azurerm:network:FrontDoorRulesEngine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFrontDoorRulesEngine gets an existing FrontDoorRulesEngine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFrontDoorRulesEngine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FrontDoorRulesEngineState, opts ...pulumi.ResourceOption) (*FrontDoorRulesEngine, error) {
	var resource FrontDoorRulesEngine
	err := ctx.ReadResource("azurerm:network:FrontDoorRulesEngine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FrontDoorRulesEngine resources.
type frontDoorRulesEngineState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the Rules Engine Configuration.
	Properties *RulesEnginePropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type FrontDoorRulesEngineState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the Rules Engine Configuration.
	Properties RulesEnginePropertiesResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (FrontDoorRulesEngineState) ElementType() reflect.Type {
	return reflect.TypeOf((*frontDoorRulesEngineState)(nil)).Elem()
}

type frontDoorRulesEngineArgs struct {
	// Name of the Front Door which is globally unique.
	FrontDoorName string `pulumi:"frontDoorName"`
	// Name of the Rules Engine which is unique within the Front Door.
	Name string `pulumi:"name"`
	// Properties of the Rules Engine Configuration.
	Properties *RulesEngineProperties `pulumi:"properties"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a FrontDoorRulesEngine resource.
type FrontDoorRulesEngineArgs struct {
	// Name of the Front Door which is globally unique.
	FrontDoorName pulumi.StringInput
	// Name of the Rules Engine which is unique within the Front Door.
	Name pulumi.StringInput
	// Properties of the Rules Engine Configuration.
	Properties RulesEnginePropertiesPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (FrontDoorRulesEngineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*frontDoorRulesEngineArgs)(nil)).Elem()
}