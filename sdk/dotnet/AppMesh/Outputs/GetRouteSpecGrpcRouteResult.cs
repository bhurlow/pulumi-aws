// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Outputs
{

    [OutputType]
    public sealed class GetRouteSpecGrpcRouteResult
    {
        /// <summary>
        /// Action to take if a match is determined.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteSpecGrpcRouteActionResult> Actions;
        /// <summary>
        /// Criteria for determining an HTTP request match.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteSpecGrpcRouteMatchResult> Matches;
        /// <summary>
        /// Retry policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteSpecGrpcRouteRetryPolicyResult> RetryPolicies;
        /// <summary>
        /// Types of timeouts.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteSpecGrpcRouteTimeoutResult> Timeouts;

        [OutputConstructor]
        private GetRouteSpecGrpcRouteResult(
            ImmutableArray<Outputs.GetRouteSpecGrpcRouteActionResult> actions,

            ImmutableArray<Outputs.GetRouteSpecGrpcRouteMatchResult> matches,

            ImmutableArray<Outputs.GetRouteSpecGrpcRouteRetryPolicyResult> retryPolicies,

            ImmutableArray<Outputs.GetRouteSpecGrpcRouteTimeoutResult> timeouts)
        {
            Actions = actions;
            Matches = matches;
            RetryPolicies = retryPolicies;
            Timeouts = timeouts;
        }
    }
}
