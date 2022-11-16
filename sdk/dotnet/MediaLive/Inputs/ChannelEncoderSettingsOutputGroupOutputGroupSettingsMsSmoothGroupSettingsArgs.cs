// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("acquisitionPointId")]
        public Input<string>? AcquisitionPointId { get; set; }

        [Input("audioOnlyTimecodecControl")]
        public Input<string>? AudioOnlyTimecodecControl { get; set; }

        /// <summary>
        /// Setting to allow self signed or verified RTMP certificates.
        /// </summary>
        [Input("certificateMode")]
        public Input<string>? CertificateMode { get; set; }

        /// <summary>
        /// Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
        /// </summary>
        [Input("connectionRetryInterval")]
        public Input<int>? ConnectionRetryInterval { get; set; }

        /// <summary>
        /// Destination address and port number for RTP or UDP packets. See Destination for more details.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs> Destination { get; set; } = null!;

        [Input("eventId")]
        public Input<int>? EventId { get; set; }

        [Input("eventIdMode")]
        public Input<string>? EventIdMode { get; set; }

        [Input("eventStopBehavior")]
        public Input<string>? EventStopBehavior { get; set; }

        [Input("filecacheDuration")]
        public Input<int>? FilecacheDuration { get; set; }

        [Input("fragmentLength")]
        public Input<int>? FragmentLength { get; set; }

        /// <summary>
        /// Specifies behavior of last resort when input video os lost.
        /// </summary>
        [Input("inputLossAction")]
        public Input<string>? InputLossAction { get; set; }

        /// <summary>
        /// Number of retry attempts.
        /// </summary>
        [Input("numRetries")]
        public Input<int>? NumRetries { get; set; }

        /// <summary>
        /// Number of seconds to wait until a restart is initiated.
        /// </summary>
        [Input("restartDelay")]
        public Input<int>? RestartDelay { get; set; }

        [Input("segmentationMode")]
        public Input<string>? SegmentationMode { get; set; }

        [Input("sendDelayMs")]
        public Input<int>? SendDelayMs { get; set; }

        [Input("sparseTrackType")]
        public Input<string>? SparseTrackType { get; set; }

        [Input("streamManifestBehavior")]
        public Input<string>? StreamManifestBehavior { get; set; }

        [Input("timestampOffset")]
        public Input<string>? TimestampOffset { get; set; }

        [Input("timestampOffsetMode")]
        public Input<string>? TimestampOffsetMode { get; set; }

        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsArgs()
        {
        }
        public static new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsArgs Empty => new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsArgs();
    }
}
