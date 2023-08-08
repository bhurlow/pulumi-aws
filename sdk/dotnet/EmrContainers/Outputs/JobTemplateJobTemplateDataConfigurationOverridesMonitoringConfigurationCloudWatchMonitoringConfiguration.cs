// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.EmrContainers.Outputs
{

    [OutputType]
    public sealed class JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationCloudWatchMonitoringConfiguration
    {
        /// <summary>
        /// The name of the log group for log publishing.
        /// </summary>
        public readonly string LogGroupName;
        /// <summary>
        /// The specified name prefix for log streams.
        /// </summary>
        public readonly string? LogStreamNamePrefix;

        [OutputConstructor]
        private JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationCloudWatchMonitoringConfiguration(
            string logGroupName,

            string? logStreamNamePrefix)
        {
            LogGroupName = logGroupName;
            LogStreamNamePrefix = logStreamNamePrefix;
        }
    }
}
