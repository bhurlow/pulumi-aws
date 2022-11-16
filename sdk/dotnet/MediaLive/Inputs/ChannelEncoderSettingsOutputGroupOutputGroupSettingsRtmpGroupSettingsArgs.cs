// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("adMarkers")]
        private InputList<string>? _adMarkers;

        /// <summary>
        /// The ad marker type for this output group.
        /// </summary>
        public InputList<string> AdMarkers
        {
            get => _adMarkers ?? (_adMarkers = new InputList<string>());
            set => _adMarkers = value;
        }

        /// <summary>
        /// Authentication scheme to use when connecting with CDN.
        /// </summary>
        [Input("authenticationScheme")]
        public Input<string>? AuthenticationScheme { get; set; }

        /// <summary>
        /// Controls behavior when content cache fills up.
        /// </summary>
        [Input("cacheFullBehavior")]
        public Input<string>? CacheFullBehavior { get; set; }

        /// <summary>
        /// Cache length in seconds, is used to calculate buffer size.
        /// </summary>
        [Input("cacheLength")]
        public Input<int>? CacheLength { get; set; }

        /// <summary>
        /// Controls the types of data that passes to onCaptionInfo outputs.
        /// </summary>
        [Input("captionData")]
        public Input<string>? CaptionData { get; set; }

        /// <summary>
        /// Specifies behavior of last resort when input video os lost.
        /// </summary>
        [Input("inputLossAction")]
        public Input<string>? InputLossAction { get; set; }

        /// <summary>
        /// Number of seconds to wait until a restart is initiated.
        /// </summary>
        [Input("restartDelay")]
        public Input<int>? RestartDelay { get; set; }

        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettingsArgs()
        {
        }
        public static new ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettingsArgs Empty => new ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettingsArgs();
    }
}
