// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents user credentials used for publishing activity
//
// Deprecated: Version v20150801 will be removed in the next major version of the provider. Upgrade to version v20150801preview or later.
type SiteInstanceDeploymentSlot struct {
	pulumi.CustomResourceState

	// Active
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Author
	Author pulumi.StringPtrOutput `pulumi:"author"`
	// AuthorEmail
	AuthorEmail pulumi.StringPtrOutput `pulumi:"authorEmail"`
	// Deployer
	Deployer pulumi.StringPtrOutput `pulumi:"deployer"`
	// Detail
	Details pulumi.StringPtrOutput `pulumi:"details"`
	// EndTime
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Message
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// Resource Name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// StartTime
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// Status
	Status pulumi.IntPtrOutput `pulumi:"status"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSiteInstanceDeploymentSlot registers a new resource with the given unique name, arguments, and options.
func NewSiteInstanceDeploymentSlot(ctx *pulumi.Context,
	name string, args *SiteInstanceDeploymentSlotArgs, opts ...pulumi.ResourceOption) (*SiteInstanceDeploymentSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	var resource SiteInstanceDeploymentSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteInstanceDeploymentSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteInstanceDeploymentSlot gets an existing SiteInstanceDeploymentSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteInstanceDeploymentSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteInstanceDeploymentSlotState, opts ...pulumi.ResourceOption) (*SiteInstanceDeploymentSlot, error) {
	var resource SiteInstanceDeploymentSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteInstanceDeploymentSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteInstanceDeploymentSlot resources.
type siteInstanceDeploymentSlotState struct {
}

type SiteInstanceDeploymentSlotState struct {
}

func (SiteInstanceDeploymentSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteInstanceDeploymentSlotState)(nil)).Elem()
}

type siteInstanceDeploymentSlotArgs struct {
	// Active
	Active *bool `pulumi:"active"`
	// Author
	Author *string `pulumi:"author"`
	// AuthorEmail
	AuthorEmail *string `pulumi:"authorEmail"`
	// Deployer
	Deployer *string `pulumi:"deployer"`
	// Detail
	Details *string `pulumi:"details"`
	// EndTime
	EndTime *string `pulumi:"endTime"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Id of web app instance
	InstanceId string `pulumi:"instanceId"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Message
	Message *string `pulumi:"message"`
	// Resource Name
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
	// StartTime
	StartTime *string `pulumi:"startTime"`
	// Status
	Status *int `pulumi:"status"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SiteInstanceDeploymentSlot resource.
type SiteInstanceDeploymentSlotArgs struct {
	// Active
	Active pulumi.BoolPtrInput
	// Author
	Author pulumi.StringPtrInput
	// AuthorEmail
	AuthorEmail pulumi.StringPtrInput
	// Deployer
	Deployer pulumi.StringPtrInput
	// Detail
	Details pulumi.StringPtrInput
	// EndTime
	EndTime pulumi.StringPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Id of web app instance
	InstanceId pulumi.StringInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Message
	Message pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringInput
	// Name of resource group
	ResourceGroupName pulumi.StringInput
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput
	// StartTime
	StartTime pulumi.StringPtrInput
	// Status
	Status pulumi.IntPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteInstanceDeploymentSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteInstanceDeploymentSlotArgs)(nil)).Elem()
}

type SiteInstanceDeploymentSlotInput interface {
	pulumi.Input

	ToSiteInstanceDeploymentSlotOutput() SiteInstanceDeploymentSlotOutput
	ToSiteInstanceDeploymentSlotOutputWithContext(ctx context.Context) SiteInstanceDeploymentSlotOutput
}

func (*SiteInstanceDeploymentSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteInstanceDeploymentSlot)(nil)).Elem()
}

func (i *SiteInstanceDeploymentSlot) ToSiteInstanceDeploymentSlotOutput() SiteInstanceDeploymentSlotOutput {
	return i.ToSiteInstanceDeploymentSlotOutputWithContext(context.Background())
}

func (i *SiteInstanceDeploymentSlot) ToSiteInstanceDeploymentSlotOutputWithContext(ctx context.Context) SiteInstanceDeploymentSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteInstanceDeploymentSlotOutput)
}

type SiteInstanceDeploymentSlotOutput struct{ *pulumi.OutputState }

func (SiteInstanceDeploymentSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteInstanceDeploymentSlot)(nil)).Elem()
}

func (o SiteInstanceDeploymentSlotOutput) ToSiteInstanceDeploymentSlotOutput() SiteInstanceDeploymentSlotOutput {
	return o
}

func (o SiteInstanceDeploymentSlotOutput) ToSiteInstanceDeploymentSlotOutputWithContext(ctx context.Context) SiteInstanceDeploymentSlotOutput {
	return o
}

// Active
func (o SiteInstanceDeploymentSlotOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// Author
func (o SiteInstanceDeploymentSlotOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Author }).(pulumi.StringPtrOutput)
}

// AuthorEmail
func (o SiteInstanceDeploymentSlotOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

// Deployer
func (o SiteInstanceDeploymentSlotOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Deployer }).(pulumi.StringPtrOutput)
}

// Detail
func (o SiteInstanceDeploymentSlotOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Details }).(pulumi.StringPtrOutput)
}

// EndTime
func (o SiteInstanceDeploymentSlotOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Kind of resource
func (o SiteInstanceDeploymentSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location
func (o SiteInstanceDeploymentSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Message
func (o SiteInstanceDeploymentSlotOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o SiteInstanceDeploymentSlotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// StartTime
func (o SiteInstanceDeploymentSlotOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Status
func (o SiteInstanceDeploymentSlotOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.IntPtrOutput { return v.Status }).(pulumi.IntPtrOutput)
}

// Resource tags
func (o SiteInstanceDeploymentSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o SiteInstanceDeploymentSlotOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteInstanceDeploymentSlotOutput{})
}
