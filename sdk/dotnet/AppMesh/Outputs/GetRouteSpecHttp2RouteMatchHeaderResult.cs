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
    public sealed class GetRouteSpecHttp2RouteMatchHeaderResult
    {
        public readonly bool Invert;
        /// <summary>
        /// Criteria for determining an HTTP request match.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteSpecHttp2RouteMatchHeaderMatchResult> Matches;
        /// <summary>
        /// Name of the route.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetRouteSpecHttp2RouteMatchHeaderResult(
            bool invert,

            ImmutableArray<Outputs.GetRouteSpecHttp2RouteMatchHeaderMatchResult> matches,

            string name)
        {
            Invert = invert;
            Matches = matches;
            Name = name;
        }
    }
}
