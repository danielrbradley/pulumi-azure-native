// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CognitiveServices
{
    public static class GetAccountPrivateEndpointConnection
    {
        public static Task<GetAccountPrivateEndpointConnectionResult> InvokeAsync(GetAccountPrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountPrivateEndpointConnectionResult>("azurerm:cognitiveservices:getAccountPrivateEndpointConnection", args ?? new GetAccountPrivateEndpointConnectionArgs(), options.WithVersion());
    }


    public sealed class GetAccountPrivateEndpointConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of Cognitive Services account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the private endpoint connection associated with the Cognitive Services Account
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountPrivateEndpointConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountPrivateEndpointConnectionResult
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource properties.
        /// </summary>
        public readonly Outputs.PrivateEndpointConnectionPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAccountPrivateEndpointConnectionResult(
            string name,

            Outputs.PrivateEndpointConnectionPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}