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
    /// Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
    /// </summary>
    public partial class SearchServicePrivateEndpointConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the private endpoint connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateEndpointConnectionPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SearchServicePrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SearchServicePrivateEndpointConnection(string name, SearchServicePrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:search:SearchServicePrivateEndpointConnection", name, args ?? new SearchServicePrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SearchServicePrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:search:SearchServicePrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SearchServicePrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SearchServicePrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SearchServicePrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class SearchServicePrivateEndpointConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the private endpoint connection to the Azure Cognitive Search service with the specified resource group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.PrivateEndpointConnectionPropertiesArgs>? Properties { get; set; }

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

        public SearchServicePrivateEndpointConnectionArgs()
        {
        }
    }
}