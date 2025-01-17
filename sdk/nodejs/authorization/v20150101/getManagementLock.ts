// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Management lock information.
 */
/** @deprecated Version v20150101 will be removed in the next major version of the provider. Upgrade to version v20170401 or later. */
export function getManagementLock(args: GetManagementLockArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementLockResult> {
    pulumi.log.warn("getManagementLock is deprecated: Version v20150101 will be removed in the next major version of the provider. Upgrade to version v20170401 or later.")
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:authorization/v20150101:getManagementLock", {
        "lockName": args.lockName,
    }, opts);
}

export interface GetManagementLockArgs {
    /**
     * Name of the management lock.
     */
    lockName: string;
}

/**
 * Management lock information.
 */
export interface GetManagementLockResult {
    /**
     * The Id of the lock.
     */
    readonly id: string;
    /**
     * The lock level of the management lock.
     */
    readonly level?: string;
    /**
     * The name of the lock.
     */
    readonly name?: string;
    /**
     * The notes of the management lock.
     */
    readonly notes?: string;
    /**
     * The type of the lock.
     */
    readonly type: string;
}

export function getManagementLockOutput(args: GetManagementLockOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetManagementLockResult> {
    return pulumi.output(args).apply(a => getManagementLock(a, opts))
}

export interface GetManagementLockOutputArgs {
    /**
     * Name of the management lock.
     */
    lockName: pulumi.Input<string>;
}
