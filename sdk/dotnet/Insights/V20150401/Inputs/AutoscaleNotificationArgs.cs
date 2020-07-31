// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20150401.Inputs
{

    /// <summary>
    /// Autoscale notification.
    /// </summary>
    public sealed class AutoscaleNotificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the email notification.
        /// </summary>
        [Input("email")]
        public Input<Inputs.EmailNotificationArgs>? Email { get; set; }

        /// <summary>
        /// the operation associated with the notification and its value must be "scale"
        /// </summary>
        [Input("operation", required: true)]
        public Input<string> Operation { get; set; } = null!;

        [Input("webhooks")]
        private InputList<Inputs.WebhookNotificationArgs>? _webhooks;

        /// <summary>
        /// the collection of webhook notifications.
        /// </summary>
        public InputList<Inputs.WebhookNotificationArgs> Webhooks
        {
            get => _webhooks ?? (_webhooks = new InputList<Inputs.WebhookNotificationArgs>());
            set => _webhooks = value;
        }

        public AutoscaleNotificationArgs()
        {
        }
    }
}