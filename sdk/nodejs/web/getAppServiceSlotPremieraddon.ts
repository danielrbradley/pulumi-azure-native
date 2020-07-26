// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getAppServiceSlotPremieraddon(args: GetAppServiceSlotPremieraddonArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServiceSlotPremieraddonResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web:getAppServiceSlotPremieraddon", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "slot": args.slot,
    }, opts);
}

export interface GetAppServiceSlotPremieraddonArgs {
    /**
     * Add-on name.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
    /**
     * Name of the deployment slot. If a slot is not specified, the API will get the named add-on for the production slot.
     */
    readonly slot: string;
}

/**
 * Premier add-on.
 */
export interface GetAppServiceSlotPremieraddonResult {
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Location.
     */
    readonly location: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * PremierAddOn resource specific properties
     */
    readonly properties: outputs.web.PremierAddOnResponseProperties;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}