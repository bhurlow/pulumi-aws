// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeCommit
{
    /// <summary>
    /// Provides a CodeCommit Trigger Resource.
    /// 
    /// &gt; **NOTE on CodeCommit**: The CodeCommit is not yet rolled out
    /// in all regions - available regions are listed
    /// [the AWS Docs](https://docs.aws.amazon.com/general/latest/gr/rande.html#codecommit_region).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codecommit_trigger.html.markdown.
    /// </summary>
    public partial class Trigger : Pulumi.CustomResource
    {
        [Output("configurationId")]
        public Output<string> ConfigurationId { get; private set; } = null!;

        /// <summary>
        /// The name for the repository. This needs to be less than 100 characters.
        /// </summary>
        [Output("repositoryName")]
        public Output<string> RepositoryName { get; private set; } = null!;

        [Output("triggers")]
        public Output<ImmutableArray<Outputs.TriggerTriggers>> Triggers { get; private set; } = null!;


        /// <summary>
        /// Create a Trigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trigger(string name, TriggerArgs args, CustomResourceOptions? options = null)
            : base("aws:codecommit/trigger:Trigger", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Trigger(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
            : base("aws:codecommit/trigger:Trigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trigger Get(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new Trigger(name, id, state, options);
        }
    }

    public sealed class TriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for the repository. This needs to be less than 100 characters.
        /// </summary>
        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        [Input("triggers", required: true)]
        private InputList<Inputs.TriggerTriggersArgs>? _triggers;
        public InputList<Inputs.TriggerTriggersArgs> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<Inputs.TriggerTriggersArgs>());
            set => _triggers = value;
        }

        public TriggerArgs()
        {
        }
    }

    public sealed class TriggerState : Pulumi.ResourceArgs
    {
        [Input("configurationId")]
        public Input<string>? ConfigurationId { get; set; }

        /// <summary>
        /// The name for the repository. This needs to be less than 100 characters.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        [Input("triggers")]
        private InputList<Inputs.TriggerTriggersGetArgs>? _triggers;
        public InputList<Inputs.TriggerTriggersGetArgs> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<Inputs.TriggerTriggersGetArgs>());
            set => _triggers = value;
        }

        public TriggerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class TriggerTriggersArgs : Pulumi.ResourceArgs
    {
        [Input("branches")]
        private InputList<string>? _branches;

        /// <summary>
        /// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
        /// </summary>
        public InputList<string> Branches
        {
            get => _branches ?? (_branches = new InputList<string>());
            set => _branches = value;
        }

        /// <summary>
        /// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
        /// </summary>
        [Input("customData")]
        public Input<string>? CustomData { get; set; }

        /// <summary>
        /// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
        /// </summary>
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        [Input("events", required: true)]
        private InputList<string>? _events;

        /// <summary>
        /// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
        /// </summary>
        public InputList<string> Events
        {
            get => _events ?? (_events = new InputList<string>());
            set => _events = value;
        }

        /// <summary>
        /// The name of the trigger.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public TriggerTriggersArgs()
        {
        }
    }

    public sealed class TriggerTriggersGetArgs : Pulumi.ResourceArgs
    {
        [Input("branches")]
        private InputList<string>? _branches;

        /// <summary>
        /// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
        /// </summary>
        public InputList<string> Branches
        {
            get => _branches ?? (_branches = new InputList<string>());
            set => _branches = value;
        }

        /// <summary>
        /// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
        /// </summary>
        [Input("customData")]
        public Input<string>? CustomData { get; set; }

        /// <summary>
        /// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
        /// </summary>
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        [Input("events", required: true)]
        private InputList<string>? _events;

        /// <summary>
        /// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
        /// </summary>
        public InputList<string> Events
        {
            get => _events ?? (_events = new InputList<string>());
            set => _events = value;
        }

        /// <summary>
        /// The name of the trigger.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public TriggerTriggersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class TriggerTriggers
    {
        /// <summary>
        /// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
        /// </summary>
        public readonly ImmutableArray<string> Branches;
        /// <summary>
        /// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
        /// </summary>
        public readonly string? CustomData;
        /// <summary>
        /// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
        /// </summary>
        public readonly string DestinationArn;
        /// <summary>
        /// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
        /// </summary>
        public readonly ImmutableArray<string> Events;
        /// <summary>
        /// The name of the trigger.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private TriggerTriggers(
            ImmutableArray<string> branches,
            string? customData,
            string destinationArn,
            ImmutableArray<string> events,
            string name)
        {
            Branches = branches;
            CustomData = customData;
            DestinationArn = destinationArn;
            Events = events;
            Name = name;
        }
    }
    }
}
