// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20151001
{
    public static class GetMediaService
    {
        public static Task<GetMediaServiceResult> InvokeAsync(GetMediaServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMediaServiceResult>("azurerm:media/v20151001:getMediaService", args ?? new GetMediaServiceArgs(), options.WithVersion());
    }


    public sealed class GetMediaServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Media Service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMediaServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMediaServiceResult
    {
        /// <summary>
        /// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The additional properties of a Media Service resource.
        /// </summary>
        public readonly Outputs.MediaServicePropertiesResponseResult Properties;
        /// <summary>
        /// Tags to help categorize the resource in the Azure portal.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMediaServiceResult(
            string? location,

            string name,

            Outputs.MediaServicePropertiesResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}