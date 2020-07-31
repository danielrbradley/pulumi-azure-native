// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20160319.Outputs
{

    [OutputType]
    public sealed class CassandraKeyspacePropertiesResponseResult
    {
        /// <summary>
        /// Name of the Cosmos DB Cassandra keyspace
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private CassandraKeyspacePropertiesResponseResult(string id)
        {
            Id = id;
        }
    }
}