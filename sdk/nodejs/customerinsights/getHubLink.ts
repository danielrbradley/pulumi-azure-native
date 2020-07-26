// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getHubLink(args: GetHubLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetHubLinkResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:customerinsights:getHubLink", {
        "hubName": args.hubName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHubLinkArgs {
    /**
     * The name of the hub.
     */
    readonly hubName: string;
    /**
     * The name of the link.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The link resource format.
 */
export interface GetHubLinkResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The definition of Link.
     */
    readonly properties: outputs.customerinsights.LinkDefinitionResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}