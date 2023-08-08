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
    public sealed class GroupTrafficSource
    {
        /// <summary>
        /// Identifies the traffic source. For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.
        /// </summary>
        public readonly string Identifier;
        /// <summary>
        /// Provides additional context for the value of Identifier.
        /// The following lists the valid values:
        /// `elb` if `identifier` is the name of a Classic Load Balancer.
        /// `elbv2` if `identifier` is the ARN of an Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target group.
        /// `vpc-lattice` if `identifier` is the ARN of a VPC Lattice target group.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GroupTrafficSource(
            string identifier,

            string? type)
        {
            Identifier = identifier;
            Type = type;
        }
    }
}
