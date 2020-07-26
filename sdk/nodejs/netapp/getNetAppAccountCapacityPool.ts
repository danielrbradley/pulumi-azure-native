// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getNetAppAccountCapacityPool(args: GetNetAppAccountCapacityPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetNetAppAccountCapacityPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:netapp:getNetAppAccountCapacityPool", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNetAppAccountCapacityPoolArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: string;
    /**
     * The name of the capacity pool
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Capacity pool resource
 */
export interface GetNetAppAccountCapacityPoolResult {
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Capacity pool properties
     */
    readonly properties: outputs.netapp.PoolPropertiesResponse;
    /**
     * Resource tags
     */
    readonly tags?: outputs.netapp.ResourceTagsResponse;
    /**
     * Resource type
     */
    readonly type: string;
}