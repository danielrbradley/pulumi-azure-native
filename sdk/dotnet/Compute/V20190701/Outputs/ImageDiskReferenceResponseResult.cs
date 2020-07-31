// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190701.Outputs
{

    [OutputType]
    public sealed class ImageDiskReferenceResponseResult
    {
        /// <summary>
        /// A relative uri containing either a Platform Image Repository or user image reference.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.
        /// </summary>
        public readonly int? Lun;

        [OutputConstructor]
        private ImageDiskReferenceResponseResult(
            string id,

            int? lun)
        {
            Id = id;
            Lun = lun;
        }
    }
}