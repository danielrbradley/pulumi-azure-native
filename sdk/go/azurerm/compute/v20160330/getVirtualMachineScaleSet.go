// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160330

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVirtualMachineScaleSet(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetResult, error) {
	var rv LookupVirtualMachineScaleSetResult
	err := ctx.Invoke("azurerm:compute/v20160330:getVirtualMachineScaleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineScaleSetArgs struct {
	// The name of the VM scale set.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a Virtual Machine Scale Set.
type LookupVirtualMachineScaleSetResult struct {
	// The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentityResponse `pulumi:"identity"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Describes the properties of a Virtual Machine Scale Set.
	Properties VirtualMachineScaleSetPropertiesResponse `pulumi:"properties"`
	// The virtual machine scale set sku.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}