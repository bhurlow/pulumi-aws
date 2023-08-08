// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Outputs
{

    [OutputType]
    public sealed class GetThemeConfigurationUiColorPaletteResult
    {
        /// <summary>
        /// Color (hexadecimal) that applies to selected states and buttons.
        /// </summary>
        public readonly string Accent;
        /// <summary>
        /// Color (hexadecimal) that applies to any text or other elements that appear over the accent color.
        /// </summary>
        public readonly string AccentForeground;
        /// <summary>
        /// Color (hexadecimal) that applies to error messages.
        /// </summary>
        public readonly string Danger;
        /// <summary>
        /// Color (hexadecimal) that applies to any text or other elements that appear over the error color.
        /// </summary>
        public readonly string DangerForeground;
        /// <summary>
        /// Color (hexadecimal) that applies to the names of fields that are identified as dimensions.
        /// </summary>
        public readonly string Dimension;
        /// <summary>
        /// Color (hexadecimal) that applies to any text or other elements that appear over the dimension color.
        /// </summary>
        public readonly string DimensionForeground;
        /// <summary>
        /// Color (hexadecimal) that applies to the names of fields that are identified as measures.
        /// </summary>
        public readonly string Measure;
        /// <summary>
        /// Color (hexadecimal) that applies to any text or other elements that appear over the measure color.
        /// </summary>
        public readonly string MeasureForeground;
        /// <summary>
        /// Color (hexadecimal) that applies to visuals and other high emphasis UI.
        /// </summary>
        public readonly string PrimaryBackground;
        /// <summary>
        /// Color (hexadecimal) of text and other foreground elements that appear over the primary background regions, such as grid lines, borders, table banding, icons, and so on.
        /// </summary>
        public readonly string PrimaryForeground;
        /// <summary>
        /// Color (hexadecimal) that applies to the sheet background and sheet controls.
        /// </summary>
        public readonly string SecondaryBackground;
        /// <summary>
        /// Color (hexadecimal) that applies to any sheet title, sheet control text, or UI that appears over the secondary background.
        /// </summary>
        public readonly string SecondaryForeground;
        /// <summary>
        /// Color (hexadecimal) that applies to success messages, for example the check mark for a successful download.
        /// </summary>
        public readonly string Success;
        /// <summary>
        /// Color (hexadecimal) that applies to any text or other elements that appear over the success color.
        /// </summary>
        public readonly string SuccessForeground;
        /// <summary>
        /// Color (hexadecimal) that applies to warning and informational messages.
        /// </summary>
        public readonly string Warning;
        /// <summary>
        /// Color (hexadecimal) that applies to any text or other elements that appear over the warning color.
        /// </summary>
        public readonly string WarningForeground;

        [OutputConstructor]
        private GetThemeConfigurationUiColorPaletteResult(
            string accent,

            string accentForeground,

            string danger,

            string dangerForeground,

            string dimension,

            string dimensionForeground,

            string measure,

            string measureForeground,

            string primaryBackground,

            string primaryForeground,

            string secondaryBackground,

            string secondaryForeground,

            string success,

            string successForeground,

            string warning,

            string warningForeground)
        {
            Accent = accent;
            AccentForeground = accentForeground;
            Danger = danger;
            DangerForeground = dangerForeground;
            Dimension = dimension;
            DimensionForeground = dimensionForeground;
            Measure = measure;
            MeasureForeground = measureForeground;
            PrimaryBackground = primaryBackground;
            PrimaryForeground = primaryForeground;
            SecondaryBackground = secondaryBackground;
            SecondaryForeground = secondaryForeground;
            Success = success;
            SuccessForeground = successForeground;
            Warning = warning;
            WarningForeground = warningForeground;
        }
    }
}
