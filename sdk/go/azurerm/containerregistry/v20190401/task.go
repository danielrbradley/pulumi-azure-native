// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The task that has the ARM resource and task properties.
// The task will have all information to schedule a run against it.
type Task struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a task.
	Properties TaskPropertiesResponseOutput `pulumi:"properties"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTask registers a new resource with the given unique name, arguments, and options.
func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.RegistryName == nil {
		return nil, errors.New("missing required argument 'RegistryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &TaskArgs{}
	}
	var resource Task
	err := ctx.RegisterResource("azurerm:containerregistry/v20190401:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTask gets an existing Task resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("azurerm:containerregistry/v20190401:Task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Task resources.
type taskState struct {
	// Identity for the resource.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties of a task.
	Properties *TaskPropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type TaskState struct {
	// Identity for the resource.
	Identity IdentityPropertiesResponsePtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties of a task.
	Properties TaskPropertiesResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (TaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskState)(nil)).Elem()
}

type taskArgs struct {
	// Identity for the resource.
	Identity *IdentityProperties `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the container registry task.
	Name string `pulumi:"name"`
	// The properties of a task.
	Properties *TaskProperties `pulumi:"properties"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Task resource.
type TaskArgs struct {
	// Identity for the resource.
	Identity IdentityPropertiesPtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringInput
	// The name of the container registry task.
	Name pulumi.StringInput
	// The properties of a task.
	Properties TaskPropertiesPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
}

func (TaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskArgs)(nil)).Elem()
}