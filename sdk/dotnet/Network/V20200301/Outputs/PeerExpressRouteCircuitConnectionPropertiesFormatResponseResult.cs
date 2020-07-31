// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200301.Outputs
{

    [OutputType]
    public sealed class PeerExpressRouteCircuitConnectionPropertiesFormatResponseResult
    {
        /// <summary>
        /// /29 IP address space to carve out Customer addresses for tunnels.
        /// </summary>
        public readonly string? AddressPrefix;
        /// <summary>
        /// The resource guid of the authorization used for the express route circuit connection.
        /// </summary>
        public readonly string? AuthResourceGuid;
        /// <summary>
        /// Express Route Circuit connection state.
        /// </summary>
        public readonly string? CircuitConnectionStatus;
        /// <summary>
        /// The name of the express route circuit connection resource.
        /// </summary>
        public readonly string? ConnectionName;
        /// <summary>
        /// Reference to Express Route Circuit Private Peering Resource of the circuit.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? ExpressRouteCircuitPeering;
        /// <summary>
        /// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? PeerExpressRouteCircuitPeering;
        /// <summary>
        /// The provisioning state of the peer express route circuit connection resource.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private PeerExpressRouteCircuitConnectionPropertiesFormatResponseResult(
            string? addressPrefix,

            string? authResourceGuid,

            string? circuitConnectionStatus,

            string? connectionName,

            Outputs.SubResourceResponseResult? expressRouteCircuitPeering,

            Outputs.SubResourceResponseResult? peerExpressRouteCircuitPeering,

            string provisioningState)
        {
            AddressPrefix = addressPrefix;
            AuthResourceGuid = authResourceGuid;
            CircuitConnectionStatus = circuitConnectionStatus;
            ConnectionName = connectionName;
            ExpressRouteCircuitPeering = expressRouteCircuitPeering;
            PeerExpressRouteCircuitPeering = peerExpressRouteCircuitPeering;
            ProvisioningState = provisioningState;
        }
    }
}