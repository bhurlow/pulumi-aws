// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.NetworkFirewall.Inputs
{

    public sealed class FirewallSubnetMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The subnet's IP address type. Valida values: `"DUALSTACK"`, `"IPV4"`.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        /// <summary>
        /// The unique identifier for the subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public FirewallSubnetMappingArgs()
        {
        }
        public static new FirewallSubnetMappingArgs Empty => new FirewallSubnetMappingArgs();
    }
}
