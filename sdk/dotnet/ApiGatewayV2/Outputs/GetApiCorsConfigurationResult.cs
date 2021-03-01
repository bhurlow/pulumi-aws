// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2.Outputs
{

    [OutputType]
    public sealed class GetApiCorsConfigurationResult
    {
        /// <summary>
        /// Whether credentials are included in the CORS request.
        /// </summary>
        public readonly bool AllowCredentials;
        /// <summary>
        /// The set of allowed HTTP headers.
        /// </summary>
        public readonly ImmutableArray<string> AllowHeaders;
        /// <summary>
        /// The set of allowed HTTP methods.
        /// </summary>
        public readonly ImmutableArray<string> AllowMethods;
        /// <summary>
        /// The set of allowed origins.
        /// </summary>
        public readonly ImmutableArray<string> AllowOrigins;
        /// <summary>
        /// The set of exposed HTTP headers.
        /// </summary>
        public readonly ImmutableArray<string> ExposeHeaders;
        /// <summary>
        /// The number of seconds that the browser should cache preflight request results.
        /// </summary>
        public readonly int MaxAge;

        [OutputConstructor]
        private GetApiCorsConfigurationResult(
            bool allowCredentials,

            ImmutableArray<string> allowHeaders,

            ImmutableArray<string> allowMethods,

            ImmutableArray<string> allowOrigins,

            ImmutableArray<string> exposeHeaders,

            int maxAge)
        {
            AllowCredentials = allowCredentials;
            AllowHeaders = allowHeaders;
            AllowMethods = allowMethods;
            AllowOrigins = allowOrigins;
            ExposeHeaders = exposeHeaders;
            MaxAge = maxAge;
        }
    }
}
