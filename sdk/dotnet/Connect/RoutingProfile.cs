// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Connect
{
    /// <summary>
    /// Provides an Amazon Connect Routing Profile resource. For more information see
    /// [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)
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
    ///     var example = new Aws.Connect.RoutingProfile("example", new()
    ///     {
    ///         DefaultOutboundQueueId = "12345678-1234-1234-1234-123456789012",
    ///         Description = "example description",
    ///         InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    ///         MediaConcurrencies = new[]
    ///         {
    ///             new Aws.Connect.Inputs.RoutingProfileMediaConcurrencyArgs
    ///             {
    ///                 Channel = "VOICE",
    ///                 Concurrency = 1,
    ///             },
    ///         },
    ///         QueueConfigs = new[]
    ///         {
    ///             new Aws.Connect.Inputs.RoutingProfileQueueConfigArgs
    ///             {
    ///                 Channel = "VOICE",
    ///                 Delay = 2,
    ///                 Priority = 1,
    ///                 QueueId = "12345678-1234-1234-1234-123456789012",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Name", "Example Routing Profile" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Amazon Connect Routing Profiles using the `instance_id` and `routing_profile_id` separated by a colon (`:`). For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:connect/routingProfile:RoutingProfile example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
    /// ```
    /// </summary>
    [AwsResourceType("aws:connect/routingProfile:RoutingProfile")]
    public partial class RoutingProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Routing Profile.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies the default outbound queue for the Routing Profile.
        /// </summary>
        [Output("defaultOutboundQueueId")]
        public Output<string> DefaultOutboundQueueId { get; private set; } = null!;

        /// <summary>
        /// Specifies the description of the Routing Profile.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the identifier of the hosting Amazon Connect Instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// One or more `media_concurrencies` blocks that specify the channels that agents can handle in the Contact Control Panel (CCP) for this Routing Profile. The `media_concurrencies` block is documented below.
        /// </summary>
        [Output("mediaConcurrencies")]
        public Output<ImmutableArray<Outputs.RoutingProfileMediaConcurrency>> MediaConcurrencies { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Routing Profile.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more `queue_configs` blocks that specify the inbound queues associated with the routing profile. If no queue is added, the agent only can make outbound calls. The `queue_configs` block is documented below.
        /// </summary>
        [Output("queueConfigs")]
        public Output<ImmutableArray<Outputs.RoutingProfileQueueConfig>> QueueConfigs { get; private set; } = null!;

        /// <summary>
        /// The identifier for the Routing Profile.
        /// </summary>
        [Output("routingProfileId")]
        public Output<string> RoutingProfileId { get; private set; } = null!;

        /// <summary>
        /// Tags to apply to the Routing Profile. If configured with a provider
        /// `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a RoutingProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoutingProfile(string name, RoutingProfileArgs args, CustomResourceOptions? options = null)
            : base("aws:connect/routingProfile:RoutingProfile", name, args ?? new RoutingProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoutingProfile(string name, Input<string> id, RoutingProfileState? state = null, CustomResourceOptions? options = null)
            : base("aws:connect/routingProfile:RoutingProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RoutingProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoutingProfile Get(string name, Input<string> id, RoutingProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new RoutingProfile(name, id, state, options);
        }
    }

    public sealed class RoutingProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the default outbound queue for the Routing Profile.
        /// </summary>
        [Input("defaultOutboundQueueId", required: true)]
        public Input<string> DefaultOutboundQueueId { get; set; } = null!;

        /// <summary>
        /// Specifies the description of the Routing Profile.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Specifies the identifier of the hosting Amazon Connect Instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("mediaConcurrencies", required: true)]
        private InputList<Inputs.RoutingProfileMediaConcurrencyArgs>? _mediaConcurrencies;

        /// <summary>
        /// One or more `media_concurrencies` blocks that specify the channels that agents can handle in the Contact Control Panel (CCP) for this Routing Profile. The `media_concurrencies` block is documented below.
        /// </summary>
        public InputList<Inputs.RoutingProfileMediaConcurrencyArgs> MediaConcurrencies
        {
            get => _mediaConcurrencies ?? (_mediaConcurrencies = new InputList<Inputs.RoutingProfileMediaConcurrencyArgs>());
            set => _mediaConcurrencies = value;
        }

        /// <summary>
        /// Specifies the name of the Routing Profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("queueConfigs")]
        private InputList<Inputs.RoutingProfileQueueConfigArgs>? _queueConfigs;

        /// <summary>
        /// One or more `queue_configs` blocks that specify the inbound queues associated with the routing profile. If no queue is added, the agent only can make outbound calls. The `queue_configs` block is documented below.
        /// </summary>
        public InputList<Inputs.RoutingProfileQueueConfigArgs> QueueConfigs
        {
            get => _queueConfigs ?? (_queueConfigs = new InputList<Inputs.RoutingProfileQueueConfigArgs>());
            set => _queueConfigs = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags to apply to the Routing Profile. If configured with a provider
        /// `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RoutingProfileArgs()
        {
        }
        public static new RoutingProfileArgs Empty => new RoutingProfileArgs();
    }

    public sealed class RoutingProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Routing Profile.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies the default outbound queue for the Routing Profile.
        /// </summary>
        [Input("defaultOutboundQueueId")]
        public Input<string>? DefaultOutboundQueueId { get; set; }

        /// <summary>
        /// Specifies the description of the Routing Profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the identifier of the hosting Amazon Connect Instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("mediaConcurrencies")]
        private InputList<Inputs.RoutingProfileMediaConcurrencyGetArgs>? _mediaConcurrencies;

        /// <summary>
        /// One or more `media_concurrencies` blocks that specify the channels that agents can handle in the Contact Control Panel (CCP) for this Routing Profile. The `media_concurrencies` block is documented below.
        /// </summary>
        public InputList<Inputs.RoutingProfileMediaConcurrencyGetArgs> MediaConcurrencies
        {
            get => _mediaConcurrencies ?? (_mediaConcurrencies = new InputList<Inputs.RoutingProfileMediaConcurrencyGetArgs>());
            set => _mediaConcurrencies = value;
        }

        /// <summary>
        /// Specifies the name of the Routing Profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("queueConfigs")]
        private InputList<Inputs.RoutingProfileQueueConfigGetArgs>? _queueConfigs;

        /// <summary>
        /// One or more `queue_configs` blocks that specify the inbound queues associated with the routing profile. If no queue is added, the agent only can make outbound calls. The `queue_configs` block is documented below.
        /// </summary>
        public InputList<Inputs.RoutingProfileQueueConfigGetArgs> QueueConfigs
        {
            get => _queueConfigs ?? (_queueConfigs = new InputList<Inputs.RoutingProfileQueueConfigGetArgs>());
            set => _queueConfigs = value;
        }

        /// <summary>
        /// The identifier for the Routing Profile.
        /// </summary>
        [Input("routingProfileId")]
        public Input<string>? RoutingProfileId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags to apply to the Routing Profile. If configured with a provider
        /// `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
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

        public RoutingProfileState()
        {
        }
        public static new RoutingProfileState Empty => new RoutingProfileState();
    }
}
