// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Registration assignment.
type RegistrationAssignment struct {
	pulumi.CustomResourceState

	// Name of the registration assignment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a registration assignment.
	Properties RegistrationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistrationAssignment registers a new resource with the given unique name, arguments, and options.
func NewRegistrationAssignment(ctx *pulumi.Context,
	name string, args *RegistrationAssignmentArgs, opts ...pulumi.ResourceOption) (*RegistrationAssignment, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &RegistrationAssignmentArgs{}
	}
	var resource RegistrationAssignment
	err := ctx.RegisterResource("azurerm:managedservices/v20190601:RegistrationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistrationAssignment gets an existing RegistrationAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistrationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationAssignmentState, opts ...pulumi.ResourceOption) (*RegistrationAssignment, error) {
	var resource RegistrationAssignment
	err := ctx.ReadResource("azurerm:managedservices/v20190601:RegistrationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistrationAssignment resources.
type registrationAssignmentState struct {
	// Name of the registration assignment.
	Name *string `pulumi:"name"`
	// Properties of a registration assignment.
	Properties *RegistrationAssignmentPropertiesResponse `pulumi:"properties"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type RegistrationAssignmentState struct {
	// Name of the registration assignment.
	Name pulumi.StringPtrInput
	// Properties of a registration assignment.
	Properties RegistrationAssignmentPropertiesResponsePtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (RegistrationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationAssignmentState)(nil)).Elem()
}

type registrationAssignmentArgs struct {
	// Guid of the registration assignment.
	Name string `pulumi:"name"`
	// Properties of a registration assignment.
	Properties *RegistrationAssignmentProperties `pulumi:"properties"`
	// Scope of the resource.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RegistrationAssignment resource.
type RegistrationAssignmentArgs struct {
	// Guid of the registration assignment.
	Name pulumi.StringInput
	// Properties of a registration assignment.
	Properties RegistrationAssignmentPropertiesPtrInput
	// Scope of the resource.
	Scope pulumi.StringInput
}

func (RegistrationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationAssignmentArgs)(nil)).Elem()
}