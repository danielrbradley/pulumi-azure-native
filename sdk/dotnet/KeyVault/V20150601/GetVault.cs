// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.KeyVault.V20150601
{
    public static class GetVault
    {
        public static Task<GetVaultResult> InvokeAsync(GetVaultArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVaultResult>("azurerm:keyvault/v20150601:getVault", args ?? new GetVaultArgs(), options.WithVersion());
    }


    public sealed class GetVaultArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the vault.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group to which the vault belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVaultArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVaultResult
    {
        /// <summary>
        /// The supported Azure location where the key vault should be created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the key vault.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the vault
        /// </summary>
        public readonly Outputs.VaultPropertiesResponseResult Properties;
        /// <summary>
        /// The tags that will be assigned to the key vault. 
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type of the key vault.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVaultResult(
            string location,

            string name,

            Outputs.VaultPropertiesResponseResult properties,

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