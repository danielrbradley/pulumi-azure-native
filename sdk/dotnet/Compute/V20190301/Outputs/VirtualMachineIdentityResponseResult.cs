// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301.Outputs
{

    [OutputType]
    public sealed class VirtualMachineIdentityResponseResult
    {
        /// <summary>
        /// The principal id of virtual machine identity. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant id associated with the virtual machine. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of identity used for the virtual machine. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The list of user identities associated with the Virtual Machine. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.VirtualMachineIdentityResponsePropertiesResult>? UserAssignedIdentities;

        [OutputConstructor]
        private VirtualMachineIdentityResponseResult(
            string principalId,

            string tenantId,

            string? type,

            ImmutableDictionary<string, Outputs.VirtualMachineIdentityResponsePropertiesResult>? userAssignedIdentities)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}