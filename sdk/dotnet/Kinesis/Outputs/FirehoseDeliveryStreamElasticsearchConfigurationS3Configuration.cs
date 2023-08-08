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
    public sealed class FirehoseDeliveryStreamElasticsearchConfigurationS3Configuration
    {
        /// <summary>
        /// The ARN of the S3 bucket
        /// </summary>
        public readonly string BucketArn;
        /// <summary>
        /// Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
        /// </summary>
        public readonly int? BufferingInterval;
        /// <summary>
        /// Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
        /// We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
        /// </summary>
        public readonly int? BufferingSize;
        /// <summary>
        /// The CloudWatch Logging Options for the delivery stream. More details are given below
        /// </summary>
        public readonly Outputs.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationCloudwatchLoggingOptions? CloudwatchLoggingOptions;
        /// <summary>
        /// The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
        /// </summary>
        public readonly string? CompressionFormat;
        /// <summary>
        /// Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
        /// </summary>
        public readonly string? ErrorOutputPrefix;
        /// <summary>
        /// Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
        /// be used.
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// The ARN of the role that provides access to the source Kinesis stream.
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private FirehoseDeliveryStreamElasticsearchConfigurationS3Configuration(
            string bucketArn,

            int? bufferingInterval,

            int? bufferingSize,

            Outputs.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationCloudwatchLoggingOptions? cloudwatchLoggingOptions,

            string? compressionFormat,

            string? errorOutputPrefix,

            string? kmsKeyArn,

            string? prefix,

            string roleArn)
        {
            BucketArn = bucketArn;
            BufferingInterval = bufferingInterval;
            BufferingSize = bufferingSize;
            CloudwatchLoggingOptions = cloudwatchLoggingOptions;
            CompressionFormat = compressionFormat;
            ErrorOutputPrefix = errorOutputPrefix;
            KmsKeyArn = kmsKeyArn;
            Prefix = prefix;
            RoleArn = roleArn;
        }
    }
}
