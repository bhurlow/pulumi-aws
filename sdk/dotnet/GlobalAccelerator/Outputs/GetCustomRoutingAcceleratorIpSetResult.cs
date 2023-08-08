// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GlobalAccelerator.Outputs
{

    [OutputType]
    public sealed class GetCustomRoutingAcceleratorIpSetResult
    {
        public readonly ImmutableArray<string> IpAddresses;
        public readonly string IpFamily;

        [OutputConstructor]
        private GetCustomRoutingAcceleratorIpSetResult(
            ImmutableArray<string> ipAddresses,

            string ipFamily)
        {
            IpAddresses = ipAddresses;
            IpFamily = ipFamily;
        }
    }
}
