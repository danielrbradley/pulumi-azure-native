// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20180501.Inputs
{

    /// <summary>
    /// Azure storage account credentials.
    /// </summary>
    public sealed class AzureStorageCredentialsInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Storage account key. One of accountKey or accountKeySecretReference must be specified.
        /// </summary>
        [Input("accountKey")]
        public Input<string>? AccountKey { get; set; }

        /// <summary>
        /// Information about KeyVault secret storing the storage account key. One of accountKey or accountKeySecretReference must be specified.
        /// </summary>
        [Input("accountKeySecretReference")]
        public Input<Inputs.KeyVaultSecretReferenceArgs>? AccountKeySecretReference { get; set; }

        public AzureStorageCredentialsInfoArgs()
        {
        }
    }
}