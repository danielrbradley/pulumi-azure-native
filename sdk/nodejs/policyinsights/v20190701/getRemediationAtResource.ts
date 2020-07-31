// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRemediationAtResource(args: GetRemediationAtResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetRemediationAtResourceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:policyinsights/v20190701:getRemediationAtResource", {
        "name": args.name,
        "resourceId": args.resourceId,
    }, opts);
}

export interface GetRemediationAtResourceArgs {
    /**
     * The name of the remediation.
     */
    readonly name: string;
    /**
     * Resource ID.
     */
    readonly resourceId: string;
}

/**
 * The remediation definition.
 */
export interface GetRemediationAtResourceResult {
    /**
     * The name of the remediation.
     */
    readonly name: string;
    /**
     * Properties for the remediation.
     */
    readonly properties: outputs.policyinsights.v20190701.RemediationPropertiesResponse;
    /**
     * The type of the remediation.
     */
    readonly type: string;
}