// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Inputs
{

    /// <summary>
    /// Describes the properties of a Virtual Machine Scale Set.
    /// </summary>
    public sealed class VirtualMachineScaleSetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
        /// </summary>
        [Input("additionalCapabilities")]
        public Input<Inputs.AdditionalCapabilitiesArgs>? AdditionalCapabilities { get; set; }

        /// <summary>
        /// Policy for automatic repairs.
        /// </summary>
        [Input("automaticRepairsPolicy")]
        public Input<Inputs.AutomaticRepairsPolicyArgs>? AutomaticRepairsPolicy { get; set; }

        /// <summary>
        /// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
        /// </summary>
        [Input("doNotRunExtensionsOnOverprovisionedVMs")]
        public Input<bool>? DoNotRunExtensionsOnOverprovisionedVMs { get; set; }

        /// <summary>
        /// Specifies information about the dedicated host group that the virtual machine scale set resides in. &lt;br&gt;&lt;br&gt;Minimum api-version: 2020-06-01.
        /// </summary>
        [Input("hostGroup")]
        public Input<Inputs.SubResourceArgs>? HostGroup { get; set; }

        /// <summary>
        /// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        /// </summary>
        [Input("overprovision")]
        public Input<bool>? Overprovision { get; set; }

        /// <summary>
        /// Fault Domain count for each placement group.
        /// </summary>
        [Input("platformFaultDomainCount")]
        public Input<int>? PlatformFaultDomainCount { get; set; }

        /// <summary>
        /// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-04-01.
        /// </summary>
        [Input("proximityPlacementGroup")]
        public Input<Inputs.SubResourceArgs>? ProximityPlacementGroup { get; set; }

        /// <summary>
        /// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
        /// </summary>
        [Input("scaleInPolicy")]
        public Input<Inputs.ScaleInPolicyArgs>? ScaleInPolicy { get; set; }

        /// <summary>
        /// When true this limits the scale set to a single placement group, of max size 100 virtual machines. NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may not be modified to true.
        /// </summary>
        [Input("singlePlacementGroup")]
        public Input<bool>? SinglePlacementGroup { get; set; }

        /// <summary>
        /// The upgrade policy.
        /// </summary>
        [Input("upgradePolicy")]
        public Input<Inputs.UpgradePolicyArgs>? UpgradePolicy { get; set; }

        /// <summary>
        /// The virtual machine profile.
        /// </summary>
        [Input("virtualMachineProfile")]
        public Input<Inputs.VirtualMachineScaleSetVMProfileArgs>? VirtualMachineProfile { get; set; }

        /// <summary>
        /// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
        /// </summary>
        [Input("zoneBalance")]
        public Input<bool>? ZoneBalance { get; set; }

        public VirtualMachineScaleSetPropertiesArgs()
        {
        }
    }
}