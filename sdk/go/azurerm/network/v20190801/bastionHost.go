// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Bastion Host resource.
type BastionHost struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Represents the bastion host resource.
	Properties BastionHostPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBastionHost registers a new resource with the given unique name, arguments, and options.
func NewBastionHost(ctx *pulumi.Context,
	name string, args *BastionHostArgs, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &BastionHostArgs{}
	}
	var resource BastionHost
	err := ctx.RegisterResource("azurerm:network/v20190801:BastionHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBastionHost gets an existing BastionHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBastionHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BastionHostState, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	var resource BastionHost
	err := ctx.ReadResource("azurerm:network/v20190801:BastionHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BastionHost resources.
type bastionHostState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Represents the bastion host resource.
	Properties *BastionHostPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type BastionHostState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Represents the bastion host resource.
	Properties BastionHostPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (BastionHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostState)(nil)).Elem()
}

type bastionHostArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the Bastion Host.
	Name string `pulumi:"name"`
	// Represents the bastion host resource.
	Properties *BastionHostPropertiesFormat `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BastionHost resource.
type BastionHostArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the Bastion Host.
	Name pulumi.StringInput
	// Represents the bastion host resource.
	Properties BastionHostPropertiesFormatPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (BastionHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostArgs)(nil)).Elem()
}