// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Inputs
{

    /// <summary>
    /// The JSON object that contains the properties required to create a Rules Engine Configuration.
    /// </summary>
    public sealed class RulesEnginePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource status.
        /// </summary>
        [Input("resourceState")]
        public Input<string>? ResourceState { get; set; }

        [Input("rules")]
        private InputList<Inputs.RulesEngineRuleArgs>? _rules;

        /// <summary>
        /// A list of rules that define a particular Rules Engine Configuration.
        /// </summary>
        public InputList<Inputs.RulesEngineRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RulesEngineRuleArgs>());
            set => _rules = value;
        }

        public RulesEnginePropertiesArgs()
        {
        }
    }
}