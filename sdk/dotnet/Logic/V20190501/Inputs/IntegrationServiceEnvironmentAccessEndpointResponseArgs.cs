// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20190501.Inputs
{

    /// <summary>
    /// The integration service environment access endpoint.
    /// </summary>
    public sealed class IntegrationServiceEnvironmentAccessEndpointResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The access endpoint type.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public IntegrationServiceEnvironmentAccessEndpointResponseArgs()
        {
        }
    }
}