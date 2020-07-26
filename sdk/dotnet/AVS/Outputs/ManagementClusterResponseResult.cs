// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AVS.Outputs
{

    [OutputType]
    public sealed class ManagementClusterResponseResult
    {
        /// <summary>
        /// The identity
        /// </summary>
        public readonly int ClusterId;
        /// <summary>
        /// The cluster size
        /// </summary>
        public readonly int ClusterSize;
        /// <summary>
        /// The hosts
        /// </summary>
        public readonly ImmutableArray<string> Hosts;

        [OutputConstructor]
        private ManagementClusterResponseResult(
            int clusterId,

            int clusterSize,

            ImmutableArray<string> hosts)
        {
            ClusterId = clusterId;
            ClusterSize = clusterSize;
            Hosts = hosts;
        }
    }
}