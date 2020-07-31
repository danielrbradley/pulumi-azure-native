// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180601
{
    public static class GetVirtualMachineScaleSet
    {
        public static Task<GetVirtualMachineScaleSetResult> InvokeAsync(GetVirtualMachineScaleSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineScaleSetResult>("azurerm:compute/v20180601:getVirtualMachineScaleSet", args ?? new GetVirtualMachineScaleSetArgs(), options.WithVersion());
    }


    public sealed class GetVirtualMachineScaleSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the VM scale set.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualMachineScaleSetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineScaleSetResult
    {
        /// <summary>
        /// The identity of the virtual machine scale set, if configured.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetIdentityResponseResult? Identity;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        public readonly Outputs.PlanResponseResult? Plan;
        /// <summary>
        /// Describes the properties of a Virtual Machine Scale Set.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetPropertiesResponseResult Properties;
        /// <summary>
        /// The virtual machine scale set sku.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetVirtualMachineScaleSetResult(
            Outputs.VirtualMachineScaleSetIdentityResponseResult? identity,

            string location,

            string name,

            Outputs.PlanResponseResult? plan,

            Outputs.VirtualMachineScaleSetPropertiesResponseResult properties,

            Outputs.SkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<string> zones)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Plan = plan;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
            Zones = zones;
        }
    }
}