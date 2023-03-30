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
    public sealed class GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustResult
    {
        /// <summary>
        /// An AWS Certificate Manager (ACM) certificate.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustAcmResult> Acms;
        /// <summary>
        /// TLS validation context trust for a local file certificate.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustFileResult> Files;
        /// <summary>
        /// TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustSdResult> Sds;

        [OutputConstructor]
        private GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustResult(
            ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustAcmResult> acms,

            ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustFileResult> files,

            ImmutableArray<Outputs.GetVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustSdResult> sds)
        {
            Acms = acms;
            Files = files;
            Sds = sds;
        }
    }
}
