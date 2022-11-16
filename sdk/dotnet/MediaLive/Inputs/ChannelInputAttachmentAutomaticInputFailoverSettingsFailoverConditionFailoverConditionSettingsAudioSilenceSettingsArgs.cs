// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsAudioSilenceSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the audio selector used as the source for this AudioDescription.
        /// </summary>
        [Input("audioSelectorName", required: true)]
        public Input<string> AudioSelectorName { get; set; } = null!;

        [Input("audioSilenceThresholdMsec")]
        public Input<int>? AudioSilenceThresholdMsec { get; set; }

        public ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsAudioSilenceSettingsArgs()
        {
        }
        public static new ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsAudioSilenceSettingsArgs Empty => new ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsAudioSilenceSettingsArgs();
    }
}
