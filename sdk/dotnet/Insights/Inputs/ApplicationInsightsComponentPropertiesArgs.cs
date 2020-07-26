// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.Inputs
{

    /// <summary>
    /// Properties that define an Application Insights component resource.
    /// </summary>
    public sealed class ApplicationInsightsComponentPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of application being monitored.
        /// </summary>
        [Input("Application_Type", required: true)]
        public Input<string> Application_Type { get; set; } = null!;

        /// <summary>
        /// Disable IP masking.
        /// </summary>
        [Input("DisableIpMasking")]
        public Input<bool>? DisableIpMasking { get; set; }

        /// <summary>
        /// Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
        /// </summary>
        [Input("Flow_Type")]
        public Input<string>? Flow_Type { get; set; }

        /// <summary>
        /// The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
        /// </summary>
        [Input("HockeyAppId")]
        public Input<string>? HockeyAppId { get; set; }

        /// <summary>
        /// Purge data immediately after 30 days.
        /// </summary>
        [Input("ImmediatePurgeDataOn30Days")]
        public Input<bool>? ImmediatePurgeDataOn30Days { get; set; }

        /// <summary>
        /// Indicates the flow of the ingestion.
        /// </summary>
        [Input("IngestionMode")]
        public Input<string>? IngestionMode { get; set; }

        /// <summary>
        /// Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
        /// </summary>
        [Input("Request_Source")]
        public Input<string>? Request_Source { get; set; }

        /// <summary>
        /// Retention period in days.
        /// </summary>
        [Input("RetentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
        /// </summary>
        [Input("SamplingPercentage")]
        public Input<double>? SamplingPercentage { get; set; }

        public ApplicationInsightsComponentPropertiesArgs()
        {
        }
    }
}