// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayFrontendIPConfigurationResponseResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the frontend IP configuration that is unique within an Application Gateway.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Properties of the application gateway frontend IP configuration.
        /// </summary>
        public readonly Outputs.ApplicationGatewayFrontendIPConfigurationPropertiesFormatResponseResult? Properties;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ApplicationGatewayFrontendIPConfigurationResponseResult(
            string etag,

            string? id,

            string? name,

            Outputs.ApplicationGatewayFrontendIPConfigurationPropertiesFormatResponseResult? properties,

            string type)
        {
            Etag = etag;
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}