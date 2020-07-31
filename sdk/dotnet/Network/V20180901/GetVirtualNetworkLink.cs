// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180901
{
    public static class GetVirtualNetworkLink
    {
        public static Task<GetVirtualNetworkLinkResult> InvokeAsync(GetVirtualNetworkLinkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkLinkResult>("azurerm:network/v20180901:getVirtualNetworkLink", args ?? new GetVirtualNetworkLinkArgs(), options.WithVersion());
    }


    public sealed class GetVirtualNetworkLinkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the virtual network link.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Private DNS zone (without a terminating dot).
        /// </summary>
        [Input("privateZoneName", required: true)]
        public string PrivateZoneName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkLinkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualNetworkLinkResult
    {
        /// <summary>
        /// The ETag of the virtual network link.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the virtual network link to the Private DNS zone.
        /// </summary>
        public readonly Outputs.VirtualNetworkLinkPropertiesResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVirtualNetworkLinkResult(
            string? etag,

            string? location,

            string name,

            Outputs.VirtualNetworkLinkPropertiesResponseResult properties,

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