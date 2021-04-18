// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Outputs
{

    [OutputType]
    public sealed class GroupWarmPool
    {
        /// <summary>
        /// Specifies the total maximum number of instances that are allowed to be in the warm pool or in any state except Terminated for the Auto Scaling group.
        /// </summary>
        public readonly int? MaxGroupPreparedCapacity;
        /// <summary>
        /// Specifies the minimum number of instances to maintain in the warm pool. This helps you to ensure that there is always a certain number of warmed instances available to handle traffic spikes. Defaults to 0 if not specified.
        /// </summary>
        public readonly int? MinSize;
        /// <summary>
        /// Sets the instance state to transition to after the lifecycle hooks finish. Valid values are: Stopped (default) or Running.
        /// </summary>
        public readonly string? PoolState;

        [OutputConstructor]
        private GroupWarmPool(
            int? maxGroupPreparedCapacity,

            int? minSize,

            string? poolState)
        {
            MaxGroupPreparedCapacity = maxGroupPreparedCapacity;
            MinSize = minSize;
            PoolState = poolState;
        }
    }
}
