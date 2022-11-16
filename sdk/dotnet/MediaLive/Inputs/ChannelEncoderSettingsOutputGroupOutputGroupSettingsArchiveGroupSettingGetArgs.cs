// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameters that control the interactions with the CDN. See Archive CDN Settings for more details.
        /// </summary>
        [Input("archiveCdnSettings")]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsGetArgs>? ArchiveCdnSettings { get; set; }

        /// <summary>
        /// Destination address and port number for RTP or UDP packets. See Destination for more details.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestinationGetArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Number of seconds to write to archive file before closing and starting a new one.
        /// </summary>
        [Input("rolloverInterval")]
        public Input<int>? RolloverInterval { get; set; }

        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingGetArgs()
        {
        }
        public static new ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingGetArgs Empty => new ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingGetArgs();
    }
}
