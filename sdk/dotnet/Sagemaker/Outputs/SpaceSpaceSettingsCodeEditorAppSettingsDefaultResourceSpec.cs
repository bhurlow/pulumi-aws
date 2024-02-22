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
    public sealed class SpaceSpaceSettingsCodeEditorAppSettingsDefaultResourceSpec
    {
        /// <summary>
        /// The instance type.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
        /// </summary>
        public readonly string? LifecycleConfigArn;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
        /// </summary>
        public readonly string? SagemakerImageArn;
        /// <summary>
        /// The SageMaker Image Version Alias.
        /// </summary>
        public readonly string? SagemakerImageVersionAlias;
        /// <summary>
        /// The ARN of the image version created on the instance.
        /// </summary>
        public readonly string? SagemakerImageVersionArn;

        [OutputConstructor]
        private SpaceSpaceSettingsCodeEditorAppSettingsDefaultResourceSpec(
            string? instanceType,

            string? lifecycleConfigArn,

            string? sagemakerImageArn,

            string? sagemakerImageVersionAlias,

            string? sagemakerImageVersionArn)
        {
            InstanceType = instanceType;
            LifecycleConfigArn = lifecycleConfigArn;
            SagemakerImageArn = sagemakerImageArn;
            SagemakerImageVersionAlias = sagemakerImageVersionAlias;
            SagemakerImageVersionArn = sagemakerImageVersionArn;
        }
    }
}
