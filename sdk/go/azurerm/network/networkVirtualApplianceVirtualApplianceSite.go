// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Virtual Appliance Site resource.
type NetworkVirtualApplianceVirtualApplianceSite struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Name of the virtual appliance site.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The properties of the Virtual Appliance Sites.
	Properties VirtualApplianceSitePropertiesResponseOutput `pulumi:"properties"`
	// Site type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkVirtualApplianceVirtualApplianceSite registers a new resource with the given unique name, arguments, and options.
func NewNetworkVirtualApplianceVirtualApplianceSite(ctx *pulumi.Context,
	name string, args *NetworkVirtualApplianceVirtualApplianceSiteArgs, opts ...pulumi.ResourceOption) (*NetworkVirtualApplianceVirtualApplianceSite, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NetworkVirtualApplianceName == nil {
		return nil, errors.New("missing required argument 'NetworkVirtualApplianceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkVirtualApplianceVirtualApplianceSiteArgs{}
	}
	var resource NetworkVirtualApplianceVirtualApplianceSite
	err := ctx.RegisterResource("azurerm:network:NetworkVirtualApplianceVirtualApplianceSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkVirtualApplianceVirtualApplianceSite gets an existing NetworkVirtualApplianceVirtualApplianceSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkVirtualApplianceVirtualApplianceSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkVirtualApplianceVirtualApplianceSiteState, opts ...pulumi.ResourceOption) (*NetworkVirtualApplianceVirtualApplianceSite, error) {
	var resource NetworkVirtualApplianceVirtualApplianceSite
	err := ctx.ReadResource("azurerm:network:NetworkVirtualApplianceVirtualApplianceSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkVirtualApplianceVirtualApplianceSite resources.
type networkVirtualApplianceVirtualApplianceSiteState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Name of the virtual appliance site.
	Name *string `pulumi:"name"`
	// The properties of the Virtual Appliance Sites.
	Properties *VirtualApplianceSitePropertiesResponse `pulumi:"properties"`
	// Site type.
	Type *string `pulumi:"type"`
}

type NetworkVirtualApplianceVirtualApplianceSiteState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Name of the virtual appliance site.
	Name pulumi.StringPtrInput
	// The properties of the Virtual Appliance Sites.
	Properties VirtualApplianceSitePropertiesResponsePtrInput
	// Site type.
	Type pulumi.StringPtrInput
}

func (NetworkVirtualApplianceVirtualApplianceSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceVirtualApplianceSiteState)(nil)).Elem()
}

type networkVirtualApplianceVirtualApplianceSiteArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the site.
	Name string `pulumi:"name"`
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	// The properties of the Virtual Appliance Sites.
	Properties *VirtualApplianceSiteProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NetworkVirtualApplianceVirtualApplianceSite resource.
type NetworkVirtualApplianceVirtualApplianceSiteArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the site.
	Name pulumi.StringInput
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringInput
	// The properties of the Virtual Appliance Sites.
	Properties VirtualApplianceSitePropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (NetworkVirtualApplianceVirtualApplianceSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceVirtualApplianceSiteArgs)(nil)).Elem()
}