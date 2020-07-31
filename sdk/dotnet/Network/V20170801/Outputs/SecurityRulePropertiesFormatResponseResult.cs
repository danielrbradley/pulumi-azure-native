// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Outputs
{

    [OutputType]
    public sealed class SecurityRulePropertiesFormatResponseResult
    {
        /// <summary>
        /// The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
        /// </summary>
        public readonly string Access;
        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
        /// </summary>
        public readonly string? DestinationAddressPrefix;
        /// <summary>
        /// The destination address prefixes. CIDR or destination IP ranges.
        /// </summary>
        public readonly ImmutableArray<string> DestinationAddressPrefixes;
        /// <summary>
        /// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
        /// </summary>
        public readonly string? DestinationPortRange;
        /// <summary>
        /// The destination port ranges.
        /// </summary>
        public readonly ImmutableArray<string> DestinationPortRanges;
        /// <summary>
        /// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
        /// </summary>
        public readonly string? SourceAddressPrefix;
        /// <summary>
        /// The CIDR or source IP ranges.
        /// </summary>
        public readonly ImmutableArray<string> SourceAddressPrefixes;
        /// <summary>
        /// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
        /// </summary>
        public readonly string? SourcePortRange;
        /// <summary>
        /// The source port ranges.
        /// </summary>
        public readonly ImmutableArray<string> SourcePortRanges;

        [OutputConstructor]
        private SecurityRulePropertiesFormatResponseResult(
            string access,

            string? description,

            string? destinationAddressPrefix,

            ImmutableArray<string> destinationAddressPrefixes,

            string? destinationPortRange,

            ImmutableArray<string> destinationPortRanges,

            string direction,

            int? priority,

            string protocol,

            string? provisioningState,

            string? sourceAddressPrefix,

            ImmutableArray<string> sourceAddressPrefixes,

            string? sourcePortRange,

            ImmutableArray<string> sourcePortRanges)
        {
            Access = access;
            Description = description;
            DestinationAddressPrefix = destinationAddressPrefix;
            DestinationAddressPrefixes = destinationAddressPrefixes;
            DestinationPortRange = destinationPortRange;
            DestinationPortRanges = destinationPortRanges;
            Direction = direction;
            Priority = priority;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            SourceAddressPrefix = sourceAddressPrefix;
            SourceAddressPrefixes = sourceAddressPrefixes;
            SourcePortRange = sourcePortRange;
            SourcePortRanges = sourcePortRanges;
        }
    }
}