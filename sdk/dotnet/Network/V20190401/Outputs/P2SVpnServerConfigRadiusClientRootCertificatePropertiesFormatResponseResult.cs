// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401.Outputs
{

    [OutputType]
    public sealed class P2SVpnServerConfigRadiusClientRootCertificatePropertiesFormatResponseResult
    {
        /// <summary>
        /// The provisioning state of the Radius client root certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The Radius client root certificate thumbprint.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private P2SVpnServerConfigRadiusClientRootCertificatePropertiesFormatResponseResult(
            string provisioningState,

            string? thumbprint)
        {
            ProvisioningState = provisioningState;
            Thumbprint = thumbprint;
        }
    }
}