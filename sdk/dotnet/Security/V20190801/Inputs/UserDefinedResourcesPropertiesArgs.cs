// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20190801.Inputs
{

    /// <summary>
    /// Properties of the IoT Security solution's user defined resources.
    /// </summary>
    public sealed class UserDefinedResourcesPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure Resource Graph query which represents the security solution's user defined resources. Required to start with "where type != "Microsoft.Devices/IotHubs""
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        [Input("querySubscriptions", required: true)]
        private InputList<string>? _querySubscriptions;

        /// <summary>
        /// List of Azure subscription ids on which the user defined resources query should be executed.
        /// </summary>
        public InputList<string> QuerySubscriptions
        {
            get => _querySubscriptions ?? (_querySubscriptions = new InputList<string>());
            set => _querySubscriptions = value;
        }

        public UserDefinedResourcesPropertiesArgs()
        {
        }
    }
}