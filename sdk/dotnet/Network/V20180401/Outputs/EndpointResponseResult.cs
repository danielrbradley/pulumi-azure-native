// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401.Outputs
{

    [OutputType]
    public sealed class EndpointResponseResult
    {
        /// <summary>
        /// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The properties of the Traffic Manager endpoint.
        /// </summary>
        public readonly Outputs.EndpointPropertiesResponseResult? Properties;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private EndpointResponseResult(
            string? id,

            string? name,

            Outputs.EndpointPropertiesResponseResult? properties,

            string? type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}