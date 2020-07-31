// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20190801.Outputs
{

    [OutputType]
    public sealed class VnetRouteResponsePropertiesResult
    {
        /// <summary>
        /// The ending address for this route. If the start address is specified in CIDR notation, this must be omitted.
        /// </summary>
        public readonly string? EndAddress;
        /// <summary>
        /// The type of route this is:
        /// DEFAULT - By default, every app has routes to the local address ranges specified by RFC1918
        /// INHERITED - Routes inherited from the real Virtual Network routes
        /// STATIC - Static route set on the app only
        /// 
        /// These values will be used for syncing an app's routes with those from a Virtual Network.
        /// </summary>
        public readonly string? RouteType;
        /// <summary>
        /// The starting address for this route. This may also include a CIDR notation, in which case the end address must not be specified.
        /// </summary>
        public readonly string? StartAddress;

        [OutputConstructor]
        private VnetRouteResponsePropertiesResult(
            string? endAddress,

            string? routeType,

            string? startAddress)
        {
            EndAddress = endAddress;
            RouteType = routeType;
            StartAddress = startAddress;
        }
    }
}