// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601.Inputs
{

    /// <summary>
    /// NetworkSecurityGroup resource
    /// </summary>
    public sealed class NetworkSecurityGroupResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public string? Etag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public string? Location { get; set; }

        /// <summary>
        /// Resource name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Network Security Group resource
        /// </summary>
        [Input("properties")]
        public Inputs.NetworkSecurityGroupPropertiesFormatResponseArgs? Properties { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public NetworkSecurityGroupResponseArgs()
        {
        }
    }
}