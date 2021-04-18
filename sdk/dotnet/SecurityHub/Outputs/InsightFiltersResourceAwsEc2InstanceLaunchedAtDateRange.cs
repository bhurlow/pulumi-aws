// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub.Outputs
{

    [OutputType]
    public sealed class InsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange
    {
        /// <summary>
        /// A date range unit for the date filter. Valid values: `DAYS`.
        /// </summary>
        public readonly string Unit;
        /// <summary>
        /// A date range value for the date filter, provided as an Integer.
        /// </summary>
        public readonly int Value;

        [OutputConstructor]
        private InsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange(
            string unit,

            int value)
        {
            Unit = unit;
            Value = value;
        }
    }
}
