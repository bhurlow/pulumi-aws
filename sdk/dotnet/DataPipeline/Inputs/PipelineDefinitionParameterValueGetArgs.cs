// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DataPipeline.Inputs
{

    public sealed class PipelineDefinitionParameterValueGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the parameter value.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Field value, expressed as a String.
        /// </summary>
        [Input("stringValue", required: true)]
        public Input<string> StringValue { get; set; } = null!;

        public PipelineDefinitionParameterValueGetArgs()
        {
        }
    }
}
