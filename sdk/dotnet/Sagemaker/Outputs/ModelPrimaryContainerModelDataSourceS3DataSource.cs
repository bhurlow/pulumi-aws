// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class ModelPrimaryContainerModelDataSourceS3DataSource
    {
        /// <summary>
        /// How the model data is prepared. Allowed values are: `None` and `Gzip`.
        /// </summary>
        public readonly string CompressionType;
        /// <summary>
        /// The type of model data to deploy. Allowed values are: `S3Object` and `S3Prefix`.
        /// </summary>
        public readonly string S3DataType;
        /// <summary>
        /// The S3 path of model data to deploy.
        /// </summary>
        public readonly string S3Uri;

        [OutputConstructor]
        private ModelPrimaryContainerModelDataSourceS3DataSource(
            string compressionType,

            string s3DataType,

            string s3Uri)
        {
            CompressionType = compressionType;
            S3DataType = s3DataType;
            S3Uri = s3Uri;
        }
    }
}
