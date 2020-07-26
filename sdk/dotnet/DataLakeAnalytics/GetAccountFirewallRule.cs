// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataLakeAnalytics
{
    public static class GetAccountFirewallRule
    {
        public static Task<GetAccountFirewallRuleResult> InvokeAsync(GetAccountFirewallRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountFirewallRuleResult>("azurerm:datalakeanalytics:getAccountFirewallRule", args ?? new GetAccountFirewallRuleArgs(), options.WithVersion());
    }


    public sealed class GetAccountFirewallRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Data Lake Analytics account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the firewall rule to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountFirewallRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountFirewallRuleResult
    {
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The firewall rule properties.
        /// </summary>
        public readonly Outputs.FirewallRulePropertiesResponseResult Properties;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAccountFirewallRuleResult(
            string name,

            Outputs.FirewallRulePropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}