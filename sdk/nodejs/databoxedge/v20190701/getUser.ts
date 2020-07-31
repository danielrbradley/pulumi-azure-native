// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:databoxedge/v20190701:getUser", {
        "deviceName": args.deviceName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetUserArgs {
    /**
     * The device name.
     */
    readonly deviceName: string;
    /**
     * The user name.
     */
    readonly name: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
 */
export interface GetUserResult {
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * The storage account credential properties.
     */
    readonly properties: outputs.databoxedge.v20190701.UserPropertiesResponse;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}