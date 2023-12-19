// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB.Inputs
{

    public sealed class LoadBalancerConnectionLogsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The S3 bucket name to store the logs in.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Boolean to enable / disable `connection_logs`. Defaults to `false`, even when `bucket` is specified.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The S3 bucket prefix. Logs are stored in the root if not configured.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public LoadBalancerConnectionLogsArgs()
        {
        }
        public static new LoadBalancerConnectionLogsArgs Empty => new LoadBalancerConnectionLogsArgs();
    }
}
