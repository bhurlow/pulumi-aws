// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Auditmanager
{
    /// <summary>
    /// Resource for managing an AWS Audit Manager Control.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Auditmanager.Control("example", new()
    ///     {
    ///         ControlMappingSources = new[]
    ///         {
    ///             new Aws.Auditmanager.Inputs.ControlControlMappingSourceArgs
    ///             {
    ///                 SourceName = "example",
    ///                 SourceSetUpOption = "Procedural_Controls_Mapping",
    ///                 SourceType = "MANUAL",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import an Audit Manager Control using the `id`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:auditmanager/control:Control example abc123-de45
    /// ```
    /// </summary>
    [AwsResourceType("aws:auditmanager/control:Control")]
    public partial class Control : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Recommended actions to carry out if the control isn't fulfilled.
        /// </summary>
        [Output("actionPlanInstructions")]
        public Output<string?> ActionPlanInstructions { get; private set; } = null!;

        /// <summary>
        /// Title of the action plan for remediating the control.
        /// </summary>
        [Output("actionPlanTitle")]
        public Output<string?> ActionPlanTitle { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the control.
        /// * `control_mapping_sources.*.source_id` - Unique identifier for the source.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Data mapping sources. See `control_mapping_sources` below.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("controlMappingSources")]
        public Output<ImmutableArray<Outputs.ControlControlMappingSource>> ControlMappingSources { get; private set; } = null!;

        /// <summary>
        /// Description of the control.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the control.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Steps to follow to determine if the control is satisfied.
        /// </summary>
        [Output("testingInformation")]
        public Output<string?> TestingInformation { get; private set; } = null!;

        /// <summary>
        /// Type of control, such as a custom control or a standard control.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Control resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Control(string name, ControlArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:auditmanager/control:Control", name, args ?? new ControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Control(string name, Input<string> id, ControlState? state = null, CustomResourceOptions? options = null)
            : base("aws:auditmanager/control:Control", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Control resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Control Get(string name, Input<string> id, ControlState? state = null, CustomResourceOptions? options = null)
        {
            return new Control(name, id, state, options);
        }
    }

    public sealed class ControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Recommended actions to carry out if the control isn't fulfilled.
        /// </summary>
        [Input("actionPlanInstructions")]
        public Input<string>? ActionPlanInstructions { get; set; }

        /// <summary>
        /// Title of the action plan for remediating the control.
        /// </summary>
        [Input("actionPlanTitle")]
        public Input<string>? ActionPlanTitle { get; set; }

        [Input("controlMappingSources")]
        private InputList<Inputs.ControlControlMappingSourceArgs>? _controlMappingSources;

        /// <summary>
        /// Data mapping sources. See `control_mapping_sources` below.
        /// 
        /// The following arguments are optional:
        /// </summary>
        public InputList<Inputs.ControlControlMappingSourceArgs> ControlMappingSources
        {
            get => _controlMappingSources ?? (_controlMappingSources = new InputList<Inputs.ControlControlMappingSourceArgs>());
            set => _controlMappingSources = value;
        }

        /// <summary>
        /// Description of the control.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the control.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Steps to follow to determine if the control is satisfied.
        /// </summary>
        [Input("testingInformation")]
        public Input<string>? TestingInformation { get; set; }

        public ControlArgs()
        {
        }
        public static new ControlArgs Empty => new ControlArgs();
    }

    public sealed class ControlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Recommended actions to carry out if the control isn't fulfilled.
        /// </summary>
        [Input("actionPlanInstructions")]
        public Input<string>? ActionPlanInstructions { get; set; }

        /// <summary>
        /// Title of the action plan for remediating the control.
        /// </summary>
        [Input("actionPlanTitle")]
        public Input<string>? ActionPlanTitle { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the control.
        /// * `control_mapping_sources.*.source_id` - Unique identifier for the source.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("controlMappingSources")]
        private InputList<Inputs.ControlControlMappingSourceGetArgs>? _controlMappingSources;

        /// <summary>
        /// Data mapping sources. See `control_mapping_sources` below.
        /// 
        /// The following arguments are optional:
        /// </summary>
        public InputList<Inputs.ControlControlMappingSourceGetArgs> ControlMappingSources
        {
            get => _controlMappingSources ?? (_controlMappingSources = new InputList<Inputs.ControlControlMappingSourceGetArgs>());
            set => _controlMappingSources = value;
        }

        /// <summary>
        /// Description of the control.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the control.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Steps to follow to determine if the control is satisfied.
        /// </summary>
        [Input("testingInformation")]
        public Input<string>? TestingInformation { get; set; }

        /// <summary>
        /// Type of control, such as a custom control or a standard control.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ControlState()
        {
        }
        public static new ControlState Empty => new ControlState();
    }
}
