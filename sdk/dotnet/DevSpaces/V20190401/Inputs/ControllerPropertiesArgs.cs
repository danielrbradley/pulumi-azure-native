// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevSpaces.V20190401.Inputs
{

    public sealed class ControllerPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Credentials of the target container host (base64).
        /// </summary>
        [Input("targetContainerHostCredentialsBase64", required: true)]
        public Input<string> TargetContainerHostCredentialsBase64 { get; set; } = null!;

        /// <summary>
        /// Resource ID of the target container host
        /// </summary>
        [Input("targetContainerHostResourceId", required: true)]
        public Input<string> TargetContainerHostResourceId { get; set; } = null!;

        public ControllerPropertiesArgs()
        {
        }
    }
}