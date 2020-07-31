// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180701.Outputs
{

    [OutputType]
    public sealed class AzureFirewallPropertiesFormatResponseResult
    {
        /// <summary>
        /// Collection of application rule collections used by a Azure Firewall.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallApplicationRuleCollectionResponseResult> ApplicationRuleCollections;
        /// <summary>
        /// IP configuration of the Azure Firewall resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallIPConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// Collection of network rule collections used by a Azure Firewall.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallNetworkRuleCollectionResponseResult> NetworkRuleCollections;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string? ProvisioningState;

        [OutputConstructor]
        private AzureFirewallPropertiesFormatResponseResult(
            ImmutableArray<Outputs.AzureFirewallApplicationRuleCollectionResponseResult> applicationRuleCollections,

            ImmutableArray<Outputs.AzureFirewallIPConfigurationResponseResult> ipConfigurations,

            ImmutableArray<Outputs.AzureFirewallNetworkRuleCollectionResponseResult> networkRuleCollections,

            string? provisioningState)
        {
            ApplicationRuleCollections = applicationRuleCollections;
            IpConfigurations = ipConfigurations;
            NetworkRuleCollections = networkRuleCollections;
            ProvisioningState = provisioningState;
        }
    }
}