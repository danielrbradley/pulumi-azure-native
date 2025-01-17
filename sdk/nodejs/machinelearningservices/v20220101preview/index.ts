// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./compute";
export * from "./getCompute";
export * from "./getPrivateEndpointConnection";
export * from "./getWorkspace";
export * from "./getWorkspaceConnection";
export * from "./listComputeKeys";
export * from "./listComputeNodes";
export * from "./listWorkspaceKeys";
export * from "./listWorkspaceNotebookAccessToken";
export * from "./listWorkspaceNotebookKeys";
export * from "./listWorkspaceStorageAccountKeys";
export * from "./privateEndpointConnection";
export * from "./workspace";
export * from "./workspaceConnection";

// Export enums:
export * from "../../types/enums/machinelearningservices/v20220101preview";

// Import resources to register:
import { Compute } from "./compute";
import { PrivateEndpointConnection } from "./privateEndpointConnection";
import { Workspace } from "./workspace";
import { WorkspaceConnection } from "./workspaceConnection";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:machinelearningservices/v20220101preview:Compute":
                return new Compute(name, <any>undefined, { urn })
            case "azure-native:machinelearningservices/v20220101preview:PrivateEndpointConnection":
                return new PrivateEndpointConnection(name, <any>undefined, { urn })
            case "azure-native:machinelearningservices/v20220101preview:Workspace":
                return new Workspace(name, <any>undefined, { urn })
            case "azure-native:machinelearningservices/v20220101preview:WorkspaceConnection":
                return new WorkspaceConnection(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "machinelearningservices/v20220101preview", _module)
