// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.V20151031.Inputs
{

    /// <summary>
    /// The parameters supplied to the create or update schedule operation.
    /// </summary>
    public sealed class ScheduleCreateOrUpdatePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the AdvancedSchedule.
        /// </summary>
        [Input("advancedSchedule")]
        public Input<Inputs.AdvancedScheduleArgs>? AdvancedSchedule { get; set; }

        /// <summary>
        /// Gets or sets the description of the schedule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Gets or sets the end time of the schedule.
        /// </summary>
        [Input("expiryTime")]
        public Input<string>? ExpiryTime { get; set; }

        /// <summary>
        /// Gets or sets the frequency of the schedule.
        /// </summary>
        [Input("frequency", required: true)]
        public Input<string> Frequency { get; set; } = null!;

        [Input("interval")]
        private InputMap<object>? _interval;

        /// <summary>
        /// Gets or sets the interval of the schedule.
        /// </summary>
        public InputMap<object> Interval
        {
            get => _interval ?? (_interval = new InputMap<object>());
            set => _interval = value;
        }

        /// <summary>
        /// Gets or sets the start time of the schedule.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        /// <summary>
        /// Gets or sets the time zone of the schedule.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public ScheduleCreateOrUpdatePropertiesArgs()
        {
        }
    }
}