// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801.Inputs
{

    public sealed class CsrPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Actual CSR string created
        /// </summary>
        [Input("csrString")]
        public Input<string>? CsrString { get; set; }

        /// <summary>
        /// Distinguished name of certificate to be created
        /// </summary>
        [Input("distinguishedName")]
        public Input<string>? DistinguishedName { get; set; }

        /// <summary>
        /// Hosting environment
        /// </summary>
        [Input("hostingEnvironment")]
        public Input<string>? HostingEnvironment { get; set; }

        /// <summary>
        /// Name used to locate CSR object
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// PFX password
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// PFX certificate of created certificate
        /// </summary>
        [Input("pfxBlob")]
        public Input<string>? PfxBlob { get; set; }

        /// <summary>
        /// Hash of the certificates public key
        /// </summary>
        [Input("publicKeyHash")]
        public Input<string>? PublicKeyHash { get; set; }

        public CsrPropertiesArgs()
        {
        }
    }
}