// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20190501.Outputs
{

    [OutputType]
    public sealed class RequestResponseResult
    {
        /// <summary>
        /// The IP or hostname and possibly port of the client connection that initiated the event. This is the RemoteAddr from the standard http request.
        /// </summary>
        public readonly string? Addr;
        /// <summary>
        /// The externally accessible hostname of the registry instance, as specified by the http host header on incoming requests.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// The ID of the request that initiated the event.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The request method that generated the event.
        /// </summary>
        public readonly string? Method;
        /// <summary>
        /// The user agent header of the request.
        /// </summary>
        public readonly string? Useragent;

        [OutputConstructor]
        private RequestResponseResult(
            string? addr,

            string? host,

            string? id,

            string? method,

            string? useragent)
        {
            Addr = addr;
            Host = host;
            Id = id;
            Method = method;
            Useragent = useragent;
        }
    }
}