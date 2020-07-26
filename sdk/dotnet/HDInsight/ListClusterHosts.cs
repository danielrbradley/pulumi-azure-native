// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HDInsight
{
    public static class ListClusterHosts
    {
        public static Task<ListClusterHostsResult> InvokeAsync(ListClusterHostsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListClusterHostsResult>("azurerm:hdinsight:listClusterHosts", args ?? new ListClusterHostsArgs(), options.WithVersion());
    }


    public sealed class ListClusterHostsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListClusterHostsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListClusterHostsResult
    {
        [OutputConstructor]
        private ListClusterHostsResult()
        {
        }
    }
}