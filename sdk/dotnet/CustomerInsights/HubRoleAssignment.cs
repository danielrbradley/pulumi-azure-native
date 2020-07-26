// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights
{
    /// <summary>
    /// The Role Assignment resource format.
    /// </summary>
    public partial class HubRoleAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Role Assignment definition.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.RoleAssignmentResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a HubRoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HubRoleAssignment(string name, HubRoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights:HubRoleAssignment", name, args ?? new HubRoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HubRoleAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights:HubRoleAssignment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HubRoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HubRoleAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HubRoleAssignment(name, id, options);
        }
    }

    public sealed class HubRoleAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// The assignment name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Role Assignment definition.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.RoleAssignmentArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public HubRoleAssignmentArgs()
        {
        }
    }
}