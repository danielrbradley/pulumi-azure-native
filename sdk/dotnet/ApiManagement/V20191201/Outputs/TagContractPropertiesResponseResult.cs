// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201.Outputs
{

    [OutputType]
    public sealed class TagContractPropertiesResponseResult
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string DisplayName;

        [OutputConstructor]
        private TagContractPropertiesResponseResult(string displayName)
        {
            DisplayName = displayName;
        }
    }
}