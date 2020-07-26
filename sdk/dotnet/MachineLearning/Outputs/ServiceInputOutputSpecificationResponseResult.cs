// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearning.Outputs
{

    [OutputType]
    public sealed class ServiceInputOutputSpecificationResponseResult
    {
        /// <summary>
        /// The description of the Swagger schema.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Specifies a collection that contains the column schema for each input or output of the web service. For more information, see the Swagger specification.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// The title of your Swagger schema.
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// The type of the entity described in swagger. Always 'object'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ServiceInputOutputSpecificationResponseResult(
            string? description,

            ImmutableDictionary<string, string> properties,

            string? title,

            string type)
        {
            Description = description;
            Properties = properties;
            Title = title;
            Type = type;
        }
    }
}