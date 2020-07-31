// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20181101.Outputs
{

    [OutputType]
    public sealed class SwiftVirtualNetworkResponsePropertiesResult
    {
        /// <summary>
        /// The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
        /// </summary>
        public readonly string? SubnetResourceId;
        /// <summary>
        /// A flag that specifies if the scale unit this Web App is on supports Swift integration.
        /// </summary>
        public readonly bool? SwiftSupported;

        [OutputConstructor]
        private SwiftVirtualNetworkResponsePropertiesResult(
            string? subnetResourceId,

            bool? swiftSupported)
        {
            SubnetResourceId = subnetResourceId;
            SwiftSupported = swiftSupported;
        }
    }
}