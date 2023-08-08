// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pipes.Inputs
{

    public sealed class PipeTargetParametersEcsTaskParametersOverridesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("containerOverrides")]
        private InputList<Inputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideGetArgs>? _containerOverrides;

        /// <summary>
        /// One or more container overrides that are sent to a task. Detailed below.
        /// </summary>
        public InputList<Inputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideGetArgs> ContainerOverrides
        {
            get => _containerOverrides ?? (_containerOverrides = new InputList<Inputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideGetArgs>());
            set => _containerOverrides = value;
        }

        /// <summary>
        /// The cpu override for the task.
        /// </summary>
        [Input("cpu")]
        public Input<string>? Cpu { get; set; }

        /// <summary>
        /// The ephemeral storage setting override for the task.  Detailed below.
        /// </summary>
        [Input("ephemeralStorage")]
        public Input<Inputs.PipeTargetParametersEcsTaskParametersOverridesEphemeralStorageGetArgs>? EphemeralStorage { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the task execution IAM role override for the task.
        /// </summary>
        [Input("executionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        [Input("inferenceAcceleratorOverrides")]
        private InputList<Inputs.PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverrideGetArgs>? _inferenceAcceleratorOverrides;

        /// <summary>
        /// List of Elastic Inference accelerator overrides for the task. Detailed below.
        /// </summary>
        public InputList<Inputs.PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverrideGetArgs> InferenceAcceleratorOverrides
        {
            get => _inferenceAcceleratorOverrides ?? (_inferenceAcceleratorOverrides = new InputList<Inputs.PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverrideGetArgs>());
            set => _inferenceAcceleratorOverrides = value;
        }

        /// <summary>
        /// The memory override for the task.
        /// </summary>
        [Input("memory")]
        public Input<string>? Memory { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that containers in this task can assume. All containers in this task are granted the permissions that are specified in this role.
        /// </summary>
        [Input("taskRoleArn")]
        public Input<string>? TaskRoleArn { get; set; }

        public PipeTargetParametersEcsTaskParametersOverridesGetArgs()
        {
        }
        public static new PipeTargetParametersEcsTaskParametersOverridesGetArgs Empty => new PipeTargetParametersEcsTaskParametersOverridesGetArgs();
    }
}
