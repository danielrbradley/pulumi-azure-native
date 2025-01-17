// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.V20180701.Outputs
{

    /// <summary>
    /// Describes all the settings to be used when analyzing a video in order to detect all the faces present.
    /// </summary>
    [OutputType]
    public sealed class FaceDetectorPresetResponse
    {
        /// <summary>
        /// Dictionary containing key value pairs for parameters not exposed in the preset itself
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ExperimentalOptions;
        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.FaceDetectorPreset'.
        /// </summary>
        public readonly string OdataType;
        /// <summary>
        /// Specifies the maximum resolution at which your video is analyzed. The default behavior is "SourceResolution," which will keep the input video at its original resolution when analyzed. Using "StandardDefinition" will resize input videos to standard definition while preserving the appropriate aspect ratio. It will only resize if the video is of higher resolution. For example, a 1920x1080 input would be scaled to 640x360 before processing. Switching to "StandardDefinition" will reduce the time it takes to process high resolution video. It may also reduce the cost of using this component (see https://azure.microsoft.com/en-us/pricing/details/media-services/#analytics for details). However, faces that end up being too small in the resized video may not be detected.
        /// </summary>
        public readonly string? Resolution;

        [OutputConstructor]
        private FaceDetectorPresetResponse(
            ImmutableDictionary<string, string>? experimentalOptions,

            string odataType,

            string? resolution)
        {
            ExperimentalOptions = experimentalOptions;
            OdataType = odataType;
            Resolution = resolution;
        }
    }
}
