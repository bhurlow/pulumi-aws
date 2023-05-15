// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelEncoderSettingsVideoDescriptionCodecSettings
    {
        /// <summary>
        /// Frame capture settings. See Frame Capture Settings for more details.
        /// </summary>
        public readonly Outputs.ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettings? FrameCaptureSettings;
        /// <summary>
        /// H264 settings. See H264 Settings for more details.
        /// </summary>
        public readonly Outputs.ChannelEncoderSettingsVideoDescriptionCodecSettingsH264Settings? H264Settings;
        public readonly Outputs.ChannelEncoderSettingsVideoDescriptionCodecSettingsH265Settings? H265Settings;

        [OutputConstructor]
        private ChannelEncoderSettingsVideoDescriptionCodecSettings(
            Outputs.ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettings? frameCaptureSettings,

            Outputs.ChannelEncoderSettingsVideoDescriptionCodecSettingsH264Settings? h264Settings,

            Outputs.ChannelEncoderSettingsVideoDescriptionCodecSettingsH265Settings? h265Settings)
        {
            FrameCaptureSettings = frameCaptureSettings;
            H264Settings = h264Settings;
            H265Settings = h265Settings;
        }
    }
}
