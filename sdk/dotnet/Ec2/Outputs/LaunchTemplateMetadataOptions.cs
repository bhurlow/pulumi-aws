// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateMetadataOptions
    {
        /// <summary>
        /// Whether the metadata service is available. Can be `"enabled"` or `"disabled"`. (Default: `"enabled"`).
        /// </summary>
        public readonly string? HttpEndpoint;
        /// <summary>
        /// The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Can be an integer from `1` to `64`. (Default: `1`).
        /// </summary>
        public readonly int? HttpPutResponseHopLimit;
        /// <summary>
        /// Whether or not the metadata service requires session tokens, also referred to as _Instance Metadata Service Version 2 (IMDSv2)_. Can be `"optional"` or `"required"`. (Default: `"optional"`).
        /// </summary>
        public readonly string? HttpTokens;

        [OutputConstructor]
        private LaunchTemplateMetadataOptions(
            string? httpEndpoint,

            int? httpPutResponseHopLimit,

            string? httpTokens)
        {
            HttpEndpoint = httpEndpoint;
            HttpPutResponseHopLimit = httpPutResponseHopLimit;
            HttpTokens = httpTokens;
        }
    }
}
