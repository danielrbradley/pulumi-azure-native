// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20171201.Inputs
{

    /// <summary>
    /// Describes the properties of an Image.
    /// </summary>
    public sealed class ImagePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source virtual machine from which Image is created.
        /// </summary>
        [Input("sourceVirtualMachine")]
        public Input<Inputs.SubResourceArgs>? SourceVirtualMachine { get; set; }

        /// <summary>
        /// Specifies the storage settings for the virtual machine disks.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.ImageStorageProfileArgs>? StorageProfile { get; set; }

        public ImagePropertiesArgs()
        {
        }
    }
}