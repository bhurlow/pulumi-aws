// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum bitrate.
        /// </summary>
        [Input("maximumBitrate")]
        public Input<int>? MaximumBitrate { get; set; }

        /// <summary>
        /// Minimum bitrate.
        /// </summary>
        [Input("minimumBitrate")]
        public Input<int>? MinimumBitrate { get; set; }

        /// <summary>
        /// Priority value.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        public MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsGetArgs()
        {
        }
        public static new MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsGetArgs Empty => new MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsGetArgs();
    }
}
