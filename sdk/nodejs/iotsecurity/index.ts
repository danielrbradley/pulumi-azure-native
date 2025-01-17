// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./defenderSetting";
export * from "./deviceGroup";
export * from "./getDefenderSetting";
export * from "./getDeviceGroup";
export * from "./getOnPremiseSensor";
export * from "./getSensor";
export * from "./getSite";
export * from "./onPremiseSensor";
export * from "./sensor";
export * from "./site";

// Export enums:
export * from "../types/enums/iotsecurity";

// Export sub-modules:
import * as v20210201preview from "./v20210201preview";

export {
    v20210201preview,
};

// Import resources to register:
import { DefenderSetting } from "./defenderSetting";
import { DeviceGroup } from "./deviceGroup";
import { OnPremiseSensor } from "./onPremiseSensor";
import { Sensor } from "./sensor";
import { Site } from "./site";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:iotsecurity:DefenderSetting":
                return new DefenderSetting(name, <any>undefined, { urn })
            case "azure-native:iotsecurity:DeviceGroup":
                return new DeviceGroup(name, <any>undefined, { urn })
            case "azure-native:iotsecurity:OnPremiseSensor":
                return new OnPremiseSensor(name, <any>undefined, { urn })
            case "azure-native:iotsecurity:Sensor":
                return new Sensor(name, <any>undefined, { urn })
            case "azure-native:iotsecurity:Site":
                return new Site(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "iotsecurity", _module)
