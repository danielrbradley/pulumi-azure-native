// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.TimeSeriesInsights.Inputs
{

    /// <summary>
    /// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
    /// </summary>
    public sealed class LocalTimestampArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An enum that represents the format of the local timestamp property that needs to be set.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// An object that represents the offset information for the local timestamp format specified. Should not be specified for LocalTimestampFormat - Embedded.
        /// </summary>
        [Input("timeZoneOffset")]
        public Input<Inputs.LocalTimestampPropertiesArgs>? TimeZoneOffset { get; set; }

        public LocalTimestampArgs()
        {
        }
    }
}