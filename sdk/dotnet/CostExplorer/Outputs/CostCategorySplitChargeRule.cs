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
    public sealed class CostCategorySplitChargeRule
    {
        /// <summary>
        /// Method that's used to define how to split your source costs across your targets. Valid values are `FIXED`, `PROPORTIONAL`, `EVEN`
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// Configuration block for the parameters for a split charge method. This is only required for the `FIXED` method. See below.
        /// </summary>
        public readonly ImmutableArray<Outputs.CostCategorySplitChargeRuleParameter> Parameters;
        /// <summary>
        /// Cost Category value that you want to split.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// Cost Category values that you want to split costs across. These values can't be used as a source in other split charge rules.
        /// </summary>
        public readonly ImmutableArray<string> Targets;

        [OutputConstructor]
        private CostCategorySplitChargeRule(
            string method,

            ImmutableArray<Outputs.CostCategorySplitChargeRuleParameter> parameters,

            string source,

            ImmutableArray<string> targets)
        {
            Method = method;
            Parameters = parameters;
            Source = source;
            Targets = targets;
        }
    }
}
