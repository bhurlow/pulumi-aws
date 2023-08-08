// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class GatewayRouteSpecHttp2RouteMatchGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("headers")]
        private InputList<Inputs.GatewayRouteSpecHttp2RouteMatchHeaderGetArgs>? _headers;

        /// <summary>
        /// Client request headers to match on.
        /// </summary>
        public InputList<Inputs.GatewayRouteSpecHttp2RouteMatchHeaderGetArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.GatewayRouteSpecHttp2RouteMatchHeaderGetArgs>());
            set => _headers = value;
        }

        /// <summary>
        /// Host name to rewrite.
        /// </summary>
        [Input("hostname")]
        public Input<Inputs.GatewayRouteSpecHttp2RouteMatchHostnameGetArgs>? Hostname { get; set; }

        /// <summary>
        /// Exact path to rewrite.
        /// </summary>
        [Input("path")]
        public Input<Inputs.GatewayRouteSpecHttp2RouteMatchPathGetArgs>? Path { get; set; }

        /// <summary>
        /// The port number that corresponds to the target for Virtual Service provider port. This is required when the provider (router or node) of the Virtual Service has multiple listeners.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Specified beginning characters to rewrite.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("queryParameters")]
        private InputList<Inputs.GatewayRouteSpecHttp2RouteMatchQueryParameterGetArgs>? _queryParameters;

        /// <summary>
        /// Client request query parameters to match on.
        /// </summary>
        public InputList<Inputs.GatewayRouteSpecHttp2RouteMatchQueryParameterGetArgs> QueryParameters
        {
            get => _queryParameters ?? (_queryParameters = new InputList<Inputs.GatewayRouteSpecHttp2RouteMatchQueryParameterGetArgs>());
            set => _queryParameters = value;
        }

        public GatewayRouteSpecHttp2RouteMatchGetArgs()
        {
        }
        public static new GatewayRouteSpecHttp2RouteMatchGetArgs Empty => new GatewayRouteSpecHttp2RouteMatchGetArgs();
    }
}
