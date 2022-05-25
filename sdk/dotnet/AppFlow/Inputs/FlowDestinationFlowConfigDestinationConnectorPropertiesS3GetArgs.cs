// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class FlowDestinationFlowConfigDestinationConnectorPropertiesS3GetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon S3 bucket name where the source files are stored.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The object key for the Amazon S3 bucket in which the source files are stored.
        /// </summary>
        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        /// <summary>
        /// The configuration that determines how Amazon AppFlow should format the flow output data when Upsolver is used as the destination. See Upsolver S3 Output Format Config for more details.
        /// </summary>
        [Input("s3OutputFormatConfig")]
        public Input<Inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigGetArgs>? S3OutputFormatConfig { get; set; }

        public FlowDestinationFlowConfigDestinationConnectorPropertiesS3GetArgs()
        {
        }
    }
}
