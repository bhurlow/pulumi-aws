// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Evidently.Inputs
{

    public sealed class LaunchExecutionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time that the launch ended.
        /// </summary>
        [Input("endedTime")]
        public Input<string>? EndedTime { get; set; }

        /// <summary>
        /// The date and time that the launch started.
        /// </summary>
        [Input("startedTime")]
        public Input<string>? StartedTime { get; set; }

        public LaunchExecutionGetArgs()
        {
        }
        public static new LaunchExecutionGetArgs Empty => new LaunchExecutionGetArgs();
    }
}
