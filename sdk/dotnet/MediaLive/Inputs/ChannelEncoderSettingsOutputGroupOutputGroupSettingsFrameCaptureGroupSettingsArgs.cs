// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination address and port number for RTP or UDP packets. See Destination for more details.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsDestinationArgs> Destination { get; set; } = null!;

        [Input("frameCaptureCdnSettings")]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettingsArgs>? FrameCaptureCdnSettings { get; set; }

        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsArgs()
        {
        }
        public static new ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsArgs Empty => new ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsArgs();
    }
}
