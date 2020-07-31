// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20161201.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayFrontendPortResponseResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Properties of Frontend port of an application gateway.
        /// </summary>
        public readonly Outputs.ApplicationGatewayFrontendPortPropertiesFormatResponseResult? Properties;

        [OutputConstructor]
        private ApplicationGatewayFrontendPortResponseResult(
            string? etag,

            string? id,

            string? name,

            Outputs.ApplicationGatewayFrontendPortPropertiesFormatResponseResult? properties)
        {
            Etag = etag;
            Id = id;
            Name = name;
            Properties = properties;
        }
    }
}