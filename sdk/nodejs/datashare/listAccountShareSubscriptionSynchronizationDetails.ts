// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function listAccountShareSubscriptionSynchronizationDetails(args: ListAccountShareSubscriptionSynchronizationDetailsArgs, opts?: pulumi.InvokeOptions): Promise<ListAccountShareSubscriptionSynchronizationDetailsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:datashare:listAccountShareSubscriptionSynchronizationDetails", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
        "shareSubscriptionName": args.shareSubscriptionName,
        "synchronizationId": args.synchronizationId,
    }, opts);
}

export interface ListAccountShareSubscriptionSynchronizationDetailsArgs {
    /**
     * The name of the share account.
     */
    readonly accountName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the share subscription.
     */
    readonly shareSubscriptionName: string;
    /**
     * Synchronization id
     */
    readonly synchronizationId: string;
}

/**
 * details of synchronization
 */
export interface ListAccountShareSubscriptionSynchronizationDetailsResult {
    /**
     * The Url of next result page.
     */
    readonly nextLink?: string;
    /**
     * Collection of items of type DataTransferObjects.
     */
    readonly value: outputs.datashare.SynchronizationDetailsResponse[];
}