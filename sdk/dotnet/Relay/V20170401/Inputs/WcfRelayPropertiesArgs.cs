// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Relay.V20170401.Inputs
{

    /// <summary>
    /// Properties of the WCF relay.
    /// </summary>
    public sealed class WcfRelayPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// WCF relay type.
        /// </summary>
        [Input("relayType")]
        public Input<string>? RelayType { get; set; }

        /// <summary>
        /// Returns true if client authorization is needed for this relay; otherwise, false.
        /// </summary>
        [Input("requiresClientAuthorization")]
        public Input<bool>? RequiresClientAuthorization { get; set; }

        /// <summary>
        /// Returns true if transport security is needed for this relay; otherwise, false.
        /// </summary>
        [Input("requiresTransportSecurity")]
        public Input<bool>? RequiresTransportSecurity { get; set; }

        /// <summary>
        /// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
        /// </summary>
        [Input("userMetadata")]
        public Input<string>? UserMetadata { get; set; }

        public WcfRelayPropertiesArgs()
        {
        }
    }
}