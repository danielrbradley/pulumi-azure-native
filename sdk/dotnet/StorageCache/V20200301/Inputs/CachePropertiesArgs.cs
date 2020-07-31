// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.V20200301.Inputs
{

    /// <summary>
    /// Properties of the Cache.
    /// </summary>
    public sealed class CachePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size of this Cache, in GB.
        /// </summary>
        [Input("cacheSizeGB")]
        public Input<int>? CacheSizeGB { get; set; }

        /// <summary>
        /// Specifies encryption settings of the cache.
        /// </summary>
        [Input("encryptionSettings")]
        public Input<Inputs.CacheEncryptionSettingsArgs>? EncryptionSettings { get; set; }

        /// <summary>
        /// Specifies network settings of the cache.
        /// </summary>
        [Input("networkSettings")]
        public Input<Inputs.CacheNetworkSettingsArgs>? NetworkSettings { get; set; }

        /// <summary>
        /// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Specifies security settings of the cache.
        /// </summary>
        [Input("securitySettings")]
        public Input<Inputs.CacheSecuritySettingsArgs>? SecuritySettings { get; set; }

        /// <summary>
        /// Subnet used for the Cache.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Upgrade status of the Cache.
        /// </summary>
        [Input("upgradeStatus")]
        public Input<Inputs.CacheUpgradeStatusArgs>? UpgradeStatus { get; set; }

        public CachePropertiesArgs()
        {
        }
    }
}