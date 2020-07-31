// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayRewriteRuleConditionResponseResult
    {
        /// <summary>
        /// Setting this paramter to truth value with force the pattern to do a case in-sensitive comparison.
        /// </summary>
        public readonly bool? IgnoreCase;
        /// <summary>
        /// Setting this value as truth will force to check the negation of the condition given by the user.
        /// </summary>
        public readonly bool? Negate;
        /// <summary>
        /// The pattern, either fixed string or regular expression, that evaluates the truthfulness of the condition.
        /// </summary>
        public readonly string? Pattern;
        /// <summary>
        /// The condition parameter of the RewriteRuleCondition.
        /// </summary>
        public readonly string? Variable;

        [OutputConstructor]
        private ApplicationGatewayRewriteRuleConditionResponseResult(
            bool? ignoreCase,

            bool? negate,

            string? pattern,

            string? variable)
        {
            IgnoreCase = ignoreCase;
            Negate = negate;
            Pattern = pattern;
            Variable = variable;
        }
    }
}