// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pipes.Inputs
{

    public sealed class PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the task's elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE. Valid Values: ENABLED, DISABLED.
        /// </summary>
        [Input("assignPublicIp")]
        public Input<string>? AssignPublicIp { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        public PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationGetArgs()
        {
        }
        public static new PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationGetArgs Empty => new PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationGetArgs();
    }
}
