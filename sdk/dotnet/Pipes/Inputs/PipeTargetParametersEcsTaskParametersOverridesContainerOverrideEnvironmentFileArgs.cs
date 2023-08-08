// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pipes.Inputs
{

    public sealed class PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of resource to assign to a container. The supported resources include GPU, MEMORY, and VCPU.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value of the key-value pair. For environment variables, this is the value of the environment variable.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFileArgs()
        {
        }
        public static new PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFileArgs Empty => new PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFileArgs();
    }
}
