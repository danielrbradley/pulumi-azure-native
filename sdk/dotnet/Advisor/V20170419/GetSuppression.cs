// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Advisor.V20170419
{
    public static class GetSuppression
    {
        public static Task<GetSuppressionResult> InvokeAsync(GetSuppressionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSuppressionResult>("azurerm:advisor/v20170419:getSuppression", args ?? new GetSuppressionArgs(), options.WithVersion());
    }


    public sealed class GetSuppressionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the suppression.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The recommendation ID.
        /// </summary>
        [Input("recommendationId", required: true)]
        public string RecommendationId { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetSuppressionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSuppressionResult
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of the suppression.
        /// </summary>
        public readonly Outputs.SuppressionPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSuppressionResult(
            string name,

            Outputs.SuppressionPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}