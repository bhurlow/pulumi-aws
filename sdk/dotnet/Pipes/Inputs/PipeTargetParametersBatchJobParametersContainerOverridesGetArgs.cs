// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pipes.Inputs
{

    public sealed class PipeTargetParametersBatchJobParametersContainerOverridesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// List of commands to send to the container that overrides the default command from the Docker image or the task definition.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("environments")]
        private InputList<Inputs.PipeTargetParametersBatchJobParametersContainerOverridesEnvironmentGetArgs>? _environments;

        /// <summary>
        /// The environment variables to send to the container. You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition. Environment variables cannot start with " AWS Batch ". This naming convention is reserved for variables that AWS Batch sets. Detailed below.
        /// </summary>
        public InputList<Inputs.PipeTargetParametersBatchJobParametersContainerOverridesEnvironmentGetArgs> Environments
        {
            get => _environments ?? (_environments = new InputList<Inputs.PipeTargetParametersBatchJobParametersContainerOverridesEnvironmentGetArgs>());
            set => _environments = value;
        }

        /// <summary>
        /// The instance type to use for a multi-node parallel job. This parameter isn't applicable to single-node container jobs or jobs that run on Fargate resources, and shouldn't be provided.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("resourceRequirements")]
        private InputList<Inputs.PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirementGetArgs>? _resourceRequirements;

        /// <summary>
        /// The type and amount of resources to assign to a container. This overrides the settings in the job definition. The supported resources include GPU, MEMORY, and VCPU. Detailed below.
        /// </summary>
        public InputList<Inputs.PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirementGetArgs> ResourceRequirements
        {
            get => _resourceRequirements ?? (_resourceRequirements = new InputList<Inputs.PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirementGetArgs>());
            set => _resourceRequirements = value;
        }

        public PipeTargetParametersBatchJobParametersContainerOverridesGetArgs()
        {
        }
        public static new PipeTargetParametersBatchJobParametersContainerOverridesGetArgs Empty => new PipeTargetParametersBatchJobParametersContainerOverridesGetArgs();
    }
}
