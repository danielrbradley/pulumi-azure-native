// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20190101
{
    public static class GetDiagnostic
    {
        public static Task<GetDiagnosticResult> InvokeAsync(GetDiagnosticArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDiagnosticResult>("azurerm:apimanagement/v20190101:getDiagnostic", args ?? new GetDiagnosticArgs(), options.WithVersion());
    }


    public sealed class GetDiagnosticArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Diagnostic identifier. Must be unique in the current API Management service instance.
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

        public GetDiagnosticArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDiagnosticResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Diagnostic entity contract properties.
        /// </summary>
        public readonly Outputs.DiagnosticContractPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDiagnosticResult(
            string name,

            Outputs.DiagnosticContractPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}