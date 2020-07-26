// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.Outputs
{

    [OutputType]
    public sealed class CacheSecuritySettingsResponseResult
    {
        /// <summary>
        /// root squash of cache property.
        /// </summary>
        public readonly bool? RootSquash;

        [OutputConstructor]
        private CacheSecuritySettingsResponseResult(bool? rootSquash)
        {
            RootSquash = rootSquash;
        }
    }
}