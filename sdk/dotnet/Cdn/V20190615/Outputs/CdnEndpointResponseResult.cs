// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20190615.Outputs
{

    [OutputType]
    public sealed class CdnEndpointResponseResult
    {
        /// <summary>
        /// ARM Resource ID string.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private CdnEndpointResponseResult(string? id)
        {
            Id = id;
        }
    }
}