// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Location
{
    /// <summary>
    /// Provides a Location Service Map.
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
    ///     var example = new Aws.Location.Map("example", new()
    ///     {
    ///         Configuration = new Aws.Location.Inputs.MapConfigurationArgs
    ///         {
    ///             Style = "VectorHereBerlin",
    ///         },
    ///         MapName = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_location_map` resources can be imported using the map name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:location/map:Map example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:location/map:Map")]
    public partial class Map : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration block with the map style selected from an available data provider. Detailed below.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.MapConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the map resource was created in ISO 8601 format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// An optional description for the map resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across all AWS.
        /// </summary>
        [Output("mapArn")]
        public Output<string> MapArn { get; private set; } = null!;

        /// <summary>
        /// The name for the map resource.
        /// </summary>
        [Output("mapName")]
        public Output<string> MapName { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the map. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the map resource was last updated in ISO 8601 format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Map resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Map(string name, MapArgs args, CustomResourceOptions? options = null)
            : base("aws:location/map:Map", name, args ?? new MapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Map(string name, Input<string> id, MapState? state = null, CustomResourceOptions? options = null)
            : base("aws:location/map:Map", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Map resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Map Get(string name, Input<string> id, MapState? state = null, CustomResourceOptions? options = null)
        {
            return new Map(name, id, state, options);
        }
    }

    public sealed class MapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block with the map style selected from an available data provider. Detailed below.
        /// </summary>
        [Input("configuration", required: true)]
        public Input<Inputs.MapConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// An optional description for the map resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name for the map resource.
        /// </summary>
        [Input("mapName", required: true)]
        public Input<string> MapName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the map. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public MapArgs()
        {
        }
        public static new MapArgs Empty => new MapArgs();
    }

    public sealed class MapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block with the map style selected from an available data provider. Detailed below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.MapConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// The timestamp for when the map resource was created in ISO 8601 format.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// An optional description for the map resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across all AWS.
        /// </summary>
        [Input("mapArn")]
        public Input<string>? MapArn { get; set; }

        /// <summary>
        /// The name for the map resource.
        /// </summary>
        [Input("mapName")]
        public Input<string>? MapName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the map. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The timestamp for when the map resource was last updated in ISO 8601 format.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public MapState()
        {
        }
        public static new MapState Empty => new MapState();
    }
}
