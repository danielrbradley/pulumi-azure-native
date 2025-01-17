// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getRoleManagementPolicyAssignment";
export * from "./roleManagementPolicyAssignment";

// Import resources to register:
import { RoleManagementPolicyAssignment } from "./roleManagementPolicyAssignment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:authorization/v20201001:RoleManagementPolicyAssignment":
                return new RoleManagementPolicyAssignment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "authorization/v20201001", _module)
