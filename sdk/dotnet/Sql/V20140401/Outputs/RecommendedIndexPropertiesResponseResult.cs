// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20140401.Outputs
{

    [OutputType]
    public sealed class RecommendedIndexPropertiesResponseResult
    {
        /// <summary>
        /// The proposed index action. You can create a missing index, drop an unused index, or rebuild an existing index to improve its performance.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Columns over which to build index
        /// </summary>
        public readonly ImmutableArray<string> Columns;
        /// <summary>
        /// The UTC datetime showing when this resource was created (ISO8601 format).
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// The estimated impact of doing recommended index action.
        /// </summary>
        public readonly ImmutableArray<Outputs.OperationImpactResponseResult> EstimatedImpact;
        /// <summary>
        /// The list of column names to be included in the index
        /// </summary>
        public readonly ImmutableArray<string> IncludedColumns;
        /// <summary>
        /// The full build index script
        /// </summary>
        public readonly string IndexScript;
        /// <summary>
        /// The type of index (CLUSTERED, NONCLUSTERED, COLUMNSTORE, CLUSTERED COLUMNSTORE)
        /// </summary>
        public readonly string IndexType;
        /// <summary>
        /// The UTC datetime of when was this resource last changed (ISO8601 format).
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// The values reported after index action is complete.
        /// </summary>
        public readonly ImmutableArray<Outputs.OperationImpactResponseResult> ReportedImpact;
        /// <summary>
        /// The schema where table to build index over resides
        /// </summary>
        public readonly string Schema;
        /// <summary>
        /// The current recommendation state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The table on which to build index.
        /// </summary>
        public readonly string Table;

        [OutputConstructor]
        private RecommendedIndexPropertiesResponseResult(
            string action,

            ImmutableArray<string> columns,

            string created,

            ImmutableArray<Outputs.OperationImpactResponseResult> estimatedImpact,

            ImmutableArray<string> includedColumns,

            string indexScript,

            string indexType,

            string lastModified,

            ImmutableArray<Outputs.OperationImpactResponseResult> reportedImpact,

            string schema,

            string state,

            string table)
        {
            Action = action;
            Columns = columns;
            Created = created;
            EstimatedImpact = estimatedImpact;
            IncludedColumns = includedColumns;
            IndexScript = indexScript;
            IndexType = indexType;
            LastModified = lastModified;
            ReportedImpact = reportedImpact;
            Schema = schema;
            State = state;
            Table = table;
        }
    }
}