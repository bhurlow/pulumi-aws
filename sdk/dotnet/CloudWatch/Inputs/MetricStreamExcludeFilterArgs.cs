// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Inputs
{

    public sealed class MetricStreamExcludeFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("metricNames")]
        private InputList<string>? _metricNames;

        /// <summary>
        /// An array that defines the metrics you want to exclude for this metric namespace
        /// </summary>
        public InputList<string> MetricNames
        {
            get => _metricNames ?? (_metricNames = new InputList<string>());
            set => _metricNames = value;
        }

        /// <summary>
        /// Name of the metric namespace in the filter.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public MetricStreamExcludeFilterArgs()
        {
        }
        public static new MetricStreamExcludeFilterArgs Empty => new MetricStreamExcludeFilterArgs();
    }
}
