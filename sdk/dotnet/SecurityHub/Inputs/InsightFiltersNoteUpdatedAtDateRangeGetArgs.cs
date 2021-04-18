// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub.Inputs
{

    public sealed class InsightFiltersNoteUpdatedAtDateRangeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A date range unit for the date filter. Valid values: `DAYS`.
        /// </summary>
        [Input("unit", required: true)]
        public Input<string> Unit { get; set; } = null!;

        /// <summary>
        /// A date range value for the date filter, provided as an Integer.
        /// </summary>
        [Input("value", required: true)]
        public Input<int> Value { get; set; } = null!;

        public InsightFiltersNoteUpdatedAtDateRangeGetArgs()
        {
        }
    }
}
