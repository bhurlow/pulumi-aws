// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ivs.Inputs
{

    public sealed class RecordingConfigurationThumbnailConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Thumbnail recording mode. Valid values: `DISABLED`, `INTERVAL`.
        /// </summary>
        [Input("recordingMode")]
        public Input<string>? RecordingMode { get; set; }

        /// <summary>
        /// - The targeted thumbnail-generation interval in seconds.
        /// </summary>
        [Input("targetIntervalSeconds")]
        public Input<int>? TargetIntervalSeconds { get; set; }

        public RecordingConfigurationThumbnailConfigurationArgs()
        {
        }
        public static new RecordingConfigurationThumbnailConfigurationArgs Empty => new RecordingConfigurationThumbnailConfigurationArgs();
    }
}
