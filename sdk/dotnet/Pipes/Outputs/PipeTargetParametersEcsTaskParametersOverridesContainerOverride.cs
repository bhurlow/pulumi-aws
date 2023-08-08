// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeTargetParametersEcsTaskParametersOverridesContainerOverride
    {
        /// <summary>
        /// List of commands to send to the container that overrides the default command from the Docker image or the task definition.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// The cpu override for the task.
        /// </summary>
        public readonly int? Cpu;
        /// <summary>
        /// A list of files containing the environment variables to pass to a container, instead of the value from the container definition. Detailed below.
        /// </summary>
        public readonly ImmutableArray<Outputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFile> EnvironmentFiles;
        /// <summary>
        /// The environment variables to send to the container. You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition. Environment variables cannot start with " AWS Batch ". This naming convention is reserved for variables that AWS Batch sets. Detailed below.
        /// </summary>
        public readonly ImmutableArray<Outputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment> Environments;
        /// <summary>
        /// The memory override for the task.
        /// </summary>
        public readonly int? Memory;
        /// <summary>
        /// The soft limit (in MiB) of memory to reserve for the container, instead of the default value from the task definition. You must also specify a container name.
        /// </summary>
        public readonly int? MemoryReservation;
        /// <summary>
        /// Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type and amount of resources to assign to a container. This overrides the settings in the job definition. The supported resources include GPU, MEMORY, and VCPU. Detailed below.
        /// </summary>
        public readonly ImmutableArray<Outputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideResourceRequirement> ResourceRequirements;

        [OutputConstructor]
        private PipeTargetParametersEcsTaskParametersOverridesContainerOverride(
            ImmutableArray<string> commands,

            int? cpu,

            ImmutableArray<Outputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFile> environmentFiles,

            ImmutableArray<Outputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment> environments,

            int? memory,

            int? memoryReservation,

            string? name,

            ImmutableArray<Outputs.PipeTargetParametersEcsTaskParametersOverridesContainerOverrideResourceRequirement> resourceRequirements)
        {
            Commands = commands;
            Cpu = cpu;
            EnvironmentFiles = environmentFiles;
            Environments = environments;
            Memory = memory;
            MemoryReservation = memoryReservation;
            Name = name;
            ResourceRequirements = resourceRequirements;
        }
    }
}
