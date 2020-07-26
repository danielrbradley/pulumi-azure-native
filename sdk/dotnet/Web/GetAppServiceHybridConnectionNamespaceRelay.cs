// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web
{
    public static class GetAppServiceHybridConnectionNamespaceRelay
    {
        public static Task<GetAppServiceHybridConnectionNamespaceRelayResult> InvokeAsync(GetAppServiceHybridConnectionNamespaceRelayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServiceHybridConnectionNamespaceRelayResult>("azurerm:web:getAppServiceHybridConnectionNamespaceRelay", args ?? new GetAppServiceHybridConnectionNamespaceRelayArgs(), options.WithVersion());
    }


    public sealed class GetAppServiceHybridConnectionNamespaceRelayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The relay name for this hybrid connection.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace for this hybrid connection.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAppServiceHybridConnectionNamespaceRelayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServiceHybridConnectionNamespaceRelayResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// HybridConnection resource specific properties
        /// </summary>
        public readonly Outputs.HybridConnectionResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAppServiceHybridConnectionNamespaceRelayResult(
            string? kind,

            string name,

            Outputs.HybridConnectionResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}