// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Orbital.V20210404Preview.Outputs
{

    /// <summary>
    /// Spacecraft Link
    /// </summary>
    [OutputType]
    public sealed class SpacecraftLinkResponse
    {
        /// <summary>
        /// Bandwidth in MHz
        /// </summary>
        public readonly double BandwidthMHz;
        /// <summary>
        /// Center Frequency in MHz
        /// </summary>
        public readonly double CenterFrequencyMHz;
        /// <summary>
        /// Direction (uplink or downlink)
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// polarization. eg (RHCP, LHCP)
        /// </summary>
        public readonly string Polarization;

        [OutputConstructor]
        private SpacecraftLinkResponse(
            double bandwidthMHz,

            double centerFrequencyMHz,

            string direction,

            string polarization)
        {
            BandwidthMHz = bandwidthMHz;
            CenterFrequencyMHz = centerFrequencyMHz;
            Direction = direction;
            Polarization = polarization;
        }
    }
}
