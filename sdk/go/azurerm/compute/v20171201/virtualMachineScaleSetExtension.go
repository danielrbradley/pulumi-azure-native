// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Virtual Machine Scale Set Extension.
type VirtualMachineScaleSetExtension struct {
	pulumi.CustomResourceState

	// The name of the extension.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Describes the properties of a Virtual Machine Scale Set Extension.
	Properties VirtualMachineScaleSetExtensionPropertiesResponseOutput `pulumi:"properties"`
}

// NewVirtualMachineScaleSetExtension registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VmScaleSetName == nil {
		return nil, errors.New("missing required argument 'VmScaleSetName'")
	}
	if args == nil {
		args = &VirtualMachineScaleSetExtensionArgs{}
	}
	var resource VirtualMachineScaleSetExtension
	err := ctx.RegisterResource("azurerm:compute/v20171201:VirtualMachineScaleSetExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetExtension gets an existing VirtualMachineScaleSetExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	var resource VirtualMachineScaleSetExtension
	err := ctx.ReadResource("azurerm:compute/v20171201:VirtualMachineScaleSetExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetExtension resources.
type virtualMachineScaleSetExtensionState struct {
	// The name of the extension.
	Name *string `pulumi:"name"`
	// Describes the properties of a Virtual Machine Scale Set Extension.
	Properties *VirtualMachineScaleSetExtensionPropertiesResponse `pulumi:"properties"`
}

type VirtualMachineScaleSetExtensionState struct {
	// The name of the extension.
	Name pulumi.StringPtrInput
	// Describes the properties of a Virtual Machine Scale Set Extension.
	Properties VirtualMachineScaleSetExtensionPropertiesResponsePtrInput
}

func (VirtualMachineScaleSetExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetExtensionArgs struct {
	// The name of the VM scale set extension.
	Name string `pulumi:"name"`
	// Describes the properties of a Virtual Machine Scale Set Extension.
	Properties *VirtualMachineScaleSetExtensionProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VM scale set where the extension should be create or updated.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
}

// The set of arguments for constructing a VirtualMachineScaleSetExtension resource.
type VirtualMachineScaleSetExtensionArgs struct {
	// The name of the VM scale set extension.
	Name pulumi.StringInput
	// Describes the properties of a Virtual Machine Scale Set Extension.
	Properties VirtualMachineScaleSetExtensionPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the VM scale set where the extension should be create or updated.
	VmScaleSetName pulumi.StringInput
}

func (VirtualMachineScaleSetExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionArgs)(nil)).Elem()
}