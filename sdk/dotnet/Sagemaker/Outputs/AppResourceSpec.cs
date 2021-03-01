// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class AppResourceSpec
    {
        /// <summary>
        /// The instance type that the image version runs on. For valid values see [Sagemaker Instance Types](https://docs.aws.amazon.com/sagemaker/latest/dg/notebooks-available-instance-types.html).
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// The ARN of the SageMaker image that the image version belongs to.
        /// </summary>
        public readonly string? SagemakerImageArn;

        [OutputConstructor]
        private AppResourceSpec(
            string? instanceType,

            string? sagemakerImageArn)
        {
            InstanceType = instanceType;
            SagemakerImageArn = sagemakerImageArn;
        }
    }
}
