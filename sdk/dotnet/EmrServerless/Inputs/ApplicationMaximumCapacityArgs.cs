// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.EmrServerless.Inputs
{

    public sealed class ApplicationMaximumCapacityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU requirements for every worker instance of the worker type.
        /// </summary>
        [Input("cpu", required: true)]
        public Input<string> Cpu { get; set; } = null!;

        /// <summary>
        /// The disk requirements for every worker instance of the worker type.
        /// </summary>
        [Input("disk")]
        public Input<string>? Disk { get; set; }

        /// <summary>
        /// The memory requirements for every worker instance of the worker type.
        /// </summary>
        [Input("memory", required: true)]
        public Input<string> Memory { get; set; } = null!;

        public ApplicationMaximumCapacityArgs()
        {
        }
    }
}
