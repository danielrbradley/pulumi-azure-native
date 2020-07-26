// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Inputs
{

    /// <summary>
    /// Gateway hostname configuration details.
    /// </summary>
    public sealed class GatewayHostnameConfigurationContractPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of Certificate entity that will be used for TLS connection establishment
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Hostname value. Supports valid domain name, partial or full wildcard
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Determines whether gateway requests client certificate
        /// </summary>
        [Input("negotiateClientCertificate")]
        public Input<bool>? NegotiateClientCertificate { get; set; }

        public GatewayHostnameConfigurationContractPropertiesArgs()
        {
        }
    }
}