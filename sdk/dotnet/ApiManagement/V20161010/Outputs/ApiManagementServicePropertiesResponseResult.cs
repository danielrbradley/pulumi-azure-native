// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20161010.Outputs
{

    [OutputType]
    public sealed class ApiManagementServicePropertiesResponseResult
    {
        /// <summary>
        /// Additional datacenter locations of the API Management service.
        /// </summary>
        public readonly ImmutableArray<Outputs.AdditionalRegionResponseResult> AdditionalLocations;
        /// <summary>
        /// Addresser email.
        /// </summary>
        public readonly string? AddresserEmail;
        /// <summary>
        /// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
        /// </summary>
        public readonly string CreatedAtUtc;
        /// <summary>
        /// Custom properties of the API Management service, like disabling TLS 1.0.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomProperties;
        /// <summary>
        /// Custom hostname configuration of the API Management service.
        /// </summary>
        public readonly ImmutableArray<Outputs.HostnameConfigurationResponseResult> HostnameConfigurations;
        /// <summary>
        /// Management API endpoint URL of the API Management service.
        /// </summary>
        public readonly string ManagementApiUrl;
        /// <summary>
        /// Publisher portal endpoint Url of the API Management service.
        /// </summary>
        public readonly string PortalUrl;
        /// <summary>
        /// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Publisher email.
        /// </summary>
        public readonly string PublisherEmail;
        /// <summary>
        /// Publisher name.
        /// </summary>
        public readonly string PublisherName;
        /// <summary>
        /// Proxy endpoint URL of the API Management service.
        /// </summary>
        public readonly string RuntimeUrl;
        /// <summary>
        /// SCM endpoint URL of the API Management service.
        /// </summary>
        public readonly string ScmUrl;
        /// <summary>
        /// Static IP addresses of the API Management service virtual machines. Available only for Standard and Premium SKU.
        /// </summary>
        public readonly ImmutableArray<string> StaticIPs;
        /// <summary>
        /// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
        /// </summary>
        public readonly string TargetProvisioningState;
        /// <summary>
        /// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
        /// </summary>
        public readonly string? VpnType;
        /// <summary>
        /// Virtual network configuration of the API Management service.
        /// </summary>
        public readonly Outputs.VirtualNetworkConfigurationResponseResult? Vpnconfiguration;

        [OutputConstructor]
        private ApiManagementServicePropertiesResponseResult(
            ImmutableArray<Outputs.AdditionalRegionResponseResult> additionalLocations,

            string? addresserEmail,

            string createdAtUtc,

            ImmutableDictionary<string, string>? customProperties,

            ImmutableArray<Outputs.HostnameConfigurationResponseResult> hostnameConfigurations,

            string managementApiUrl,

            string portalUrl,

            string provisioningState,

            string publisherEmail,

            string publisherName,

            string runtimeUrl,

            string scmUrl,

            ImmutableArray<string> staticIPs,

            string targetProvisioningState,

            string? vpnType,

            Outputs.VirtualNetworkConfigurationResponseResult? vpnconfiguration)
        {
            AdditionalLocations = additionalLocations;
            AddresserEmail = addresserEmail;
            CreatedAtUtc = createdAtUtc;
            CustomProperties = customProperties;
            HostnameConfigurations = hostnameConfigurations;
            ManagementApiUrl = managementApiUrl;
            PortalUrl = portalUrl;
            ProvisioningState = provisioningState;
            PublisherEmail = publisherEmail;
            PublisherName = publisherName;
            RuntimeUrl = runtimeUrl;
            ScmUrl = scmUrl;
            StaticIPs = staticIPs;
            TargetProvisioningState = targetProvisioningState;
            VpnType = vpnType;
            Vpnconfiguration = vpnconfiguration;
        }
    }
}