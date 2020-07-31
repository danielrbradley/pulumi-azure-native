// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Outputs
{

    [OutputType]
    public sealed class AzureFirewallPropertiesFormatResponseResult
    {
        /// <summary>
        /// Collection of application rule collections used by Azure Firewall.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallApplicationRuleCollectionResponseResult> ApplicationRuleCollections;
        /// <summary>
        /// The firewallPolicy associated with this azure firewall.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? FirewallPolicy;
        /// <summary>
        /// IP addresses associated with AzureFirewall.
        /// </summary>
        public readonly Outputs.HubIPAddressesResponseResult HubIpAddresses;
        /// <summary>
        /// IP configuration of the Azure Firewall resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallIPConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// Collection of NAT rule collections used by Azure Firewall.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallNatRuleCollectionResponseResult> NatRuleCollections;
        /// <summary>
        /// Collection of network rule collections used by Azure Firewall.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallNetworkRuleCollectionResponseResult> NetworkRuleCollections;
        /// <summary>
        /// The provisioning state of the Azure firewall resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The operation mode for Threat Intelligence.
        /// </summary>
        public readonly string? ThreatIntelMode;
        /// <summary>
        /// The virtualHub to which the firewall belongs.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? VirtualHub;

        [OutputConstructor]
        private AzureFirewallPropertiesFormatResponseResult(
            ImmutableArray<Outputs.AzureFirewallApplicationRuleCollectionResponseResult> applicationRuleCollections,

            Outputs.SubResourceResponseResult? firewallPolicy,

            Outputs.HubIPAddressesResponseResult hubIpAddresses,

            ImmutableArray<Outputs.AzureFirewallIPConfigurationResponseResult> ipConfigurations,

            ImmutableArray<Outputs.AzureFirewallNatRuleCollectionResponseResult> natRuleCollections,

            ImmutableArray<Outputs.AzureFirewallNetworkRuleCollectionResponseResult> networkRuleCollections,

            string? provisioningState,

            string? threatIntelMode,

            Outputs.SubResourceResponseResult? virtualHub)
        {
            ApplicationRuleCollections = applicationRuleCollections;
            FirewallPolicy = firewallPolicy;
            HubIpAddresses = hubIpAddresses;
            IpConfigurations = ipConfigurations;
            NatRuleCollections = natRuleCollections;
            NetworkRuleCollections = networkRuleCollections;
            ProvisioningState = provisioningState;
            ThreatIntelMode = threatIntelMode;
            VirtualHub = virtualHub;
        }
    }
}