// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201
{
    public static class ListWebAppFunctionSecretsSlot
    {
        public static Task<ListWebAppFunctionSecretsSlotResult> InvokeAsync(ListWebAppFunctionSecretsSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWebAppFunctionSecretsSlotResult>("azurerm:web/v20180201:listWebAppFunctionSecretsSlot", args ?? new ListWebAppFunctionSecretsSlotArgs(), options.WithVersion());
    }


    public sealed class ListWebAppFunctionSecretsSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public ListWebAppFunctionSecretsSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWebAppFunctionSecretsSlotResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// FunctionSecrets resource specific properties
        /// </summary>
        public readonly Outputs.FunctionSecretsResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppFunctionSecretsSlotResult(
            string? kind,

            string name,

            Outputs.FunctionSecretsResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}