// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelEncoderSettingsOutputGroupOutputGroupSettings
    {
        /// <summary>
        /// Archive group settings. See Archive Group Settings for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSetting> ArchiveGroupSettings;
        public readonly Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettings? FrameCaptureGroupSettings;
        public readonly Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettings? HlsGroupSettings;
        /// <summary>
        /// Media package group settings. See Media Package Group Settings for more details.
        /// </summary>
        public readonly Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettings? MediaPackageGroupSettings;
        public readonly Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettings? MsSmoothGroupSettings;
        public readonly Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMultiplexGroupSettings? MultiplexGroupSettings;
        /// <summary>
        /// RTMP group settings. See RTMP Group Settings for more details.
        /// </summary>
        public readonly Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettings? RtmpGroupSettings;
        public readonly Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsUdpGroupSettings? UdpGroupSettings;

        [OutputConstructor]
        private ChannelEncoderSettingsOutputGroupOutputGroupSettings(
            ImmutableArray<Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSetting> archiveGroupSettings,

            Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettings? frameCaptureGroupSettings,

            Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettings? hlsGroupSettings,

            Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettings? mediaPackageGroupSettings,

            Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettings? msSmoothGroupSettings,

            Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMultiplexGroupSettings? multiplexGroupSettings,

            Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettings? rtmpGroupSettings,

            Outputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsUdpGroupSettings? udpGroupSettings)
        {
            ArchiveGroupSettings = archiveGroupSettings;
            FrameCaptureGroupSettings = frameCaptureGroupSettings;
            HlsGroupSettings = hlsGroupSettings;
            MediaPackageGroupSettings = mediaPackageGroupSettings;
            MsSmoothGroupSettings = msSmoothGroupSettings;
            MultiplexGroupSettings = multiplexGroupSettings;
            RtmpGroupSettings = rtmpGroupSettings;
            UdpGroupSettings = udpGroupSettings;
        }
    }
}
