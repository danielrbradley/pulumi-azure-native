// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20170901.Inputs
{

    public sealed class LinuxUserConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The uid and gid properties must be specified together or not at all. If not specified the underlying operating system picks the gid.
        /// </summary>
        [Input("gid")]
        public Input<int>? Gid { get; set; }

        /// <summary>
        /// The private key must not be password protected. The private key is used to automatically configure asymmetric-key based authentication for SSH between nodes in a Linux pool when the pool's enableInterNodeCommunication property is true (it is ignored if enableInterNodeCommunication is false). It does this by placing the key pair into the user's .ssh directory. If not specified, password-less SSH is not configured between nodes (no modification of the user's .ssh directory is done).
        /// </summary>
        [Input("sshPrivateKey")]
        public Input<string>? SshPrivateKey { get; set; }

        /// <summary>
        /// The uid and gid properties must be specified together or not at all. If not specified the underlying operating system picks the uid.
        /// </summary>
        [Input("uid")]
        public Input<int>? Uid { get; set; }

        public LinuxUserConfigurationArgs()
        {
        }
    }
}