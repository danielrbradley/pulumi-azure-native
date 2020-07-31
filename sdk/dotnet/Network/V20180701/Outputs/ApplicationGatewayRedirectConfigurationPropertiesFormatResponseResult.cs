// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180701.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayRedirectConfigurationPropertiesFormatResponseResult
    {
        /// <summary>
        /// Include path in the redirected url.
        /// </summary>
        public readonly bool? IncludePath;
        /// <summary>
        /// Include query string in the redirected url.
        /// </summary>
        public readonly bool? IncludeQueryString;
        /// <summary>
        /// Path rules specifying redirect configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> PathRules;
        /// <summary>
        /// Supported http redirection types - Permanent, Temporary, Found, SeeOther.
        /// </summary>
        public readonly string? RedirectType;
        /// <summary>
        /// Request routing specifying redirect configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> RequestRoutingRules;
        /// <summary>
        /// Reference to a listener to redirect the request to.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? TargetListener;
        /// <summary>
        /// Url to redirect the request to.
        /// </summary>
        public readonly string? TargetUrl;
        /// <summary>
        /// Url path maps specifying default redirect configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> UrlPathMaps;

        [OutputConstructor]
        private ApplicationGatewayRedirectConfigurationPropertiesFormatResponseResult(
            bool? includePath,

            bool? includeQueryString,

            ImmutableArray<Outputs.SubResourceResponseResult> pathRules,

            string? redirectType,

            ImmutableArray<Outputs.SubResourceResponseResult> requestRoutingRules,

            Outputs.SubResourceResponseResult? targetListener,

            string? targetUrl,

            ImmutableArray<Outputs.SubResourceResponseResult> urlPathMaps)
        {
            IncludePath = includePath;
            IncludeQueryString = includeQueryString;
            PathRules = pathRules;
            RedirectType = redirectType;
            RequestRoutingRules = requestRoutingRules;
            TargetListener = targetListener;
            TargetUrl = targetUrl;
            UrlPathMaps = urlPathMaps;
        }
    }
}