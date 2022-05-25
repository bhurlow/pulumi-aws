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
    public sealed class GetTagsFilterOrTagsResult
    {
        /// <summary>
        /// key that's used to sort the data. Valid values are: `BlendedCost`,  `UnblendedCost`, `AmortizedCost`, `NetAmortizedCost`, `NetUnblendedCost`, `UsageQuantity`, `NormalizedUsageAmount`.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// Match options that you can use to filter your results. MatchOptions is only applicable for actions related to cost category. The default values for MatchOptions is `EQUALS` and `CASE_SENSITIVE`. Valid values are: `EQUALS`,  `ABSENT`, `STARTS_WITH`, `ENDS_WITH`, `CONTAINS`, `CASE_SENSITIVE`, `CASE_INSENSITIVE`.
        /// </summary>
        public readonly ImmutableArray<string> MatchOptions;
        /// <summary>
        /// Specific value of the Cost Category.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetTagsFilterOrTagsResult(
            string? key,

            ImmutableArray<string> matchOptions,

            ImmutableArray<string> values)
        {
            Key = key;
            MatchOptions = matchOptions;
            Values = values;
        }
    }
}
