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
    public sealed class RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatement
    {
        /// <summary>
        /// The part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
        /// </summary>
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatementFieldToMatch? FieldToMatch;
        /// <summary>
        /// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. See Text Transformation below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatement(
            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatementTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
