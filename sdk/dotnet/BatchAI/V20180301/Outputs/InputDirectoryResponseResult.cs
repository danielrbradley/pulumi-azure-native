// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20180301.Outputs
{

    [OutputType]
    public sealed class InputDirectoryResponseResult
    {
        /// <summary>
        /// The path of the input directory will be available as a value of an environment variable with AZ_BATCHAI_INPUT_&lt;id&gt; name, where &lt;id&gt; is the value of id attribute.
        /// </summary>
        public readonly string Id;
        public readonly string Path;

        [OutputConstructor]
        private InputDirectoryResponseResult(
            string id,

            string path)
        {
            Id = id;
            Path = path;
        }
    }
}