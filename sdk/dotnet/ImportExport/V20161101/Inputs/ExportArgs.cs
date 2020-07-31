// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ImportExport.V20161101.Inputs
{

    /// <summary>
    /// A property containing information about the blobs to be exported for an export job. This property is required for export jobs, but must not be specified for import jobs.
    /// </summary>
    public sealed class ExportArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A list of the blobs to be exported.
        /// </summary>
        [Input("blobList")]
        public Input<Inputs.ExportPropertiesArgs>? BlobList { get; set; }

        /// <summary>
        /// The relative URI to the block blob that contains the list of blob paths or blob path prefixes as defined above, beginning with the container name. If the blob is in root container, the URI must begin with $root. 
        /// </summary>
        [Input("blobListblobPath")]
        public Input<string>? BlobListblobPath { get; set; }

        public ExportArgs()
        {
        }
    }
}