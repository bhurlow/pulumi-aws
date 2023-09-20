// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fis
{
    /// <summary>
    /// Provides an FIS Experiment Template, which can be used to run an experiment.
    /// An experiment template contains one or more actions to run on specified targets during an experiment.
    /// It also contains the stop conditions that prevent the experiment from going out of bounds.
    /// See [Amazon Fault Injection Simulator](https://docs.aws.amazon.com/fis/index.html)
    /// for more information.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Fis.ExperimentTemplate("example", new()
    ///     {
    ///         Description = "example",
    ///         RoleArn = aws_iam_role.Example.Arn,
    ///         StopConditions = new[]
    ///         {
    ///             new Aws.Fis.Inputs.ExperimentTemplateStopConditionArgs
    ///             {
    ///                 Source = "none",
    ///             },
    ///         },
    ///         Actions = new[]
    ///         {
    ///             new Aws.Fis.Inputs.ExperimentTemplateActionArgs
    ///             {
    ///                 Name = "example-action",
    ///                 ActionId = "aws:ec2:terminate-instances",
    ///                 Target = new Aws.Fis.Inputs.ExperimentTemplateActionTargetArgs
    ///                 {
    ///                     Key = "Instances",
    ///                     Value = "example-target",
    ///                 },
    ///             },
    ///         },
    ///         Targets = new[]
    ///         {
    ///             new Aws.Fis.Inputs.ExperimentTemplateTargetArgs
    ///             {
    ///                 Name = "example-target",
    ///                 ResourceType = "aws:ec2:instance",
    ///                 SelectionMode = "COUNT(1)",
    ///                 ResourceTags = new[]
    ///                 {
    ///                     new Aws.Fis.Inputs.ExperimentTemplateTargetResourceTagArgs
    ///                     {
    ///                         Key = "env",
    ///                         Value = "example",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import FIS Experiment Templates using the `id`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:fis/experimentTemplate:ExperimentTemplate template EXT123AbCdEfGhIjK
    /// ```
    /// </summary>
    [AwsResourceType("aws:fis/experimentTemplate:ExperimentTemplate")]
    public partial class ExperimentTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to be performed during an experiment. See below.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.ExperimentTemplateAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// Description for the experiment template.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The configuration for experiment logging. See below.
        /// </summary>
        [Output("logConfiguration")]
        public Output<Outputs.ExperimentTemplateLogConfiguration?> LogConfiguration { get; private set; } = null!;

        /// <summary>
        /// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// When an ongoing experiment should be stopped. See below.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("stopConditions")]
        public Output<ImmutableArray<Outputs.ExperimentTemplateStopCondition>> StopConditions { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Target of an action. See below.
        /// </summary>
        [Output("targets")]
        public Output<ImmutableArray<Outputs.ExperimentTemplateTarget>> Targets { get; private set; } = null!;


        /// <summary>
        /// Create a ExperimentTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExperimentTemplate(string name, ExperimentTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws:fis/experimentTemplate:ExperimentTemplate", name, args ?? new ExperimentTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExperimentTemplate(string name, Input<string> id, ExperimentTemplateState? state = null, CustomResourceOptions? options = null)
            : base("aws:fis/experimentTemplate:ExperimentTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExperimentTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExperimentTemplate Get(string name, Input<string> id, ExperimentTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ExperimentTemplate(name, id, state, options);
        }
    }

    public sealed class ExperimentTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.ExperimentTemplateActionArgs>? _actions;

        /// <summary>
        /// Action to be performed during an experiment. See below.
        /// </summary>
        public InputList<Inputs.ExperimentTemplateActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.ExperimentTemplateActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Description for the experiment template.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The configuration for experiment logging. See below.
        /// </summary>
        [Input("logConfiguration")]
        public Input<Inputs.ExperimentTemplateLogConfigurationArgs>? LogConfiguration { get; set; }

        /// <summary>
        /// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("stopConditions", required: true)]
        private InputList<Inputs.ExperimentTemplateStopConditionArgs>? _stopConditions;

        /// <summary>
        /// When an ongoing experiment should be stopped. See below.
        /// 
        /// The following arguments are optional:
        /// </summary>
        public InputList<Inputs.ExperimentTemplateStopConditionArgs> StopConditions
        {
            get => _stopConditions ?? (_stopConditions = new InputList<Inputs.ExperimentTemplateStopConditionArgs>());
            set => _stopConditions = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targets")]
        private InputList<Inputs.ExperimentTemplateTargetArgs>? _targets;

        /// <summary>
        /// Target of an action. See below.
        /// </summary>
        public InputList<Inputs.ExperimentTemplateTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.ExperimentTemplateTargetArgs>());
            set => _targets = value;
        }

        public ExperimentTemplateArgs()
        {
        }
        public static new ExperimentTemplateArgs Empty => new ExperimentTemplateArgs();
    }

    public sealed class ExperimentTemplateState : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.ExperimentTemplateActionGetArgs>? _actions;

        /// <summary>
        /// Action to be performed during an experiment. See below.
        /// </summary>
        public InputList<Inputs.ExperimentTemplateActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.ExperimentTemplateActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Description for the experiment template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The configuration for experiment logging. See below.
        /// </summary>
        [Input("logConfiguration")]
        public Input<Inputs.ExperimentTemplateLogConfigurationGetArgs>? LogConfiguration { get; set; }

        /// <summary>
        /// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("stopConditions")]
        private InputList<Inputs.ExperimentTemplateStopConditionGetArgs>? _stopConditions;

        /// <summary>
        /// When an ongoing experiment should be stopped. See below.
        /// 
        /// The following arguments are optional:
        /// </summary>
        public InputList<Inputs.ExperimentTemplateStopConditionGetArgs> StopConditions
        {
            get => _stopConditions ?? (_stopConditions = new InputList<Inputs.ExperimentTemplateStopConditionGetArgs>());
            set => _stopConditions = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        [Input("targets")]
        private InputList<Inputs.ExperimentTemplateTargetGetArgs>? _targets;

        /// <summary>
        /// Target of an action. See below.
        /// </summary>
        public InputList<Inputs.ExperimentTemplateTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.ExperimentTemplateTargetGetArgs>());
            set => _targets = value;
        }

        public ExperimentTemplateState()
        {
        }
        public static new ExperimentTemplateState Empty => new ExperimentTemplateState();
    }
}
