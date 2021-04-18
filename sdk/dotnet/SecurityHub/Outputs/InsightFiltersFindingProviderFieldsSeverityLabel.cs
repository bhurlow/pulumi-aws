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
    public sealed class InsightFiltersFindingProviderFieldsSeverityLabel
    {
        /// <summary>
        /// The condition to apply to a string value when querying for findings. Valid values include: `EQUALS` and `NOT_EQUALS`.
        /// </summary>
        public readonly string Comparison;
        /// <summary>
        /// A date range value for the date filter, provided as an Integer.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private InsightFiltersFindingProviderFieldsSeverityLabel(
            string comparison,

            string value)
        {
            Comparison = comparison;
            Value = value;
        }
    }
}
