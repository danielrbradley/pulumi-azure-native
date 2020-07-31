// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Tag Contract details.
type TagByOperation struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tag entity contract properties.
	Properties TagContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagByOperation registers a new resource with the given unique name, arguments, and options.
func NewTagByOperation(ctx *pulumi.Context,
	name string, args *TagByOperationArgs, opts ...pulumi.ResourceOption) (*TagByOperation, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.OperationId == nil {
		return nil, errors.New("missing required argument 'OperationId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &TagByOperationArgs{}
	}
	var resource TagByOperation
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:TagByOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagByOperation gets an existing TagByOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagByOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByOperationState, opts ...pulumi.ResourceOption) (*TagByOperation, error) {
	var resource TagByOperation
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:TagByOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagByOperation resources.
type tagByOperationState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Tag entity contract properties.
	Properties *TagContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type TagByOperationState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Tag entity contract properties.
	Properties TagContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (TagByOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByOperationState)(nil)).Elem()
}

type tagByOperationArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Tag identifier. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a TagByOperation resource.
type TagByOperationArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	Name pulumi.StringInput
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (TagByOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByOperationArgs)(nil)).Elem()
}