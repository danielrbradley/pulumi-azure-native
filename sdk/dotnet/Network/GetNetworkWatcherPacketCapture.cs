// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network
{
    public static class GetNetworkWatcherPacketCapture
    {
        public static Task<GetNetworkWatcherPacketCaptureResult> InvokeAsync(GetNetworkWatcherPacketCaptureArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkWatcherPacketCaptureResult>("azurerm:network:getNetworkWatcherPacketCapture", args ?? new GetNetworkWatcherPacketCaptureArgs(), options.WithVersion());
    }


    public sealed class GetNetworkWatcherPacketCaptureArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the packet capture session.
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

        public GetNetworkWatcherPacketCaptureArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkWatcherPacketCaptureResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Name of the packet capture session.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the packet capture result.
        /// </summary>
        public readonly Outputs.PacketCaptureResultPropertiesResponseResult Properties;

        [OutputConstructor]
        private GetNetworkWatcherPacketCaptureResult(
            string etag,

            string name,

            Outputs.PacketCaptureResultPropertiesResponseResult properties)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
        }
    }
}