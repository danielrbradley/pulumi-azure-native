// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.Outputs
{

    [OutputType]
    public sealed class SnapshotPropertiesResponseResult
    {
        /// <summary>
        /// The creation date of the snapshot
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// UUID v4 used to identify the FileSystem
        /// </summary>
        public readonly string? FileSystemId;
        /// <summary>
        /// Azure lifecycle management
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// UUID v4 used to identify the Snapshot
        /// </summary>
        public readonly string SnapshotId;

        [OutputConstructor]
        private SnapshotPropertiesResponseResult(
            string created,

            string? fileSystemId,

            string provisioningState,

            string snapshotId)
        {
            Created = created;
            FileSystemId = fileSystemId;
            ProvisioningState = provisioningState;
            SnapshotId = snapshotId;
        }
    }
}