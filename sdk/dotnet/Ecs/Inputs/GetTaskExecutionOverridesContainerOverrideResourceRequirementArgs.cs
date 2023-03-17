// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class GetTaskExecutionOverridesContainerOverrideResourceRequirementInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of resource to assign to a container. Valid values are `GPU` or `InferenceAccelerator`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value for the specified resource type. If the `GPU` type is used, the value is the number of physical GPUs the Amazon ECS container agent reserves for the container. The number of GPUs that's reserved for all containers in a task can't exceed the number of available GPUs on the container instance that the task is launched on. If the `InferenceAccelerator` type is used, the value matches the `deviceName` for an InferenceAccelerator specified in a task definition.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GetTaskExecutionOverridesContainerOverrideResourceRequirementInputArgs()
        {
        }
        public static new GetTaskExecutionOverridesContainerOverrideResourceRequirementInputArgs Empty => new GetTaskExecutionOverridesContainerOverrideResourceRequirementInputArgs();
    }
}
