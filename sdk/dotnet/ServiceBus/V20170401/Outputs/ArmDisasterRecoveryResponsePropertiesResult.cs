// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus.V20170401.Outputs
{

    [OutputType]
    public sealed class ArmDisasterRecoveryResponsePropertiesResult
    {
        /// <summary>
        /// Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        /// </summary>
        public readonly string? AlternateName;
        /// <summary>
        /// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        /// </summary>
        public readonly string? PartnerNamespace;
        /// <summary>
        /// Number of entities pending to be replicated.
        /// </summary>
        public readonly int PendingReplicationOperationsCount;
        /// <summary>
        /// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private ArmDisasterRecoveryResponsePropertiesResult(
            string? alternateName,

            string? partnerNamespace,

            int pendingReplicationOperationsCount,

            string provisioningState,

            string role)
        {
            AlternateName = alternateName;
            PartnerNamespace = partnerNamespace;
            PendingReplicationOperationsCount = pendingReplicationOperationsCount;
            ProvisioningState = provisioningState;
            Role = role;
        }
    }
}