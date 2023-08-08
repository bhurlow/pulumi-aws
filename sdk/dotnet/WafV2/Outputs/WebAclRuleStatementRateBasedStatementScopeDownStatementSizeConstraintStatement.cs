// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Outputs
{

    [OutputType]
    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement
    {
        /// <summary>
        /// Operator to use to compare the request part to the size setting. Valid values include: `EQ`, `NE`, `LE`, `LT`, `GE`, or `GT`.
        /// </summary>
        public readonly string ComparisonOperator;
        /// <summary>
        /// Part of a web request that you want AWS WAF to inspect. See `field_to_match` below for details.
        /// </summary>
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatementFieldToMatch? FieldToMatch;
        /// <summary>
        /// Size, in bytes, to compare to the request part, after any transformations. Valid values are integers between 0 and 21474836480, inclusive.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. At least one transformation is required. See `text_transformation` below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement(
            string comparisonOperator,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatementFieldToMatch? fieldToMatch,

            int size,

            ImmutableArray<Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatementTextTransformation> textTransformations)
        {
            ComparisonOperator = comparisonOperator;
            FieldToMatch = fieldToMatch;
            Size = size;
            TextTransformations = textTransformations;
        }
    }
}
