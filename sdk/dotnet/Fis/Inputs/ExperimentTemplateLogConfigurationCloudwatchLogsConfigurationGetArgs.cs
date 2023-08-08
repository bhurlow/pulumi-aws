// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fis.Inputs
{

    public sealed class ExperimentTemplateLogConfigurationCloudwatchLogsConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the destination Amazon CloudWatch Logs log group.
        /// </summary>
        [Input("logGroupArn", required: true)]
        public Input<string> LogGroupArn { get; set; } = null!;

        public ExperimentTemplateLogConfigurationCloudwatchLogsConfigurationGetArgs()
        {
        }
        public static new ExperimentTemplateLogConfigurationCloudwatchLogsConfigurationGetArgs Empty => new ExperimentTemplateLogConfigurationCloudwatchLogsConfigurationGetArgs();
    }
}
