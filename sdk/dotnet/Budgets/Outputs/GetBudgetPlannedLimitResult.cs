// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Budgets.Outputs
{

    [OutputType]
    public sealed class GetBudgetPlannedLimitResult
    {
        /// <summary>
        /// The cost or usage amount that's associated with a budget forecast, actual spend, or budget threshold. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
        /// </summary>
        public readonly string Amount;
        /// <summary>
        /// (Required) The start time of the budget limit. Format: `2017-01-01_12:00`. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The unit of measurement that's used for the budget forecast, actual spend, or budget threshold, such as USD or GBP. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
        /// </summary>
        public readonly string Unit;

        [OutputConstructor]
        private GetBudgetPlannedLimitResult(
            string amount,

            string startTime,

            string unit)
        {
            Amount = amount;
            StartTime = startTime;
            Unit = unit;
        }
    }
}
