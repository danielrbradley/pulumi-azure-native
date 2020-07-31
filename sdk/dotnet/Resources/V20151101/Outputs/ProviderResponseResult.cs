// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Resources.V20151101.Outputs
{

    [OutputType]
    public sealed class ProviderResponseResult
    {
        /// <summary>
        /// Gets or sets the provider id.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Gets or sets the namespace of the provider.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// Gets or sets the registration state of the provider.
        /// </summary>
        public readonly string? RegistrationState;
        /// <summary>
        /// Gets or sets the collection of provider resource types.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProviderResourceTypeResponseResult> ResourceTypes;

        [OutputConstructor]
        private ProviderResponseResult(
            string? id,

            string? @namespace,

            string? registrationState,

            ImmutableArray<Outputs.ProviderResourceTypeResponseResult> resourceTypes)
        {
            Id = id;
            Namespace = @namespace;
            RegistrationState = registrationState;
            ResourceTypes = resourceTypes;
        }
    }
}