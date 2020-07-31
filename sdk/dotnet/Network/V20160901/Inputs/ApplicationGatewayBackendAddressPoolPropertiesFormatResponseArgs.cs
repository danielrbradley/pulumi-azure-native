// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160901.Inputs
{

    /// <summary>
    /// Properties of Backend Address Pool of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayBackendAddressPoolPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        [Input("backendAddresses")]
        private List<Inputs.ApplicationGatewayBackendAddressResponseArgs>? _backendAddresses;

        /// <summary>
        /// Backend addresses
        /// </summary>
        public List<Inputs.ApplicationGatewayBackendAddressResponseArgs> BackendAddresses
        {
            get => _backendAddresses ?? (_backendAddresses = new List<Inputs.ApplicationGatewayBackendAddressResponseArgs>());
            set => _backendAddresses = value;
        }

        [Input("backendIPConfigurations")]
        private List<Inputs.NetworkInterfaceIPConfigurationResponseArgs>? _backendIPConfigurations;

        /// <summary>
        /// Collection of references to IPs defined in network interfaces.
        /// </summary>
        public List<Inputs.NetworkInterfaceIPConfigurationResponseArgs> BackendIPConfigurations
        {
            get => _backendIPConfigurations ?? (_backendIPConfigurations = new List<Inputs.NetworkInterfaceIPConfigurationResponseArgs>());
            set => _backendIPConfigurations = value;
        }

        /// <summary>
        /// Provisioning state of the backend address pool resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        public ApplicationGatewayBackendAddressPoolPropertiesFormatResponseArgs()
        {
        }
    }
}