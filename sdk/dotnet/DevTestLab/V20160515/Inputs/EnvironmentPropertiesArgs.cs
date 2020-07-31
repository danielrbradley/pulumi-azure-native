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
    /// Properties of an environment.
    /// </summary>
    public sealed class EnvironmentPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the Azure Resource Manager template that produced the environment.
        /// </summary>
        [Input("armTemplateDisplayName")]
        public Input<string>? ArmTemplateDisplayName { get; set; }

        /// <summary>
        /// The deployment properties of the environment.
        /// </summary>
        [Input("deploymentProperties")]
        public Input<Inputs.EnvironmentDeploymentPropertiesArgs>? DeploymentProperties { get; set; }

        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        [Input("uniqueIdentifier")]
        public Input<string>? UniqueIdentifier { get; set; }

        public EnvironmentPropertiesArgs()
        {
        }
    }
}