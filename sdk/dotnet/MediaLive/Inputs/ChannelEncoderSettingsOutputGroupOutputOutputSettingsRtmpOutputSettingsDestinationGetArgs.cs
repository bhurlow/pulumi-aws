// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsDestinationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reference ID for the destination.
        /// </summary>
        [Input("destinationRefId", required: true)]
        public Input<string> DestinationRefId { get; set; } = null!;

        public ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsDestinationGetArgs()
        {
        }
        public static new ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsDestinationGetArgs Empty => new ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsDestinationGetArgs();
    }
}
