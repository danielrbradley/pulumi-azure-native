// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic
{
    public static class GetIntegrationAccountAgreement
    {
        public static Task<GetIntegrationAccountAgreementResult> InvokeAsync(GetIntegrationAccountAgreementArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationAccountAgreementResult>("azurerm:logic:getIntegrationAccountAgreement", args ?? new GetIntegrationAccountAgreementArgs(), options.WithVersion());
    }


    public sealed class GetIntegrationAccountAgreementArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public string IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The integration account agreement name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIntegrationAccountAgreementArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationAccountAgreementResult
    {
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Gets the resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The integration account agreement properties.
        /// </summary>
        public readonly Outputs.IntegrationAccountAgreementPropertiesResponseResult Properties;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Gets the resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIntegrationAccountAgreementResult(
            string? location,

            string name,

            Outputs.IntegrationAccountAgreementPropertiesResponseResult properties,

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