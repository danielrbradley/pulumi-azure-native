// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180401
{
    public static class GetVirtualMachine
    {
        public static Task<GetVirtualMachineResult> InvokeAsync(GetVirtualMachineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineResult>("azurerm:compute/v20180401:getVirtualMachine", args ?? new GetVirtualMachineArgs(), options.WithVersion());
    }


    public sealed class GetVirtualMachineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualMachineArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineResult
    {
        /// <summary>
        /// The identity of the virtual machine, if configured.
        /// </summary>
        public readonly Outputs.VirtualMachineIdentityResponseResult? Identity;
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
        /// Describes the properties of a Virtual Machine.
        /// </summary>
        public readonly Outputs.VirtualMachinePropertiesResponseResult Properties;
        /// <summary>
        /// The virtual machine child extension resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineExtensionResponseResult> Resources;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The virtual machine zones.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetVirtualMachineResult(
            Outputs.VirtualMachineIdentityResponseResult? identity,

            string location,

            string name,

            Outputs.PlanResponseResult? plan,

            Outputs.VirtualMachinePropertiesResponseResult properties,

            ImmutableArray<Outputs.VirtualMachineExtensionResponseResult> resources,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<string> zones)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Plan = plan;
            Properties = properties;
            Resources = resources;
            Tags = tags;
            Type = type;
            Zones = zones;
        }
    }
}