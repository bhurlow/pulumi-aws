// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateIamInstanceProfile
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the instance profile.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The name of the instance profile.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private LaunchTemplateIamInstanceProfile(
            string? arn,

            string? name)
        {
            Arn = arn;
            Name = name;
        }
    }
}