// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Inputs
{

    public sealed class FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the S3 bucket
        /// </summary>
        [Input("bucketArn", required: true)]
        public Input<string> BucketArn { get; set; } = null!;

        /// <summary>
        /// Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
        /// </summary>
        [Input("bufferingInterval")]
        public Input<int>? BufferingInterval { get; set; }

        /// <summary>
        /// Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
        /// We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
        /// </summary>
        [Input("bufferingSize")]
        public Input<int>? BufferingSize { get; set; }

        /// <summary>
        /// The CloudWatch Logging Options for the delivery stream. More details are given below
        /// </summary>
        [Input("cloudwatchLoggingOptions")]
        public Input<Inputs.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsGetArgs>? CloudwatchLoggingOptions { get; set; }

        /// <summary>
        /// The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
        /// </summary>
        [Input("compressionFormat")]
        public Input<string>? CompressionFormat { get; set; }

        /// <summary>
        /// Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
        /// </summary>
        [Input("errorOutputPrefix")]
        public Input<string>? ErrorOutputPrefix { get; set; }

        /// <summary>
        /// Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
        /// be used.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The ARN of the role that provides access to the source Kinesis stream.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationGetArgs()
        {
        }
        public static new FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationGetArgs Empty => new FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationGetArgs();
    }
}
