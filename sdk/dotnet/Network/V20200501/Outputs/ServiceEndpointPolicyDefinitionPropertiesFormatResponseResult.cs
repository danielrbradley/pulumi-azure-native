// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Outputs
{

    [OutputType]
    public sealed class ServiceEndpointPolicyDefinitionPropertiesFormatResponseResult
    {
        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The provisioning state of the service endpoint policy definition resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Service endpoint name.
        /// </summary>
        public readonly string? Service;
        /// <summary>
        /// A list of service resources.
        /// </summary>
        public readonly ImmutableArray<string> ServiceResources;

        [OutputConstructor]
        private ServiceEndpointPolicyDefinitionPropertiesFormatResponseResult(
            string? description,

            string provisioningState,

            string? service,

            ImmutableArray<string> serviceResources)
        {
            Description = description;
            ProvisioningState = provisioningState;
            Service = service;
            ServiceResources = serviceResources;
        }
    }
}