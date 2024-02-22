// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class EndpointConfigurationShadowProductionVariantRoutingConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets how the endpoint routes incoming traffic. Valid values are `LEAST_OUTSTANDING_REQUESTS` and `RANDOM`. `LEAST_OUTSTANDING_REQUESTS` routes requests to the specific instances that have more capacity to process them. `RANDOM` routes each request to a randomly chosen instance.
        /// </summary>
        [Input("routingStrategy", required: true)]
        public Input<string> RoutingStrategy { get; set; } = null!;

        public EndpointConfigurationShadowProductionVariantRoutingConfigGetArgs()
        {
        }
        public static new EndpointConfigurationShadowProductionVariantRoutingConfigGetArgs Empty => new EndpointConfigurationShadowProductionVariantRoutingConfigGetArgs();
    }
}
