// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRecordSet(args: GetRecordSetArgs, opts?: pulumi.InvokeOptions): Promise<GetRecordSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20160401:getRecordSet", {
        "name": args.name,
        "recordType": args.recordType,
        "resourceGroupName": args.resourceGroupName,
        "zoneName": args.zoneName,
    }, opts);
}

export interface GetRecordSetArgs {
    /**
     * The name of the record set, relative to the name of the zone.
     */
    readonly name: string;
    /**
     * The type of DNS record in this record set.
     */
    readonly recordType: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the DNS zone (without a terminating dot).
     */
    readonly zoneName: string;
}

/**
 * Describes a DNS record set (a collection of DNS records with the same name and type).
 */
export interface GetRecordSetResult {
    /**
     * The etag of the record set.
     */
    readonly etag?: string;
    /**
     * The name of the record set.
     */
    readonly name?: string;
    /**
     * The properties of the record set.
     */
    readonly properties: outputs.network.v20160401.RecordSetPropertiesResponse;
    /**
     * The type of the record set.
     */
    readonly type?: string;
}