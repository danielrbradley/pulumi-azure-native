// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150615.Inputs
{

    /// <summary>
    /// Azure resource manager sub resource properties.
    /// </summary>
    public sealed class SubResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource Identifier.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public SubResourceArgs()
        {
        }
    }
}