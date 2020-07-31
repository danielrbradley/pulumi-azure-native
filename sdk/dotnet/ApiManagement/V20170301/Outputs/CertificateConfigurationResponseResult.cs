// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20170301.Outputs
{

    [OutputType]
    public sealed class CertificateConfigurationResponseResult
    {
        /// <summary>
        /// Certificate information.
        /// </summary>
        public readonly Outputs.CertificateInformationResponseResult Certificate;
        /// <summary>
        /// Certificate Password.
        /// </summary>
        public readonly string? CertificatePassword;
        /// <summary>
        /// Base64 Encoded certificate.
        /// </summary>
        public readonly string? EncodedCertificate;
        /// <summary>
        /// The local certificate store location. Only Root and CertificateAuthority are valid locations.
        /// </summary>
        public readonly string StoreName;

        [OutputConstructor]
        private CertificateConfigurationResponseResult(
            Outputs.CertificateInformationResponseResult certificate,

            string? certificatePassword,

            string? encodedCertificate,

            string storeName)
        {
            Certificate = certificate;
            CertificatePassword = certificatePassword;
            EncodedCertificate = encodedCertificate;
            StoreName = storeName;
        }
    }
}