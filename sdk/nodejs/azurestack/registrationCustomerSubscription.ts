// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Customer subscription.
 */
export class RegistrationCustomerSubscription extends pulumi.CustomResource {
    /**
     * Get an existing RegistrationCustomerSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegistrationCustomerSubscription {
        return new RegistrationCustomerSubscription(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:azurestack:RegistrationCustomerSubscription';

    /**
     * Returns true if the given object is an instance of RegistrationCustomerSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistrationCustomerSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistrationCustomerSubscription.__pulumiType;
    }

    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Customer subscription properties.
     */
    public readonly properties!: pulumi.Output<outputs.azurestack.CustomerSubscriptionPropertiesResponse>;
    /**
     * Type of Resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RegistrationCustomerSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistrationCustomerSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistrationCustomerSubscriptionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RegistrationCustomerSubscriptionArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.registrationName === undefined) {
                throw new Error("Missing required property 'registrationName'");
            }
            if (!args || args.resourceGroup === undefined) {
                throw new Error("Missing required property 'resourceGroup'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["registrationName"] = args ? args.registrationName : undefined;
            inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegistrationCustomerSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegistrationCustomerSubscription resource.
 */
export interface RegistrationCustomerSubscriptionArgs {
    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Name of the product.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Customer subscription properties.
     */
    readonly properties?: pulumi.Input<inputs.azurestack.CustomerSubscriptionProperties>;
    /**
     * Name of the Azure Stack registration.
     */
    readonly registrationName: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    readonly resourceGroup: pulumi.Input<string>;
}