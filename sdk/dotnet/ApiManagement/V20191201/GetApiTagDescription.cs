// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201
{
    public static class GetApiTagDescription
    {
        public static Task<GetApiTagDescriptionResult> InvokeAsync(GetApiTagDescriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiTagDescriptionResult>("azurerm:apimanagement/v20191201:getApiTagDescription", args ?? new GetApiTagDescriptionArgs(), options.WithVersion());
    }


    public sealed class GetApiTagDescriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// Tag description identifier. Used when creating tagDescription for API/Tag association. Based on API and Tag names.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetApiTagDescriptionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiTagDescriptionResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// TagDescription entity contract properties.
        /// </summary>
        public readonly Outputs.TagDescriptionContractPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiTagDescriptionResult(
            string name,

            Outputs.TagDescriptionContractPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}