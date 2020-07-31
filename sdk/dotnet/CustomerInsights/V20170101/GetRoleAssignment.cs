// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.V20170101
{
    public static class GetRoleAssignment
    {
        public static Task<GetRoleAssignmentResult> InvokeAsync(GetRoleAssignmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleAssignmentResult>("azurerm:customerinsights/v20170101:getRoleAssignment", args ?? new GetRoleAssignmentArgs(), options.WithVersion());
    }


    public sealed class GetRoleAssignmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public string HubName { get; set; } = null!;

        /// <summary>
        /// The name of the role assignment.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRoleAssignmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRoleAssignmentResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Role Assignment definition.
        /// </summary>
        public readonly Outputs.RoleAssignmentResponseResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRoleAssignmentResult(
            string name,

            Outputs.RoleAssignmentResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}