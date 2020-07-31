// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201
{
    public static class GetApiManagementService
    {
        public static Task<GetApiManagementServiceResult> InvokeAsync(GetApiManagementServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiManagementServiceResult>("azurerm:apimanagement/v20191201:getApiManagementService", args ?? new GetApiManagementServiceArgs(), options.WithVersion());
    }


    public sealed class GetApiManagementServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApiManagementServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiManagementServiceResult
    {
        /// <summary>
        /// ETag of the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Managed service identity of the Api Management service.
        /// </summary>
        public readonly Outputs.ApiManagementServiceIdentityResponseResult? Identity;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the API Management service.
        /// </summary>
        public readonly Outputs.ApiManagementServicePropertiesResponseResult Properties;
        /// <summary>
        /// SKU properties of the API Management service.
        /// </summary>
        public readonly Outputs.ApiManagementServiceSkuPropertiesResponseResult Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type for API Management resource is set to Microsoft.ApiManagement.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiManagementServiceResult(
            string etag,

            Outputs.ApiManagementServiceIdentityResponseResult? identity,

            string location,

            string name,

            Outputs.ApiManagementServicePropertiesResponseResult properties,

            Outputs.ApiManagementServiceSkuPropertiesResponseResult sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}