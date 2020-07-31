// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus.V20150801
{
    /// <summary>
    /// Description of subscription resource.
    /// </summary>
    public partial class Subscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Description of Subscription Resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SubscriptionPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20150801:Subscription", name, args ?? new SubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20150801:Subscription", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, options);
        }
    }

    public sealed class SubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Subscription data center location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The subscription name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Description of Subscription Resource.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.SubscriptionPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The topic name.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        /// <summary>
        /// Resource manager type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SubscriptionArgs()
        {
        }
    }
}