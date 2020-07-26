// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn
{
    public static class GetProfileEndpoint
    {
        public static Task<GetProfileEndpointResult> InvokeAsync(GetProfileEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProfileEndpointResult>("azurerm:cdn:getProfileEndpoint", args ?? new GetProfileEndpointArgs(), options.WithVersion());
    }


    public sealed class GetProfileEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetProfileEndpointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProfileEndpointResult
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The JSON object that contains the properties required to create an endpoint.
        /// </summary>
        public readonly Outputs.EndpointPropertiesResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProfileEndpointResult(
            string location,

            string name,

            Outputs.EndpointPropertiesResponseResult properties,

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