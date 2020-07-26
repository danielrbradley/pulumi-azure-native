// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageSync.Inputs
{

    /// <summary>
    /// CloudEndpoint Properties object.
    /// </summary>
    public sealed class CloudEndpointCreateParametersPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure file share name
        /// </summary>
        [Input("azureFileShareName")]
        public Input<string>? AzureFileShareName { get; set; }

        /// <summary>
        /// Friendly Name
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// Storage Account Resource Id
        /// </summary>
        [Input("storageAccountResourceId")]
        public Input<string>? StorageAccountResourceId { get; set; }

        /// <summary>
        /// Storage Account Tenant Id
        /// </summary>
        [Input("storageAccountTenantId")]
        public Input<string>? StorageAccountTenantId { get; set; }

        public CloudEndpointCreateParametersPropertiesArgs()
        {
        }
    }
}