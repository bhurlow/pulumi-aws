// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Outputs
{

    [OutputType]
    public sealed class FirehoseDeliveryStreamHttpEndpointConfiguration
    {
        /// <summary>
        /// The access key required for Kinesis Firehose to authenticate with the HTTP endpoint selected as the destination.
        /// </summary>
        public readonly string? AccessKey;
        /// <summary>
        /// Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination. The default value is 300 (5 minutes).
        /// </summary>
        public readonly int? BufferingInterval;
        /// <summary>
        /// Buffer incoming data to the specified size, in MBs, before delivering it to the destination. The default value is 5.
        /// </summary>
        public readonly int? BufferingSize;
        /// <summary>
        /// The CloudWatch Logging Options for the delivery stream. More details are given below.
        /// </summary>
        public readonly Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationCloudwatchLoggingOptions? CloudwatchLoggingOptions;
        /// <summary>
        /// The HTTP endpoint name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The data processing configuration.  More details are given below.
        /// </summary>
        public readonly Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfiguration? ProcessingConfiguration;
        /// <summary>
        /// The request configuration.  More details are given below.
        /// </summary>
        public readonly Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfiguration? RequestConfiguration;
        /// <summary>
        /// Total amount of seconds Firehose spends on retries. This duration starts after the initial attempt fails, It does not include the time periods during which Firehose waits for acknowledgment from the specified destination after each attempt. Valid values between `0` and `7200`. Default is `300`.
        /// </summary>
        public readonly int? RetryDuration;
        /// <summary>
        /// Kinesis Data Firehose uses this IAM role for all the permissions that the delivery stream needs. The pattern needs to be `arn:.*`.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// Defines how documents should be delivered to Amazon S3.  Valid values are `FailedDataOnly` and `AllData`.  Default value is `FailedDataOnly`.
        /// </summary>
        public readonly string? S3BackupMode;
        /// <summary>
        /// The S3 Configuration. See s3_configuration for more details.
        /// </summary>
        public readonly Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationS3Configuration S3Configuration;
        /// <summary>
        /// The HTTP endpoint URL to which Kinesis Firehose sends your data.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private FirehoseDeliveryStreamHttpEndpointConfiguration(
            string? accessKey,

            int? bufferingInterval,

            int? bufferingSize,

            Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationCloudwatchLoggingOptions? cloudwatchLoggingOptions,

            string? name,

            Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfiguration? processingConfiguration,

            Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfiguration? requestConfiguration,

            int? retryDuration,

            string? roleArn,

            string? s3BackupMode,

            Outputs.FirehoseDeliveryStreamHttpEndpointConfigurationS3Configuration s3Configuration,

            string url)
        {
            AccessKey = accessKey;
            BufferingInterval = bufferingInterval;
            BufferingSize = bufferingSize;
            CloudwatchLoggingOptions = cloudwatchLoggingOptions;
            Name = name;
            ProcessingConfiguration = processingConfiguration;
            RequestConfiguration = requestConfiguration;
            RetryDuration = retryDuration;
            RoleArn = roleArn;
            S3BackupMode = s3BackupMode;
            S3Configuration = s3Configuration;
            Url = url;
        }
    }
}
