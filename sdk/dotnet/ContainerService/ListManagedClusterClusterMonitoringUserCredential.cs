// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService
{
    public static class ListManagedClusterClusterMonitoringUserCredential
    {
        public static Task<ListManagedClusterClusterMonitoringUserCredentialResult> InvokeAsync(ListManagedClusterClusterMonitoringUserCredentialArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListManagedClusterClusterMonitoringUserCredentialResult>("azurerm:containerservice:listManagedClusterClusterMonitoringUserCredential", args ?? new ListManagedClusterClusterMonitoringUserCredentialArgs(), options.WithVersion());
    }


    public sealed class ListManagedClusterClusterMonitoringUserCredentialArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public ListManagedClusterClusterMonitoringUserCredentialArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListManagedClusterClusterMonitoringUserCredentialResult
    {
        /// <summary>
        /// Base64-encoded Kubernetes configuration file.
        /// </summary>
        public readonly ImmutableArray<Outputs.CredentialResultResponseResult> Kubeconfigs;

        [OutputConstructor]
        private ListManagedClusterClusterMonitoringUserCredentialResult(ImmutableArray<Outputs.CredentialResultResponseResult> kubeconfigs)
        {
            Kubeconfigs = kubeconfigs;
        }
    }
}