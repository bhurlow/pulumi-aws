// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class InstanceInstanceMarketOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of market for the instance. Valid value is `spot`. Defaults to `spot`.
        /// </summary>
        [Input("marketType")]
        public Input<string>? MarketType { get; set; }

        /// <summary>
        /// Block to configure the options for Spot Instances. See Spot Options below for details on attributes.
        /// </summary>
        [Input("spotOptions")]
        public Input<Inputs.InstanceInstanceMarketOptionsSpotOptionsGetArgs>? SpotOptions { get; set; }

        public InstanceInstanceMarketOptionsGetArgs()
        {
        }
        public static new InstanceInstanceMarketOptionsGetArgs Empty => new InstanceInstanceMarketOptionsGetArgs();
    }
}
