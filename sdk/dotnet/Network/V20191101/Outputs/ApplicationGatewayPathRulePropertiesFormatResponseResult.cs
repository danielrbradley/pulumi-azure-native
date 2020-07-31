// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayPathRulePropertiesFormatResponseResult
    {
        /// <summary>
        /// Backend address pool resource of URL path map path rule.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? BackendAddressPool;
        /// <summary>
        /// Backend http settings resource of URL path map path rule.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? BackendHttpSettings;
        /// <summary>
        /// Reference to the FirewallPolicy resource.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? FirewallPolicy;
        /// <summary>
        /// Path rules of URL path map.
        /// </summary>
        public readonly ImmutableArray<string> Paths;
        /// <summary>
        /// The provisioning state of the path rule resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Redirect configuration resource of URL path map path rule.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? RedirectConfiguration;
        /// <summary>
        /// Rewrite rule set resource of URL path map path rule.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? RewriteRuleSet;

        [OutputConstructor]
        private ApplicationGatewayPathRulePropertiesFormatResponseResult(
            Outputs.SubResourceResponseResult? backendAddressPool,

            Outputs.SubResourceResponseResult? backendHttpSettings,

            Outputs.SubResourceResponseResult? firewallPolicy,

            ImmutableArray<string> paths,

            string provisioningState,

            Outputs.SubResourceResponseResult? redirectConfiguration,

            Outputs.SubResourceResponseResult? rewriteRuleSet)
        {
            BackendAddressPool = backendAddressPool;
            BackendHttpSettings = backendHttpSettings;
            FirewallPolicy = firewallPolicy;
            Paths = paths;
            ProvisioningState = provisioningState;
            RedirectConfiguration = redirectConfiguration;
            RewriteRuleSet = rewriteRuleSet;
        }
    }
}