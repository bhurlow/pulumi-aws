// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsAudioDescriptionCodecSettingsWavSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bitDepth")]
        public Input<double>? BitDepth { get; set; }

        [Input("codingMode")]
        public Input<string>? CodingMode { get; set; }

        [Input("sampleRate")]
        public Input<double>? SampleRate { get; set; }

        public ChannelEncoderSettingsAudioDescriptionCodecSettingsWavSettingsGetArgs()
        {
        }
        public static new ChannelEncoderSettingsAudioDescriptionCodecSettingsWavSettingsGetArgs Empty => new ChannelEncoderSettingsAudioDescriptionCodecSettingsWavSettingsGetArgs();
    }
}
