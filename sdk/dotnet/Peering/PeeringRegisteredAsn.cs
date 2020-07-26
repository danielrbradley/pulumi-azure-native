// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering
{
    /// <summary>
    /// The customer's ASN that is registered by the peering service provider.
    /// </summary>
    public partial class PeeringRegisteredAsn : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties that define a registered ASN.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PeeringRegisteredAsnPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PeeringRegisteredAsn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PeeringRegisteredAsn(string name, PeeringRegisteredAsnArgs args, CustomResourceOptions? options = null)
            : base("azurerm:peering:PeeringRegisteredAsn", name, args ?? new PeeringRegisteredAsnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PeeringRegisteredAsn(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:peering:PeeringRegisteredAsn", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PeeringRegisteredAsn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PeeringRegisteredAsn Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PeeringRegisteredAsn(name, id, options);
        }
    }

    public sealed class PeeringRegisteredAsnArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the ASN.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public Input<string> PeeringName { get; set; } = null!;

        /// <summary>
        /// The properties that define a registered ASN.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.PeeringRegisteredAsnPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public PeeringRegisteredAsnArgs()
        {
        }
    }
}