// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GlobalAccelerator.Outputs
{

    [OutputType]
    public sealed class CustomRoutingAcceleratorAttributes
    {
        /// <summary>
        /// Indicates whether flow logs are enabled. Defaults to `false`. Valid values: `true`, `false`.
        /// </summary>
        public readonly bool? FlowLogsEnabled;
        /// <summary>
        /// The name of the Amazon S3 bucket for the flow logs. Required if `flow_logs_enabled` is `true`.
        /// </summary>
        public readonly string? FlowLogsS3Bucket;
        /// <summary>
        /// The prefix for the location in the Amazon S3 bucket for the flow logs. Required if `flow_logs_enabled` is `true`.
        /// </summary>
        public readonly string? FlowLogsS3Prefix;

        [OutputConstructor]
        private CustomRoutingAcceleratorAttributes(
            bool? flowLogsEnabled,

            string? flowLogsS3Bucket,

            string? flowLogsS3Prefix)
        {
            FlowLogsEnabled = flowLogsEnabled;
            FlowLogsS3Bucket = flowLogsS3Bucket;
            FlowLogsS3Prefix = flowLogsS3Prefix;
        }
    }
}
