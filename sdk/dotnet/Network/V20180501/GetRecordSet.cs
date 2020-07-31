// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180501
{
    public static class GetRecordSet
    {
        public static Task<GetRecordSetResult> InvokeAsync(GetRecordSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRecordSetResult>("azurerm:network/v20180501:getRecordSet", args ?? new GetRecordSetArgs(), options.WithVersion());
    }


    public sealed class GetRecordSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the record set, relative to the name of the zone.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The type of DNS record in this record set.
        /// </summary>
        [Input("recordType", required: true)]
        public string RecordType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the DNS zone (without a terminating dot).
        /// </summary>
        [Input("zoneName", required: true)]
        public string ZoneName { get; set; } = null!;

        public GetRecordSetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRecordSetResult
    {
        /// <summary>
        /// The etag of the record set.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The name of the record set.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of the record set.
        /// </summary>
        public readonly Outputs.RecordSetPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the record set.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRecordSetResult(
            string? etag,

            string name,

            Outputs.RecordSetPropertiesResponseResult properties,

            string type)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}