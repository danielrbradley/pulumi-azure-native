// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20160515.Outputs
{

    [OutputType]
    public sealed class EnvironmentDeploymentPropertiesResponseResult
    {
        /// <summary>
        /// The Azure Resource Manager template's identifier.
        /// </summary>
        public readonly string? ArmTemplateId;
        /// <summary>
        /// The parameters of the Azure Resource Manager template.
        /// </summary>
        public readonly ImmutableArray<Outputs.ArmTemplateParameterPropertiesResponseResult> Parameters;

        [OutputConstructor]
        private EnvironmentDeploymentPropertiesResponseResult(
            string? armTemplateId,

            ImmutableArray<Outputs.ArmTemplateParameterPropertiesResponseResult> parameters)
        {
            ArmTemplateId = armTemplateId;
            Parameters = parameters;
        }
    }
}