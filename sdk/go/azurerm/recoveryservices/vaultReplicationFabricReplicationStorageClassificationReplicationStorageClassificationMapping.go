// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Storage mapping object.
type VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the storage mapping object.
	Properties StorageClassificationMappingPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping registers a new resource with the given unique name, arguments, and options.
func NewVaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping(ctx *pulumi.Context,
	name string, args *VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingArgs, opts ...pulumi.ResourceOption) (*VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping, error) {
	if args == nil || args.FabricName == nil {
		return nil, errors.New("missing required argument 'FabricName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil || args.StorageClassificationName == nil {
		return nil, errors.New("missing required argument 'StorageClassificationName'")
	}
	if args == nil {
		args = &VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingArgs{}
	}
	var resource VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping
	err := ctx.RegisterResource("azurerm:recoveryservices:VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping gets an existing VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingState, opts ...pulumi.ResourceOption) (*VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping, error) {
	var resource VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping
	err := ctx.ReadResource("azurerm:recoveryservices:VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping resources.
type vaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingState struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Properties of the storage mapping object.
	Properties *StorageClassificationMappingPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingState struct {
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// Properties of the storage mapping object.
	Properties StorageClassificationMappingPropertiesResponsePtrInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingState)(nil)).Elem()
}

type vaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// Storage classification mapping name.
	Name string `pulumi:"name"`
	// Storage mapping input properties.
	Properties *StorageMappingInputProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
	// Storage classification name.
	StorageClassificationName string `pulumi:"storageClassificationName"`
}

// The set of arguments for constructing a VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMapping resource.
type VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// Storage classification mapping name.
	Name pulumi.StringInput
	// Storage mapping input properties.
	Properties StorageMappingInputPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
	// Storage classification name.
	StorageClassificationName pulumi.StringInput
}

func (VaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultReplicationFabricReplicationStorageClassificationReplicationStorageClassificationMappingArgs)(nil)).Elem()
}