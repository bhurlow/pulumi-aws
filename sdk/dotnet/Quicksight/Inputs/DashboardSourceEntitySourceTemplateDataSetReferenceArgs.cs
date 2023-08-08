// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Inputs
{

    public sealed class DashboardSourceEntitySourceTemplateDataSetReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dataset Amazon Resource Name (ARN).
        /// </summary>
        [Input("dataSetArn", required: true)]
        public Input<string> DataSetArn { get; set; } = null!;

        /// <summary>
        /// Dataset placeholder.
        /// </summary>
        [Input("dataSetPlaceholder", required: true)]
        public Input<string> DataSetPlaceholder { get; set; } = null!;

        public DashboardSourceEntitySourceTemplateDataSetReferenceArgs()
        {
        }
        public static new DashboardSourceEntitySourceTemplateDataSetReferenceArgs Empty => new DashboardSourceEntitySourceTemplateDataSetReferenceArgs();
    }
}
