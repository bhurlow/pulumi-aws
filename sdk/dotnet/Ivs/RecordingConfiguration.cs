// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ivs
{
    /// <summary>
    /// Resource for managing an AWS IVS (Interactive Video) Recording Configuration.
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
    ///     var example = new Aws.Ivs.RecordingConfiguration("example", new()
    ///     {
    ///         DestinationConfiguration = new Aws.Ivs.Inputs.RecordingConfigurationDestinationConfigurationArgs
    ///         {
    ///             S3 = new Aws.Ivs.Inputs.RecordingConfigurationDestinationConfigurationS3Args
    ///             {
    ///                 BucketName = "ivs-stream-archive",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_ivs_recording_configuration.example
    /// 
    ///  id = "arn:aws:ivs:us-west-2:326937407773:recording-configuration/KAk1sHBl2L47" } Using `pulumi import`, import IVS (Interactive Video) Recording Configuration using the ARN. For exampleconsole % pulumi import aws_ivs_recording_configuration.example arn:aws:ivs:us-west-2:326937407773:recording-configuration/KAk1sHBl2L47
    /// </summary>
    [AwsResourceType("aws:ivs/recordingConfiguration:RecordingConfiguration")]
    public partial class RecordingConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the Recording Configuration.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Object containing destination configuration for where recorded video will be stored.
        /// </summary>
        [Output("destinationConfiguration")]
        public Output<Outputs.RecordingConfigurationDestinationConfiguration> DestinationConfiguration { get; private set; } = null!;

        /// <summary>
        /// Recording Configuration name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
        /// </summary>
        [Output("recordingReconnectWindowSeconds")]
        public Output<int> RecordingReconnectWindowSeconds { get; private set; } = null!;

        /// <summary>
        /// The current state of the Recording Configuration.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
        /// </summary>
        [Output("thumbnailConfiguration")]
        public Output<Outputs.RecordingConfigurationThumbnailConfiguration> ThumbnailConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a RecordingConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecordingConfiguration(string name, RecordingConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:ivs/recordingConfiguration:RecordingConfiguration", name, args ?? new RecordingConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecordingConfiguration(string name, Input<string> id, RecordingConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ivs/recordingConfiguration:RecordingConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RecordingConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecordingConfiguration Get(string name, Input<string> id, RecordingConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new RecordingConfiguration(name, id, state, options);
        }
    }

    public sealed class RecordingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Object containing destination configuration for where recorded video will be stored.
        /// </summary>
        [Input("destinationConfiguration", required: true)]
        public Input<Inputs.RecordingConfigurationDestinationConfigurationArgs> DestinationConfiguration { get; set; } = null!;

        /// <summary>
        /// Recording Configuration name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
        /// </summary>
        [Input("recordingReconnectWindowSeconds")]
        public Input<int>? RecordingReconnectWindowSeconds { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
        /// </summary>
        [Input("thumbnailConfiguration")]
        public Input<Inputs.RecordingConfigurationThumbnailConfigurationArgs>? ThumbnailConfiguration { get; set; }

        public RecordingConfigurationArgs()
        {
        }
        public static new RecordingConfigurationArgs Empty => new RecordingConfigurationArgs();
    }

    public sealed class RecordingConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the Recording Configuration.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Object containing destination configuration for where recorded video will be stored.
        /// </summary>
        [Input("destinationConfiguration")]
        public Input<Inputs.RecordingConfigurationDestinationConfigurationGetArgs>? DestinationConfiguration { get; set; }

        /// <summary>
        /// Recording Configuration name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
        /// </summary>
        [Input("recordingReconnectWindowSeconds")]
        public Input<int>? RecordingReconnectWindowSeconds { get; set; }

        /// <summary>
        /// The current state of the Recording Configuration.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
        /// </summary>
        [Input("thumbnailConfiguration")]
        public Input<Inputs.RecordingConfigurationThumbnailConfigurationGetArgs>? ThumbnailConfiguration { get; set; }

        public RecordingConfigurationState()
        {
        }
        public static new RecordingConfigurationState Empty => new RecordingConfigurationState();
    }
}
