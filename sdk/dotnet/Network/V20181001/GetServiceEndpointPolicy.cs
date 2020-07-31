// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181001
{
    public static class GetServiceEndpointPolicy
    {
        public static Task<GetServiceEndpointPolicyResult> InvokeAsync(GetServiceEndpointPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceEndpointPolicyResult>("azurerm:network/v20181001:getServiceEndpointPolicy", args ?? new GetServiceEndpointPolicyArgs(), options.WithVersion());
    }


    public sealed class GetServiceEndpointPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the service endpoint policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServiceEndpointPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceEndpointPolicyResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the service end point policy
        /// </summary>
        public readonly Outputs.ServiceEndpointPolicyPropertiesFormatResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServiceEndpointPolicyResult(
            string? etag,

            string? location,

            string name,

            Outputs.ServiceEndpointPolicyPropertiesFormatResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}