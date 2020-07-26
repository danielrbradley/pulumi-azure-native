// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Agent Pool.
type ManagedClusterAgentPool struct {
	pulumi.CustomResourceState

	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of an agent pool.
	Properties ManagedClusterAgentPoolProfilePropertiesResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedClusterAgentPool registers a new resource with the given unique name, arguments, and options.
func NewManagedClusterAgentPool(ctx *pulumi.Context,
	name string, args *ManagedClusterAgentPoolArgs, opts ...pulumi.ResourceOption) (*ManagedClusterAgentPool, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &ManagedClusterAgentPoolArgs{}
	}
	var resource ManagedClusterAgentPool
	err := ctx.RegisterResource("azurerm:containerservice:ManagedClusterAgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedClusterAgentPool gets an existing ManagedClusterAgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedClusterAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterAgentPoolState, opts ...pulumi.ResourceOption) (*ManagedClusterAgentPool, error) {
	var resource ManagedClusterAgentPool
	err := ctx.ReadResource("azurerm:containerservice:ManagedClusterAgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedClusterAgentPool resources.
type managedClusterAgentPoolState struct {
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of an agent pool.
	Properties *ManagedClusterAgentPoolProfilePropertiesResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ManagedClusterAgentPoolState struct {
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of an agent pool.
	Properties ManagedClusterAgentPoolProfilePropertiesResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ManagedClusterAgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterAgentPoolState)(nil)).Elem()
}

type managedClusterAgentPoolArgs struct {
	// The name of the agent pool.
	Name string `pulumi:"name"`
	// Properties of an agent pool.
	Properties *ManagedClusterAgentPoolProfileProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ManagedClusterAgentPool resource.
type ManagedClusterAgentPoolArgs struct {
	// The name of the agent pool.
	Name pulumi.StringInput
	// Properties of an agent pool.
	Properties ManagedClusterAgentPoolProfilePropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
}

func (ManagedClusterAgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterAgentPoolArgs)(nil)).Elem()
}