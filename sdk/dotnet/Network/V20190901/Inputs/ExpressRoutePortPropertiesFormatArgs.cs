// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901.Inputs
{

    /// <summary>
    /// Properties specific to ExpressRoutePort resources.
    /// </summary>
    public sealed class ExpressRoutePortPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bandwidth of procured ports in Gbps.
        /// </summary>
        [Input("bandwidthInGbps")]
        public Input<int>? BandwidthInGbps { get; set; }

        /// <summary>
        /// Encapsulation method on physical ports.
        /// </summary>
        [Input("encapsulation")]
        public Input<string>? Encapsulation { get; set; }

        [Input("links")]
        private InputList<Inputs.ExpressRouteLinkArgs>? _links;

        /// <summary>
        /// The set of physical links of the ExpressRoutePort resource.
        /// </summary>
        public InputList<Inputs.ExpressRouteLinkArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.ExpressRouteLinkArgs>());
            set => _links = value;
        }

        /// <summary>
        /// The name of the peering location that the ExpressRoutePort is mapped to physically.
        /// </summary>
        [Input("peeringLocation")]
        public Input<string>? PeeringLocation { get; set; }

        public ExpressRoutePortPropertiesFormatArgs()
        {
        }
    }
}