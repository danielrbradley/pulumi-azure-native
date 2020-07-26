// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search
{
    public static class GetSearchServicePrivateEndpointConnection
    {
        public static Task<GetSearchServicePrivateEndpointConnectionResult> InvokeAsync(GetSearchServicePrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSearchServicePrivateEndpointConnectionResult>("azurerm:search:getSearchServicePrivateEndpointConnection", args ?? new GetSearchServicePrivateEndpointConnectionArgs(), options.WithVersion());
    }


    public sealed class GetSearchServicePrivateEndpointConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the private endpoint connection to the Azure Cognitive Search service with the specified resource group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Cognitive Search service associated with the specified resource group.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public string SearchServiceName { get; set; } = null!;

        public GetSearchServicePrivateEndpointConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSearchServicePrivateEndpointConnectionResult
    {
        /// <summary>
        /// The name of the private endpoint connection.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
        /// </summary>
        public readonly Outputs.PrivateEndpointConnectionPropertiesResponseResult Properties;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSearchServicePrivateEndpointConnectionResult(
            string name,

            Outputs.PrivateEndpointConnectionPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}