// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Inputs
{

    public sealed class FirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationProcessorGetArgs : Pulumi.ResourceArgs
    {
        [Input("parameters")]
        private InputList<Inputs.FirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationProcessorParameterGetArgs>? _parameters;

        /// <summary>
        /// Array of processor parameters. More details are given below
        /// </summary>
        public InputList<Inputs.FirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationProcessorParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.FirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationProcessorParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The type of processor. Valid Values: `RecordDeAggregation`, `Lambda`, `MetadataExtraction`, `AppendDelimiterToRecord`. Validation is done against [AWS SDK constants](https://docs.aws.amazon.com/sdk-for-go/api/service/firehose/#pkg-constants); so that values not explicitly listed may also work.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationProcessorGetArgs()
        {
        }
    }
}
