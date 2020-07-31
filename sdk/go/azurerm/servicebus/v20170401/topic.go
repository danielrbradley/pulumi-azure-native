// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of topic resource.
type Topic struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of topic resource.
	Properties SBTopicPropertiesResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &TopicArgs{}
	}
	var resource Topic
	err := ctx.RegisterResource("azurerm:servicebus/v20170401:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azurerm:servicebus/v20170401:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// Properties of topic resource.
	Properties *SBTopicPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type TopicState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// Properties of topic resource.
	Properties SBTopicPropertiesResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// The topic name.
	Name string `pulumi:"name"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Properties of topic resource.
	Properties *SBTopicProperties `pulumi:"properties"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The topic name.
	Name pulumi.StringInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Properties of topic resource.
	Properties SBTopicPropertiesPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}