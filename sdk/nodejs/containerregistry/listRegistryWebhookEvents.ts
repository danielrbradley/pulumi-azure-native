// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function listRegistryWebhookEvents(args: ListRegistryWebhookEventsArgs, opts?: pulumi.InvokeOptions): Promise<ListRegistryWebhookEventsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:containerregistry:listRegistryWebhookEvents", {
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
        "webhookName": args.webhookName,
    }, opts);
}

export interface ListRegistryWebhookEventsArgs {
    /**
     * The name of the container registry.
     */
    readonly registryName: string;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the webhook.
     */
    readonly webhookName: string;
}

/**
 * The result of a request to list events for a webhook.
 */
export interface ListRegistryWebhookEventsResult {
    /**
     * The URI that can be used to request the next list of events.
     */
    readonly nextLink?: string;
    /**
     * The list of events. Since this list may be incomplete, the nextLink field should be used to request the next list of events.
     */
    readonly value?: outputs.containerregistry.EventResponse[];
}