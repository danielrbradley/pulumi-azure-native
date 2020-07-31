// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801
{
    /// <summary>
    /// Represents a web app
    /// </summary>
    public partial class Site : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.SiteResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Site resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Site(string name, SiteArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/v20150801:Site", name, args ?? new SiteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Site(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/v20150801:Site", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Site resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Site Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Site(name, id, options);
        }
    }

    public sealed class SiteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, web app hostname is force registered with DNS
        /// </summary>
        [Input("forceDnsRegistration")]
        public Input<string>? ForceDnsRegistration { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("properties")]
        public Input<Inputs.SitePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// If true, custom (non *.azurewebsites.net) domains associated with web app are not verified.
        /// </summary>
        [Input("skipCustomDomainVerification")]
        public Input<string>? SkipCustomDomainVerification { get; set; }

        /// <summary>
        /// If true web app hostname is not registered with DNS on creation. This parameter is
        ///             only used for app creation
        /// </summary>
        [Input("skipDnsRegistration")]
        public Input<string>? SkipDnsRegistration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Time to live in seconds for web app's default domain name
        /// </summary>
        [Input("ttlInSeconds")]
        public Input<string>? TtlInSeconds { get; set; }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SiteArgs()
        {
        }
    }
}