// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforPostgreSQL.V20171201.Outputs
{

    [OutputType]
    public sealed class ServerPropertiesResponseResult
    {
        /// <summary>
        /// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// Status showing whether the server data encryption is enabled with customer-managed keys.
        /// </summary>
        public readonly string ByokEnforcement;
        /// <summary>
        /// Earliest restore point creation time (ISO8601 format)
        /// </summary>
        public readonly string? EarliestRestoreDate;
        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        public readonly string? FullyQualifiedDomainName;
        /// <summary>
        /// Status showing whether the server enabled infrastructure encryption.
        /// </summary>
        public readonly string? InfrastructureEncryption;
        /// <summary>
        /// The master server id of a replica server.
        /// </summary>
        public readonly string? MasterServerId;
        /// <summary>
        /// Enforce a minimal Tls version for the server.
        /// </summary>
        public readonly string? MinimalTlsVersion;
        /// <summary>
        /// List of private endpoint connections on a server
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponseResult> PrivateEndpointConnections;
        /// <summary>
        /// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The maximum number of replicas that a master server can have.
        /// </summary>
        public readonly int? ReplicaCapacity;
        /// <summary>
        /// The replication role of the server.
        /// </summary>
        public readonly string? ReplicationRole;
        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        public readonly string? SslEnforcement;
        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        public readonly Outputs.StorageProfileResponseResult? StorageProfile;
        /// <summary>
        /// A state of a server that is visible to user.
        /// </summary>
        public readonly string? UserVisibleState;
        /// <summary>
        /// Server version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ServerPropertiesResponseResult(
            string? administratorLogin,

            string byokEnforcement,

            string? earliestRestoreDate,

            string? fullyQualifiedDomainName,

            string? infrastructureEncryption,

            string? masterServerId,

            string? minimalTlsVersion,

            ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponseResult> privateEndpointConnections,

            string? publicNetworkAccess,

            int? replicaCapacity,

            string? replicationRole,

            string? sslEnforcement,

            Outputs.StorageProfileResponseResult? storageProfile,

            string? userVisibleState,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            ByokEnforcement = byokEnforcement;
            EarliestRestoreDate = earliestRestoreDate;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            InfrastructureEncryption = infrastructureEncryption;
            MasterServerId = masterServerId;
            MinimalTlsVersion = minimalTlsVersion;
            PrivateEndpointConnections = privateEndpointConnections;
            PublicNetworkAccess = publicNetworkAccess;
            ReplicaCapacity = replicaCapacity;
            ReplicationRole = replicationRole;
            SslEnforcement = sslEnforcement;
            StorageProfile = storageProfile;
            UserVisibleState = userVisibleState;
            Version = version;
        }
    }
}