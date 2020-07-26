// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.Outputs
{

    [OutputType]
    public sealed class NetworkConfigurationResponseResult
    {
        /// <summary>
        /// The access endpoint.
        /// </summary>
        public readonly Outputs.IntegrationServiceEnvironmentAccessEndpointResponseResult? AccessEndpoint;
        /// <summary>
        /// The subnets.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceReferenceResponseResult> Subnets;
        /// <summary>
        /// Gets the virtual network address space.
        /// </summary>
        public readonly string? VirtualNetworkAddressSpace;

        [OutputConstructor]
        private NetworkConfigurationResponseResult(
            Outputs.IntegrationServiceEnvironmentAccessEndpointResponseResult? accessEndpoint,

            ImmutableArray<Outputs.ResourceReferenceResponseResult> subnets,

            string? virtualNetworkAddressSpace)
        {
            AccessEndpoint = accessEndpoint;
            Subnets = subnets;
            VirtualNetworkAddressSpace = virtualNetworkAddressSpace;
        }
    }
}