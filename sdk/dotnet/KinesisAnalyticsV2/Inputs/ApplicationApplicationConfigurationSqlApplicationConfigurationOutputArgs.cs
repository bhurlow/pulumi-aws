// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.KinesisAnalyticsV2.Inputs
{

    public sealed class ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the data format when records are written to the destination.
        /// </summary>
        [Input("destinationSchema", required: true)]
        public Input<Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs> DestinationSchema { get; set; } = null!;

        /// <summary>
        /// Identifies a Kinesis Data Firehose delivery stream as the destination.
        /// </summary>
        [Input("kinesisFirehoseOutput")]
        public Input<Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutputArgs>? KinesisFirehoseOutput { get; set; }

        /// <summary>
        /// Identifies a Kinesis data stream as the destination.
        /// </summary>
        [Input("kinesisStreamsOutput")]
        public Input<Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisStreamsOutputArgs>? KinesisStreamsOutput { get; set; }

        /// <summary>
        /// Identifies a Lambda function as the destination.
        /// </summary>
        [Input("lambdaOutput")]
        public Input<Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutputArgs>? LambdaOutput { get; set; }

        /// <summary>
        /// The name of the in-application stream.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("outputId")]
        public Input<string>? OutputId { get; set; }

        public ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs()
        {
        }
    }
}
