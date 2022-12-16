// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class EndpointConfigurationProductionVariantCoreDumpConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL for S3 location where the captured data is stored.
        /// </summary>
        [Input("destinationS3Uri", required: true)]
        public Input<string> DestinationS3Uri { get; set; } = null!;

        /// <summary>
        /// The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        public EndpointConfigurationProductionVariantCoreDumpConfigArgs()
        {
        }
        public static new EndpointConfigurationProductionVariantCoreDumpConfigArgs Empty => new EndpointConfigurationProductionVariantCoreDumpConfigArgs();
    }
}
