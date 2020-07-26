// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Databricks
{
    public static class GetWorkspaceVirtualNetworkPeering
    {
        public static Task<GetWorkspaceVirtualNetworkPeeringResult> InvokeAsync(GetWorkspaceVirtualNetworkPeeringArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceVirtualNetworkPeeringResult>("azurerm:databricks:getWorkspaceVirtualNetworkPeering", args ?? new GetWorkspaceVirtualNetworkPeeringArgs(), options.WithVersion());
    }


    public sealed class GetWorkspaceVirtualNetworkPeeringArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the workspace vNet peering.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetWorkspaceVirtualNetworkPeeringArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceVirtualNetworkPeeringResult
    {
        /// <summary>
        /// Name of the virtual network peering resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of properties for vNet Peering
        /// </summary>
        public readonly Outputs.VirtualNetworkPeeringPropertiesFormatResponseResult Properties;
        /// <summary>
        /// type of the virtual network peering resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWorkspaceVirtualNetworkPeeringResult(
            string name,

            Outputs.VirtualNetworkPeeringPropertiesFormatResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}