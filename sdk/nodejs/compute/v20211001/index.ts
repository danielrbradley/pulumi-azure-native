// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./gallery";
export * from "./galleryApplication";
export * from "./galleryApplicationVersion";
export * from "./galleryImage";
export * from "./galleryImageVersion";
export * from "./getGallery";
export * from "./getGalleryApplication";
export * from "./getGalleryApplicationVersion";
export * from "./getGalleryImage";
export * from "./getGalleryImageVersion";

// Export enums:
export * from "../../types/enums/compute/v20211001";

// Import resources to register:
import { Gallery } from "./gallery";
import { GalleryApplication } from "./galleryApplication";
import { GalleryApplicationVersion } from "./galleryApplicationVersion";
import { GalleryImage } from "./galleryImage";
import { GalleryImageVersion } from "./galleryImageVersion";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:compute/v20211001:Gallery":
                return new Gallery(name, <any>undefined, { urn })
            case "azure-native:compute/v20211001:GalleryApplication":
                return new GalleryApplication(name, <any>undefined, { urn })
            case "azure-native:compute/v20211001:GalleryApplicationVersion":
                return new GalleryApplicationVersion(name, <any>undefined, { urn })
            case "azure-native:compute/v20211001:GalleryImage":
                return new GalleryImage(name, <any>undefined, { urn })
            case "azure-native:compute/v20211001:GalleryImageVersion":
                return new GalleryImageVersion(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "compute/v20211001", _module)
