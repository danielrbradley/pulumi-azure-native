// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWebAppSiteExtension(args: GetWebAppSiteExtensionArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAppSiteExtensionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web/v20181101:getWebAppSiteExtension", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetWebAppSiteExtensionArgs {
    /**
     * Site extension name.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Site Extension Information.
 */
export interface GetWebAppSiteExtensionResult {
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * SiteExtensionInfo resource specific properties
     */
    readonly properties: outputs.web.v20181101.SiteExtensionInfoResponseProperties;
    /**
     * Resource type.
     */
    readonly type: string;
}