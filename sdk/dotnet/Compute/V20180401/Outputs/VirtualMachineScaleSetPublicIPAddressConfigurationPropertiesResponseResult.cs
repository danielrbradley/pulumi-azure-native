// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180401.Outputs
{

    [OutputType]
    public sealed class VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesResponseResult
    {
        /// <summary>
        /// The dns settings to be applied on the publicIP addresses .
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseResult? DnsSettings;
        /// <summary>
        /// The idle timeout of the public IP address.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// The list of IP tags associated with the public IP address.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineScaleSetIpTagResponseResult> IpTags;

        [OutputConstructor]
        private VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesResponseResult(
            Outputs.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseResult? dnsSettings,

            int? idleTimeoutInMinutes,

            ImmutableArray<Outputs.VirtualMachineScaleSetIpTagResponseResult> ipTags)
        {
            DnsSettings = dnsSettings;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpTags = ipTags;
        }
    }
}