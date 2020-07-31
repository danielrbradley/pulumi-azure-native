// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Relay.V20170401
{
    public static class GetHybridConnection
    {
        public static Task<GetHybridConnectionResult> InvokeAsync(GetHybridConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHybridConnectionResult>("azurerm:relay/v20170401:getHybridConnection", args ?? new GetHybridConnectionArgs(), options.WithVersion());
    }


    public sealed class GetHybridConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hybrid connection name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetHybridConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHybridConnectionResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the HybridConnection.
        /// </summary>
        public readonly Outputs.HybridConnectionResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetHybridConnectionResult(
            string name,

            Outputs.HybridConnectionResponsePropertiesResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}