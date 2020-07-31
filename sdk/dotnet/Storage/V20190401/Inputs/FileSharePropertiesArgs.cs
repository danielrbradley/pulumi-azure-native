// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20190401.Inputs
{

    /// <summary>
    /// The properties of the file share.
    /// </summary>
    public sealed class FileSharePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A name-value pair to associate with the share as metadata.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120).
        /// </summary>
        [Input("shareQuota")]
        public Input<int>? ShareQuota { get; set; }

        public FileSharePropertiesArgs()
        {
        }
    }
}