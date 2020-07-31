// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20191109
{
    public static class GetCluster
    {
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azurerm:kusto/v20191109:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The identity of the cluster, if configured.
        /// </summary>
        public readonly Outputs.IdentityResponseResult? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The cluster properties.
        /// </summary>
        public readonly Outputs.ClusterPropertiesResponseResult Properties;
        /// <summary>
        /// The SKU of the cluster.
        /// </summary>
        public readonly Outputs.AzureSkuResponseResult Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The availability zones of the cluster.
        /// </summary>
        public readonly Outputs.ZonesResponseResult? Zones;

        [OutputConstructor]
        private GetClusterResult(
            Outputs.IdentityResponseResult? identity,

            string location,

            string name,

            Outputs.ClusterPropertiesResponseResult properties,

            Outputs.AzureSkuResponseResult sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.ZonesResponseResult? zones)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
            Zones = zones;
        }
    }
}