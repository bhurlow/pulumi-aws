// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsAudioDescriptionCodecSettingsAacSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("bitrate")]
        public Input<double>? Bitrate { get; set; }

        [Input("codingMode")]
        public Input<string>? CodingMode { get; set; }

        [Input("inputType")]
        public Input<string>? InputType { get; set; }

        [Input("profile")]
        public Input<string>? Profile { get; set; }

        [Input("rawFormat")]
        public Input<string>? RawFormat { get; set; }

        [Input("sampleRate")]
        public Input<double>? SampleRate { get; set; }

        [Input("spec")]
        public Input<string>? Spec { get; set; }

        [Input("vbrQuality")]
        public Input<string>? VbrQuality { get; set; }

        public ChannelEncoderSettingsAudioDescriptionCodecSettingsAacSettingsArgs()
        {
        }
        public static new ChannelEncoderSettingsAudioDescriptionCodecSettingsAacSettingsArgs Empty => new ChannelEncoderSettingsAudioDescriptionCodecSettingsAacSettingsArgs();
    }
}
