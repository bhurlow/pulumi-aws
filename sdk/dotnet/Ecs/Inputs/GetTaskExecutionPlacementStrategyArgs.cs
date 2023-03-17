// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class GetTaskExecutionPlacementStrategyInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field to apply the placement strategy against.
        /// </summary>
        [Input("field")]
        public Input<string>? Field { get; set; }

        /// <summary>
        /// The type of placement strategy. Valid values are `random`, `spread`, and `binpack`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetTaskExecutionPlacementStrategyInputArgs()
        {
        }
        public static new GetTaskExecutionPlacementStrategyInputArgs Empty => new GetTaskExecutionPlacementStrategyInputArgs();
    }
}
