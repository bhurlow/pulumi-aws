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
    public sealed class ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsFilterSettingsTemporalFilterSettings
    {
        /// <summary>
        /// Post filter sharpening.
        /// </summary>
        public readonly string? PostFilterSharpening;
        /// <summary>
        /// Filter strength.
        /// </summary>
        public readonly string? Strength;

        [OutputConstructor]
        private ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsFilterSettingsTemporalFilterSettings(
            string? postFilterSharpening,

            string? strength)
        {
            PostFilterSharpening = postFilterSharpening;
            Strength = strength;
        }
    }
}
