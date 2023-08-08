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
    public sealed class ChannelEncoderSettingsCaptionDescription
    {
        /// <summary>
        /// Indicates whether the caption track implements accessibility features such as written descriptions of spoken dialog, music, and sounds.
        /// </summary>
        public readonly string? Accessibility;
        /// <summary>
        /// Specifies which input caption selector to use as a caption source when generating output captions. This field should match a captionSelector name.
        /// </summary>
        public readonly string CaptionSelectorName;
        /// <summary>
        /// Additional settings for captions destination that depend on the destination type. See Destination Settings for more details.
        /// </summary>
        public readonly Outputs.ChannelEncoderSettingsCaptionDescriptionDestinationSettings? DestinationSettings;
        /// <summary>
        /// ISO 639-2 three-digit code.
        /// </summary>
        public readonly string? LanguageCode;
        /// <summary>
        /// Human readable information to indicate captions available for players (eg. English, or Spanish).
        /// </summary>
        public readonly string? LanguageDescription;
        /// <summary>
        /// Name of the caption description. Used to associate a caption description with an output. Names must be unique within an event.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ChannelEncoderSettingsCaptionDescription(
            string? accessibility,

            string captionSelectorName,

            Outputs.ChannelEncoderSettingsCaptionDescriptionDestinationSettings? destinationSettings,

            string? languageCode,

            string? languageDescription,

            string name)
        {
            Accessibility = accessibility;
            CaptionSelectorName = captionSelectorName;
            DestinationSettings = destinationSettings;
            LanguageCode = languageCode;
            LanguageDescription = languageDescription;
            Name = name;
        }
    }
}
