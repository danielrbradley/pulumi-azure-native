// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.Outputs
{

    [OutputType]
    public sealed class EncryptionPropertiesResponseResult
    {
        /// <summary>
        /// Type of the key source.
        /// </summary>
        public readonly string? KeySource;
        /// <summary>
        /// Additional details when using Microsoft.KeyVault
        /// </summary>
        public readonly Outputs.KeyVaultPropertiesResponseResult? KeyVaultProperties;

        [OutputConstructor]
        private EncryptionPropertiesResponseResult(
            string? keySource,

            Outputs.KeyVaultPropertiesResponseResult? keyVaultProperties)
        {
            KeySource = keySource;
            KeyVaultProperties = keyVaultProperties;
        }
    }
}