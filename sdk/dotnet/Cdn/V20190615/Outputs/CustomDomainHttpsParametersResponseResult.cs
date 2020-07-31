// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20190615.Outputs
{

    [OutputType]
    public sealed class CustomDomainHttpsParametersResponseResult
    {
        /// <summary>
        /// Defines the source of the SSL certificate.
        /// </summary>
        public readonly string CertificateSource;
        /// <summary>
        /// TLS protocol version that will be used for Https
        /// </summary>
        public readonly string? MinimumTlsVersion;
        /// <summary>
        /// Defines the TLS extension protocol that is used for secure delivery.
        /// </summary>
        public readonly string ProtocolType;

        [OutputConstructor]
        private CustomDomainHttpsParametersResponseResult(
            string certificateSource,

            string? minimumTlsVersion,

            string protocolType)
        {
            CertificateSource = certificateSource;
            MinimumTlsVersion = minimumTlsVersion;
            ProtocolType = protocolType;
        }
    }
}