// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20190601.Outputs
{

    [OutputType]
    public sealed class ManagementPolicyDefinitionResponseResult
    {
        /// <summary>
        /// An object that defines the action set.
        /// </summary>
        public readonly Outputs.ManagementPolicyActionResponseResult Actions;
        /// <summary>
        /// An object that defines the filter set.
        /// </summary>
        public readonly Outputs.ManagementPolicyFilterResponseResult? Filters;

        [OutputConstructor]
        private ManagementPolicyDefinitionResponseResult(
            Outputs.ManagementPolicyActionResponseResult actions,

            Outputs.ManagementPolicyFilterResponseResult? filters)
        {
            Actions = actions;
            Filters = filters;
        }
    }
}