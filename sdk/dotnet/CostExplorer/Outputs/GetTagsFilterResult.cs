// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CostExplorer.Outputs
{

    [OutputType]
    public sealed class GetTagsFilterResult
    {
        /// <summary>
        /// Return results that match both `Dimension` objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTagsFilterAndResult> Ands;
        /// <summary>
        /// Configuration block for the filter that's based on `CostCategory` values. See below.
        /// </summary>
        public readonly Outputs.GetTagsFilterCostCategoryResult? CostCategory;
        /// <summary>
        /// Configuration block for the specific `Dimension` to use for `Expression`. See below.
        /// </summary>
        public readonly Outputs.GetTagsFilterDimensionResult? Dimension;
        /// <summary>
        /// Return results that match both `Dimension` object.
        /// </summary>
        public readonly Outputs.GetTagsFilterNotResult? Not;
        /// <summary>
        /// Return results that match both `Dimension` object.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTagsFilterOrResult> Ors;
        /// <summary>
        /// Tags that match your request.
        /// </summary>
        public readonly Outputs.GetTagsFilterTagsResult? Tags;

        [OutputConstructor]
        private GetTagsFilterResult(
            ImmutableArray<Outputs.GetTagsFilterAndResult> ands,

            Outputs.GetTagsFilterCostCategoryResult? costCategory,

            Outputs.GetTagsFilterDimensionResult? dimension,

            Outputs.GetTagsFilterNotResult? not,

            ImmutableArray<Outputs.GetTagsFilterOrResult> ors,

            Outputs.GetTagsFilterTagsResult? tags)
        {
            Ands = ands;
            CostCategory = costCategory;
            Dimension = dimension;
            Not = not;
            Ors = ors;
            Tags = tags;
        }
    }
}
