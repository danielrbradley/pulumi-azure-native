// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20200601.Outputs
{

    [OutputType]
    public sealed class ExportDeliveryDestinationResponseResult
    {
        /// <summary>
        /// The name of the container where exports will be uploaded.
        /// </summary>
        public readonly string Container;
        /// <summary>
        /// The resource id of the storage account where exports will be delivered.
        /// </summary>
        public readonly string ResourceId;
        /// <summary>
        /// The name of the directory where exports will be uploaded.
        /// </summary>
        public readonly string? RootFolderPath;

        [OutputConstructor]
        private ExportDeliveryDestinationResponseResult(
            string container,

            string resourceId,

            string? rootFolderPath)
        {
            Container = container;
            ResourceId = resourceId;
            RootFolderPath = rootFolderPath;
        }
    }
}