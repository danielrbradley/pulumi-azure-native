// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20190101.Inputs
{

    /// <summary>
    /// API Management service resource SKU properties.
    /// </summary>
    public sealed class ApiManagementServiceSkuPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Capacity of the SKU (number of deployed units of the SKU).
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Name of the Sku.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ApiManagementServiceSkuPropertiesArgs()
        {
        }
    }
}