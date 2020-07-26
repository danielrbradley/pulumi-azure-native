// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.Outputs
{

    [OutputType]
    public sealed class NetworkInterfacePropertiesResponseResult
    {
        /// <summary>
        /// The DNS name.
        /// </summary>
        public readonly string? DnsName;
        /// <summary>
        /// The private IP address.
        /// </summary>
        public readonly string? PrivateIpAddress;
        /// <summary>
        /// The public IP address.
        /// </summary>
        public readonly string? PublicIpAddress;
        /// <summary>
        /// The resource ID of the public IP address.
        /// </summary>
        public readonly string? PublicIpAddressId;
        /// <summary>
        /// The RdpAuthority property is a server DNS host name or IP address followed by the service port number for RDP (Remote Desktop Protocol).
        /// </summary>
        public readonly string? RdpAuthority;
        /// <summary>
        /// The configuration for sharing a public IP address across multiple virtual machines.
        /// </summary>
        public readonly Outputs.SharedPublicIpAddressConfigurationResponseResult? SharedPublicIpAddressConfiguration;
        /// <summary>
        /// The SshAuthority property is a server DNS host name or IP address followed by the service port number for SSH.
        /// </summary>
        public readonly string? SshAuthority;
        /// <summary>
        /// The resource ID of the sub net.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The resource ID of the virtual network.
        /// </summary>
        public readonly string? VirtualNetworkId;

        [OutputConstructor]
        private NetworkInterfacePropertiesResponseResult(
            string? dnsName,

            string? privateIpAddress,

            string? publicIpAddress,

            string? publicIpAddressId,

            string? rdpAuthority,

            Outputs.SharedPublicIpAddressConfigurationResponseResult? sharedPublicIpAddressConfiguration,

            string? sshAuthority,

            string? subnetId,

            string? virtualNetworkId)
        {
            DnsName = dnsName;
            PrivateIpAddress = privateIpAddress;
            PublicIpAddress = publicIpAddress;
            PublicIpAddressId = publicIpAddressId;
            RdpAuthority = rdpAuthority;
            SharedPublicIpAddressConfiguration = sharedPublicIpAddressConfiguration;
            SshAuthority = sshAuthority;
            SubnetId = subnetId;
            VirtualNetworkId = virtualNetworkId;
        }
    }
}