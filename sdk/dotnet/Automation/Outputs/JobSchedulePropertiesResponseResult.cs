// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.Outputs
{

    [OutputType]
    public sealed class JobSchedulePropertiesResponseResult
    {
        /// <summary>
        /// Gets or sets the id of job schedule.
        /// </summary>
        public readonly string? JobScheduleId;
        /// <summary>
        /// Gets or sets the parameters of the job schedule.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Parameters;
        /// <summary>
        /// Gets or sets the hybrid worker group that the scheduled job should run on.
        /// </summary>
        public readonly string? RunOn;
        /// <summary>
        /// Gets or sets the runbook.
        /// </summary>
        public readonly Outputs.RunbookAssociationPropertyResponseResult? Runbook;
        /// <summary>
        /// Gets or sets the schedule.
        /// </summary>
        public readonly Outputs.ScheduleAssociationPropertyResponseResult? Schedule;

        [OutputConstructor]
        private JobSchedulePropertiesResponseResult(
            string? jobScheduleId,

            ImmutableDictionary<string, string>? parameters,

            string? runOn,

            Outputs.RunbookAssociationPropertyResponseResult? runbook,

            Outputs.ScheduleAssociationPropertyResponseResult? schedule)
        {
            JobScheduleId = jobScheduleId;
            Parameters = parameters;
            RunOn = runOn;
            Runbook = runbook;
            Schedule = schedule;
        }
    }
}