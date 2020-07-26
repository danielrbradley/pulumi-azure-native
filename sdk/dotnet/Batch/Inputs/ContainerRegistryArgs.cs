// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.Inputs
{

    public sealed class ContainerRegistryArgs : Pulumi.ResourceArgs
    {
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// If omitted, the default is "docker.io".
        /// </summary>
        [Input("registryServer")]
        public Input<string>? RegistryServer { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ContainerRegistryArgs()
        {
        }
    }
}