// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub.Inputs
{

    public sealed class InsightFiltersNetworkDestinationIpv6GetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A finding's CIDR value.
        /// </summary>
        [Input("cidr", required: true)]
        public Input<string> Cidr { get; set; } = null!;

        public InsightFiltersNetworkDestinationIpv6GetArgs()
        {
        }
    }
}
