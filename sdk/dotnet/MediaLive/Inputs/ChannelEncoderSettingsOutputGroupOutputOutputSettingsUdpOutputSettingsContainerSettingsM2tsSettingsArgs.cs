// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("absentInputAudioBehavior")]
        public Input<string>? AbsentInputAudioBehavior { get; set; }

        [Input("arib")]
        public Input<string>? Arib { get; set; }

        [Input("aribCaptionsPid")]
        public Input<string>? AribCaptionsPid { get; set; }

        [Input("aribCaptionsPidControl")]
        public Input<string>? AribCaptionsPidControl { get; set; }

        [Input("audioBufferModel")]
        public Input<string>? AudioBufferModel { get; set; }

        [Input("audioFramesPerPes")]
        public Input<int>? AudioFramesPerPes { get; set; }

        [Input("audioPids")]
        public Input<string>? AudioPids { get; set; }

        [Input("audioStreamType")]
        public Input<string>? AudioStreamType { get; set; }

        [Input("bitrate")]
        public Input<int>? Bitrate { get; set; }

        [Input("bufferModel")]
        public Input<string>? BufferModel { get; set; }

        [Input("ccDescriptor")]
        public Input<string>? CcDescriptor { get; set; }

        [Input("dvbNitSettings")]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsArgs>? DvbNitSettings { get; set; }

        [Input("dvbSdtSettings")]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsArgs>? DvbSdtSettings { get; set; }

        [Input("dvbSubPids")]
        public Input<string>? DvbSubPids { get; set; }

        [Input("dvbTdtSettings")]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsArgs>? DvbTdtSettings { get; set; }

        [Input("dvbTeletextPid")]
        public Input<string>? DvbTeletextPid { get; set; }

        [Input("ebif")]
        public Input<string>? Ebif { get; set; }

        [Input("ebpAudioInterval")]
        public Input<string>? EbpAudioInterval { get; set; }

        [Input("ebpLookaheadMs")]
        public Input<int>? EbpLookaheadMs { get; set; }

        [Input("ebpPlacement")]
        public Input<string>? EbpPlacement { get; set; }

        [Input("ecmPid")]
        public Input<string>? EcmPid { get; set; }

        [Input("esRateInPes")]
        public Input<string>? EsRateInPes { get; set; }

        [Input("etvPlatformPid")]
        public Input<string>? EtvPlatformPid { get; set; }

        [Input("etvSignalPid")]
        public Input<string>? EtvSignalPid { get; set; }

        [Input("fragmentTime")]
        public Input<double>? FragmentTime { get; set; }

        [Input("klv")]
        public Input<string>? Klv { get; set; }

        [Input("klvDataPids")]
        public Input<string>? KlvDataPids { get; set; }

        [Input("nielsenId3Behavior")]
        public Input<string>? NielsenId3Behavior { get; set; }

        [Input("nullPacketBitrate")]
        public Input<double>? NullPacketBitrate { get; set; }

        [Input("patInterval")]
        public Input<int>? PatInterval { get; set; }

        [Input("pcrControl")]
        public Input<string>? PcrControl { get; set; }

        [Input("pcrPeriod")]
        public Input<int>? PcrPeriod { get; set; }

        [Input("pcrPid")]
        public Input<string>? PcrPid { get; set; }

        [Input("pmtInterval")]
        public Input<int>? PmtInterval { get; set; }

        [Input("pmtPid")]
        public Input<string>? PmtPid { get; set; }

        [Input("programNum")]
        public Input<int>? ProgramNum { get; set; }

        [Input("rateMode")]
        public Input<string>? RateMode { get; set; }

        [Input("scte27Pids")]
        public Input<string>? Scte27Pids { get; set; }

        [Input("scte35Control")]
        public Input<string>? Scte35Control { get; set; }

        [Input("scte35Pid")]
        public Input<string>? Scte35Pid { get; set; }

        [Input("segmentationMarkers")]
        public Input<string>? SegmentationMarkers { get; set; }

        [Input("segmentationStyle")]
        public Input<string>? SegmentationStyle { get; set; }

        [Input("segmentationTime")]
        public Input<double>? SegmentationTime { get; set; }

        [Input("timedMetadataBehavior")]
        public Input<string>? TimedMetadataBehavior { get; set; }

        [Input("timedMetadataPid")]
        public Input<string>? TimedMetadataPid { get; set; }

        [Input("transportStreamId")]
        public Input<int>? TransportStreamId { get; set; }

        [Input("videoPid")]
        public Input<string>? VideoPid { get; set; }

        public ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsArgs()
        {
        }
        public static new ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsArgs Empty => new ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsArgs();
    }
}
