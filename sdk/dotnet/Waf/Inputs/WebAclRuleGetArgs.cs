// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf.Inputs
{

    public sealed class WebAclRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule. Not used if `type` is `GROUP`.
        /// </summary>
        [Input("action")]
        public Input<Inputs.WebAclRuleActionGetArgs>? Action { get; set; }

        /// <summary>
        /// Override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule. Only used if `type` is `GROUP`.
        /// </summary>
        [Input("overrideAction")]
        public Input<Inputs.WebAclRuleOverrideActionGetArgs>? OverrideAction { get; set; }

        /// <summary>
        /// Specifies the order in which the rules in a WebACL are evaluated.
        /// Rules with a lower value are evaluated before rules with a higher value.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// ID of the associated WAF (Global) rule (e.g., `aws.waf.Rule`). WAF (Regional) rules cannot be used.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        /// <summary>
        /// The rule type, either `REGULAR`, as defined by [Rule](http://docs.aws.amazon.com/waf/latest/APIReference/API_Rule.html), `RATE_BASED`, as defined by [RateBasedRule](http://docs.aws.amazon.com/waf/latest/APIReference/API_RateBasedRule.html), or `GROUP`, as defined by [RuleGroup](https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html). The default is REGULAR. If you add a RATE_BASED rule, you need to set `type` as `RATE_BASED`. If you add a GROUP rule, you need to set `type` as `GROUP`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public WebAclRuleGetArgs()
        {
        }
        public static new WebAclRuleGetArgs Empty => new WebAclRuleGetArgs();
    }
}
