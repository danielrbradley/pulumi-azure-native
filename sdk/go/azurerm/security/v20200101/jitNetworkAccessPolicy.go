// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type JitNetworkAccessPolicy struct {
	pulumi.CustomResourceState

	// Kind of the resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Location where the resource is stored
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name       pulumi.StringOutput                            `pulumi:"name"`
	Properties JitNetworkAccessPolicyPropertiesResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJitNetworkAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, args *JitNetworkAccessPolicyArgs, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	if args == nil || args.AscLocation == nil {
		return nil, errors.New("missing required argument 'AscLocation'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &JitNetworkAccessPolicyArgs{}
	}
	var resource JitNetworkAccessPolicy
	err := ctx.RegisterResource("azurerm:security/v20200101:JitNetworkAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJitNetworkAccessPolicy gets an existing JitNetworkAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JitNetworkAccessPolicyState, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	var resource JitNetworkAccessPolicy
	err := ctx.ReadResource("azurerm:security/v20200101:JitNetworkAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JitNetworkAccessPolicy resources.
type jitNetworkAccessPolicyState struct {
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name       *string                                   `pulumi:"name"`
	Properties *JitNetworkAccessPolicyPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type JitNetworkAccessPolicyState struct {
	// Kind of the resource
	Kind pulumi.StringPtrInput
	// Location where the resource is stored
	Location pulumi.StringPtrInput
	// Resource name
	Name       pulumi.StringPtrInput
	Properties JitNetworkAccessPolicyPropertiesResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (JitNetworkAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyState)(nil)).Elem()
}

type jitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Name of a Just-in-Time access configuration policy.
	Name       string                           `pulumi:"name"`
	Properties JitNetworkAccessPolicyProperties `pulumi:"properties"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a JitNetworkAccessPolicy resource.
type JitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation pulumi.StringInput
	// Kind of the resource
	Kind pulumi.StringPtrInput
	// Name of a Just-in-Time access configuration policy.
	Name       pulumi.StringInput
	Properties JitNetworkAccessPolicyPropertiesInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (JitNetworkAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyArgs)(nil)).Elem()
}