// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20160301.Outputs
{

    [OutputType]
    public sealed class AlertRuleResponseResult
    {
        /// <summary>
        /// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleActionResponseResult> Actions;
        /// <summary>
        /// the condition that results in the alert rule being activated.
        /// </summary>
        public readonly Outputs.RuleConditionResponseResult Condition;
        /// <summary>
        /// the description of the alert rule that will be included in the alert email.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// the flag that indicates whether the alert rule is enabled.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// Last time the rule was updated in ISO8601 format.
        /// </summary>
        public readonly string LastUpdatedTime;
        /// <summary>
        /// the name of the alert rule.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private AlertRuleResponseResult(
            ImmutableArray<Outputs.RuleActionResponseResult> actions,

            Outputs.RuleConditionResponseResult condition,

            string? description,

            bool isEnabled,

            string lastUpdatedTime,

            string name)
        {
            Actions = actions;
            Condition = condition;
            Description = description;
            IsEnabled = isEnabled;
            LastUpdatedTime = lastUpdatedTime;
            Name = name;
        }
    }
}