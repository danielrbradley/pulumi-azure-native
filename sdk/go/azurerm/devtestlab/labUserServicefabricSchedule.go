// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A schedule.
type LabUserServicefabricSchedule struct {
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

// NewLabUserServicefabricSchedule registers a new resource with the given unique name, arguments, and options.
func NewLabUserServicefabricSchedule(ctx *pulumi.Context,
	name string, args *LabUserServicefabricScheduleArgs, opts ...pulumi.ResourceOption) (*LabUserServicefabricSchedule, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
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
	if args == nil || args.ServiceFabricName == nil {
		return nil, errors.New("missing required argument 'ServiceFabricName'")
	}
	if args == nil || args.UserName == nil {
		return nil, errors.New("missing required argument 'UserName'")
	}
	if args == nil {
		args = &LabUserServicefabricScheduleArgs{}
	}
	var resource LabUserServicefabricSchedule
	err := ctx.RegisterResource("azurerm:devtestlab:LabUserServicefabricSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLabUserServicefabricSchedule gets an existing LabUserServicefabricSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLabUserServicefabricSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabUserServicefabricScheduleState, opts ...pulumi.ResourceOption) (*LabUserServicefabricSchedule, error) {
	var resource LabUserServicefabricSchedule
	err := ctx.ReadResource("azurerm:devtestlab:LabUserServicefabricSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LabUserServicefabricSchedule resources.
type labUserServicefabricScheduleState struct {
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

type LabUserServicefabricScheduleState struct {
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

func (LabUserServicefabricScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*labUserServicefabricScheduleState)(nil)).Elem()
}

type labUserServicefabricScheduleArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the schedule.
	Name string `pulumi:"name"`
	// The properties of the resource.
	Properties ScheduleProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service fabric.
	ServiceFabricName string `pulumi:"serviceFabricName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the user profile.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a LabUserServicefabricSchedule resource.
type LabUserServicefabricScheduleArgs struct {
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the schedule.
	Name pulumi.StringInput
	// The properties of the resource.
	Properties SchedulePropertiesInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the service fabric.
	ServiceFabricName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The name of the user profile.
	UserName pulumi.StringInput
}

func (LabUserServicefabricScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labUserServicefabricScheduleArgs)(nil)).Elem()
}