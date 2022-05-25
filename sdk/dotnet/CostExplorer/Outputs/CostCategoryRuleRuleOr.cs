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
    public sealed class CostCategoryRuleRuleOr
    {
        /// <summary>
        /// Configuration block for the filter that's based on `CostCategory` values. See below.
        /// </summary>
        public readonly Outputs.CostCategoryRuleRuleOrCostCategory? CostCategory;
        /// <summary>
        /// Configuration block for the specific `Dimension` to use for `Expression`. See below.
        /// </summary>
        public readonly Outputs.CostCategoryRuleRuleOrDimension? Dimension;
        public readonly Outputs.CostCategoryRuleRuleOrTags? Tags;

        [OutputConstructor]
        private CostCategoryRuleRuleOr(
            Outputs.CostCategoryRuleRuleOrCostCategory? costCategory,

            Outputs.CostCategoryRuleRuleOrDimension? dimension,

            Outputs.CostCategoryRuleRuleOrTags? tags)
        {
            CostCategory = costCategory;
            Dimension = dimension;
            Tags = tags;
        }
    }
}
