// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualMachineExtension(args: GetVirtualMachineExtensionArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineExtensionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:compute/v20200601:getVirtualMachineExtension", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "vmName": args.vmName,
    }, opts);
}

export interface GetVirtualMachineExtensionArgs {
    /**
     * The name of the virtual machine extension.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the virtual machine containing the extension.
     */
    readonly vmName: string;
}

/**
 * Describes a Virtual Machine Extension.
 */
export interface GetVirtualMachineExtensionResult {
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Describes the properties of a Virtual Machine Extension.
     */
    readonly properties: outputs.compute.v20200601.VirtualMachineExtensionPropertiesResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}