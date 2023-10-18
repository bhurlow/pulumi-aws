// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Outputs
{

    [OutputType]
    public sealed class TableImportTableS3BucketSource
    {
        /// <summary>
        /// The S3 bucket that is being imported from.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The account number of the S3 bucket that is being imported from.
        /// </summary>
        public readonly string? BucketOwner;
        /// <summary>
        /// The key prefix shared by all S3 Objects that are being imported.
        /// </summary>
        public readonly string? KeyPrefix;

        [OutputConstructor]
        private TableImportTableS3BucketSource(
            string bucket,

            string? bucketOwner,

            string? keyPrefix)
        {
            Bucket = bucket;
            BucketOwner = bucketOwner;
            KeyPrefix = keyPrefix;
        }
    }
}
