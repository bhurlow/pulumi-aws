// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsFilterSettingsTemporalFilterSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Post filter sharpening.
        /// </summary>
        [Input("postFilterSharpening")]
        public Input<string>? PostFilterSharpening { get; set; }

        /// <summary>
        /// Filter strength.
        /// </summary>
        [Input("strength")]
        public Input<string>? Strength { get; set; }

        public ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsFilterSettingsTemporalFilterSettingsGetArgs()
        {
        }
        public static new ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsFilterSettingsTemporalFilterSettingsGetArgs Empty => new ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsFilterSettingsTemporalFilterSettingsGetArgs();
    }
}
