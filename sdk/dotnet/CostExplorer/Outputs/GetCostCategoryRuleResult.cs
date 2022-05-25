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
    public sealed class GetCostCategoryRuleResult
    {
        /// <summary>
        /// Configuration block for the value the line item is categorized as if the line item contains the matched dimension. See below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCostCategoryRuleInheritedValueResult> InheritedValues;
        /// <summary>
        /// Configuration block for the `Expression` object used to categorize costs. See below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCostCategoryRuleRuleResult> Rules;
        /// <summary>
        /// Parameter type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Default value for the cost category.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetCostCategoryRuleResult(
            ImmutableArray<Outputs.GetCostCategoryRuleInheritedValueResult> inheritedValues,

            ImmutableArray<Outputs.GetCostCategoryRuleRuleResult> rules,

            string type,

            string value)
        {
            InheritedValues = inheritedValues;
            Rules = rules;
            Type = type;
            Value = value;
        }
    }
}
