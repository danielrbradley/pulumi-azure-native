// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StreamAnalytics.Outputs
{

    [OutputType]
    public sealed class OutputResponseResult
    {
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The properties that are associated with an output. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        public readonly Outputs.OutputPropertiesResponseResult? Properties;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private OutputResponseResult(
            string id,

            string? name,

            Outputs.OutputPropertiesResponseResult? properties,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}