// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180710.Outputs
{

    [OutputType]
    public sealed class InnerHealthErrorResponseResult
    {
        /// <summary>
        /// Error creation time (UTC)
        /// </summary>
        public readonly string? CreationTimeUtc;
        /// <summary>
        /// ID of the entity.
        /// </summary>
        public readonly string? EntityId;
        /// <summary>
        /// Category of error.
        /// </summary>
        public readonly string? ErrorCategory;
        /// <summary>
        /// Error code.
        /// </summary>
        public readonly string? ErrorCode;
        /// <summary>
        /// Level of error.
        /// </summary>
        public readonly string? ErrorLevel;
        /// <summary>
        /// Error message.
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// Source of error.
        /// </summary>
        public readonly string? ErrorSource;
        /// <summary>
        /// Type of error.
        /// </summary>
        public readonly string? ErrorType;
        /// <summary>
        /// Possible causes of error.
        /// </summary>
        public readonly string? PossibleCauses;
        /// <summary>
        /// Recommended action to resolve error.
        /// </summary>
        public readonly string? RecommendedAction;
        /// <summary>
        /// DRA error message.
        /// </summary>
        public readonly string? RecoveryProviderErrorMessage;
        /// <summary>
        /// Summary message of the entity.
        /// </summary>
        public readonly string? SummaryMessage;

        [OutputConstructor]
        private InnerHealthErrorResponseResult(
            string? creationTimeUtc,

            string? entityId,

            string? errorCategory,

            string? errorCode,

            string? errorLevel,

            string? errorMessage,

            string? errorSource,

            string? errorType,

            string? possibleCauses,

            string? recommendedAction,

            string? recoveryProviderErrorMessage,

            string? summaryMessage)
        {
            CreationTimeUtc = creationTimeUtc;
            EntityId = entityId;
            ErrorCategory = errorCategory;
            ErrorCode = errorCode;
            ErrorLevel = errorLevel;
            ErrorMessage = errorMessage;
            ErrorSource = errorSource;
            ErrorType = errorType;
            PossibleCauses = possibleCauses;
            RecommendedAction = recommendedAction;
            RecoveryProviderErrorMessage = recoveryProviderErrorMessage;
            SummaryMessage = summaryMessage;
        }
    }
}