// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate
{
    public static class GetAssessmentProjectHypervcollector
    {
        public static Task<GetAssessmentProjectHypervcollectorResult> InvokeAsync(GetAssessmentProjectHypervcollectorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAssessmentProjectHypervcollectorResult>("azurerm:migrate:getAssessmentProjectHypervcollector", args ?? new GetAssessmentProjectHypervcollectorArgs(), options.WithVersion());
    }


    public sealed class GetAssessmentProjectHypervcollectorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique name of a Hyper-V collector within a project.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public string ProjectName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAssessmentProjectHypervcollectorArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAssessmentProjectHypervcollectorResult
    {
        public readonly string? ETag;
        public readonly string Name;
        public readonly Outputs.CollectorPropertiesResponseResult Properties;
        public readonly string Type;

        [OutputConstructor]
        private GetAssessmentProjectHypervcollectorResult(
            string? eTag,

            string name,

            Outputs.CollectorPropertiesResponseResult properties,

            string type)
        {
            ETag = eTag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}