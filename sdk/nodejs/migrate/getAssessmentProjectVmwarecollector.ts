// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getAssessmentProjectVmwarecollector(args: GetAssessmentProjectVmwarecollectorArgs, opts?: pulumi.InvokeOptions): Promise<GetAssessmentProjectVmwarecollectorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:migrate:getAssessmentProjectVmwarecollector", {
        "name": args.name,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAssessmentProjectVmwarecollectorArgs {
    /**
     * Unique name of a VMware collector within a project.
     */
    readonly name: string;
    /**
     * Name of the Azure Migrate project.
     */
    readonly projectName: string;
    /**
     * Name of the Azure Resource Group that project is part of.
     */
    readonly resourceGroupName: string;
}

export interface GetAssessmentProjectVmwarecollectorResult {
    readonly eTag?: string;
    readonly name: string;
    readonly properties: outputs.migrate.CollectorPropertiesResponse;
    readonly type: string;
}