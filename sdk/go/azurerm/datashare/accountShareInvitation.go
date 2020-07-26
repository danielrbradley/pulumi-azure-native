// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Invitation data transfer object.
type AccountShareInvitation struct {
	pulumi.CustomResourceState

	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties on the Invitation
	Properties InvitationPropertiesResponseOutput `pulumi:"properties"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccountShareInvitation registers a new resource with the given unique name, arguments, and options.
func NewAccountShareInvitation(ctx *pulumi.Context,
	name string, args *AccountShareInvitationArgs, opts ...pulumi.ResourceOption) (*AccountShareInvitation, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ShareName == nil {
		return nil, errors.New("missing required argument 'ShareName'")
	}
	if args == nil {
		args = &AccountShareInvitationArgs{}
	}
	var resource AccountShareInvitation
	err := ctx.RegisterResource("azurerm:datashare:AccountShareInvitation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountShareInvitation gets an existing AccountShareInvitation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountShareInvitation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountShareInvitationState, opts ...pulumi.ResourceOption) (*AccountShareInvitation, error) {
	var resource AccountShareInvitation
	err := ctx.ReadResource("azurerm:datashare:AccountShareInvitation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountShareInvitation resources.
type accountShareInvitationState struct {
	// Name of the azure resource
	Name *string `pulumi:"name"`
	// Properties on the Invitation
	Properties *InvitationPropertiesResponse `pulumi:"properties"`
	// Type of the azure resource
	Type *string `pulumi:"type"`
}

type AccountShareInvitationState struct {
	// Name of the azure resource
	Name pulumi.StringPtrInput
	// Properties on the Invitation
	Properties InvitationPropertiesResponsePtrInput
	// Type of the azure resource
	Type pulumi.StringPtrInput
}

func (AccountShareInvitationState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountShareInvitationState)(nil)).Elem()
}

type accountShareInvitationArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the invitation.
	Name string `pulumi:"name"`
	// Properties on the Invitation
	Properties *InvitationProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share to send the invitation for.
	ShareName string `pulumi:"shareName"`
}

// The set of arguments for constructing a AccountShareInvitation resource.
type AccountShareInvitationArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the invitation.
	Name pulumi.StringInput
	// Properties on the Invitation
	Properties InvitationPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share to send the invitation for.
	ShareName pulumi.StringInput
}

func (AccountShareInvitationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountShareInvitationArgs)(nil)).Elem()
}