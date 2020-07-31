// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20180901
{
    public static class GetActionGroup
    {
        public static Task<GetActionGroupResult> InvokeAsync(GetActionGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetActionGroupResult>("azurerm:insights/v20180901:getActionGroup", args ?? new GetActionGroupArgs(), options.WithVersion());
    }


    public sealed class GetActionGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the action group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetActionGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetActionGroupResult
    {
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The action groups properties of the resource.
        /// </summary>
        public readonly Outputs.ActionGroupResponseResult Properties;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetActionGroupResult(
            string location,

            string name,

            Outputs.ActionGroupResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}