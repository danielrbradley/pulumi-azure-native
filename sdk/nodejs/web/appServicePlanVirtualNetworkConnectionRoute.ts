// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Virtual Network route contract used to pass routing information for a Virtual Network.
 */
export class AppServicePlanVirtualNetworkConnectionRoute extends pulumi.CustomResource {
    /**
     * Get an existing AppServicePlanVirtualNetworkConnectionRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppServicePlanVirtualNetworkConnectionRoute {
        return new AppServicePlanVirtualNetworkConnectionRoute(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web:AppServicePlanVirtualNetworkConnectionRoute';

    /**
     * Returns true if the given object is an instance of AppServicePlanVirtualNetworkConnectionRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppServicePlanVirtualNetworkConnectionRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppServicePlanVirtualNetworkConnectionRoute.__pulumiType;
    }

    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * VnetRoute resource specific properties
     */
    public readonly properties!: pulumi.Output<outputs.web.VnetRouteResponseProperties>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AppServicePlanVirtualNetworkConnectionRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppServicePlanVirtualNetworkConnectionRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppServicePlanVirtualNetworkConnectionRouteArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AppServicePlanVirtualNetworkConnectionRouteArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vnetName === undefined) {
                throw new Error("Missing required property 'vnetName'");
            }
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["vnetName"] = args ? args.vnetName : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AppServicePlanVirtualNetworkConnectionRoute.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppServicePlanVirtualNetworkConnectionRoute resource.
 */
export interface AppServicePlanVirtualNetworkConnectionRouteArgs {
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the Virtual Network route.
     */
    readonly name: pulumi.Input<string>;
    /**
     * VnetRoute resource specific properties
     */
    readonly properties?: pulumi.Input<inputs.web.VnetRouteProperties>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the Virtual Network.
     */
    readonly vnetName: pulumi.Input<string>;
}