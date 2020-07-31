// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20170301.Outputs
{

    [OutputType]
    public sealed class LoggerSamplingContractResponseResult
    {
        /// <summary>
        /// Sampling settings entity contract properties.
        /// </summary>
        public readonly Outputs.LoggerSamplingPropertiesResponseResult? Properties;

        [OutputConstructor]
        private LoggerSamplingContractResponseResult(Outputs.LoggerSamplingPropertiesResponseResult? properties)
        {
            Properties = properties;
        }
    }
}