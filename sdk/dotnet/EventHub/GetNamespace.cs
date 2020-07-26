// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventHub
{
    public static class GetNamespace
    {
        public static Task<GetNamespaceResult> InvokeAsync(GetNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceResult>("azurerm:eventhub:getNamespace", args ?? new GetNamespaceArgs(), options.WithVersion());
    }


    public sealed class GetNamespaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceResult
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespace properties supplied for create namespace operation.
        /// </summary>
        public readonly Outputs.EHNamespaceResponsePropertiesResult Properties;
        /// <summary>
        /// Properties of sku resource
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNamespaceResult(
            string? location,

            string name,

            Outputs.EHNamespaceResponsePropertiesResult properties,

            Outputs.SkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}