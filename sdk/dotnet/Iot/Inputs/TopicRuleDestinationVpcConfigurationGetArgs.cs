// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Inputs
{

    public sealed class TopicRuleDestinationVpcConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of a role that has permission to create and attach to elastic network interfaces (ENIs).
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The security groups of the VPC destination.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The subnet IDs of the VPC destination.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public TopicRuleDestinationVpcConfigurationGetArgs()
        {
        }
    }
}
