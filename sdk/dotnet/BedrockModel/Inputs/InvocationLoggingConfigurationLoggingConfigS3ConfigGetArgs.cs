// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.BedrockModel.Inputs
{

    public sealed class InvocationLoggingConfigurationLoggingConfigS3ConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// S3 bucket name.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// S3 prefix.
        /// </summary>
        [Input("keyPrefix")]
        public Input<string>? KeyPrefix { get; set; }

        public InvocationLoggingConfigurationLoggingConfigS3ConfigGetArgs()
        {
        }
        public static new InvocationLoggingConfigurationLoggingConfigS3ConfigGetArgs Empty => new InvocationLoggingConfigurationLoggingConfigS3ConfigGetArgs();
    }
}
