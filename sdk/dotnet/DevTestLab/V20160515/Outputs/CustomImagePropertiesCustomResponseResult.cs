// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20160515.Outputs
{

    [OutputType]
    public sealed class CustomImagePropertiesCustomResponseResult
    {
        /// <summary>
        /// The image name.
        /// </summary>
        public readonly string? ImageName;
        /// <summary>
        /// The OS type of the custom image (i.e. Windows, Linux)
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// Indicates whether sysprep has been run on the VHD.
        /// </summary>
        public readonly bool? SysPrep;

        [OutputConstructor]
        private CustomImagePropertiesCustomResponseResult(
            string? imageName,

            string osType,

            bool? sysPrep)
        {
            ImageName = imageName;
            OsType = osType;
            SysPrep = sysPrep;
        }
    }
}