// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601.Inputs
{

    /// <summary>
    /// Properties of Authentication certificates of application gateway
    /// </summary>
    public sealed class ApplicationGatewayAuthenticationCertificatePropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate public data 
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Provisioning state of the authentication certificate resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public ApplicationGatewayAuthenticationCertificatePropertiesFormatArgs()
        {
        }
    }
}