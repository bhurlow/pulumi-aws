// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Evidently.Inputs
{

    public sealed class LaunchScheduledSplitsConfigStepSegmentOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a number indicating the order to use to evaluate segment overrides, if there are more than one. Segment overrides with lower numbers are evaluated first.
        /// </summary>
        [Input("evaluationOrder", required: true)]
        public Input<int> EvaluationOrder { get; set; } = null!;

        /// <summary>
        /// The name or ARN of the segment to use.
        /// </summary>
        [Input("segment", required: true)]
        public Input<string> Segment { get; set; } = null!;

        [Input("weights", required: true)]
        private InputMap<int>? _weights;

        /// <summary>
        /// The traffic allocation percentages among the feature variations to assign to this segment. This is a set of key-value pairs. The keys are variation names. The values represent the amount of traffic to allocate to that variation for this segment. This is expressed in thousandths of a percent, so a weight of 50000 represents 50% of traffic.
        /// </summary>
        public InputMap<int> Weights
        {
            get => _weights ?? (_weights = new InputMap<int>());
            set => _weights = value;
        }

        public LaunchScheduledSplitsConfigStepSegmentOverrideArgs()
        {
        }
        public static new LaunchScheduledSplitsConfigStepSegmentOverrideArgs Empty => new LaunchScheduledSplitsConfigStepSegmentOverrideArgs();
    }
}
