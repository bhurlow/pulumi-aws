// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("nielsenWatermarksSettings")]
        public Input<Inputs.ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsGetArgs>? NielsenWatermarksSettings { get; set; }

        public ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsGetArgs()
        {
        }
        public static new ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsGetArgs Empty => new ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsGetArgs();
    }
}
