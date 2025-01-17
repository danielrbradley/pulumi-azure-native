// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./customLocation";
export * from "./getCustomLocation";
export * from "./getResourceSyncRule";
export * from "./resourceSyncRule";

// Export enums:
export * from "../../types/enums/extendedlocation/v20210831preview";

// Import resources to register:
import { CustomLocation } from "./customLocation";
import { ResourceSyncRule } from "./resourceSyncRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:extendedlocation/v20210831preview:CustomLocation":
                return new CustomLocation(name, <any>undefined, { urn })
            case "azure-native:extendedlocation/v20210831preview:ResourceSyncRule":
                return new ResourceSyncRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "extendedlocation/v20210831preview", _module)
