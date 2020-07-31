// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getLogger(args: GetLoggerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoggerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20191201:getLogger", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetLoggerArgs {
    /**
     * Logger identifier. Must be unique in the API Management service instance.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
}

/**
 * Logger details.
 */
export interface GetLoggerResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Logger entity contract properties.
     */
    readonly properties: outputs.apimanagement.v20191201.LoggerContractPropertiesResponse;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}