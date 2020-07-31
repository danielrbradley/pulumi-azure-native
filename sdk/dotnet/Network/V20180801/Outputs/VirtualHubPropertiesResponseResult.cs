// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801.Outputs
{

    [OutputType]
    public sealed class VirtualHubPropertiesResponseResult
    {
        /// <summary>
        /// Address-prefix for this VirtualHub.
        /// </summary>
        public readonly string? AddressPrefix;
        /// <summary>
        /// The expressRouteGateway associated with this VirtualHub
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? ExpressRouteGateway;
        /// <summary>
        /// The P2SVpnGateway associated with this VirtualHub
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? P2SVpnGateway;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The routeTable associated with this virtual hub.
        /// </summary>
        public readonly Outputs.VirtualHubRouteTableResponseResult? RouteTable;
        /// <summary>
        /// list of all vnet connections with this VirtualHub.
        /// </summary>
        public readonly ImmutableArray<Outputs.HubVirtualNetworkConnectionResponseResult> VirtualNetworkConnections;
        /// <summary>
        /// The VirtualWAN to which the VirtualHub belongs
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? VirtualWan;
        /// <summary>
        /// The VpnGateway associated with this VirtualHub
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? VpnGateway;

        [OutputConstructor]
        private VirtualHubPropertiesResponseResult(
            string? addressPrefix,

            Outputs.SubResourceResponseResult? expressRouteGateway,

            Outputs.SubResourceResponseResult? p2SVpnGateway,

            string? provisioningState,

            Outputs.VirtualHubRouteTableResponseResult? routeTable,

            ImmutableArray<Outputs.HubVirtualNetworkConnectionResponseResult> virtualNetworkConnections,

            Outputs.SubResourceResponseResult? virtualWan,

            Outputs.SubResourceResponseResult? vpnGateway)
        {
            AddressPrefix = addressPrefix;
            ExpressRouteGateway = expressRouteGateway;
            P2SVpnGateway = p2SVpnGateway;
            ProvisioningState = provisioningState;
            RouteTable = routeTable;
            VirtualNetworkConnections = virtualNetworkConnections;
            VirtualWan = virtualWan;
            VpnGateway = vpnGateway;
        }
    }
}