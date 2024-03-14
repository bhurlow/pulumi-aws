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
    public sealed class RuleGroupRuleStatementRateBasedStatement
    {
        /// <summary>
        /// Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `CUSTOM_KEYS`, `FORWARDED_IP` or `IP`. Default: `IP`.
        /// </summary>
        public readonly string? AggregateKeyType;
        /// <summary>
        /// Aggregate the request counts using one or more web request components as the aggregate keys. See `custom_key` below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupRuleStatementRateBasedStatementCustomKey> CustomKeys;
        /// <summary>
        /// The amount of time, in seconds, that AWS WAF should include in its request counts, looking back from the current time. Valid values are `60`, `120`, `300`, and `600`. Defaults to `300` (5 minutes).
        /// 
        /// **NOTE:** This setting doesn't determine how often AWS WAF checks the rate, but how far back it looks each time it checks. AWS WAF checks the rate about every 10 seconds.
        /// </summary>
        public readonly int? EvaluationWindowSec;
        /// <summary>
        /// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See Forwarded IP Config below for details.
        /// </summary>
        public readonly Outputs.RuleGroupRuleStatementRateBasedStatementForwardedIpConfig? ForwardedIpConfig;
        /// <summary>
        /// The limit on requests per 5-minute period for a single originating IP address.
        /// </summary>
        public readonly int Limit;
        /// <summary>
        /// An optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See Statement above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
        /// </summary>
        public readonly Outputs.RuleGroupRuleStatementRateBasedStatementScopeDownStatement? ScopeDownStatement;

        [OutputConstructor]
        private RuleGroupRuleStatementRateBasedStatement(
            string? aggregateKeyType,

            ImmutableArray<Outputs.RuleGroupRuleStatementRateBasedStatementCustomKey> customKeys,

            int? evaluationWindowSec,

            Outputs.RuleGroupRuleStatementRateBasedStatementForwardedIpConfig? forwardedIpConfig,

            int limit,

            Outputs.RuleGroupRuleStatementRateBasedStatementScopeDownStatement? scopeDownStatement)
        {
            AggregateKeyType = aggregateKeyType;
            CustomKeys = customKeys;
            EvaluationWindowSec = evaluationWindowSec;
            ForwardedIpConfig = forwardedIpConfig;
            Limit = limit;
            ScopeDownStatement = scopeDownStatement;
        }
    }
}
