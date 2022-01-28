// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DataPipeline
{
    /// <summary>
    /// Provides a DataPipeline Pipeline Definition resource.
    /// 
    /// ## Import
    /// 
    /// `aws_datapipeline_pipeline_definition` can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:datapipeline/pipelineDefinition:PipelineDefinition example df-1234567890
    /// ```
    /// </summary>
    [AwsResourceType("aws:datapipeline/pipelineDefinition:PipelineDefinition")]
    public partial class PipelineDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration block for the parameter objects used in the pipeline definition. See below
        /// </summary>
        [Output("parameterObjects")]
        public Output<ImmutableArray<Outputs.PipelineDefinitionParameterObject>> ParameterObjects { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the parameter values used in the pipeline definition. See below
        /// </summary>
        [Output("parameterValues")]
        public Output<ImmutableArray<Outputs.PipelineDefinitionParameterValue>> ParameterValues { get; private set; } = null!;

        /// <summary>
        /// ID of the pipeline.
        /// </summary>
        [Output("pipelineId")]
        public Output<string> PipelineId { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the objects that define the pipeline. See below
        /// </summary>
        [Output("pipelineObjects")]
        public Output<ImmutableArray<Outputs.PipelineDefinitionPipelineObject>> PipelineObjects { get; private set; } = null!;


        /// <summary>
        /// Create a PipelineDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PipelineDefinition(string name, PipelineDefinitionArgs args, CustomResourceOptions? options = null)
            : base("aws:datapipeline/pipelineDefinition:PipelineDefinition", name, args ?? new PipelineDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PipelineDefinition(string name, Input<string> id, PipelineDefinitionState? state = null, CustomResourceOptions? options = null)
            : base("aws:datapipeline/pipelineDefinition:PipelineDefinition", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PipelineDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PipelineDefinition Get(string name, Input<string> id, PipelineDefinitionState? state = null, CustomResourceOptions? options = null)
        {
            return new PipelineDefinition(name, id, state, options);
        }
    }

    public sealed class PipelineDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("parameterObjects")]
        private InputList<Inputs.PipelineDefinitionParameterObjectArgs>? _parameterObjects;

        /// <summary>
        /// Configuration block for the parameter objects used in the pipeline definition. See below
        /// </summary>
        public InputList<Inputs.PipelineDefinitionParameterObjectArgs> ParameterObjects
        {
            get => _parameterObjects ?? (_parameterObjects = new InputList<Inputs.PipelineDefinitionParameterObjectArgs>());
            set => _parameterObjects = value;
        }

        [Input("parameterValues")]
        private InputList<Inputs.PipelineDefinitionParameterValueArgs>? _parameterValues;

        /// <summary>
        /// Configuration block for the parameter values used in the pipeline definition. See below
        /// </summary>
        public InputList<Inputs.PipelineDefinitionParameterValueArgs> ParameterValues
        {
            get => _parameterValues ?? (_parameterValues = new InputList<Inputs.PipelineDefinitionParameterValueArgs>());
            set => _parameterValues = value;
        }

        /// <summary>
        /// ID of the pipeline.
        /// </summary>
        [Input("pipelineId", required: true)]
        public Input<string> PipelineId { get; set; } = null!;

        [Input("pipelineObjects", required: true)]
        private InputList<Inputs.PipelineDefinitionPipelineObjectArgs>? _pipelineObjects;

        /// <summary>
        /// Configuration block for the objects that define the pipeline. See below
        /// </summary>
        public InputList<Inputs.PipelineDefinitionPipelineObjectArgs> PipelineObjects
        {
            get => _pipelineObjects ?? (_pipelineObjects = new InputList<Inputs.PipelineDefinitionPipelineObjectArgs>());
            set => _pipelineObjects = value;
        }

        public PipelineDefinitionArgs()
        {
        }
    }

    public sealed class PipelineDefinitionState : Pulumi.ResourceArgs
    {
        [Input("parameterObjects")]
        private InputList<Inputs.PipelineDefinitionParameterObjectGetArgs>? _parameterObjects;

        /// <summary>
        /// Configuration block for the parameter objects used in the pipeline definition. See below
        /// </summary>
        public InputList<Inputs.PipelineDefinitionParameterObjectGetArgs> ParameterObjects
        {
            get => _parameterObjects ?? (_parameterObjects = new InputList<Inputs.PipelineDefinitionParameterObjectGetArgs>());
            set => _parameterObjects = value;
        }

        [Input("parameterValues")]
        private InputList<Inputs.PipelineDefinitionParameterValueGetArgs>? _parameterValues;

        /// <summary>
        /// Configuration block for the parameter values used in the pipeline definition. See below
        /// </summary>
        public InputList<Inputs.PipelineDefinitionParameterValueGetArgs> ParameterValues
        {
            get => _parameterValues ?? (_parameterValues = new InputList<Inputs.PipelineDefinitionParameterValueGetArgs>());
            set => _parameterValues = value;
        }

        /// <summary>
        /// ID of the pipeline.
        /// </summary>
        [Input("pipelineId")]
        public Input<string>? PipelineId { get; set; }

        [Input("pipelineObjects")]
        private InputList<Inputs.PipelineDefinitionPipelineObjectGetArgs>? _pipelineObjects;

        /// <summary>
        /// Configuration block for the objects that define the pipeline. See below
        /// </summary>
        public InputList<Inputs.PipelineDefinitionPipelineObjectGetArgs> PipelineObjects
        {
            get => _pipelineObjects ?? (_pipelineObjects = new InputList<Inputs.PipelineDefinitionPipelineObjectGetArgs>());
            set => _pipelineObjects = value;
        }

        public PipelineDefinitionState()
        {
        }
    }
}
