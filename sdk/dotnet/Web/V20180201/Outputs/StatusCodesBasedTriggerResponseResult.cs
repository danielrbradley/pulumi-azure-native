// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201.Outputs
{

    [OutputType]
    public sealed class StatusCodesBasedTriggerResponseResult
    {
        /// <summary>
        /// Request Count.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// HTTP status code.
        /// </summary>
        public readonly int? Status;
        /// <summary>
        /// Request Sub Status.
        /// </summary>
        public readonly int? SubStatus;
        /// <summary>
        /// Time interval.
        /// </summary>
        public readonly string? TimeInterval;
        /// <summary>
        /// Win32 error code.
        /// </summary>
        public readonly int? Win32Status;

        [OutputConstructor]
        private StatusCodesBasedTriggerResponseResult(
            int? count,

            int? status,

            int? subStatus,

            string? timeInterval,

            int? win32Status)
        {
            Count = count;
            Status = status;
            SubStatus = subStatus;
            TimeInterval = timeInterval;
            Win32Status = win32Status;
        }
    }
}