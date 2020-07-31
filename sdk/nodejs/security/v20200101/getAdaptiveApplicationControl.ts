// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getAdaptiveApplicationControl(args: GetAdaptiveApplicationControlArgs, opts?: pulumi.InvokeOptions): Promise<GetAdaptiveApplicationControlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:security/v20200101:getAdaptiveApplicationControl", {
        "ascLocation": args.ascLocation,
        "name": args.name,
    }, opts);
}

export interface GetAdaptiveApplicationControlArgs {
    /**
     * The location where ASC stores the data of the subscription. can be retrieved from Get locations
     */
    readonly ascLocation: string;
    /**
     * Name of an application control VM/server group
     */
    readonly name: string;
}

export interface GetAdaptiveApplicationControlResult {
    /**
     * Location where the resource is stored
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Represents a VM/server group and set of rules to be allowed running on a machine
     */
    readonly properties: outputs.security.v20200101.AppWhitelistingGroupDataResponse;
    /**
     * Resource type
     */
    readonly type: string;
}