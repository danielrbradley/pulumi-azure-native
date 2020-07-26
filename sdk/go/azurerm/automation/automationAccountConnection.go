// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the connection.
type AutomationAccountConnection struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the properties of the connection.
	Properties ConnectionPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAutomationAccountConnection registers a new resource with the given unique name, arguments, and options.
func NewAutomationAccountConnection(ctx *pulumi.Context,
	name string, args *AutomationAccountConnectionArgs, opts ...pulumi.ResourceOption) (*AutomationAccountConnection, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
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
		args = &AutomationAccountConnectionArgs{}
	}
	var resource AutomationAccountConnection
	err := ctx.RegisterResource("azurerm:automation:AutomationAccountConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationAccountConnection gets an existing AutomationAccountConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationAccountConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationAccountConnectionState, opts ...pulumi.ResourceOption) (*AutomationAccountConnection, error) {
	var resource AutomationAccountConnection
	err := ctx.ReadResource("azurerm:automation:AutomationAccountConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationAccountConnection resources.
type automationAccountConnectionState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the properties of the connection.
	Properties *ConnectionPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type AutomationAccountConnectionState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the properties of the connection.
	Properties ConnectionPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (AutomationAccountConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountConnectionState)(nil)).Elem()
}

type automationAccountConnectionArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The parameters supplied to the create or update connection operation.
	Name string `pulumi:"name"`
	// Gets or sets the properties of the connection.
	Properties ConnectionCreateOrUpdateProperties `pulumi:"properties"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AutomationAccountConnection resource.
type AutomationAccountConnectionArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The parameters supplied to the create or update connection operation.
	Name pulumi.StringInput
	// Gets or sets the properties of the connection.
	Properties ConnectionCreateOrUpdatePropertiesInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (AutomationAccountConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountConnectionArgs)(nil)).Elem()
}