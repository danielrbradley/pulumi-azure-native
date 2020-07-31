// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180201.Inputs
{

    /// <summary>
    /// Parameters that define the operation to create a connection monitor.
    /// </summary>
    public sealed class ConnectionMonitorParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the connection monitor will start automatically once created.
        /// </summary>
        [Input("autoStart")]
        public Input<bool>? AutoStart { get; set; }

        /// <summary>
        /// Describes the destination of connection monitor.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ConnectionMonitorDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Monitoring interval in seconds.
        /// </summary>
        [Input("monitoringIntervalInSeconds")]
        public Input<int>? MonitoringIntervalInSeconds { get; set; }

        /// <summary>
        /// Describes the source of connection monitor.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.ConnectionMonitorSourceArgs> Source { get; set; } = null!;

        public ConnectionMonitorParametersArgs()
        {
        }
    }
}