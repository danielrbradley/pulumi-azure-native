// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180401.Inputs
{

    /// <summary>
    /// Describes the properties of a Proximity Placement Group.
    /// </summary>
    public sealed class ProximityPlacementGroupPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the type of the proximity placement group. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **Standard** : Co-locate resources within an Azure region or Availability Zone. &lt;br&gt;&lt;br&gt; **Ultra** : For future use.
        /// </summary>
        [Input("proximityPlacementGroupType")]
        public Input<string>? ProximityPlacementGroupType { get; set; }

        public ProximityPlacementGroupPropertiesArgs()
        {
        }
    }
}