// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("outputRectangle")]
        public Input<Inputs.ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsOutputRectangleGetArgs>? OutputRectangle { get; set; }

        [Input("pageNumber")]
        public Input<string>? PageNumber { get; set; }

        public ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsGetArgs()
        {
        }
        public static new ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsGetArgs Empty => new ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsGetArgs();
    }
}
