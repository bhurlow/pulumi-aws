// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Inputs
{

    public sealed class ClusterServerlessv2ScalingConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum capacity for an Aurora DB cluster in `serverless` DB engine mode. The maximum capacity must be greater than or equal to the minimum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `16`.
        /// </summary>
        [Input("maxCapacity", required: true)]
        public Input<double> MaxCapacity { get; set; } = null!;

        /// <summary>
        /// Minimum capacity for an Aurora DB cluster in `serverless` DB engine mode. The minimum capacity must be lesser than or equal to the maximum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `1`.
        /// </summary>
        [Input("minCapacity", required: true)]
        public Input<double> MinCapacity { get; set; } = null!;

        public ClusterServerlessv2ScalingConfigurationGetArgs()
        {
        }
        public static new ClusterServerlessv2ScalingConfigurationGetArgs Empty => new ClusterServerlessv2ScalingConfigurationGetArgs();
    }
}
