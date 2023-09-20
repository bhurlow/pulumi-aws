// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive
{
    /// <summary>
    /// Resource for managing an AWS MediaLive Channel.
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
    ///     var example = new Aws.MediaLive.Channel("example", new()
    ///     {
    ///         ChannelClass = "STANDARD",
    ///         RoleArn = aws_iam_role.Example.Arn,
    ///         InputSpecification = new Aws.MediaLive.Inputs.ChannelInputSpecificationArgs
    ///         {
    ///             Codec = "AVC",
    ///             InputResolution = "HD",
    ///             MaximumBitrate = "MAX_20_MBPS",
    ///         },
    ///         InputAttachments = new[]
    ///         {
    ///             new Aws.MediaLive.Inputs.ChannelInputAttachmentArgs
    ///             {
    ///                 InputAttachmentName = "example-input",
    ///                 InputId = aws_medialive_input.Example.Id,
    ///             },
    ///         },
    ///         Destinations = new[]
    ///         {
    ///             new Aws.MediaLive.Inputs.ChannelDestinationArgs
    ///             {
    ///                 Id = "destination",
    ///                 Settings = new[]
    ///                 {
    ///                     new Aws.MediaLive.Inputs.ChannelDestinationSettingArgs
    ///                     {
    ///                         Url = $"s3://{aws_s3_bucket.Main.Id}/test1",
    ///                     },
    ///                     new Aws.MediaLive.Inputs.ChannelDestinationSettingArgs
    ///                     {
    ///                         Url = $"s3://{aws_s3_bucket.Main2.Id}/test2",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         EncoderSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsArgs
    ///         {
    ///             TimecodeConfig = new Aws.MediaLive.Inputs.ChannelEncoderSettingsTimecodeConfigArgs
    ///             {
    ///                 Source = "EMBEDDED",
    ///             },
    ///             AudioDescriptions = new[]
    ///             {
    ///                 new Aws.MediaLive.Inputs.ChannelEncoderSettingsAudioDescriptionArgs
    ///                 {
    ///                     AudioSelectorName = "example audio selector",
    ///                     Name = "audio-selector",
    ///                 },
    ///             },
    ///             VideoDescriptions = new[]
    ///             {
    ///                 new Aws.MediaLive.Inputs.ChannelEncoderSettingsVideoDescriptionArgs
    ///                 {
    ///                     Name = "example-video",
    ///                 },
    ///             },
    ///             OutputGroups = new[]
    ///             {
    ///                 new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupArgs
    ///                 {
    ///                     OutputGroupSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArgs
    ///                     {
    ///                         ArchiveGroupSettings = new[]
    ///                         {
    ///                             new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArgs
    ///                             {
    ///                                 Destination = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestinationArgs
    ///                                 {
    ///                                     DestinationRefId = "destination",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     Outputs = new[]
    ///                     {
    ///                         new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputArgs
    ///                         {
    ///                             OutputName = "example-name",
    ///                             VideoDescriptionName = "example-video",
    ///                             AudioDescriptionNames = new[]
    ///                             {
    ///                                 "audio-selector",
    ///                             },
    ///                             OutputSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs
    ///                             {
    ///                                 ArchiveOutputSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsArgs
    ///                                 {
    ///                                     NameModifier = "_1",
    ///                                     Extension = "m2ts",
    ///                                     ContainerSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsArgs
    ///                                     {
    ///                                         M2tsSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsArgs
    ///                                         {
    ///                                             AudioBufferModel = "ATSC",
    ///                                             BufferModel = "MULTIPLEX",
    ///                                             RateMode = "CBR",
    ///                                         },
    ///                                     },
    ///                                 },
    ///                             },
    ///                         },
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
    /// Using `pulumi import`, import MediaLive Channel using the `channel_id`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:medialive/channel:Channel example 1234567
    /// ```
    /// </summary>
    [AwsResourceType("aws:medialive/channel:Channel")]
    public partial class Channel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the Channel.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specification of CDI inputs for this channel. See CDI Input Specification for more details.
        /// </summary>
        [Output("cdiInputSpecification")]
        public Output<Outputs.ChannelCdiInputSpecification?> CdiInputSpecification { get; private set; } = null!;

        /// <summary>
        /// Concise argument description.
        /// </summary>
        [Output("channelClass")]
        public Output<string> ChannelClass { get; private set; } = null!;

        /// <summary>
        /// ID of the channel in MediaPackage that is the destination for this output group.
        /// </summary>
        [Output("channelId")]
        public Output<string> ChannelId { get; private set; } = null!;

        /// <summary>
        /// Destinations for channel. See Destinations for more details.
        /// </summary>
        [Output("destinations")]
        public Output<ImmutableArray<Outputs.ChannelDestination>> Destinations { get; private set; } = null!;

        /// <summary>
        /// Encoder settings. See Encoder Settings for more details.
        /// </summary>
        [Output("encoderSettings")]
        public Output<Outputs.ChannelEncoderSettings> EncoderSettings { get; private set; } = null!;

        /// <summary>
        /// Input attachments for the channel. See Input Attachments for more details.
        /// </summary>
        [Output("inputAttachments")]
        public Output<ImmutableArray<Outputs.ChannelInputAttachment>> InputAttachments { get; private set; } = null!;

        /// <summary>
        /// Specification of network and file inputs for the channel.
        /// </summary>
        [Output("inputSpecification")]
        public Output<Outputs.ChannelInputSpecification> InputSpecification { get; private set; } = null!;

        /// <summary>
        /// The log level to write to Cloudwatch logs.
        /// </summary>
        [Output("logLevel")]
        public Output<string> LogLevel { get; private set; } = null!;

        /// <summary>
        /// Maintenance settings for this channel. See Maintenance for more details.
        /// </summary>
        [Output("maintenance")]
        public Output<Outputs.ChannelMaintenance> Maintenance { get; private set; } = null!;

        /// <summary>
        /// Name of the Channel.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Concise argument description.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Whether to start/stop channel. Default: `false`
        /// </summary>
        [Output("startChannel")]
        public Output<bool?> StartChannel { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Settings for the VPC outputs.
        /// </summary>
        [Output("vpc")]
        public Output<Outputs.ChannelVpc?> Vpc { get; private set; } = null!;


        /// <summary>
        /// Create a Channel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Channel(string name, ChannelArgs args, CustomResourceOptions? options = null)
            : base("aws:medialive/channel:Channel", name, args ?? new ChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Channel(string name, Input<string> id, ChannelState? state = null, CustomResourceOptions? options = null)
            : base("aws:medialive/channel:Channel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Channel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Channel Get(string name, Input<string> id, ChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new Channel(name, id, state, options);
        }
    }

    public sealed class ChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specification of CDI inputs for this channel. See CDI Input Specification for more details.
        /// </summary>
        [Input("cdiInputSpecification")]
        public Input<Inputs.ChannelCdiInputSpecificationArgs>? CdiInputSpecification { get; set; }

        /// <summary>
        /// Concise argument description.
        /// </summary>
        [Input("channelClass", required: true)]
        public Input<string> ChannelClass { get; set; } = null!;

        [Input("destinations", required: true)]
        private InputList<Inputs.ChannelDestinationArgs>? _destinations;

        /// <summary>
        /// Destinations for channel. See Destinations for more details.
        /// </summary>
        public InputList<Inputs.ChannelDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.ChannelDestinationArgs>());
            set => _destinations = value;
        }

        /// <summary>
        /// Encoder settings. See Encoder Settings for more details.
        /// </summary>
        [Input("encoderSettings", required: true)]
        public Input<Inputs.ChannelEncoderSettingsArgs> EncoderSettings { get; set; } = null!;

        [Input("inputAttachments", required: true)]
        private InputList<Inputs.ChannelInputAttachmentArgs>? _inputAttachments;

        /// <summary>
        /// Input attachments for the channel. See Input Attachments for more details.
        /// </summary>
        public InputList<Inputs.ChannelInputAttachmentArgs> InputAttachments
        {
            get => _inputAttachments ?? (_inputAttachments = new InputList<Inputs.ChannelInputAttachmentArgs>());
            set => _inputAttachments = value;
        }

        /// <summary>
        /// Specification of network and file inputs for the channel.
        /// </summary>
        [Input("inputSpecification", required: true)]
        public Input<Inputs.ChannelInputSpecificationArgs> InputSpecification { get; set; } = null!;

        /// <summary>
        /// The log level to write to Cloudwatch logs.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// Maintenance settings for this channel. See Maintenance for more details.
        /// </summary>
        [Input("maintenance")]
        public Input<Inputs.ChannelMaintenanceArgs>? Maintenance { get; set; }

        /// <summary>
        /// Name of the Channel.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Concise argument description.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Whether to start/stop channel. Default: `false`
        /// </summary>
        [Input("startChannel")]
        public Input<bool>? StartChannel { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Settings for the VPC outputs.
        /// </summary>
        [Input("vpc")]
        public Input<Inputs.ChannelVpcArgs>? Vpc { get; set; }

        public ChannelArgs()
        {
        }
        public static new ChannelArgs Empty => new ChannelArgs();
    }

    public sealed class ChannelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the Channel.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specification of CDI inputs for this channel. See CDI Input Specification for more details.
        /// </summary>
        [Input("cdiInputSpecification")]
        public Input<Inputs.ChannelCdiInputSpecificationGetArgs>? CdiInputSpecification { get; set; }

        /// <summary>
        /// Concise argument description.
        /// </summary>
        [Input("channelClass")]
        public Input<string>? ChannelClass { get; set; }

        /// <summary>
        /// ID of the channel in MediaPackage that is the destination for this output group.
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        [Input("destinations")]
        private InputList<Inputs.ChannelDestinationGetArgs>? _destinations;

        /// <summary>
        /// Destinations for channel. See Destinations for more details.
        /// </summary>
        public InputList<Inputs.ChannelDestinationGetArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.ChannelDestinationGetArgs>());
            set => _destinations = value;
        }

        /// <summary>
        /// Encoder settings. See Encoder Settings for more details.
        /// </summary>
        [Input("encoderSettings")]
        public Input<Inputs.ChannelEncoderSettingsGetArgs>? EncoderSettings { get; set; }

        [Input("inputAttachments")]
        private InputList<Inputs.ChannelInputAttachmentGetArgs>? _inputAttachments;

        /// <summary>
        /// Input attachments for the channel. See Input Attachments for more details.
        /// </summary>
        public InputList<Inputs.ChannelInputAttachmentGetArgs> InputAttachments
        {
            get => _inputAttachments ?? (_inputAttachments = new InputList<Inputs.ChannelInputAttachmentGetArgs>());
            set => _inputAttachments = value;
        }

        /// <summary>
        /// Specification of network and file inputs for the channel.
        /// </summary>
        [Input("inputSpecification")]
        public Input<Inputs.ChannelInputSpecificationGetArgs>? InputSpecification { get; set; }

        /// <summary>
        /// The log level to write to Cloudwatch logs.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// Maintenance settings for this channel. See Maintenance for more details.
        /// </summary>
        [Input("maintenance")]
        public Input<Inputs.ChannelMaintenanceGetArgs>? Maintenance { get; set; }

        /// <summary>
        /// Name of the Channel.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Concise argument description.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Whether to start/stop channel. Default: `false`
        /// </summary>
        [Input("startChannel")]
        public Input<bool>? StartChannel { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// Settings for the VPC outputs.
        /// </summary>
        [Input("vpc")]
        public Input<Inputs.ChannelVpcGetArgs>? Vpc { get; set; }

        public ChannelState()
        {
        }
        public static new ChannelState Empty => new ChannelState();
    }
}
