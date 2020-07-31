// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Inputs
{

    /// <summary>
    /// Parameters for VpnConnection.
    /// </summary>
    public sealed class VpnSiteLinkConnectionPropertiesResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expected bandwidth in MBPS.
        /// </summary>
        [Input("connectionBandwidth")]
        public int? ConnectionBandwidth { get; set; }

        /// <summary>
        /// The connection status.
        /// </summary>
        [Input("connectionStatus")]
        public string? ConnectionStatus { get; set; }

        /// <summary>
        /// Egress bytes transferred.
        /// </summary>
        [Input("egressBytesTransferred", required: true)]
        public int EgressBytesTransferred { get; set; }

        /// <summary>
        /// EnableBgp flag.
        /// </summary>
        [Input("enableBgp")]
        public bool? EnableBgp { get; set; }

        /// <summary>
        /// EnableBgp flag.
        /// </summary>
        [Input("enableRateLimiting")]
        public bool? EnableRateLimiting { get; set; }

        /// <summary>
        /// Ingress bytes transferred.
        /// </summary>
        [Input("ingressBytesTransferred", required: true)]
        public int IngressBytesTransferred { get; set; }

        [Input("ipsecPolicies")]
        private List<Inputs.IpsecPolicyResponseArgs>? _ipsecPolicies;

        /// <summary>
        /// The IPSec Policies to be considered by this connection.
        /// </summary>
        public List<Inputs.IpsecPolicyResponseArgs> IpsecPolicies
        {
            get => _ipsecPolicies ?? (_ipsecPolicies = new List<Inputs.IpsecPolicyResponseArgs>());
            set => _ipsecPolicies = value;
        }

        /// <summary>
        /// The provisioning state of the VPN site link connection resource.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        /// <summary>
        /// Routing weight for vpn connection.
        /// </summary>
        [Input("routingWeight")]
        public int? RoutingWeight { get; set; }

        /// <summary>
        /// SharedKey for the vpn connection.
        /// </summary>
        [Input("sharedKey")]
        public string? SharedKey { get; set; }

        /// <summary>
        /// Use local azure ip to initiate connection.
        /// </summary>
        [Input("useLocalAzureIpAddress")]
        public bool? UseLocalAzureIpAddress { get; set; }

        /// <summary>
        /// Enable policy-based traffic selectors.
        /// </summary>
        [Input("usePolicyBasedTrafficSelectors")]
        public bool? UsePolicyBasedTrafficSelectors { get; set; }

        /// <summary>
        /// Connection protocol used for this connection.
        /// </summary>
        [Input("vpnConnectionProtocolType")]
        public string? VpnConnectionProtocolType { get; set; }

        /// <summary>
        /// Id of the connected vpn site link.
        /// </summary>
        [Input("vpnSiteLink")]
        public Inputs.SubResourceResponseArgs? VpnSiteLink { get; set; }

        public VpnSiteLinkConnectionPropertiesResponseArgs()
        {
        }
    }
}