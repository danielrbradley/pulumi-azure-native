// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20160601.Outputs
{

    [OutputType]
    public sealed class RosettaNetPipAcknowledgmentOfReceiptSettingsResponseResult
    {
        /// <summary>
        /// The non-repudiation is required or not.
        /// </summary>
        public readonly bool IsNonRepudiationRequired;
        /// <summary>
        /// The time to acknowledge in seconds.
        /// </summary>
        public readonly int TimeToAcknowledgeInSeconds;

        [OutputConstructor]
        private RosettaNetPipAcknowledgmentOfReceiptSettingsResponseResult(
            bool isNonRepudiationRequired,

            int timeToAcknowledgeInSeconds)
        {
            IsNonRepudiationRequired = isNonRepudiationRequired;
            TimeToAcknowledgeInSeconds = timeToAcknowledgeInSeconds;
        }
    }
}