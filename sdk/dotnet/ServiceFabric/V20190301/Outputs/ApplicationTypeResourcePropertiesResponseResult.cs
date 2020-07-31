// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301.Outputs
{

    [OutputType]
    public sealed class ApplicationTypeResourcePropertiesResponseResult
    {
        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private ApplicationTypeResourcePropertiesResponseResult(string provisioningState)
        {
            ProvisioningState = provisioningState;
        }
    }
}