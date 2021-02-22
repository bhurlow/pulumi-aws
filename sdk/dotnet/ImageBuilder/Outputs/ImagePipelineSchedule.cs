// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ImageBuilder.Outputs
{

    [OutputType]
    public sealed class ImagePipelineSchedule
    {
        /// <summary>
        /// Condition when the pipeline should trigger a new image build. Valid values are `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE` and `EXPRESSION_MATCH_ONLY`. Defaults to `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE`.
        /// </summary>
        public readonly string? PipelineExecutionStartCondition;
        /// <summary>
        /// Cron expression of how often the pipeline start condition is evaluated. For example, `cron(0 0 * * ? *)` is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as `cron(0 0 * * *)`, must be updated to the six field syntax. For more information, see the [Image Builder User Guide](https://docs.aws.amazon.com/imagebuilder/latest/userguide/cron-expressions.html).
        /// </summary>
        public readonly string ScheduleExpression;

        [OutputConstructor]
        private ImagePipelineSchedule(
            string? pipelineExecutionStartCondition,

            string scheduleExpression)
        {
            PipelineExecutionStartCondition = pipelineExecutionStartCondition;
            ScheduleExpression = scheduleExpression;
        }
    }
}
