// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaPackage
{
    /// <summary>
    /// Provides an AWS Elemental MediaPackage Channel.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/media_package_channel.html.markdown.
    /// </summary>
    public partial class Channel : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the channel
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A unique identifier describing the channel
        /// </summary>
        [Output("channelId")]
        public Output<string> ChannelId { get; private set; } = null!;

        /// <summary>
        /// A description of the channel
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A single item list of HLS ingest information
        /// </summary>
        [Output("hlsIngests")]
        public Output<ImmutableArray<Outputs.ChannelHlsIngests>> HlsIngests { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Channel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Channel(string name, ChannelArgs args, CustomResourceOptions? options = null)
            : base("aws:mediapackage/channel:Channel", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Channel(string name, Input<string> id, ChannelState? state = null, CustomResourceOptions? options = null)
            : base("aws:mediapackage/channel:Channel", name, state, MakeResourceOptions(options, id))
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

    public sealed class ChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier describing the channel
        /// </summary>
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        /// <summary>
        /// A description of the channel
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ChannelArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class ChannelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the channel
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A unique identifier describing the channel
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        /// <summary>
        /// A description of the channel
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hlsIngests")]
        private InputList<Inputs.ChannelHlsIngestsGetArgs>? _hlsIngests;

        /// <summary>
        /// A single item list of HLS ingest information
        /// </summary>
        public InputList<Inputs.ChannelHlsIngestsGetArgs> HlsIngests
        {
            get => _hlsIngests ?? (_hlsIngests = new InputList<Inputs.ChannelHlsIngestsGetArgs>());
            set => _hlsIngests = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ChannelState()
        {
            Description = "Managed by Pulumi";
        }
    }

    namespace Inputs
    {

    public sealed class ChannelHlsIngestsGetArgs : Pulumi.ResourceArgs
    {
        [Input("ingestEndpoints")]
        private InputList<ChannelHlsIngestsIngestEndpointsGetArgs>? _ingestEndpoints;

        /// <summary>
        /// A list of the ingest endpoints
        /// </summary>
        public InputList<ChannelHlsIngestsIngestEndpointsGetArgs> IngestEndpoints
        {
            get => _ingestEndpoints ?? (_ingestEndpoints = new InputList<ChannelHlsIngestsIngestEndpointsGetArgs>());
            set => _ingestEndpoints = value;
        }

        public ChannelHlsIngestsGetArgs()
        {
        }
    }

    public sealed class ChannelHlsIngestsIngestEndpointsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The URL
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The username
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ChannelHlsIngestsIngestEndpointsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ChannelHlsIngests
    {
        /// <summary>
        /// A list of the ingest endpoints
        /// </summary>
        public readonly ImmutableArray<ChannelHlsIngestsIngestEndpoints> IngestEndpoints;

        [OutputConstructor]
        private ChannelHlsIngests(ImmutableArray<ChannelHlsIngestsIngestEndpoints> ingestEndpoints)
        {
            IngestEndpoints = ingestEndpoints;
        }
    }

    [OutputType]
    public sealed class ChannelHlsIngestsIngestEndpoints
    {
        /// <summary>
        /// The password
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The URL
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// The username
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private ChannelHlsIngestsIngestEndpoints(
            string password,
            string url,
            string username)
        {
            Password = password;
            Url = url;
            Username = username;
        }
    }
    }
}
