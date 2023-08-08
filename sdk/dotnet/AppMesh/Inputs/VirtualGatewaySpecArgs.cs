// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class VirtualGatewaySpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults for backends.
        /// </summary>
        [Input("backendDefaults")]
        public Input<Inputs.VirtualGatewaySpecBackendDefaultsArgs>? BackendDefaults { get; set; }

        [Input("listeners", required: true)]
        private InputList<Inputs.VirtualGatewaySpecListenerArgs>? _listeners;

        /// <summary>
        /// Listeners that the mesh endpoint is expected to receive inbound traffic from. You can specify one listener.
        /// </summary>
        public InputList<Inputs.VirtualGatewaySpecListenerArgs> Listeners
        {
            get => _listeners ?? (_listeners = new InputList<Inputs.VirtualGatewaySpecListenerArgs>());
            set => _listeners = value;
        }

        /// <summary>
        /// Inbound and outbound access logging information for the virtual gateway.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.VirtualGatewaySpecLoggingArgs>? Logging { get; set; }

        public VirtualGatewaySpecArgs()
        {
        }
        public static new VirtualGatewaySpecArgs Empty => new VirtualGatewaySpecArgs();
    }
}
