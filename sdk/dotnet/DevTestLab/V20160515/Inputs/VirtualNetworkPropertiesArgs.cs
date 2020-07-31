// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20160515.Inputs
{

    /// <summary>
    /// Properties of a virtual network.
    /// </summary>
    public sealed class VirtualNetworkPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("allowedSubnets")]
        private InputList<Inputs.SubnetArgs>? _allowedSubnets;

        /// <summary>
        /// The allowed subnets of the virtual network.
        /// </summary>
        public InputList<Inputs.SubnetArgs> AllowedSubnets
        {
            get => _allowedSubnets ?? (_allowedSubnets = new InputList<Inputs.SubnetArgs>());
            set => _allowedSubnets = value;
        }

        /// <summary>
        /// The description of the virtual network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Microsoft.Network resource identifier of the virtual network.
        /// </summary>
        [Input("externalProviderResourceId")]
        public Input<string>? ExternalProviderResourceId { get; set; }

        [Input("externalSubnets")]
        private InputList<Inputs.ExternalSubnetArgs>? _externalSubnets;

        /// <summary>
        /// The external subnet properties.
        /// </summary>
        public InputList<Inputs.ExternalSubnetArgs> ExternalSubnets
        {
            get => _externalSubnets ?? (_externalSubnets = new InputList<Inputs.ExternalSubnetArgs>());
            set => _externalSubnets = value;
        }

        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        [Input("subnetOverrides")]
        private InputList<Inputs.SubnetOverrideArgs>? _subnetOverrides;

        /// <summary>
        /// The subnet overrides of the virtual network.
        /// </summary>
        public InputList<Inputs.SubnetOverrideArgs> SubnetOverrides
        {
            get => _subnetOverrides ?? (_subnetOverrides = new InputList<Inputs.SubnetOverrideArgs>());
            set => _subnetOverrides = value;
        }

        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        [Input("uniqueIdentifier")]
        public Input<string>? UniqueIdentifier { get; set; }

        public VirtualNetworkPropertiesArgs()
        {
        }
    }
}