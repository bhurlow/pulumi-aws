// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.FinSpace.Inputs
{

    public sealed class KxClusterAutoScalingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Metric your cluster will track in order to scale in and out. For example, CPU_UTILIZATION_PERCENTAGE is the average CPU usage across all nodes in a cluster.
        /// </summary>
        [Input("autoScalingMetric", required: true)]
        public Input<string> AutoScalingMetric { get; set; } = null!;

        /// <summary>
        /// Highest number of nodes to scale. Cannot be greater than 5
        /// </summary>
        [Input("maxNodeCount", required: true)]
        public Input<int> MaxNodeCount { get; set; } = null!;

        /// <summary>
        /// Desired value of chosen `auto_scaling_metric`. When metric drops below this value, cluster will scale in. When metric goes above this value, cluster will scale out. Can be set between 0 and 100 percent.
        /// </summary>
        [Input("metricTarget", required: true)]
        public Input<double> MetricTarget { get; set; } = null!;

        /// <summary>
        /// Lowest number of nodes to scale. Must be at least 1 and less than the `max_node_count`. If nodes in cluster belong to multiple availability zones, then `min_node_count` must be at least 3.
        /// </summary>
        [Input("minNodeCount", required: true)]
        public Input<int> MinNodeCount { get; set; } = null!;

        /// <summary>
        /// Duration in seconds that FinSpace will wait after a scale in event before initiating another scaling event.
        /// </summary>
        [Input("scaleInCooldownSeconds", required: true)]
        public Input<double> ScaleInCooldownSeconds { get; set; } = null!;

        /// <summary>
        /// Duration in seconds that FinSpace will wait after a scale out event before initiating another scaling event.
        /// </summary>
        [Input("scaleOutCooldownSeconds", required: true)]
        public Input<double> ScaleOutCooldownSeconds { get; set; } = null!;

        public KxClusterAutoScalingConfigurationArgs()
        {
        }
        public static new KxClusterAutoScalingConfigurationArgs Empty => new KxClusterAutoScalingConfigurationArgs();
    }
}
