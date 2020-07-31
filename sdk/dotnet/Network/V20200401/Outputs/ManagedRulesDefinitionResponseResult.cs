// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Outputs
{

    [OutputType]
    public sealed class ManagedRulesDefinitionResponseResult
    {
        /// <summary>
        /// The Exclusions that are applied on the policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.OwaspCrsExclusionEntryResponseResult> Exclusions;
        /// <summary>
        /// The managed rule sets that are associated with the policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedRuleSetResponseResult> ManagedRuleSets;

        [OutputConstructor]
        private ManagedRulesDefinitionResponseResult(
            ImmutableArray<Outputs.OwaspCrsExclusionEntryResponseResult> exclusions,

            ImmutableArray<Outputs.ManagedRuleSetResponseResult> managedRuleSets)
        {
            Exclusions = exclusions;
            ManagedRuleSets = managedRuleSets;
        }
    }
}