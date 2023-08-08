// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Outputs
{

    [OutputType]
    public sealed class CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig
    {
        /// <summary>
        /// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for `cookie_behavior` are `none`, `whitelist`, `allExcept`, and `all`.
        /// </summary>
        public readonly string CookieBehavior;
        /// <summary>
        /// Object that contains a list of cookie names. See Items for more information.
        /// </summary>
        public readonly Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies? Cookies;

        [OutputConstructor]
        private CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig(
            string cookieBehavior,

            Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies? cookies)
        {
            CookieBehavior = cookieBehavior;
            Cookies = cookies;
        }
    }
}
