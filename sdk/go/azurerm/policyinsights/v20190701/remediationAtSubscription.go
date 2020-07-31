// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The remediation definition.
type RemediationAtSubscription struct {
	pulumi.CustomResourceState

	// The name of the remediation.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties for the remediation.
	Properties RemediationPropertiesResponseOutput `pulumi:"properties"`
	// The type of the remediation.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRemediationAtSubscription registers a new resource with the given unique name, arguments, and options.
func NewRemediationAtSubscription(ctx *pulumi.Context,
	name string, args *RemediationAtSubscriptionArgs, opts ...pulumi.ResourceOption) (*RemediationAtSubscription, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &RemediationAtSubscriptionArgs{}
	}
	var resource RemediationAtSubscription
	err := ctx.RegisterResource("azurerm:policyinsights/v20190701:RemediationAtSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediationAtSubscription gets an existing RemediationAtSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediationAtSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtSubscriptionState, opts ...pulumi.ResourceOption) (*RemediationAtSubscription, error) {
	var resource RemediationAtSubscription
	err := ctx.ReadResource("azurerm:policyinsights/v20190701:RemediationAtSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemediationAtSubscription resources.
type remediationAtSubscriptionState struct {
	// The name of the remediation.
	Name *string `pulumi:"name"`
	// Properties for the remediation.
	Properties *RemediationPropertiesResponse `pulumi:"properties"`
	// The type of the remediation.
	Type *string `pulumi:"type"`
}

type RemediationAtSubscriptionState struct {
	// The name of the remediation.
	Name pulumi.StringPtrInput
	// Properties for the remediation.
	Properties RemediationPropertiesResponsePtrInput
	// The type of the remediation.
	Type pulumi.StringPtrInput
}

func (RemediationAtSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtSubscriptionState)(nil)).Elem()
}

type remediationAtSubscriptionArgs struct {
	// The name of the remediation.
	Name string `pulumi:"name"`
	// Properties for the remediation.
	Properties *RemediationProperties `pulumi:"properties"`
}

// The set of arguments for constructing a RemediationAtSubscription resource.
type RemediationAtSubscriptionArgs struct {
	// The name of the remediation.
	Name pulumi.StringInput
	// Properties for the remediation.
	Properties RemediationPropertiesPtrInput
}

func (RemediationAtSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtSubscriptionArgs)(nil)).Elem()
}