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
    public sealed class VirtualNodeSpecListenerConnectionPool
    {
        /// <summary>
        /// Connection pool information for gRPC listeners.
        /// </summary>
        public readonly Outputs.VirtualNodeSpecListenerConnectionPoolGrpc? Grpc;
        /// <summary>
        /// Connection pool information for HTTP listeners.
        /// </summary>
        public readonly Outputs.VirtualNodeSpecListenerConnectionPoolHttp? Http;
        /// <summary>
        /// Connection pool information for HTTP2 listeners.
        /// </summary>
        public readonly Outputs.VirtualNodeSpecListenerConnectionPoolHttp2? Http2;
        /// <summary>
        /// Connection pool information for TCP listeners.
        /// </summary>
        public readonly Outputs.VirtualNodeSpecListenerConnectionPoolTcp? Tcp;

        [OutputConstructor]
        private VirtualNodeSpecListenerConnectionPool(
            Outputs.VirtualNodeSpecListenerConnectionPoolGrpc? grpc,

            Outputs.VirtualNodeSpecListenerConnectionPoolHttp? http,

            Outputs.VirtualNodeSpecListenerConnectionPoolHttp2? http2,

            Outputs.VirtualNodeSpecListenerConnectionPoolTcp? tcp)
        {
            Grpc = grpc;
            Http = http;
            Http2 = http2;
            Tcp = tcp;
        }
    }
}
