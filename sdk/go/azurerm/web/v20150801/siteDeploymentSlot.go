// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents user credentials used for publishing activity
type SiteDeploymentSlot struct {
	pulumi.CustomResourceState

	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name       pulumi.StringPtrOutput             `pulumi:"name"`
	Properties DeploymentResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSiteDeploymentSlot registers a new resource with the given unique name, arguments, and options.
func NewSiteDeploymentSlot(ctx *pulumi.Context,
	name string, args *SiteDeploymentSlotArgs, opts ...pulumi.ResourceOption) (*SiteDeploymentSlot, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &SiteDeploymentSlotArgs{}
	}
	var resource SiteDeploymentSlot
	err := ctx.RegisterResource("azurerm:web/v20150801:SiteDeploymentSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteDeploymentSlot gets an existing SiteDeploymentSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteDeploymentSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteDeploymentSlotState, opts ...pulumi.ResourceOption) (*SiteDeploymentSlot, error) {
	var resource SiteDeploymentSlot
	err := ctx.ReadResource("azurerm:web/v20150801:SiteDeploymentSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteDeploymentSlot resources.
type siteDeploymentSlotState struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name       *string                       `pulumi:"name"`
	Properties *DeploymentResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SiteDeploymentSlotState struct {
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name       pulumi.StringPtrInput
	Properties DeploymentResponsePropertiesPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteDeploymentSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDeploymentSlotState)(nil)).Elem()
}

type siteDeploymentSlotArgs struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Id
	Name       string                `pulumi:"name"`
	Properties *DeploymentProperties `pulumi:"properties"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SiteDeploymentSlot resource.
type SiteDeploymentSlotArgs struct {
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Resource Id
	Name       pulumi.StringInput
	Properties DeploymentPropertiesPtrInput
	// Name of resource group
	ResourceGroupName pulumi.StringInput
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteDeploymentSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDeploymentSlotArgs)(nil)).Elem()
}