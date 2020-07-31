// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20181001.Inputs
{

    /// <summary>
    /// The instance view of a resource.
    /// </summary>
    public sealed class AvailabilitySetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fault Domain count.
        /// </summary>
        [Input("platformFaultDomainCount")]
        public Input<int>? PlatformFaultDomainCount { get; set; }

        /// <summary>
        /// Update Domain count.
        /// </summary>
        [Input("platformUpdateDomainCount")]
        public Input<int>? PlatformUpdateDomainCount { get; set; }

        /// <summary>
        /// Specifies information about the proximity placement group that the availability set should be assigned to. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-04-01.
        /// </summary>
        [Input("proximityPlacementGroup")]
        public Input<Inputs.SubResourceArgs>? ProximityPlacementGroup { get; set; }

        [Input("virtualMachines")]
        private InputList<Inputs.SubResourceArgs>? _virtualMachines;

        /// <summary>
        /// A list of references to all virtual machines in the availability set.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> VirtualMachines
        {
            get => _virtualMachines ?? (_virtualMachines = new InputList<Inputs.SubResourceArgs>());
            set => _virtualMachines = value;
        }

        public AvailabilitySetPropertiesArgs()
        {
        }
    }
}