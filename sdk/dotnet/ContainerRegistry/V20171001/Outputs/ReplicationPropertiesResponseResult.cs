// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20171001.Outputs
{

    [OutputType]
    public sealed class ReplicationPropertiesResponseResult
    {
        /// <summary>
        /// The provisioning state of the replication at the time the operation was called.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The status of the replication at the time the operation was called.
        /// </summary>
        public readonly Outputs.StatusResponseResult Status;

        [OutputConstructor]
        private ReplicationPropertiesResponseResult(
            string provisioningState,

            Outputs.StatusResponseResult status)
        {
            ProvisioningState = provisioningState;
            Status = status;
        }
    }
}