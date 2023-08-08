// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class InstanceInstanceMarketOptionsSpotOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The behavior when a Spot Instance is interrupted. Valid values include `hibernate`, `stop`, `terminate` . The default is `terminate`.
        /// </summary>
        [Input("instanceInterruptionBehavior")]
        public Input<string>? InstanceInterruptionBehavior { get; set; }

        /// <summary>
        /// The maximum hourly price that you're willing to pay for a Spot Instance.
        /// </summary>
        [Input("maxPrice")]
        public Input<string>? MaxPrice { get; set; }

        /// <summary>
        /// The Spot Instance request type. Valid values include `one-time`, `persistent`. Persistent Spot Instance requests are only supported when the instance interruption behavior is either hibernate or stop. The default is `one-time`.
        /// </summary>
        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        /// <summary>
        /// The end date of the request, in UTC format (YYYY-MM-DDTHH:MM:SSZ). Supported only for persistent requests.
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public InstanceInstanceMarketOptionsSpotOptionsArgs()
        {
        }
        public static new InstanceInstanceMarketOptionsSpotOptionsArgs Empty => new InstanceInstanceMarketOptionsSpotOptionsArgs();
    }
}
