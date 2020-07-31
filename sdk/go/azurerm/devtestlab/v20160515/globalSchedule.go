// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160515

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A schedule.
type GlobalSchedule struct {
	pulumi.CustomResourceState

	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the resource.
	Properties SchedulePropertiesResponseOutput `pulumi:"properties"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGlobalSchedule registers a new resource with the given unique name, arguments, and options.
func NewGlobalSchedule(ctx *pulumi.Context,
	name string, args *GlobalScheduleArgs, opts ...pulumi.ResourceOption) (*GlobalSchedule, error) {
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
		args = &GlobalScheduleArgs{}
	}
	var resource GlobalSchedule
	err := ctx.RegisterResource("azurerm:devtestlab/v20160515:GlobalSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalSchedule gets an existing GlobalSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalScheduleState, opts ...pulumi.ResourceOption) (*GlobalSchedule, error) {
	var resource GlobalSchedule
	err := ctx.ReadResource("azurerm:devtestlab/v20160515:GlobalSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalSchedule resources.
type globalScheduleState struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties of the resource.
	Properties *SchedulePropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type GlobalScheduleState struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties of the resource.
	Properties SchedulePropertiesResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (GlobalScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalScheduleState)(nil)).Elem()
}

type globalScheduleArgs struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the schedule.
	Name string `pulumi:"name"`
	// The properties of the resource.
	Properties ScheduleProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GlobalSchedule resource.
type GlobalScheduleArgs struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the schedule.
	Name pulumi.StringInput
	// The properties of the resource.
	Properties SchedulePropertiesInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
}

func (GlobalScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalScheduleArgs)(nil)).Elem()
}