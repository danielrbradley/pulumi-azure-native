// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201.Inputs
{

    /// <summary>
    /// Schema Document Properties.
    /// </summary>
    public sealed class SchemaDocumentPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("definitions")]
        private InputMap<object>? _definitions;

        /// <summary>
        /// Types definitions. Used for Swagger/OpenAPI schemas only, null otherwise.
        /// </summary>
        public InputMap<object> Definitions
        {
            get => _definitions ?? (_definitions = new InputMap<object>());
            set => _definitions = value;
        }

        /// <summary>
        /// Json escaped string defining the document representing the Schema. Used for schemas other than Swagger/OpenAPI.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public SchemaDocumentPropertiesArgs()
        {
        }
    }
}