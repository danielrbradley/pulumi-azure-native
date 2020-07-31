// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSmartDetectorAlertRule(args: GetSmartDetectorAlertRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetSmartDetectorAlertRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:alertsmanagement/v20190301:getSmartDetectorAlertRule", {
        "expandDetector": args.expandDetector,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetSmartDetectorAlertRuleArgs {
    /**
     * Indicates if Smart Detector should be expanded.
     */
    readonly expandDetector?: boolean;
    /**
     * The name of the alert rule.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The alert rule information
 */
export interface GetSmartDetectorAlertRuleResult {
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * The properties of the alert rule.
     */
    readonly properties: outputs.alertsmanagement.v20190301.AlertRulePropertiesResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}