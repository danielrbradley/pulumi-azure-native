// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ImportExport.V20161101.Inputs
{

    /// <summary>
    /// Specifies the return address information for the job.
    /// </summary>
    public sealed class ReturnAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The city name to use when returning the drives.
        /// </summary>
        [Input("city", required: true)]
        public Input<string> City { get; set; } = null!;

        /// <summary>
        /// The country or region to use when returning the drives. 
        /// </summary>
        [Input("countryOrRegion", required: true)]
        public Input<string> CountryOrRegion { get; set; } = null!;

        /// <summary>
        /// Email address of the recipient of the returned drives.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// Phone number of the recipient of the returned drives.
        /// </summary>
        [Input("phone", required: true)]
        public Input<string> Phone { get; set; } = null!;

        /// <summary>
        /// The postal code to use when returning the drives.
        /// </summary>
        [Input("postalCode", required: true)]
        public Input<string> PostalCode { get; set; } = null!;

        /// <summary>
        /// The name of the recipient who will receive the hard drives when they are returned. 
        /// </summary>
        [Input("recipientName", required: true)]
        public Input<string> RecipientName { get; set; } = null!;

        /// <summary>
        /// The state or province to use when returning the drives.
        /// </summary>
        [Input("stateOrProvince")]
        public Input<string>? StateOrProvince { get; set; }

        /// <summary>
        /// The first line of the street address to use when returning the drives. 
        /// </summary>
        [Input("streetAddress1", required: true)]
        public Input<string> StreetAddress1 { get; set; } = null!;

        /// <summary>
        /// The second line of the street address to use when returning the drives. 
        /// </summary>
        [Input("streetAddress2")]
        public Input<string>? StreetAddress2 { get; set; }

        public ReturnAddressArgs()
        {
        }
    }
}