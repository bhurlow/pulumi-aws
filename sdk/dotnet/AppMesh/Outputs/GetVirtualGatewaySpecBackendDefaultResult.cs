// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Outputs
{

    [OutputType]
    public sealed class GetVirtualGatewaySpecBackendDefaultResult
    {
        /// <summary>
        /// Default client policy for virtual gateway backends.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyResult> ClientPolicies;

        [OutputConstructor]
        private GetVirtualGatewaySpecBackendDefaultResult(ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyResult> clientPolicies)
        {
            ClientPolicies = clientPolicies;
        }
    }
}
