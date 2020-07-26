// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network
{
    public static class GetNetworkWatcherFlowLog
    {
        public static Task<GetNetworkWatcherFlowLogResult> InvokeAsync(GetNetworkWatcherFlowLogArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkWatcherFlowLogResult>("azurerm:network:getNetworkWatcherFlowLog", args ?? new GetNetworkWatcherFlowLogArgs(), options.WithVersion());
    }


    public sealed class GetNetworkWatcherFlowLogArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the flow log resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public string NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkWatcherFlowLogArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkWatcherFlowLogResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the flow log.
        /// </summary>
        public readonly Outputs.FlowLogPropertiesFormatResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNetworkWatcherFlowLogResult(
            string etag,

            string? location,

            string name,

            Outputs.FlowLogPropertiesFormatResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}