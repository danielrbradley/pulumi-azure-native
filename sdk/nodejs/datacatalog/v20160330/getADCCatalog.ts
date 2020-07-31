// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getADCCatalog(args: GetADCCatalogArgs, opts?: pulumi.InvokeOptions): Promise<GetADCCatalogResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:datacatalog/v20160330:getADCCatalog", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetADCCatalogArgs {
    /**
     * The name of the data catalog in the specified subscription and resource group.
     */
    readonly name: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Azure Data Catalog.
 */
export interface GetADCCatalogResult {
    /**
     * Resource etag
     */
    readonly etag?: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Azure Data Catalog properties.
     */
    readonly properties: outputs.datacatalog.v20160330.ADCCatalogPropertiesResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}