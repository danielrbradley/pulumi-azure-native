// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20200501
{
    public static class GetApplication
    {
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("azurerm:batch/v20200501:getApplication", args ?? new GetApplicationArgs(), options.WithVersion());
    }


    public sealed class GetApplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the application. This must be unique within the account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The ETag of the resource, used for concurrency statements.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties associated with the Application.
        /// </summary>
        public readonly Outputs.ApplicationPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationResult(
            string etag,

            string name,

            Outputs.ApplicationPropertiesResponseResult properties,

            string type)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}