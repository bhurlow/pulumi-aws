// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dlm.Outputs
{

    [OutputType]
    public sealed class LifecyclePolicyPolicyDetailsScheduleDeprecateRule
    {
        /// <summary>
        /// Specifies the number of oldest AMIs to deprecate. Must be an integer between `1` and `1000`. Conflicts with `interval` and `interval_unit`.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
        /// </summary>
        public readonly string? IntervalUnit;

        [OutputConstructor]
        private LifecyclePolicyPolicyDetailsScheduleDeprecateRule(
            int? count,

            int? interval,

            string? intervalUnit)
        {
            Count = count;
            Interval = interval;
            IntervalUnit = intervalUnit;
        }
    }
}
