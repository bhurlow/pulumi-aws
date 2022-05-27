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
    public sealed class ResponseHeadersPolicyServerTimingHeadersConfig
    {
        /// <summary>
        /// A Boolean that determines whether CloudFront adds the `Server-Timing` header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// A number 0–100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the Server-Timing header to. Valid range: Minimum value of 0.0. Maximum value of 100.0.
        /// </summary>
        public readonly double SamplingRate;

        [OutputConstructor]
        private ResponseHeadersPolicyServerTimingHeadersConfig(
            bool enabled,

            double samplingRate)
        {
            Enabled = enabled;
            SamplingRate = samplingRate;
        }
    }
}
