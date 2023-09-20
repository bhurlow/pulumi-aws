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
    /// Resource for managing an AWS Location Geofence Collection.
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
    ///     var example = new Aws.Location.GeofenceCollection("example", new()
    ///     {
    ///         CollectionName = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Location Geofence Collection using the `collection_name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:location/geofenceCollection:GeofenceCollection example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:location/geofenceCollection:GeofenceCollection")]
    public partial class GeofenceCollection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS.
        /// </summary>
        [Output("collectionArn")]
        public Output<string> CollectionArn { get; private set; } = null!;

        /// <summary>
        /// The name of the geofence collection.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("collectionName")]
        public Output<string> CollectionName { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the geofence collection resource was created in ISO 8601 format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The optional description for the geofence collection.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the geofence collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the geofence collection resource was last updated in ISO 8601 format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GeofenceCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeofenceCollection(string name, GeofenceCollectionArgs args, CustomResourceOptions? options = null)
            : base("aws:location/geofenceCollection:GeofenceCollection", name, args ?? new GeofenceCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeofenceCollection(string name, Input<string> id, GeofenceCollectionState? state = null, CustomResourceOptions? options = null)
            : base("aws:location/geofenceCollection:GeofenceCollection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GeofenceCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeofenceCollection Get(string name, Input<string> id, GeofenceCollectionState? state = null, CustomResourceOptions? options = null)
        {
            return new GeofenceCollection(name, id, state, options);
        }
    }

    public sealed class GeofenceCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the geofence collection.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("collectionName", required: true)]
        public Input<string> CollectionName { get; set; } = null!;

        /// <summary>
        /// The optional description for the geofence collection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the geofence collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GeofenceCollectionArgs()
        {
        }
        public static new GeofenceCollectionArgs Empty => new GeofenceCollectionArgs();
    }

    public sealed class GeofenceCollectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS.
        /// </summary>
        [Input("collectionArn")]
        public Input<string>? CollectionArn { get; set; }

        /// <summary>
        /// The name of the geofence collection.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("collectionName")]
        public Input<string>? CollectionName { get; set; }

        /// <summary>
        /// The timestamp for when the geofence collection resource was created in ISO 8601 format.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The optional description for the geofence collection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the geofence collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// The timestamp for when the geofence collection resource was last updated in ISO 8601 format.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public GeofenceCollectionState()
        {
        }
        public static new GeofenceCollectionState Empty => new GeofenceCollectionState();
    }
}
