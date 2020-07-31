// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20200501.Outputs
{

    [OutputType]
    public sealed class KeyVaultPropertiesResponseResult
    {
        /// <summary>
        /// The current key used to encrypt the Media Services account, including the key version.
        /// </summary>
        public readonly string CurrentKeyIdentifier;
        /// <summary>
        /// The URL of the Key Vault key used to encrypt the account. The key may either be versioned (for example https://vault/keys/mykey/version1) or reference a key without a version (for example https://vault/keys/mykey).
        /// </summary>
        public readonly string? KeyIdentifier;

        [OutputConstructor]
        private KeyVaultPropertiesResponseResult(
            string currentKeyIdentifier,

            string? keyIdentifier)
        {
            CurrentKeyIdentifier = currentKeyIdentifier;
            KeyIdentifier = keyIdentifier;
        }
    }
}