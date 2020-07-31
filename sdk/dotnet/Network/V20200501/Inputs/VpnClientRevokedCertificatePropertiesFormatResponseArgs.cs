// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Inputs
{

    /// <summary>
    /// Properties of the revoked VPN client certificate of virtual network gateway.
    /// </summary>
    public sealed class VpnClientRevokedCertificatePropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The provisioning state of the VPN client revoked certificate resource.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        /// <summary>
        /// The revoked VPN client certificate thumbprint.
        /// </summary>
        [Input("thumbprint")]
        public string? Thumbprint { get; set; }

        public VpnClientRevokedCertificatePropertiesFormatResponseArgs()
        {
        }
    }
}