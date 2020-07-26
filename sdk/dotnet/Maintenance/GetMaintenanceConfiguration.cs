// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Maintenance
{
    public static class GetMaintenanceConfiguration
    {
        public static Task<GetMaintenanceConfigurationResult> InvokeAsync(GetMaintenanceConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMaintenanceConfigurationResult>("azurerm:maintenance:getMaintenanceConfiguration", args ?? new GetMaintenanceConfigurationArgs(), options.WithVersion());
    }


    public sealed class GetMaintenanceConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Resource Identifier
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMaintenanceConfigurationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMaintenanceConfigurationResult
    {
        /// <summary>
        /// Gets or sets location of the resource
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets properties of the resource
        /// </summary>
        public readonly Outputs.MaintenanceConfigurationPropertiesResponseResult Properties;
        /// <summary>
        /// Gets or sets tags of the resource
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Type of the resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMaintenanceConfigurationResult(
            string? location,

            string name,

            Outputs.MaintenanceConfigurationPropertiesResponseResult properties,

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