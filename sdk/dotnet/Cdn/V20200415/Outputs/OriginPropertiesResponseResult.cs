// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20200415.Outputs
{

    [OutputType]
    public sealed class OriginPropertiesResponseResult
    {
        /// <summary>
        /// Origin is enabled for load balancing or not
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The value of the HTTP port. Must be between 1 and 65535.
        /// </summary>
        public readonly int? HttpPort;
        /// <summary>
        /// The value of the HTTPS port. Must be between 1 and 65535.
        /// </summary>
        public readonly int? HttpsPort;
        /// <summary>
        /// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
        /// </summary>
        public readonly string? OriginHostHeader;
        /// <summary>
        /// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The approval status for the connection to the Private Link
        /// </summary>
        public readonly string PrivateEndpointStatus;
        /// <summary>
        /// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
        /// </summary>
        public readonly string? PrivateLinkAlias;
        /// <summary>
        /// A custom message to be included in the approval request to connect to the Private Link.
        /// </summary>
        public readonly string? PrivateLinkApprovalMessage;
        /// <summary>
        /// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
        /// </summary>
        public readonly string? PrivateLinkLocation;
        /// <summary>
        /// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
        /// </summary>
        public readonly string? PrivateLinkResourceId;
        /// <summary>
        /// Provisioning status of the origin.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource status of the origin.
        /// </summary>
        public readonly string ResourceState;
        /// <summary>
        /// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private OriginPropertiesResponseResult(
            bool? enabled,

            string hostName,

            int? httpPort,

            int? httpsPort,

            string? originHostHeader,

            int? priority,

            string privateEndpointStatus,

            string? privateLinkAlias,

            string? privateLinkApprovalMessage,

            string? privateLinkLocation,

            string? privateLinkResourceId,

            string provisioningState,

            string resourceState,

            int? weight)
        {
            Enabled = enabled;
            HostName = hostName;
            HttpPort = httpPort;
            HttpsPort = httpsPort;
            OriginHostHeader = originHostHeader;
            Priority = priority;
            PrivateEndpointStatus = privateEndpointStatus;
            PrivateLinkAlias = privateLinkAlias;
            PrivateLinkApprovalMessage = privateLinkApprovalMessage;
            PrivateLinkLocation = privateLinkLocation;
            PrivateLinkResourceId = privateLinkResourceId;
            ProvisioningState = provisioningState;
            ResourceState = resourceState;
            Weight = weight;
        }
    }
}