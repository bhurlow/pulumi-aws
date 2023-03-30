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
    public sealed class GetRouteSpecHttp2RouteTimeoutResult
    {
        public readonly ImmutableArray<Outputs.GetRouteSpecHttp2RouteTimeoutIdleResult> Idles;
        public readonly ImmutableArray<Outputs.GetRouteSpecHttp2RouteTimeoutPerRequestResult> PerRequests;

        [OutputConstructor]
        private GetRouteSpecHttp2RouteTimeoutResult(
            ImmutableArray<Outputs.GetRouteSpecHttp2RouteTimeoutIdleResult> idles,

            ImmutableArray<Outputs.GetRouteSpecHttp2RouteTimeoutPerRequestResult> perRequests)
        {
            Idles = idles;
            PerRequests = perRequests;
        }
    }
}
