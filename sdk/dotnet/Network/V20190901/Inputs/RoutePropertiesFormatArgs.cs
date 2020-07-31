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
    /// Route resource.
    /// </summary>
    public sealed class RoutePropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR to which the route applies.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        /// </summary>
        [Input("nextHopIpAddress")]
        public Input<string>? NextHopIpAddress { get; set; }

        /// <summary>
        /// The type of Azure hop the packet should be sent to.
        /// </summary>
        [Input("nextHopType", required: true)]
        public Input<string> NextHopType { get; set; } = null!;

        public RoutePropertiesFormatArgs()
        {
        }
    }
}