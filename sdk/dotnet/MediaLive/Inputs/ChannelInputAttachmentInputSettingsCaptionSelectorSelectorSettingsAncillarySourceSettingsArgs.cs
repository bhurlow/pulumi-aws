// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsAncillarySourceSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceAncillaryChannelNumber")]
        public Input<int>? SourceAncillaryChannelNumber { get; set; }

        public ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsAncillarySourceSettingsArgs()
        {
        }
        public static new ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsAncillarySourceSettingsArgs Empty => new ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsAncillarySourceSettingsArgs();
    }
}
