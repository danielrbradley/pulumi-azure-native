// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventGrid Domain.
//
// Deprecated: Version v20190601 will be removed in the next major version of the provider. Upgrade to version v20200401preview or later.
type Domain struct {
	pulumi.CustomResourceState

	// Endpoint for the domain.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the domain.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure-native:eventgrid/v20190601:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:eventgrid/v20190601:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// Endpoint for the domain.
func (o DomainOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Location of the resource.
func (o DomainOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource.
func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the domain.
func (o DomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Tags of the resource.
func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o DomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
