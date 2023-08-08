// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Inputs
{

    public sealed class FirehoseDeliveryStreamOpensearchConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
        /// </summary>
        [Input("bufferingInterval")]
        public Input<int>? BufferingInterval { get; set; }

        /// <summary>
        /// Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
        /// </summary>
        [Input("bufferingSize")]
        public Input<int>? BufferingSize { get; set; }

        /// <summary>
        /// The CloudWatch Logging Options for the delivery stream. More details are given below
        /// </summary>
        [Input("cloudwatchLoggingOptions")]
        public Input<Inputs.FirehoseDeliveryStreamOpensearchConfigurationCloudwatchLoggingOptionsArgs>? CloudwatchLoggingOptions { get; set; }

        /// <summary>
        /// The endpoint to use when communicating with the cluster. Conflicts with `domain_arn`.
        /// </summary>
        [Input("clusterEndpoint")]
        public Input<string>? ClusterEndpoint { get; set; }

        /// <summary>
        /// The ARN of the Amazon ES domain.  The pattern needs to be `arn:.*`.  Conflicts with `cluster_endpoint`.
        /// </summary>
        [Input("domainArn")]
        public Input<string>? DomainArn { get; set; }

        /// <summary>
        /// The Opensearch index name.
        /// </summary>
        [Input("indexName", required: true)]
        public Input<string> IndexName { get; set; } = null!;

        /// <summary>
        /// The Opensearch index rotation period.  Index rotation appends a timestamp to the IndexName to facilitate expiration of old data.  Valid values are `NoRotation`, `OneHour`, `OneDay`, `OneWeek`, and `OneMonth`.  The default value is `OneDay`.
        /// </summary>
        [Input("indexRotationPeriod")]
        public Input<string>? IndexRotationPeriod { get; set; }

        /// <summary>
        /// The data processing configuration.  More details are given below.
        /// </summary>
        [Input("processingConfiguration")]
        public Input<Inputs.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationArgs>? ProcessingConfiguration { get; set; }

        /// <summary>
        /// After an initial failure to deliver to Amazon OpenSearch, the total amount of time, in seconds between 0 to 7200, during which Firehose re-attempts delivery (including the first attempt).  After this time has elapsed, the failed documents are written to Amazon S3.  The default value is 300s.  There will be no retry if the value is 0.
        /// </summary>
        [Input("retryDuration")]
        public Input<int>? RetryDuration { get; set; }

        /// <summary>
        /// The ARN of the IAM role to be assumed by Firehose for calling the Amazon ES Configuration API and for indexing documents.  The IAM role must have permission for `DescribeDomain`, `DescribeDomains`, and `DescribeDomainConfig`.  The pattern needs to be `arn:.*`.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Defines how documents should be delivered to Amazon S3.  Valid values are `FailedDocumentsOnly` and `AllDocuments`.  Default value is `FailedDocumentsOnly`.
        /// </summary>
        [Input("s3BackupMode")]
        public Input<string>? S3BackupMode { get; set; }

        /// <summary>
        /// The S3 Configuration. See s3_configuration for more details.
        /// </summary>
        [Input("s3Configuration", required: true)]
        public Input<Inputs.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs> S3Configuration { get; set; } = null!;

        /// <summary>
        /// The Elasticsearch type name with maximum length of 100 characters. Types are deprecated in OpenSearch_1.1. TypeName must be empty.
        /// </summary>
        [Input("typeName")]
        public Input<string>? TypeName { get; set; }

        /// <summary>
        /// The VPC configuration for the delivery stream to connect to OpenSearch associated with the VPC. More details are given below
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.FirehoseDeliveryStreamOpensearchConfigurationVpcConfigArgs>? VpcConfig { get; set; }

        public FirehoseDeliveryStreamOpensearchConfigurationArgs()
        {
        }
        public static new FirehoseDeliveryStreamOpensearchConfigurationArgs Empty => new FirehoseDeliveryStreamOpensearchConfigurationArgs();
    }
}
