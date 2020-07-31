// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190201.Outputs
{

    [OutputType]
    public sealed class NetworkInterfacePropertiesFormatResponseResult
    {
        /// <summary>
        /// The DNS settings in network interface.
        /// </summary>
        public readonly Outputs.NetworkInterfaceDnsSettingsResponseResult? DnsSettings;
        /// <summary>
        /// If the network interface is accelerated networking enabled.
        /// </summary>
        public readonly bool? EnableAcceleratedNetworking;
        /// <summary>
        /// Indicates whether IP forwarding is enabled on this network interface.
        /// </summary>
        public readonly bool? EnableIPForwarding;
        /// <summary>
        /// A list of references to linked BareMetal resources
        /// </summary>
        public readonly ImmutableArray<string> HostedWorkloads;
        /// <summary>
        /// A reference to the interface endpoint to which the network interface is linked.
        /// </summary>
        public readonly Outputs.InterfaceEndpointResponseResult InterfaceEndpoint;
        /// <summary>
        /// A list of IPConfigurations of the network interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        public readonly Outputs.NetworkSecurityGroupResponseResult? NetworkSecurityGroup;
        /// <summary>
        /// Gets whether this is a primary network interface on a virtual machine.
        /// </summary>
        public readonly bool? Primary;
        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The resource GUID property of the network interface resource.
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// A list of TapConfigurations of the network interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceTapConfigurationResponseResult> TapConfigurations;
        /// <summary>
        /// The reference of a virtual machine.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult VirtualMachine;

        [OutputConstructor]
        private NetworkInterfacePropertiesFormatResponseResult(
            Outputs.NetworkInterfaceDnsSettingsResponseResult? dnsSettings,

            bool? enableAcceleratedNetworking,

            bool? enableIPForwarding,

            ImmutableArray<string> hostedWorkloads,

            Outputs.InterfaceEndpointResponseResult interfaceEndpoint,

            ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponseResult> ipConfigurations,

            string? macAddress,

            Outputs.NetworkSecurityGroupResponseResult? networkSecurityGroup,

            bool? primary,

            string? provisioningState,

            string? resourceGuid,

            ImmutableArray<Outputs.NetworkInterfaceTapConfigurationResponseResult> tapConfigurations,

            Outputs.SubResourceResponseResult virtualMachine)
        {
            DnsSettings = dnsSettings;
            EnableAcceleratedNetworking = enableAcceleratedNetworking;
            EnableIPForwarding = enableIPForwarding;
            HostedWorkloads = hostedWorkloads;
            InterfaceEndpoint = interfaceEndpoint;
            IpConfigurations = ipConfigurations;
            MacAddress = macAddress;
            NetworkSecurityGroup = networkSecurityGroup;
            Primary = primary;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            TapConfigurations = tapConfigurations;
            VirtualMachine = virtualMachine;
        }
    }
}