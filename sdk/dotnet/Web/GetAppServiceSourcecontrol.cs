// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web
{
    public static class GetAppServiceSourcecontrol
    {
        public static Task<GetAppServiceSourcecontrolResult> InvokeAsync(GetAppServiceSourcecontrolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServiceSourcecontrolResult>("azurerm:web:getAppServiceSourcecontrol", args ?? new GetAppServiceSourcecontrolArgs(), options.WithVersion());
    }


    public sealed class GetAppServiceSourcecontrolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAppServiceSourcecontrolArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServiceSourcecontrolResult
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
        /// SiteSourceControl resource specific properties
        /// </summary>
        public readonly Outputs.SiteSourceControlResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAppServiceSourcecontrolResult(
            string? kind,

            string name,

            Outputs.SiteSourceControlResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}