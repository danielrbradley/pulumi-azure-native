// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getvNetPeering(args: GetvNetPeeringArgs, opts?: pulumi.InvokeOptions): Promise<GetvNetPeeringResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:databricks/v20180401:getvNetPeering", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetvNetPeeringArgs {
    /**
     * The name of the workspace vNet peering.
     */
    readonly name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: string;
}

/**
 * Peerings in a VirtualNetwork resource
 */
export interface GetvNetPeeringResult {
    /**
     * Name of the virtual network peering resource
     */
    readonly name: string;
    /**
     * List of properties for vNet Peering
     */
    readonly properties: outputs.databricks.v20180401.VirtualNetworkPeeringPropertiesFormatResponse;
    /**
     * type of the virtual network peering resource
     */
    readonly type: string;
}