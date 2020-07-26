// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge
{
    public static class GetDataBoxEdgeDeviceUser
    {
        public static Task<GetDataBoxEdgeDeviceUserResult> InvokeAsync(GetDataBoxEdgeDeviceUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataBoxEdgeDeviceUserResult>("azurerm:databoxedge:getDataBoxEdgeDeviceUser", args ?? new GetDataBoxEdgeDeviceUserArgs(), options.WithVersion());
    }


    public sealed class GetDataBoxEdgeDeviceUserArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public string DeviceName { get; set; } = null!;

        /// <summary>
        /// The user name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDataBoxEdgeDeviceUserArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataBoxEdgeDeviceUserResult
    {
        /// <summary>
        /// The object name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The storage account credential properties.
        /// </summary>
        public readonly Outputs.UserPropertiesResponseResult Properties;
        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDataBoxEdgeDeviceUserResult(
            string name,

            Outputs.UserPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}