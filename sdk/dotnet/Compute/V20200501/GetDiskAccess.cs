// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20200501
{
    public static class GetDiskAccess
    {
        public static Task<GetDiskAccessResult> InvokeAsync(GetDiskAccessArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDiskAccessResult>("azurerm:compute/v20200501:getDiskAccess", args ?? new GetDiskAccessArgs(), options.WithVersion());
    }


    public sealed class GetDiskAccessArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDiskAccessArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDiskAccessResult
    {
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        public readonly Outputs.DiskAccessPropertiesResponseResult Properties;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDiskAccessResult(
            string location,

            string name,

            Outputs.DiskAccessPropertiesResponseResult properties,

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