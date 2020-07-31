// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190513

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Base class for backup items.
type ProtectedItem struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ProtectedItemResource properties
	Properties ProtectedItemResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProtectedItem registers a new resource with the given unique name, arguments, and options.
func NewProtectedItem(ctx *pulumi.Context,
	name string, args *ProtectedItemArgs, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	if args == nil || args.ContainerName == nil {
		return nil, errors.New("missing required argument 'ContainerName'")
	}
	if args == nil || args.FabricName == nil {
		return nil, errors.New("missing required argument 'FabricName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VaultName == nil {
		return nil, errors.New("missing required argument 'VaultName'")
	}
	if args == nil {
		args = &ProtectedItemArgs{}
	}
	var resource ProtectedItem
	err := ctx.RegisterResource("azurerm:recoveryservices/v20190513:ProtectedItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectedItem gets an existing ProtectedItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectedItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectedItemState, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	var resource ProtectedItem
	err := ctx.ReadResource("azurerm:recoveryservices/v20190513:ProtectedItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectedItem resources.
type protectedItemState struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name associated with the resource.
	Name *string `pulumi:"name"`
	// ProtectedItemResource properties
	Properties *ProtectedItemResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type *string `pulumi:"type"`
}

type ProtectedItemState struct {
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name associated with the resource.
	Name pulumi.StringPtrInput
	// ProtectedItemResource properties
	Properties ProtectedItemResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringPtrInput
}

func (ProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemState)(nil)).Elem()
}

type protectedItemArgs struct {
	// Container name associated with the backup item.
	ContainerName string `pulumi:"containerName"`
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Fabric name associated with the backup item.
	FabricName string `pulumi:"fabricName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Item name to be backed up.
	Name string `pulumi:"name"`
	// ProtectedItemResource properties
	Properties *ProtectedItemDefinition `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a ProtectedItem resource.
type ProtectedItemArgs struct {
	// Container name associated with the backup item.
	ContainerName pulumi.StringInput
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Fabric name associated with the backup item.
	FabricName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Item name to be backed up.
	Name pulumi.StringInput
	// ProtectedItemResource properties
	Properties ProtectedItemDefinitionPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the recovery services vault.
	VaultName pulumi.StringInput
}

func (ProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemArgs)(nil)).Elem()
}