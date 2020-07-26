// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayPrivateLinkConfigurationPropertiesResponseResult
    {
        /// <summary>
        /// An array of application gateway private link ip configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayPrivateLinkIpConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// The provisioning state of the application gateway private link configuration.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private ApplicationGatewayPrivateLinkConfigurationPropertiesResponseResult(
            ImmutableArray<Outputs.ApplicationGatewayPrivateLinkIpConfigurationResponseResult> ipConfigurations,

            string provisioningState)
        {
            IpConfigurations = ipConfigurations;
            ProvisioningState = provisioningState;
        }
    }
}