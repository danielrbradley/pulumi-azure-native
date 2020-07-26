// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web
{
    public static class ListAppServiceSlotFunctionKeys
    {
        public static Task<ListAppServiceSlotFunctionKeysResult> InvokeAsync(ListAppServiceSlotFunctionKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListAppServiceSlotFunctionKeysResult>("azurerm:web:listAppServiceSlotFunctionKeys", args ?? new ListAppServiceSlotFunctionKeysArgs(), options.WithVersion());
    }


    public sealed class ListAppServiceSlotFunctionKeysArgs : Pulumi.InvokeArgs
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

        public ListAppServiceSlotFunctionKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListAppServiceSlotFunctionKeysResult
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
        /// Settings.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListAppServiceSlotFunctionKeysResult(
            string? kind,

            string name,

            ImmutableDictionary<string, string> properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}