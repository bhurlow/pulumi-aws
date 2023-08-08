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
    public sealed class ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettings
    {
        /// <summary>
        /// If no explicit xPosition or yPosition is provided, setting alignment to centered will place the captions at the bottom center of the output. Similarly, setting a left alignment will align captions to the bottom left of the output. If x and y positions are given in conjunction with the alignment parameter, the font will be justified (either left or centered) relative to those coordinates. Selecting “smart” justification will left-justify live subtitles and center-justify pre-recorded subtitles. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly string? Alignment;
        /// <summary>
        /// Specifies the color of the rectangle behind the captions. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly string? BackgroundColor;
        /// <summary>
        /// Specifies the opacity of the background rectangle. 255 is opaque; 0 is transparent. Leaving this parameter out is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? BackgroundOpacity;
        /// <summary>
        /// External font file used for caption burn-in. File extension must be ‘ttf’ or ‘tte’. Although the user can select output fonts for many different types of input captions, embedded, STL and teletext sources use a strict grid system. Using external fonts with these caption sources could cause unexpected display of proportional fonts. All burn-in and DVB-Sub font settings must match. See Font for more details.
        /// </summary>
        public readonly Outputs.ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFont? Font;
        /// <summary>
        /// Specifies the color of the burned-in captions. This option is not valid for source captions that are STL, 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly string? FontColor;
        /// <summary>
        /// Specifies the opacity of the burned-in captions. 255 is opaque; 0 is transparent. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? FontOpacity;
        /// <summary>
        /// Font resolution in DPI (dots per inch); default is 96 dpi. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? FontResolution;
        /// <summary>
        /// When set to ‘auto’ fontSize will scale depending on the size of the output. Giving a positive integer will specify the exact font size in points. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly string? FontSize;
        /// <summary>
        /// Specifies font outline color. This option is not valid for source captions that are either 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly string OutlineColor;
        /// <summary>
        /// Specifies font outline size in pixels. This option is not valid for source captions that are either 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? OutlineSize;
        /// <summary>
        /// Specifies the color of the shadow cast by the captions. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly string? ShadowColor;
        /// <summary>
        /// Specifies the opacity of the shadow. 255 is opaque; 0 is transparent. Leaving this parameter out is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? ShadowOpacity;
        /// <summary>
        /// Specifies the horizontal offset of the shadow relative to the captions in pixels. A value of -2 would result in a shadow offset 2 pixels to the left. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? ShadowXOffset;
        /// <summary>
        /// Specifies the vertical offset of the shadow relative to the captions in pixels. A value of -2 would result in a shadow offset 2 pixels above the text. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? ShadowYOffset;
        /// <summary>
        /// Controls whether a fixed grid size will be used to generate the output subtitles bitmap. Only applicable for Teletext inputs and DVB-Sub/Burn-in outputs.
        /// </summary>
        public readonly string TeletextGridControl;
        /// <summary>
        /// Specifies the horizontal position of the caption relative to the left side of the output in pixels. A value of 10 would result in the captions starting 10 pixels from the left of the output. If no explicit xPosition is provided, the horizontal caption position will be determined by the alignment parameter. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? XPosition;
        /// <summary>
        /// Specifies the vertical position of the caption relative to the top of the output in pixels. A value of 10 would result in the captions starting 10 pixels from the top of the output. If no explicit yPosition is provided, the caption will be positioned towards the bottom of the output. All burn-in and DVB-Sub font settings must match.
        /// </summary>
        public readonly int? YPosition;

        [OutputConstructor]
        private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettings(
            string? alignment,

            string? backgroundColor,

            int? backgroundOpacity,

            Outputs.ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFont? font,

            string? fontColor,

            int? fontOpacity,

            int? fontResolution,

            string? fontSize,

            string outlineColor,

            int? outlineSize,

            string? shadowColor,

            int? shadowOpacity,

            int? shadowXOffset,

            int? shadowYOffset,

            string teletextGridControl,

            int? xPosition,

            int? yPosition)
        {
            Alignment = alignment;
            BackgroundColor = backgroundColor;
            BackgroundOpacity = backgroundOpacity;
            Font = font;
            FontColor = fontColor;
            FontOpacity = fontOpacity;
            FontResolution = fontResolution;
            FontSize = fontSize;
            OutlineColor = outlineColor;
            OutlineSize = outlineSize;
            ShadowColor = shadowColor;
            ShadowOpacity = shadowOpacity;
            ShadowXOffset = shadowXOffset;
            ShadowYOffset = shadowYOffset;
            TeletextGridControl = teletextGridControl;
            XPosition = xPosition;
            YPosition = yPosition;
        }
    }
}
