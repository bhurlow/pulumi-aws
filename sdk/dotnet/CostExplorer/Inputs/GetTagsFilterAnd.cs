// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CostExplorer.Inputs
{

    public sealed class GetTagsFilterAndArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Configuration block for the filter that's based on `CostCategory` values. See below.
        /// </summary>
        [Input("costCategory")]
        public Inputs.GetTagsFilterAndCostCategoryArgs? CostCategory { get; set; }

        /// <summary>
        /// Configuration block for the specific `Dimension` to use for `Expression`. See below.
        /// </summary>
        [Input("dimension")]
        public Inputs.GetTagsFilterAndDimensionArgs? Dimension { get; set; }

        /// <summary>
        /// Tags that match your request.
        /// </summary>
        [Input("tags")]
        public Inputs.GetTagsFilterAndTagsArgs? Tags { get; set; }

        public GetTagsFilterAndArgs()
        {
        }
    }
}
