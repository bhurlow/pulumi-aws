// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CostExplorer.Inputs
{

    public sealed class CostCategoryRuleRuleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ands")]
        private InputList<Inputs.CostCategoryRuleRuleAndGetArgs>? _ands;

        /// <summary>
        /// Return results that match both `Dimension` objects.
        /// </summary>
        public InputList<Inputs.CostCategoryRuleRuleAndGetArgs> Ands
        {
            get => _ands ?? (_ands = new InputList<Inputs.CostCategoryRuleRuleAndGetArgs>());
            set => _ands = value;
        }

        /// <summary>
        /// Configuration block for the filter that's based on `CostCategory` values. See below.
        /// </summary>
        [Input("costCategory")]
        public Input<Inputs.CostCategoryRuleRuleCostCategoryGetArgs>? CostCategory { get; set; }

        /// <summary>
        /// Configuration block for the specific `Dimension` to use for `Expression`. See below.
        /// </summary>
        [Input("dimension")]
        public Input<Inputs.CostCategoryRuleRuleDimensionGetArgs>? Dimension { get; set; }

        /// <summary>
        /// Return results that match both `Dimension` object.
        /// </summary>
        [Input("not")]
        public Input<Inputs.CostCategoryRuleRuleNotGetArgs>? Not { get; set; }

        [Input("ors")]
        private InputList<Inputs.CostCategoryRuleRuleOrGetArgs>? _ors;

        /// <summary>
        /// Return results that match both `Dimension` object.
        /// </summary>
        public InputList<Inputs.CostCategoryRuleRuleOrGetArgs> Ors
        {
            get => _ors ?? (_ors = new InputList<Inputs.CostCategoryRuleRuleOrGetArgs>());
            set => _ors = value;
        }

        /// <summary>
        /// Configuration block for the specific `Tag` to use for `Expression`. See below.
        /// </summary>
        [Input("tags")]
        public Input<Inputs.CostCategoryRuleRuleTagsGetArgs>? Tags { get; set; }

        public CostCategoryRuleRuleGetArgs()
        {
        }
        public static new CostCategoryRuleRuleGetArgs Empty => new CostCategoryRuleRuleGetArgs();
    }
}
