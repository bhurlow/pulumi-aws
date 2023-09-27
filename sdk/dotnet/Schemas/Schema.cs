// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Schemas
{
    /// <summary>
    /// Provides an EventBridge Schema resource.
    /// 
    /// &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testRegistry = new Aws.Schemas.Registry("testRegistry");
    /// 
    ///     var testSchema = new Aws.Schemas.Schema("testSchema", new()
    ///     {
    ///         RegistryName = testRegistry.Name,
    ///         Type = "OpenApi3",
    ///         Description = "The schema definition for my event",
    ///         Content = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["openapi"] = "3.0.0",
    ///             ["info"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["version"] = "1.0.0",
    ///                 ["title"] = "Event",
    ///             },
    ///             ["paths"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///             },
    ///             ["components"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["schemas"] = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Event"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["type"] = "object",
    ///                         ["properties"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["name"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["type"] = "string",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import EventBridge schema using the `name` and `registry_name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:schemas/schema:Schema test name/registry
    /// ```
    /// </summary>
    [AwsResourceType("aws:schemas/schema:Schema")]
    public partial class Schema : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the discoverer.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The schema specification. Must be a valid Open API 3.0 spec.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The description of the schema. Maximum of 256 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The last modified date of the schema.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the registry in which this schema belongs.
        /// </summary>
        [Output("registryName")]
        public Output<string> RegistryName { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The version of the schema.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The created date of the version of the schema.
        /// </summary>
        [Output("versionCreatedDate")]
        public Output<string> VersionCreatedDate { get; private set; } = null!;


        /// <summary>
        /// Create a Schema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schema(string name, SchemaArgs args, CustomResourceOptions? options = null)
            : base("aws:schemas/schema:Schema", name, args ?? new SchemaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Schema(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
            : base("aws:schemas/schema:Schema", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Schema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schema Get(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
        {
            return new Schema(name, id, state, options);
        }
    }

    public sealed class SchemaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The schema specification. Must be a valid Open API 3.0 spec.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The description of the schema. Maximum of 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the registry in which this schema belongs.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

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
        /// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SchemaArgs()
        {
        }
        public static new SchemaArgs Empty => new SchemaArgs();
    }

    public sealed class SchemaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the discoverer.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The schema specification. Must be a valid Open API 3.0 spec.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The description of the schema. Maximum of 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The last modified date of the schema.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        /// <summary>
        /// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the registry in which this schema belongs.
        /// </summary>
        [Input("registryName")]
        public Input<string>? RegistryName { get; set; }

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

        /// <summary>
        /// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The version of the schema.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The created date of the version of the schema.
        /// </summary>
        [Input("versionCreatedDate")]
        public Input<string>? VersionCreatedDate { get; set; }

        public SchemaState()
        {
        }
        public static new SchemaState Empty => new SchemaState();
    }
}
