// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Batch.Inputs
{

    public sealed class JobQueueComputeEnvironmentOrderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the compute environment.
        /// </summary>
        [Input("computeEnvironment", required: true)]
        public Input<string> ComputeEnvironment { get; set; } = null!;

        /// <summary>
        /// The order of the compute environment. Compute environments are tried in ascending order. For example, if two compute environments are associated with a job queue, the compute environment with a lower order integer value is tried for job placement first.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        public JobQueueComputeEnvironmentOrderArgs()
        {
        }
        public static new JobQueueComputeEnvironmentOrderArgs Empty => new JobQueueComputeEnvironmentOrderArgs();
    }
}
