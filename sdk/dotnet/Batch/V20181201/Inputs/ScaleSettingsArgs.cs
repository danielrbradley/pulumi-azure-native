// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20181201.Inputs
{

    /// <summary>
    /// Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
    /// </summary>
    public sealed class ScaleSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This property and fixedScale are mutually exclusive and one of the properties must be specified.
        /// </summary>
        [Input("autoScale")]
        public Input<Inputs.AutoScaleSettingsArgs>? AutoScale { get; set; }

        /// <summary>
        /// This property and autoScale are mutually exclusive and one of the properties must be specified.
        /// </summary>
        [Input("fixedScale")]
        public Input<Inputs.FixedScaleSettingsArgs>? FixedScale { get; set; }

        public ScaleSettingsArgs()
        {
        }
    }
}