// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("captureInterval")]
        public Input<int>? CaptureInterval { get; set; }

        [Input("captureIntervalUnits")]
        public Input<string>? CaptureIntervalUnits { get; set; }

        public ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettingsArgs()
        {
        }
        public static new ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettingsArgs Empty => new ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettingsArgs();
    }
}
