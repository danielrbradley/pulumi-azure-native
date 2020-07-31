// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801
{
    public static class GetSiteSourceControl
    {
        public static Task<GetSiteSourceControlResult> InvokeAsync(GetSiteSourceControlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSiteSourceControlResult>("azurerm:web/v20150801:getSiteSourceControl", args ?? new GetSiteSourceControlArgs(), options.WithVersion());
    }


    public sealed class GetSiteSourceControlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of web app
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSiteSourceControlArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSiteSourceControlResult
    {
        /// <summary>
        /// Kind of resource
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string? Name;
        public readonly Outputs.SiteSourceControlResponsePropertiesResult Properties;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetSiteSourceControlResult(
            string? kind,

            string location,

            string? name,

            Outputs.SiteSourceControlResponsePropertiesResult properties,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}