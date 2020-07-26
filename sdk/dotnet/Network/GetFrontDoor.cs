// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network
{
    public static class GetFrontDoor
    {
        public static Task<GetFrontDoorResult> InvokeAsync(GetFrontDoorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFrontDoorResult>("azurerm:network:getFrontDoor", args ?? new GetFrontDoorArgs(), options.WithVersion());
    }


    public sealed class GetFrontDoorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Front Door which is globally unique.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFrontDoorArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFrontDoorResult
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the Front Door Load Balancer
        /// </summary>
        public readonly Outputs.FrontDoorPropertiesResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetFrontDoorResult(
            string? location,

            string name,

            Outputs.FrontDoorPropertiesResponseResult properties,

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