// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web
{
    public static class GetStaticSite
    {
        public static Task<GetStaticSiteResult> InvokeAsync(GetStaticSiteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStaticSiteResult>("azurerm:web:getStaticSite", args ?? new GetStaticSiteArgs(), options.WithVersion());
    }


    public sealed class GetStaticSiteArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the static site.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetStaticSiteArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStaticSiteResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Core resource properties
        /// </summary>
        public readonly Outputs.StaticSiteResponseResult Properties;
        /// <summary>
        /// Description of a SKU for a scalable resource.
        /// </summary>
        public readonly Outputs.SkuDescriptionResponseResult? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetStaticSiteResult(
            string? kind,

            string location,

            string name,

            Outputs.StaticSiteResponseResult properties,

            Outputs.SkuDescriptionResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}