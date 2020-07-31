// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getConnectionMonitor(args: GetConnectionMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionMonitorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20171101:getConnectionMonitor", {
        "name": args.name,
        "networkWatcherName": args.networkWatcherName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConnectionMonitorArgs {
    /**
     * The name of the connection monitor.
     */
    readonly name: string;
    /**
     * The name of the Network Watcher resource.
     */
    readonly networkWatcherName: string;
    /**
     * The name of the resource group containing Network Watcher.
     */
    readonly resourceGroupName: string;
}

/**
 * Information about the connection monitor.
 */
export interface GetConnectionMonitorResult {
    readonly etag?: string;
    /**
     * Connection monitor location.
     */
    readonly location?: string;
    /**
     * Name of the connection monitor.
     */
    readonly name: string;
    /**
     * Describes the properties of a connection monitor.
     */
    readonly properties: outputs.network.v20171101.ConnectionMonitorResultPropertiesResponse;
    /**
     * Connection monitor tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Connection monitor type.
     */
    readonly type: string;
}