// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelInputAttachmentAutomaticInputFailoverSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This clear time defines the requirement a recovered input must meet to be considered healthy. The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input\_preference for the failover pair is set to PRIMARY\_INPUT\_PREFERRED, because after this time, MediaLive will switch back to the primary input.
        /// </summary>
        [Input("errorClearTimeMsec")]
        public Input<int>? ErrorClearTimeMsec { get; set; }

        [Input("failoverConditions")]
        private InputList<Inputs.ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionGetArgs>? _failoverConditions;

        /// <summary>
        /// A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input. See Failover Condition Block for more details.
        /// </summary>
        public InputList<Inputs.ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionGetArgs> FailoverConditions
        {
            get => _failoverConditions ?? (_failoverConditions = new InputList<Inputs.ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionGetArgs>());
            set => _failoverConditions = value;
        }

        /// <summary>
        /// Input preference when deciding which input to make active when a previously failed input has recovered.
        /// </summary>
        [Input("inputPreference")]
        public Input<string>? InputPreference { get; set; }

        /// <summary>
        /// The input ID of the secondary input in the automatic input failover pair.
        /// </summary>
        [Input("secondaryInputId", required: true)]
        public Input<string> SecondaryInputId { get; set; } = null!;

        public ChannelInputAttachmentAutomaticInputFailoverSettingsGetArgs()
        {
        }
        public static new ChannelInputAttachmentAutomaticInputFailoverSettingsGetArgs Empty => new ChannelInputAttachmentAutomaticInputFailoverSettingsGetArgs();
    }
}
