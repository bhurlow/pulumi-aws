// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DataPipeline.Inputs
{

    public sealed class PipelineDefinitionPipelineObjectArgs : Pulumi.ResourceArgs
    {
        [Input("fields")]
        private InputList<Inputs.PipelineDefinitionPipelineObjectFieldArgs>? _fields;

        /// <summary>
        /// Configuration block for Key-value pairs that define the properties of the object. See below
        /// </summary>
        public InputList<Inputs.PipelineDefinitionPipelineObjectFieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.PipelineDefinitionPipelineObjectFieldArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// ID of the parameter value.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// ARN of the storage connector.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public PipelineDefinitionPipelineObjectArgs()
        {
        }
    }
}
