// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf.Outputs
{

    [OutputType]
    public sealed class RuleGroupActivatedRule
    {
        /// <summary>
        /// Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
        /// </summary>
        public readonly Outputs.RuleGroupActivatedRuleAction Action;
        /// <summary>
        /// Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The ID of a rule
        /// </summary>
        public readonly string RuleId;
        /// <summary>
        /// The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private RuleGroupActivatedRule(
            Outputs.RuleGroupActivatedRuleAction action,

            int priority,

            string ruleId,

            string? type)
        {
            Action = action;
            Priority = priority;
            RuleId = ruleId;
            Type = type;
        }
    }
}
