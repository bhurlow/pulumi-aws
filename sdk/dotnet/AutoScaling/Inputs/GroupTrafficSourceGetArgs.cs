// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class GroupTrafficSourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies the traffic source. For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        /// <summary>
        /// Provides additional context for the value of Identifier.
        /// The following lists the valid values:
        /// `elb` if `identifier` is the name of a Classic Load Balancer.
        /// `elbv2` if `identifier` is the ARN of an Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target group.
        /// `vpc-lattice` if `identifier` is the ARN of a VPC Lattice target group.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GroupTrafficSourceGetArgs()
        {
        }
        public static new GroupTrafficSourceGetArgs Empty => new GroupTrafficSourceGetArgs();
    }
}
