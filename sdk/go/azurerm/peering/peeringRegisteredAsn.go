// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package peering

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The customer's ASN that is registered by the peering service provider.
type PeeringRegisteredAsn struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties that define a registered ASN.
	Properties PeeringRegisteredAsnPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPeeringRegisteredAsn registers a new resource with the given unique name, arguments, and options.
func NewPeeringRegisteredAsn(ctx *pulumi.Context,
	name string, args *PeeringRegisteredAsnArgs, opts ...pulumi.ResourceOption) (*PeeringRegisteredAsn, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PeeringName == nil {
		return nil, errors.New("missing required argument 'PeeringName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PeeringRegisteredAsnArgs{}
	}
	var resource PeeringRegisteredAsn
	err := ctx.RegisterResource("azurerm:peering:PeeringRegisteredAsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeeringRegisteredAsn gets an existing PeeringRegisteredAsn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeeringRegisteredAsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringRegisteredAsnState, opts ...pulumi.ResourceOption) (*PeeringRegisteredAsn, error) {
	var resource PeeringRegisteredAsn
	err := ctx.ReadResource("azurerm:peering:PeeringRegisteredAsn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeeringRegisteredAsn resources.
type peeringRegisteredAsnState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties that define a registered ASN.
	Properties *PeeringRegisteredAsnPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type PeeringRegisteredAsnState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties that define a registered ASN.
	Properties PeeringRegisteredAsnPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (PeeringRegisteredAsnState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringRegisteredAsnState)(nil)).Elem()
}

type peeringRegisteredAsnArgs struct {
	// The name of the ASN.
	Name string `pulumi:"name"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The properties that define a registered ASN.
	Properties *PeeringRegisteredAsnProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PeeringRegisteredAsn resource.
type PeeringRegisteredAsnArgs struct {
	// The name of the ASN.
	Name pulumi.StringInput
	// The name of the peering.
	PeeringName pulumi.StringInput
	// The properties that define a registered ASN.
	Properties PeeringRegisteredAsnPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (PeeringRegisteredAsnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringRegisteredAsnArgs)(nil)).Elem()
}