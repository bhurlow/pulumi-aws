// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pipes.Inputs
{

    public sealed class PipeEnrichmentParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge ApiDestination. If you specify an API Gateway REST API or EventBridge ApiDestination as a target, you can use this parameter to specify headers, path parameters, and query string keys/values as part of your target invoking request. If you're using ApiDestinations, the corresponding Connection can also have these values configured. In case of any conflicting keys, values from the Connection take precedence. Detailed below.
        /// </summary>
        [Input("httpParameters")]
        public Input<Inputs.PipeEnrichmentParametersHttpParametersArgs>? HttpParameters { get; set; }

        /// <summary>
        /// Valid JSON text passed to the target. In this case, nothing from the event itself is passed to the target. Maximum length of 8192 characters.
        /// </summary>
        [Input("inputTemplate")]
        public Input<string>? InputTemplate { get; set; }

        public PipeEnrichmentParametersArgs()
        {
        }
        public static new PipeEnrichmentParametersArgs Empty => new PipeEnrichmentParametersArgs();
    }
}
