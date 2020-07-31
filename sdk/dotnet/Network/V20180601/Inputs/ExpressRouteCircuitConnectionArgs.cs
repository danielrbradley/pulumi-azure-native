// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180601.Inputs
{

    /// <summary>
    /// Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
    /// </summary>
    public sealed class ExpressRouteCircuitConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties")]
        public Input<Inputs.ExpressRouteCircuitConnectionPropertiesFormatArgs>? Properties { get; set; }

        public ExpressRouteCircuitConnectionArgs()
        {
        }
    }
}