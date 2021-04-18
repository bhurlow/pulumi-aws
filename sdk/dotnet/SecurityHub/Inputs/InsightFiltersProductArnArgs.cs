// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub.Inputs
{

    public sealed class InsightFiltersProductArnArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The condition to apply to a string value when querying for findings. Valid values include: `EQUALS` and `NOT_EQUALS`.
        /// </summary>
        [Input("comparison", required: true)]
        public Input<string> Comparison { get; set; } = null!;

        /// <summary>
        /// A date range value for the date filter, provided as an Integer.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public InsightFiltersProductArnArgs()
        {
        }
    }
}
