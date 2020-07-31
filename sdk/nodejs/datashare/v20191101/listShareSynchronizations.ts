// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listShareSynchronizations(args: ListShareSynchronizationsArgs, opts?: pulumi.InvokeOptions): Promise<ListShareSynchronizationsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:datashare/v20191101:listShareSynchronizations", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
        "shareName": args.shareName,
    }, opts);
}

export interface ListShareSynchronizationsArgs {
    /**
     * The name of the share account.
     */
    readonly accountName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the share.
     */
    readonly shareName: string;
}

/**
 * List response for get ShareSynchronization.
 */
export interface ListShareSynchronizationsResult {
    /**
     * The Url of next result page.
     */
    readonly nextLink?: string;
    /**
     * Collection of items of type DataTransferObjects.
     */
    readonly value: outputs.datashare.v20191101.ShareSynchronizationResponse[];
}