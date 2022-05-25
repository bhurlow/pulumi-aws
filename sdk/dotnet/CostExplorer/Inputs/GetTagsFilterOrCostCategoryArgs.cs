// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CostExplorer.Inputs
{

    public sealed class GetTagsFilterOrCostCategoryInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// key that's used to sort the data. Valid values are: `BlendedCost`,  `UnblendedCost`, `AmortizedCost`, `NetAmortizedCost`, `NetUnblendedCost`, `UsageQuantity`, `NormalizedUsageAmount`.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("matchOptions")]
        private InputList<string>? _matchOptions;

        /// <summary>
        /// Match options that you can use to filter your results. MatchOptions is only applicable for actions related to cost category. The default values for MatchOptions is `EQUALS` and `CASE_SENSITIVE`. Valid values are: `EQUALS`,  `ABSENT`, `STARTS_WITH`, `ENDS_WITH`, `CONTAINS`, `CASE_SENSITIVE`, `CASE_INSENSITIVE`.
        /// </summary>
        public InputList<string> MatchOptions
        {
            get => _matchOptions ?? (_matchOptions = new InputList<string>());
            set => _matchOptions = value;
        }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Specific value of the Cost Category.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GetTagsFilterOrCostCategoryInputArgs()
        {
        }
    }
}
