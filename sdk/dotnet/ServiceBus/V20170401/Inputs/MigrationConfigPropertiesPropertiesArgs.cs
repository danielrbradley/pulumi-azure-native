// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus.V20170401.Inputs
{

    /// <summary>
    /// Properties required to the Create Migration Configuration
    /// </summary>
    public sealed class MigrationConfigPropertiesPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name to access Standard Namespace after migration
        /// </summary>
        [Input("postMigrationName", required: true)]
        public Input<string> PostMigrationName { get; set; } = null!;

        /// <summary>
        /// Existing premium Namespace ARM Id name which has no entities, will be used for migration
        /// </summary>
        [Input("targetNamespace", required: true)]
        public Input<string> TargetNamespace { get; set; } = null!;

        public MigrationConfigPropertiesPropertiesArgs()
        {
        }
    }
}