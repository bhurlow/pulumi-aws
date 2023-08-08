// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Inputs
{

    public sealed class ClusterScalingConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable automatic pause. A DB cluster can be paused only when it's idle (it has no connections). If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it. Defaults to `true`.
        /// </summary>
        [Input("autoPause")]
        public Input<bool>? AutoPause { get; set; }

        /// <summary>
        /// Maximum capacity for an Aurora DB cluster in `serverless` DB engine mode. The maximum capacity must be greater than or equal to the minimum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `16`.
        /// </summary>
        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        /// <summary>
        /// Minimum capacity for an Aurora DB cluster in `serverless` DB engine mode. The minimum capacity must be lesser than or equal to the maximum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `1`.
        /// </summary>
        [Input("minCapacity")]
        public Input<int>? MinCapacity { get; set; }

        /// <summary>
        /// Time, in seconds, before an Aurora DB cluster in serverless mode is paused. Valid values are `300` through `86400`. Defaults to `300`.
        /// </summary>
        [Input("secondsUntilAutoPause")]
        public Input<int>? SecondsUntilAutoPause { get; set; }

        /// <summary>
        /// Action to take when the timeout is reached. Valid values: `ForceApplyCapacityChange`, `RollbackCapacityChange`. Defaults to `RollbackCapacityChange`. See [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.timeout-action).
        /// </summary>
        [Input("timeoutAction")]
        public Input<string>? TimeoutAction { get; set; }

        public ClusterScalingConfigurationGetArgs()
        {
        }
        public static new ClusterScalingConfigurationGetArgs Empty => new ClusterScalingConfigurationGetArgs();
    }
}
