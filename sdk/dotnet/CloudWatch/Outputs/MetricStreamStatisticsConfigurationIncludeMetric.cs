// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Outputs
{

    [OutputType]
    public sealed class MetricStreamStatisticsConfigurationIncludeMetric
    {
        /// <summary>
        /// The name of the metric.
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// The namespace of the metric.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private MetricStreamStatisticsConfigurationIncludeMetric(
            string metricName,

            string @namespace)
        {
            MetricName = metricName;
            Namespace = @namespace;
        }
    }
}
