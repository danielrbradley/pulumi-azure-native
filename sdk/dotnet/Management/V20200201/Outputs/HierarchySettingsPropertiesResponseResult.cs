// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.V20200201.Outputs
{

    [OutputType]
    public sealed class HierarchySettingsPropertiesResponseResult
    {
        /// <summary>
        /// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
        /// </summary>
        public readonly string? DefaultManagementGroup;
        /// <summary>
        /// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
        /// </summary>
        public readonly bool? RequireAuthorizationForGroupCreation;
        /// <summary>
        /// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
        /// </summary>
        public readonly string? TenantId;

        [OutputConstructor]
        private HierarchySettingsPropertiesResponseResult(
            string? defaultManagementGroup,

            bool? requireAuthorizationForGroupCreation,

            string? tenantId)
        {
            DefaultManagementGroup = defaultManagementGroup;
            RequireAuthorizationForGroupCreation = requireAuthorizationForGroupCreation;
            TenantId = tenantId;
        }
    }
}