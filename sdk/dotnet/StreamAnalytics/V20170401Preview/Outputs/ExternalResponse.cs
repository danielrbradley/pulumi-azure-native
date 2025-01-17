// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StreamAnalytics.V20170401Preview.Outputs
{

    /// <summary>
    /// The storage account where the custom code artifacts are located.
    /// </summary>
    [OutputType]
    public sealed class ExternalResponse
    {
        public readonly string? Container;
        public readonly string? Path;
        /// <summary>
        /// The properties that are associated with an Azure Storage account
        /// </summary>
        public readonly Outputs.StorageAccountResponse? StorageAccount;

        [OutputConstructor]
        private ExternalResponse(
            string? container,

            string? path,

            Outputs.StorageAccountResponse? storageAccount)
        {
            Container = container;
            Path = path;
            StorageAccount = storageAccount;
        }
    }
}
