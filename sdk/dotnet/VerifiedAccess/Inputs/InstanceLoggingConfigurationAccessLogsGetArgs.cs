// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.VerifiedAccess.Inputs
{

    public sealed class InstanceLoggingConfigurationAccessLogsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A block that specifies configures sending Verified Access logs to CloudWatch Logs. Detailed below.
        /// </summary>
        [Input("cloudwatchLogs")]
        public Input<Inputs.InstanceLoggingConfigurationAccessLogsCloudwatchLogsGetArgs>? CloudwatchLogs { get; set; }

        /// <summary>
        /// Include trust data sent by trust providers into the logs.
        /// </summary>
        [Input("includeTrustContext")]
        public Input<bool>? IncludeTrustContext { get; set; }

        /// <summary>
        /// A block that specifies configures sending Verified Access logs to Kinesis. Detailed below.
        /// </summary>
        [Input("kinesisDataFirehose")]
        public Input<Inputs.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseGetArgs>? KinesisDataFirehose { get; set; }

        /// <summary>
        /// The logging version to use. Refer to [VerifiedAccessLogOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessLogOptions.html) for the allowed values.
        /// </summary>
        [Input("logVersion")]
        public Input<string>? LogVersion { get; set; }

        /// <summary>
        /// A block that specifies configures sending Verified Access logs to S3. Detailed below.
        /// </summary>
        [Input("s3")]
        public Input<Inputs.InstanceLoggingConfigurationAccessLogsS3GetArgs>? S3 { get; set; }

        public InstanceLoggingConfigurationAccessLogsGetArgs()
        {
        }
        public static new InstanceLoggingConfigurationAccessLogsGetArgs Empty => new InstanceLoggingConfigurationAccessLogsGetArgs();
    }
}
