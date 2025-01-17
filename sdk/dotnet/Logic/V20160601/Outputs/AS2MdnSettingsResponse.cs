// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.V20160601.Outputs
{

    /// <summary>
    /// The AS2 agreement mdn settings.
    /// </summary>
    [OutputType]
    public sealed class AS2MdnSettingsResponse
    {
        /// <summary>
        /// The disposition notification to header value.
        /// </summary>
        public readonly string? DispositionNotificationTo;
        /// <summary>
        /// The MDN text.
        /// </summary>
        public readonly string? MdnText;
        /// <summary>
        /// The signing or hashing algorithm.
        /// </summary>
        public readonly string MicHashingAlgorithm;
        /// <summary>
        /// The value indicating whether to send or request a MDN.
        /// </summary>
        public readonly bool NeedMdn;
        /// <summary>
        /// The receipt delivery URL.
        /// </summary>
        public readonly string? ReceiptDeliveryUrl;
        /// <summary>
        /// The value indicating whether to send inbound MDN to message box.
        /// </summary>
        public readonly bool SendInboundMdnToMessageBox;
        /// <summary>
        /// The value indicating whether to send the asynchronous MDN.
        /// </summary>
        public readonly bool SendMdnAsynchronously;
        /// <summary>
        /// The value indicating whether the MDN needs to be signed or not.
        /// </summary>
        public readonly bool SignMdn;
        /// <summary>
        /// The value indicating whether to sign the outbound MDN if optional.
        /// </summary>
        public readonly bool SignOutboundMdnIfOptional;

        [OutputConstructor]
        private AS2MdnSettingsResponse(
            string? dispositionNotificationTo,

            string? mdnText,

            string micHashingAlgorithm,

            bool needMdn,

            string? receiptDeliveryUrl,

            bool sendInboundMdnToMessageBox,

            bool sendMdnAsynchronously,

            bool signMdn,

            bool signOutboundMdnIfOptional)
        {
            DispositionNotificationTo = dispositionNotificationTo;
            MdnText = mdnText;
            MicHashingAlgorithm = micHashingAlgorithm;
            NeedMdn = needMdn;
            ReceiptDeliveryUrl = receiptDeliveryUrl;
            SendInboundMdnToMessageBox = sendInboundMdnToMessageBox;
            SendMdnAsynchronously = sendMdnAsynchronously;
            SignMdn = signMdn;
            SignOutboundMdnIfOptional = signOutboundMdnIfOptional;
        }
    }
}
