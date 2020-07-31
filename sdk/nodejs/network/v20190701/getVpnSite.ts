// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVpnSite(args: GetVpnSiteArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnSiteResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20190701:getVpnSite", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVpnSiteArgs {
    /**
     * The name of the VpnSite being retrieved.
     */
    readonly name: string;
    /**
     * The resource group name of the VpnSite.
     */
    readonly resourceGroupName: string;
}

/**
 * VpnSite Resource.
 */
export interface GetVpnSiteResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Properties of the VPN site.
     */
    readonly properties: outputs.network.v20190701.VpnSitePropertiesResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}