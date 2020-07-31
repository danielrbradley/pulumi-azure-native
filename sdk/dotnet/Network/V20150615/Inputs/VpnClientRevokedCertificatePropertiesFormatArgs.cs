// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150615.Inputs
{

    /// <summary>
    /// Properties of the revoked VPN client certificate of virtual network gateway.
    /// </summary>
    public sealed class VpnClientRevokedCertificatePropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The provisioning state of the VPN client revoked certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The revoked VPN client certificate thumbprint.
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        public VpnClientRevokedCertificatePropertiesFormatArgs()
        {
        }
    }
}