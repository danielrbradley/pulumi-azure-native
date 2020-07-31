// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.KeyVault.V20190901.Outputs
{

    [OutputType]
    public sealed class PrivateLinkServiceConnectionStateResponseResult
    {
        /// <summary>
        /// A message indicating if changes on the service provider require any updates on the consumer.
        /// </summary>
        public readonly string? ActionRequired;
        /// <summary>
        /// The reason for approval or rejection.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Indicates whether the connection has been approved, rejected or removed by the key vault owner.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private PrivateLinkServiceConnectionStateResponseResult(
            string? actionRequired,

            string? description,

            string? status)
        {
            ActionRequired = actionRequired;
            Description = description;
            Status = status;
        }
    }
}