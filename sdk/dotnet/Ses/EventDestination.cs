// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides an SES event destination
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_event_destination.html.markdown.
    /// </summary>
    public partial class EventDestination : Pulumi.CustomResource
    {
        /// <summary>
        /// CloudWatch destination for the events
        /// </summary>
        [Output("cloudwatchDestinations")]
        public Output<ImmutableArray<Outputs.EventDestinationCloudwatchDestinations>> CloudwatchDestinations { get; private set; } = null!;

        /// <summary>
        /// The name of the configuration set
        /// </summary>
        [Output("configurationSetName")]
        public Output<string> ConfigurationSetName { get; private set; } = null!;

        /// <summary>
        /// If true, the event destination will be enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Send the events to a kinesis firehose destination
        /// </summary>
        [Output("kinesisDestination")]
        public Output<Outputs.EventDestinationKinesisDestination?> KinesisDestination { get; private set; } = null!;

        /// <summary>
        /// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
        /// </summary>
        [Output("matchingTypes")]
        public Output<ImmutableArray<string>> MatchingTypes { get; private set; } = null!;

        /// <summary>
        /// The name of the event destination
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Send the events to an SNS Topic destination
        /// </summary>
        [Output("snsDestination")]
        public Output<Outputs.EventDestinationSnsDestination?> SnsDestination { get; private set; } = null!;


        /// <summary>
        /// Create a EventDestination resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventDestination(string name, EventDestinationArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/eventDestination:EventDestination", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EventDestination(string name, Input<string> id, EventDestinationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/eventDestination:EventDestination", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventDestination resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventDestination Get(string name, Input<string> id, EventDestinationState? state = null, CustomResourceOptions? options = null)
        {
            return new EventDestination(name, id, state, options);
        }
    }

    public sealed class EventDestinationArgs : Pulumi.ResourceArgs
    {
        [Input("cloudwatchDestinations")]
        private InputList<Inputs.EventDestinationCloudwatchDestinationsArgs>? _cloudwatchDestinations;

        /// <summary>
        /// CloudWatch destination for the events
        /// </summary>
        public InputList<Inputs.EventDestinationCloudwatchDestinationsArgs> CloudwatchDestinations
        {
            get => _cloudwatchDestinations ?? (_cloudwatchDestinations = new InputList<Inputs.EventDestinationCloudwatchDestinationsArgs>());
            set => _cloudwatchDestinations = value;
        }

        /// <summary>
        /// The name of the configuration set
        /// </summary>
        [Input("configurationSetName", required: true)]
        public Input<string> ConfigurationSetName { get; set; } = null!;

        /// <summary>
        /// If true, the event destination will be enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Send the events to a kinesis firehose destination
        /// </summary>
        [Input("kinesisDestination")]
        public Input<Inputs.EventDestinationKinesisDestinationArgs>? KinesisDestination { get; set; }

        [Input("matchingTypes", required: true)]
        private InputList<string>? _matchingTypes;

        /// <summary>
        /// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
        /// </summary>
        public InputList<string> MatchingTypes
        {
            get => _matchingTypes ?? (_matchingTypes = new InputList<string>());
            set => _matchingTypes = value;
        }

        /// <summary>
        /// The name of the event destination
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Send the events to an SNS Topic destination
        /// </summary>
        [Input("snsDestination")]
        public Input<Inputs.EventDestinationSnsDestinationArgs>? SnsDestination { get; set; }

        public EventDestinationArgs()
        {
        }
    }

    public sealed class EventDestinationState : Pulumi.ResourceArgs
    {
        [Input("cloudwatchDestinations")]
        private InputList<Inputs.EventDestinationCloudwatchDestinationsGetArgs>? _cloudwatchDestinations;

        /// <summary>
        /// CloudWatch destination for the events
        /// </summary>
        public InputList<Inputs.EventDestinationCloudwatchDestinationsGetArgs> CloudwatchDestinations
        {
            get => _cloudwatchDestinations ?? (_cloudwatchDestinations = new InputList<Inputs.EventDestinationCloudwatchDestinationsGetArgs>());
            set => _cloudwatchDestinations = value;
        }

        /// <summary>
        /// The name of the configuration set
        /// </summary>
        [Input("configurationSetName")]
        public Input<string>? ConfigurationSetName { get; set; }

        /// <summary>
        /// If true, the event destination will be enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Send the events to a kinesis firehose destination
        /// </summary>
        [Input("kinesisDestination")]
        public Input<Inputs.EventDestinationKinesisDestinationGetArgs>? KinesisDestination { get; set; }

        [Input("matchingTypes")]
        private InputList<string>? _matchingTypes;

        /// <summary>
        /// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
        /// </summary>
        public InputList<string> MatchingTypes
        {
            get => _matchingTypes ?? (_matchingTypes = new InputList<string>());
            set => _matchingTypes = value;
        }

        /// <summary>
        /// The name of the event destination
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Send the events to an SNS Topic destination
        /// </summary>
        [Input("snsDestination")]
        public Input<Inputs.EventDestinationSnsDestinationGetArgs>? SnsDestination { get; set; }

        public EventDestinationState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EventDestinationCloudwatchDestinationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default value for the event
        /// </summary>
        [Input("defaultValue", required: true)]
        public Input<string> DefaultValue { get; set; } = null!;

        /// <summary>
        /// The name for the dimension
        /// </summary>
        [Input("dimensionName", required: true)]
        public Input<string> DimensionName { get; set; } = null!;

        /// <summary>
        /// The source for the value. It can be either `"messageTag"` or `"emailHeader"`
        /// </summary>
        [Input("valueSource", required: true)]
        public Input<string> ValueSource { get; set; } = null!;

        public EventDestinationCloudwatchDestinationsArgs()
        {
        }
    }

    public sealed class EventDestinationCloudwatchDestinationsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default value for the event
        /// </summary>
        [Input("defaultValue", required: true)]
        public Input<string> DefaultValue { get; set; } = null!;

        /// <summary>
        /// The name for the dimension
        /// </summary>
        [Input("dimensionName", required: true)]
        public Input<string> DimensionName { get; set; } = null!;

        /// <summary>
        /// The source for the value. It can be either `"messageTag"` or `"emailHeader"`
        /// </summary>
        [Input("valueSource", required: true)]
        public Input<string> ValueSource { get; set; } = null!;

        public EventDestinationCloudwatchDestinationsGetArgs()
        {
        }
    }

    public sealed class EventDestinationKinesisDestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the role that has permissions to access the Kinesis Stream
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The ARN of the Kinesis Stream
        /// </summary>
        [Input("streamArn", required: true)]
        public Input<string> StreamArn { get; set; } = null!;

        public EventDestinationKinesisDestinationArgs()
        {
        }
    }

    public sealed class EventDestinationKinesisDestinationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the role that has permissions to access the Kinesis Stream
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The ARN of the Kinesis Stream
        /// </summary>
        [Input("streamArn", required: true)]
        public Input<string> StreamArn { get; set; } = null!;

        public EventDestinationKinesisDestinationGetArgs()
        {
        }
    }

    public sealed class EventDestinationSnsDestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the SNS topic
        /// </summary>
        [Input("topicArn", required: true)]
        public Input<string> TopicArn { get; set; } = null!;

        public EventDestinationSnsDestinationArgs()
        {
        }
    }

    public sealed class EventDestinationSnsDestinationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the SNS topic
        /// </summary>
        [Input("topicArn", required: true)]
        public Input<string> TopicArn { get; set; } = null!;

        public EventDestinationSnsDestinationGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EventDestinationCloudwatchDestinations
    {
        /// <summary>
        /// The default value for the event
        /// </summary>
        public readonly string DefaultValue;
        /// <summary>
        /// The name for the dimension
        /// </summary>
        public readonly string DimensionName;
        /// <summary>
        /// The source for the value. It can be either `"messageTag"` or `"emailHeader"`
        /// </summary>
        public readonly string ValueSource;

        [OutputConstructor]
        private EventDestinationCloudwatchDestinations(
            string defaultValue,
            string dimensionName,
            string valueSource)
        {
            DefaultValue = defaultValue;
            DimensionName = dimensionName;
            ValueSource = valueSource;
        }
    }

    [OutputType]
    public sealed class EventDestinationKinesisDestination
    {
        /// <summary>
        /// The ARN of the role that has permissions to access the Kinesis Stream
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The ARN of the Kinesis Stream
        /// </summary>
        public readonly string StreamArn;

        [OutputConstructor]
        private EventDestinationKinesisDestination(
            string roleArn,
            string streamArn)
        {
            RoleArn = roleArn;
            StreamArn = streamArn;
        }
    }

    [OutputType]
    public sealed class EventDestinationSnsDestination
    {
        /// <summary>
        /// The ARN of the SNS topic
        /// </summary>
        public readonly string TopicArn;

        [OutputConstructor]
        private EventDestinationSnsDestination(string topicArn)
        {
            TopicArn = topicArn;
        }
    }
    }
}
