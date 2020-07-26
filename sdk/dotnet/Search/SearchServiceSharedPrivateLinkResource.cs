// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search
{
    /// <summary>
    /// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
    /// </summary>
    public partial class SearchServiceSharedPrivateLinkResource : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the shared private link resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Describes the properties of a Shared Private Link Resource managed by the Azure Cognitive Search service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SharedPrivateLinkResourcePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SearchServiceSharedPrivateLinkResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SearchServiceSharedPrivateLinkResource(string name, SearchServiceSharedPrivateLinkResourceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:search:SearchServiceSharedPrivateLinkResource", name, args ?? new SearchServiceSharedPrivateLinkResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SearchServiceSharedPrivateLinkResource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:search:SearchServiceSharedPrivateLinkResource", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SearchServiceSharedPrivateLinkResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SearchServiceSharedPrivateLinkResource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SearchServiceSharedPrivateLinkResource(name, id, options);
        }
    }

    public sealed class SearchServiceSharedPrivateLinkResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the shared private link resource managed by the Azure Cognitive Search service within the specified resource group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Describes the properties of a Shared Private Link Resource managed by the Azure Cognitive Search service.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.SharedPrivateLinkResourcePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Cognitive Search service associated with the specified resource group.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public Input<string> SearchServiceName { get; set; } = null!;

        public SearchServiceSharedPrivateLinkResourceArgs()
        {
        }
    }
}