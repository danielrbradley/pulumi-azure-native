// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180401.Outputs
{

    [OutputType]
    public sealed class DiskPropertiesResponseResult
    {
        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        public readonly Outputs.CreationDataResponseResult CreationData;
        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        public readonly int? DiskSizeGB;
        /// <summary>
        /// Encryption settings for disk or snapshot
        /// </summary>
        public readonly Outputs.EncryptionSettingsResponseResult? EncryptionSettings;
        /// <summary>
        /// The Operating System type.
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// The disk provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The time when the disk was created.
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private DiskPropertiesResponseResult(
            Outputs.CreationDataResponseResult creationData,

            int? diskSizeGB,

            Outputs.EncryptionSettingsResponseResult? encryptionSettings,

            string? osType,

            string provisioningState,

            string timeCreated)
        {
            CreationData = creationData;
            DiskSizeGB = diskSizeGB;
            EncryptionSettings = encryptionSettings;
            OsType = osType;
            ProvisioningState = provisioningState;
            TimeCreated = timeCreated;
        }
    }
}