// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The relationship resource format.
type HubRelationship struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The definition of Relationship.
	Properties RelationshipDefinitionResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHubRelationship registers a new resource with the given unique name, arguments, and options.
func NewHubRelationship(ctx *pulumi.Context,
	name string, args *HubRelationshipArgs, opts ...pulumi.ResourceOption) (*HubRelationship, error) {
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &HubRelationshipArgs{}
	}
	var resource HubRelationship
	err := ctx.RegisterResource("azurerm:customerinsights:HubRelationship", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHubRelationship gets an existing HubRelationship resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHubRelationship(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubRelationshipState, opts ...pulumi.ResourceOption) (*HubRelationship, error) {
	var resource HubRelationship
	err := ctx.ReadResource("azurerm:customerinsights:HubRelationship", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HubRelationship resources.
type hubRelationshipState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The definition of Relationship.
	Properties *RelationshipDefinitionResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type HubRelationshipState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The definition of Relationship.
	Properties RelationshipDefinitionResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (HubRelationshipState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubRelationshipState)(nil)).Elem()
}

type hubRelationshipArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the Relationship.
	Name string `pulumi:"name"`
	// The definition of Relationship.
	Properties *RelationshipDefinition `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a HubRelationship resource.
type HubRelationshipArgs struct {
	// The name of the hub.
	HubName pulumi.StringInput
	// The name of the Relationship.
	Name pulumi.StringInput
	// The definition of Relationship.
	Properties RelationshipDefinitionPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (HubRelationshipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubRelationshipArgs)(nil)).Elem()
}