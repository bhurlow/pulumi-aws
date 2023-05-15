// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppAutoScaling.Outputs
{

    [OutputType]
    public sealed class PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetric
    {
        /// <summary>
        /// Math expression used on the returned metric. You must specify either `expression` or `metric_stat`, but not both.
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// Short name for the metric used in target tracking scaling policy.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Human-readable label for this metric or expression.
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// Structure that defines CloudWatch metric to be used in target tracking scaling policy. You must specify either `expression` or `metric_stat`, but not both.
        /// </summary>
        public readonly Outputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStat? MetricStat;
        /// <summary>
        /// Boolean that indicates whether to return the timestamps and raw data values of this metric, the default is true
        /// </summary>
        public readonly bool? ReturnData;

        [OutputConstructor]
        private PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetric(
            string? expression,

            string id,

            string? label,

            Outputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStat? metricStat,

            bool? returnData)
        {
            Expression = expression;
            Id = id;
            Label = label;
            MetricStat = metricStat;
            ReturnData = returnData;
        }
    }
}
