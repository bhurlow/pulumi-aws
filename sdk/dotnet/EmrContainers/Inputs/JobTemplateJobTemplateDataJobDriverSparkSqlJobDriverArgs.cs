// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.EmrContainers.Inputs
{

    public sealed class JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SQL file to be executed.
        /// </summary>
        [Input("entryPoint")]
        public Input<string>? EntryPoint { get; set; }

        /// <summary>
        /// The Spark parameters to be included in the Spark SQL command.
        /// </summary>
        [Input("sparkSqlParameters")]
        public Input<string>? SparkSqlParameters { get; set; }

        public JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs()
        {
        }
        public static new JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs Empty => new JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs();
    }
}
