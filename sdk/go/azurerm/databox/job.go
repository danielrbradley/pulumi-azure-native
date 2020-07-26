// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databox

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Job Resource.
type Job struct {
	pulumi.CustomResourceState

	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a job.
	Properties JobPropertiesResponseOutput `pulumi:"properties"`
	// The sku type.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
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
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &JobArgs{}
	}
	var resource Job
	err := ctx.RegisterResource("azurerm:databox:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azurerm:databox:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location *string `pulumi:"location"`
	// Name of the object.
	Name *string `pulumi:"name"`
	// Properties of a job.
	Properties *JobPropertiesResponse `pulumi:"properties"`
	// The sku type.
	Sku *SkuResponse `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags map[string]string `pulumi:"tags"`
	// Type of the object.
	Type *string `pulumi:"type"`
}

type JobState struct {
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location pulumi.StringPtrInput
	// Name of the object.
	Name pulumi.StringPtrInput
	// Properties of a job.
	Properties JobPropertiesResponsePtrInput
	// The sku type.
	Sku SkuResponsePtrInput
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags pulumi.StringMapInput
	// Type of the object.
	Type pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location string `pulumi:"location"`
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	Name string `pulumi:"name"`
	// Properties of a job.
	Properties JobProperties `pulumi:"properties"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku type.
	Sku Sku `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location pulumi.StringInput
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	Name pulumi.StringInput
	// Properties of a job.
	Properties JobPropertiesInput
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput
	// The sku type.
	Sku SkuInput
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags pulumi.StringMapInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}