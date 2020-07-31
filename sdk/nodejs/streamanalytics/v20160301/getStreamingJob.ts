// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getStreamingJob(args: GetStreamingJobArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamingJobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:streamanalytics/v20160301:getStreamingJob", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetStreamingJobArgs {
    /**
     * The name of the streaming job.
     */
    readonly name: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
}

/**
 * A streaming job object, containing all information associated with the named streaming job.
 */
export interface GetStreamingJobResult {
    /**
     * Resource location. Required on PUT (CreateOrReplace) requests.
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
     */
    readonly properties: outputs.streamanalytics.v20160301.StreamingJobPropertiesResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}