// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.V20170426.Outputs
{

    [OutputType]
    public sealed class KpiDefinitionResponseResult
    {
        /// <summary>
        /// The aliases.
        /// </summary>
        public readonly ImmutableArray<Outputs.KpiAliasResponseResult> Aliases;
        /// <summary>
        /// The calculation window.
        /// </summary>
        public readonly string CalculationWindow;
        /// <summary>
        /// Name of calculation window field.
        /// </summary>
        public readonly string? CalculationWindowFieldName;
        /// <summary>
        /// Localized description for the KPI.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Description;
        /// <summary>
        /// Localized display name for the KPI.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DisplayName;
        /// <summary>
        /// The mapping entity type.
        /// </summary>
        public readonly string EntityType;
        /// <summary>
        /// The mapping entity name.
        /// </summary>
        public readonly string EntityTypeName;
        /// <summary>
        /// The computation expression for the KPI.
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// The KPI extracts.
        /// </summary>
        public readonly ImmutableArray<Outputs.KpiExtractResponseResult> Extracts;
        /// <summary>
        /// The filter expression for the KPI.
        /// </summary>
        public readonly string? Filter;
        /// <summary>
        /// The computation function for the KPI.
        /// </summary>
        public readonly string Function;
        /// <summary>
        /// the group by properties for the KPI.
        /// </summary>
        public readonly ImmutableArray<string> GroupBy;
        /// <summary>
        /// The KPI GroupByMetadata.
        /// </summary>
        public readonly ImmutableArray<Outputs.KpiGroupByMetadataResponseResult> GroupByMetadata;
        /// <summary>
        /// The KPI name.
        /// </summary>
        public readonly string KpiName;
        /// <summary>
        /// The participant profiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.KpiParticipantProfilesMetadataResponseResult> ParticipantProfilesMetadata;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The hub name.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The KPI thresholds.
        /// </summary>
        public readonly Outputs.KpiThresholdsResponseResult? ThresHolds;
        /// <summary>
        /// The unit of measurement for the KPI.
        /// </summary>
        public readonly string? Unit;

        [OutputConstructor]
        private KpiDefinitionResponseResult(
            ImmutableArray<Outputs.KpiAliasResponseResult> aliases,

            string calculationWindow,

            string? calculationWindowFieldName,

            ImmutableDictionary<string, string>? description,

            ImmutableDictionary<string, string>? displayName,

            string entityType,

            string entityTypeName,

            string expression,

            ImmutableArray<Outputs.KpiExtractResponseResult> extracts,

            string? filter,

            string function,

            ImmutableArray<string> groupBy,

            ImmutableArray<Outputs.KpiGroupByMetadataResponseResult> groupByMetadata,

            string kpiName,

            ImmutableArray<Outputs.KpiParticipantProfilesMetadataResponseResult> participantProfilesMetadata,

            string provisioningState,

            string tenantId,

            Outputs.KpiThresholdsResponseResult? thresHolds,

            string? unit)
        {
            Aliases = aliases;
            CalculationWindow = calculationWindow;
            CalculationWindowFieldName = calculationWindowFieldName;
            Description = description;
            DisplayName = displayName;
            EntityType = entityType;
            EntityTypeName = entityTypeName;
            Expression = expression;
            Extracts = extracts;
            Filter = filter;
            Function = function;
            GroupBy = groupBy;
            GroupByMetadata = groupByMetadata;
            KpiName = kpiName;
            ParticipantProfilesMetadata = participantProfilesMetadata;
            ProvisioningState = provisioningState;
            TenantId = tenantId;
            ThresHolds = thresHolds;
            Unit = unit;
        }
    }
}