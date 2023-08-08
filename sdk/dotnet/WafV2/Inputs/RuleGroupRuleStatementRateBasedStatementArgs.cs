// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementRateBasedStatementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `FORWARDED_IP` or `IP`. Default: `IP`.
        /// </summary>
        [Input("aggregateKeyType")]
        public Input<string>? AggregateKeyType { get; set; }

        /// <summary>
        /// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See Forwarded IP Config below for details.
        /// </summary>
        [Input("forwardedIpConfig")]
        public Input<Inputs.RuleGroupRuleStatementRateBasedStatementForwardedIpConfigArgs>? ForwardedIpConfig { get; set; }

        /// <summary>
        /// The limit on requests per 5-minute period for a single originating IP address.
        /// </summary>
        [Input("limit", required: true)]
        public Input<int> Limit { get; set; } = null!;

        /// <summary>
        /// An optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See Statement above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
        /// </summary>
        [Input("scopeDownStatement")]
        public Input<Inputs.RuleGroupRuleStatementRateBasedStatementScopeDownStatementArgs>? ScopeDownStatement { get; set; }

        public RuleGroupRuleStatementRateBasedStatementArgs()
        {
        }
        public static new RuleGroupRuleStatementRateBasedStatementArgs Empty => new RuleGroupRuleStatementRateBasedStatementArgs();
    }
}
