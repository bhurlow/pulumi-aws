// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsNielsenNaesIiNwSettingGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("checkDigitString", required: true)]
        public Input<string> CheckDigitString { get; set; } = null!;

        /// <summary>
        /// The Nielsen Source ID to include in the watermark.
        /// </summary>
        [Input("sid", required: true)]
        public Input<double> Sid { get; set; } = null!;

        public ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsNielsenNaesIiNwSettingGetArgs()
        {
        }
        public static new ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsNielsenNaesIiNwSettingGetArgs Empty => new ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsNielsenNaesIiNwSettingGetArgs();
    }
}
