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
    public sealed class RuleGroupRuleStatementOrStatementStatementNotStatementStatementGeoMatchStatement
    {
        /// <summary>
        /// An array of two-character country codes, for example, [ "US", "CN" ], from the alpha-2 country ISO codes of the `ISO 3166` international standard. See the [documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_GeoMatchStatement.html) for valid values.
        /// </summary>
        public readonly ImmutableArray<string> CountryCodes;

        [OutputConstructor]
        private RuleGroupRuleStatementOrStatementStatementNotStatementStatementGeoMatchStatement(ImmutableArray<string> countryCodes)
        {
            CountryCodes = countryCodes;
        }
    }
}
