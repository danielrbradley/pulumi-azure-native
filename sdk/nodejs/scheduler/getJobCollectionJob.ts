// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getJobCollectionJob(args: GetJobCollectionJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobCollectionJobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:scheduler:getJobCollectionJob", {
        "jobCollectionName": args.jobCollectionName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetJobCollectionJobArgs {
    /**
     * The job collection name.
     */
    readonly jobCollectionName: string;
    /**
     * The job name.
     */
    readonly name: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

export interface GetJobCollectionJobResult {
    /**
     * Gets the job resource name.
     */
    readonly name: string;
    /**
     * Gets or sets the job properties.
     */
    readonly properties: outputs.scheduler.JobPropertiesResponse;
    /**
     * Gets the job resource type.
     */
    readonly type: string;
}