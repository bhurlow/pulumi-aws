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
    public sealed class InsightFiltersNetworkSourcePort
    {
        /// <summary>
        /// The equal-to condition to be applied to a single field when querying for findings, provided as a String.
        /// </summary>
        public readonly string? Eq;
        /// <summary>
        /// The greater-than-equal condition to be applied to a single field when querying for findings, provided as a String.
        /// </summary>
        public readonly string? Gte;
        /// <summary>
        /// The less-than-equal condition to be applied to a single field when querying for findings, provided as a String.
        /// </summary>
        public readonly string? Lte;

        [OutputConstructor]
        private InsightFiltersNetworkSourcePort(
            string? eq,

            string? gte,

            string? lte)
        {
            Eq = eq;
            Gte = gte;
            Lte = lte;
        }
    }
}
