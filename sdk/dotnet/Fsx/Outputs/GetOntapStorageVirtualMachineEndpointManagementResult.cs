// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fsx.Outputs
{

    [OutputType]
    public sealed class GetOntapStorageVirtualMachineEndpointManagementResult
    {
        public readonly string DnsName;
        public readonly ImmutableArray<string> IpAddresses;

        [OutputConstructor]
        private GetOntapStorageVirtualMachineEndpointManagementResult(
            string dnsName,

            ImmutableArray<string> ipAddresses)
        {
            DnsName = dnsName;
            IpAddresses = ipAddresses;
        }
    }
}
