// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class FlowSourceFlowConfigSourceConnectorPropertiesServiceNowGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object specified in the Veeva flow source.
        /// </summary>
        [Input("object", required: true)]
        public Input<string> Object { get; set; } = null!;

        public FlowSourceFlowConfigSourceConnectorPropertiesServiceNowGetArgs()
        {
        }
    }
}
