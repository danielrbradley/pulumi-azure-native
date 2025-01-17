// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.V20220401Preview.Outputs
{

    /// <summary>
    /// Cni configuration
    /// </summary>
    [OutputType]
    public sealed class CniConfigResponse
    {
        /// <summary>
        /// ComponentType of the Kubernetes node.
        /// </summary>
        public readonly string ComponentType;
        /// <summary>
        /// Pod Subnet
        /// </summary>
        public readonly string PodSubnet;
        /// <summary>
        /// Service subnet
        /// </summary>
        public readonly string ServiceSubnet;
        /// <summary>
        /// Cni type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Cni version
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private CniConfigResponse(
            string componentType,

            string podSubnet,

            string serviceSubnet,

            string type,

            string version)
        {
            ComponentType = componentType;
            PodSubnet = podSubnet;
            ServiceSubnet = serviceSubnet;
            Type = type;
            Version = version;
        }
    }
}
