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
    public sealed class PushSettingsResponsePropertiesResult
    {
        /// <summary>
        /// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
        /// </summary>
        public readonly string? DynamicTagsJson;
        /// <summary>
        /// Gets or sets a flag indicating whether the Push endpoint is enabled.
        /// </summary>
        public readonly bool IsPushEnabled;
        /// <summary>
        /// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
        /// </summary>
        public readonly string? TagWhitelistJson;
        /// <summary>
        /// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
        /// Tags can consist of alphanumeric characters and the following:
        /// '_', '@', '#', '.', ':', '-'. 
        /// Validation should be performed at the PushRequestHandler.
        /// </summary>
        public readonly string? TagsRequiringAuth;

        [OutputConstructor]
        private PushSettingsResponsePropertiesResult(
            string? dynamicTagsJson,

            bool isPushEnabled,

            string? tagWhitelistJson,

            string? tagsRequiringAuth)
        {
            DynamicTagsJson = dynamicTagsJson;
            IsPushEnabled = isPushEnabled;
            TagWhitelistJson = tagWhitelistJson;
            TagsRequiringAuth = tagsRequiringAuth;
        }
    }
}