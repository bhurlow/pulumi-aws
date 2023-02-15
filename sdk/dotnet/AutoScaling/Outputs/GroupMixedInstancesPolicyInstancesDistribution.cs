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
    public sealed class GroupMixedInstancesPolicyInstancesDistribution
    {
        /// <summary>
        /// Strategy to use when launching on-demand instances. Valid values: `prioritized`, `lowest-price`. Default: `prioritized`.
        /// </summary>
        public readonly string? OnDemandAllocationStrategy;
        /// <summary>
        /// Absolute minimum amount of desired capacity that must be fulfilled by on-demand instances. Default: `0`.
        /// </summary>
        public readonly int? OnDemandBaseCapacity;
        /// <summary>
        /// Percentage split between on-demand and Spot instances above the base on-demand capacity. Default: `100`.
        /// </summary>
        public readonly int? OnDemandPercentageAboveBaseCapacity;
        /// <summary>
        /// How to allocate capacity across the Spot pools. Valid values: `lowest-price`, `capacity-optimized`, `capacity-optimized-prioritized`, and `price-capacity-optimized`. Default: `lowest-price`.
        /// </summary>
        public readonly string? SpotAllocationStrategy;
        /// <summary>
        /// Number of Spot pools per availability zone to allocate capacity. EC2 Auto Scaling selects the cheapest Spot pools and evenly allocates Spot capacity across the number of Spot pools that you specify. Only available with `spot_allocation_strategy` set to `lowest-price`. Otherwise it must be set to `0`, if it has been defined before. Default: `2`.
        /// </summary>
        public readonly int? SpotInstancePools;
        /// <summary>
        /// Maximum price per unit hour that the user is willing to pay for the Spot instances. Default: an empty string which means the on-demand price.
        /// </summary>
        public readonly string? SpotMaxPrice;

        [OutputConstructor]
        private GroupMixedInstancesPolicyInstancesDistribution(
            string? onDemandAllocationStrategy,

            int? onDemandBaseCapacity,

            int? onDemandPercentageAboveBaseCapacity,

            string? spotAllocationStrategy,

            int? spotInstancePools,

            string? spotMaxPrice)
        {
            OnDemandAllocationStrategy = onDemandAllocationStrategy;
            OnDemandBaseCapacity = onDemandBaseCapacity;
            OnDemandPercentageAboveBaseCapacity = onDemandPercentageAboveBaseCapacity;
            SpotAllocationStrategy = spotAllocationStrategy;
            SpotInstancePools = spotInstancePools;
            SpotMaxPrice = spotMaxPrice;
        }
    }
}
