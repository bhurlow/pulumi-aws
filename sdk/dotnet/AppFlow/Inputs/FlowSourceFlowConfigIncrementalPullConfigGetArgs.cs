// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class FlowSourceFlowConfigIncrementalPullConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A field that specifies the date time or timestamp field as the criteria to use when importing incremental records from the source.
        /// </summary>
        [Input("datetimeTypeFieldName")]
        public Input<string>? DatetimeTypeFieldName { get; set; }

        public FlowSourceFlowConfigIncrementalPullConfigGetArgs()
        {
        }
    }
}
