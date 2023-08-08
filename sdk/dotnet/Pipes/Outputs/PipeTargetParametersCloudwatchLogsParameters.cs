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
    public sealed class PipeTargetParametersCloudwatchLogsParameters
    {
        /// <summary>
        /// The name of the log stream.
        /// </summary>
        public readonly string? LogStreamName;
        /// <summary>
        /// The time the event occurred, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC. This is the JSON path to the field in the event e.g. $.detail.timestamp
        /// </summary>
        public readonly string? Timestamp;

        [OutputConstructor]
        private PipeTargetParametersCloudwatchLogsParameters(
            string? logStreamName,

            string? timestamp)
        {
            LogStreamName = logStreamName;
            Timestamp = timestamp;
        }
    }
}
