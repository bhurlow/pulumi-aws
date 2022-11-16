// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("certficateMode")]
        public Input<string>? CertficateMode { get; set; }

        /// <summary>
        /// Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
        /// </summary>
        [Input("connectionRetryInterval")]
        public Input<int>? ConnectionRetryInterval { get; set; }

        /// <summary>
        /// Destination address and port number for RTP or UDP packets. See Destination for more details.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Number of retry attempts.
        /// </summary>
        [Input("numRetries")]
        public Input<int>? NumRetries { get; set; }

        public ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsArgs()
        {
        }
        public static new ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsArgs Empty => new ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsArgs();
    }
}
