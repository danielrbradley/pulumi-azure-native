// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170601.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayBackendAddressPoolPropertiesFormatResponseResult
    {
        /// <summary>
        /// Backend addresses
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayBackendAddressResponseResult> BackendAddresses;
        /// <summary>
        /// Collection of references to IPs defined in network interfaces.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponseResult> BackendIPConfigurations;
        /// <summary>
        /// Provisioning state of the backend address pool resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;

        [OutputConstructor]
        private ApplicationGatewayBackendAddressPoolPropertiesFormatResponseResult(
            ImmutableArray<Outputs.ApplicationGatewayBackendAddressResponseResult> backendAddresses,

            ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponseResult> backendIPConfigurations,

            string? provisioningState)
        {
            BackendAddresses = backendAddresses;
            BackendIPConfigurations = backendIPConfigurations;
            ProvisioningState = provisioningState;
        }
    }
}