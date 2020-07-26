// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDomainDomainOwnershipIdentifier(args: GetDomainDomainOwnershipIdentifierArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainDomainOwnershipIdentifierResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:domainregistration:getDomainDomainOwnershipIdentifier", {
        "domainName": args.domainName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDomainDomainOwnershipIdentifierArgs {
    /**
     * Name of domain.
     */
    readonly domainName: string;
    /**
     * Name of identifier.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Domain ownership Identifier.
 */
export interface GetDomainDomainOwnershipIdentifierResult {
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * DomainOwnershipIdentifier resource specific properties
     */
    readonly properties: outputs.domainregistration.DomainOwnershipIdentifierResponseProperties;
    /**
     * Resource type.
     */
    readonly type: string;
}