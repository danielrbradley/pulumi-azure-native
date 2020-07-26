// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a link to virtual network for a Private DNS zone.
type PrivateDnsZoneVirtualNetworkLink struct {
	pulumi.CustomResourceState

	// The ETag of the virtual network link.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the virtual network link to the Private DNS zone.
	Properties VirtualNetworkLinkPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateDnsZoneVirtualNetworkLink registers a new resource with the given unique name, arguments, and options.
func NewPrivateDnsZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *PrivateDnsZoneVirtualNetworkLinkArgs, opts ...pulumi.ResourceOption) (*PrivateDnsZoneVirtualNetworkLink, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PrivateZoneName == nil {
		return nil, errors.New("missing required argument 'PrivateZoneName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PrivateDnsZoneVirtualNetworkLinkArgs{}
	}
	var resource PrivateDnsZoneVirtualNetworkLink
	err := ctx.RegisterResource("azurerm:network:PrivateDnsZoneVirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDnsZoneVirtualNetworkLink gets an existing PrivateDnsZoneVirtualNetworkLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDnsZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDnsZoneVirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*PrivateDnsZoneVirtualNetworkLink, error) {
	var resource PrivateDnsZoneVirtualNetworkLink
	err := ctx.ReadResource("azurerm:network:PrivateDnsZoneVirtualNetworkLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDnsZoneVirtualNetworkLink resources.
type privateDnsZoneVirtualNetworkLinkState struct {
	// The ETag of the virtual network link.
	Etag *string `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Properties of the virtual network link to the Private DNS zone.
	Properties *VirtualNetworkLinkPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `pulumi:"type"`
}

type PrivateDnsZoneVirtualNetworkLinkState struct {
	// The ETag of the virtual network link.
	Etag pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Properties of the virtual network link to the Private DNS zone.
	Properties VirtualNetworkLinkPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type pulumi.StringPtrInput
}

func (PrivateDnsZoneVirtualNetworkLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneVirtualNetworkLinkState)(nil)).Elem()
}

type privateDnsZoneVirtualNetworkLinkArgs struct {
	// The ETag of the virtual network link.
	Etag *string `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the virtual network link.
	Name string `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName string `pulumi:"privateZoneName"`
	// Properties of the virtual network link to the Private DNS zone.
	Properties *VirtualNetworkLinkProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateDnsZoneVirtualNetworkLink resource.
type PrivateDnsZoneVirtualNetworkLinkArgs struct {
	// The ETag of the virtual network link.
	Etag pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the virtual network link.
	Name pulumi.StringInput
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName pulumi.StringInput
	// Properties of the virtual network link to the Private DNS zone.
	Properties VirtualNetworkLinkPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateDnsZoneVirtualNetworkLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneVirtualNetworkLinkArgs)(nil)).Elem()
}