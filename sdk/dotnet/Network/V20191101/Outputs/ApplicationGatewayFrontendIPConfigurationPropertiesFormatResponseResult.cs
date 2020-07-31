// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayFrontendIPConfigurationPropertiesFormatResponseResult
    {
        /// <summary>
        /// PrivateIPAddress of the network interface IP Configuration.
        /// </summary>
        public readonly string? PrivateIPAddress;
        /// <summary>
        /// The private IP address allocation method.
        /// </summary>
        public readonly string? PrivateIPAllocationMethod;
        /// <summary>
        /// The provisioning state of the frontend IP configuration resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Reference to the PublicIP resource.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? PublicIPAddress;
        /// <summary>
        /// Reference to the subnet resource.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? Subnet;

        [OutputConstructor]
        private ApplicationGatewayFrontendIPConfigurationPropertiesFormatResponseResult(
            string? privateIPAddress,

            string? privateIPAllocationMethod,

            string provisioningState,

            Outputs.SubResourceResponseResult? publicIPAddress,

            Outputs.SubResourceResponseResult? subnet)
        {
            PrivateIPAddress = privateIPAddress;
            PrivateIPAllocationMethod = privateIPAllocationMethod;
            ProvisioningState = provisioningState;
            PublicIPAddress = publicIPAddress;
            Subnet = subnet;
        }
    }
}