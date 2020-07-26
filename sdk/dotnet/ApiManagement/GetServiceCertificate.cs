// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement
{
    public static class GetServiceCertificate
    {
        public static Task<GetServiceCertificateResult> InvokeAsync(GetServiceCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceCertificateResult>("azurerm:apimanagement:getServiceCertificate", args ?? new GetServiceCertificateArgs(), options.WithVersion());
    }


    public sealed class GetServiceCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the certificate entity. Must be unique in the current API Management service instance.
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

        public GetServiceCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceCertificateResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Certificate properties details.
        /// </summary>
        public readonly Outputs.CertificateContractPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServiceCertificateResult(
            string name,

            Outputs.CertificateContractPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}