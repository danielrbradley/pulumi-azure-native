// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// NetApp account resource
type NetAppAccount struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// NetApp Account properties
	Properties AccountPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags ResourceTagsResponsePtrOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetAppAccount registers a new resource with the given unique name, arguments, and options.
func NewNetAppAccount(ctx *pulumi.Context,
	name string, args *NetAppAccountArgs, opts ...pulumi.ResourceOption) (*NetAppAccount, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetAppAccountArgs{}
	}
	var resource NetAppAccount
	err := ctx.RegisterResource("azurerm:netapp:NetAppAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetAppAccount gets an existing NetAppAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetAppAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetAppAccountState, opts ...pulumi.ResourceOption) (*NetAppAccount, error) {
	var resource NetAppAccount
	err := ctx.ReadResource("azurerm:netapp:NetAppAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetAppAccount resources.
type netAppAccountState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// NetApp Account properties
	Properties *AccountPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags *ResourceTagsResponse `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type NetAppAccountState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// NetApp Account properties
	Properties AccountPropertiesResponsePtrInput
	// Resource tags
	Tags ResourceTagsResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (NetAppAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*netAppAccountState)(nil)).Elem()
}

type netAppAccountArgs struct {
	// Resource location
	Location string `pulumi:"location"`
	// The name of the NetApp account
	Name string `pulumi:"name"`
	// NetApp Account properties
	Properties *AccountProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags *ResourceTags `pulumi:"tags"`
}

// The set of arguments for constructing a NetAppAccount resource.
type NetAppAccountArgs struct {
	// Resource location
	Location pulumi.StringInput
	// The name of the NetApp account
	Name pulumi.StringInput
	// NetApp Account properties
	Properties AccountPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags ResourceTagsPtrInput
}

func (NetAppAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*netAppAccountArgs)(nil)).Elem()
}