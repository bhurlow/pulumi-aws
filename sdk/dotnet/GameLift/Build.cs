// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift
{
    /// <summary>
    /// Provides an GameLift Build resource.
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
    ///     var test = new Aws.GameLift.Build("test", new()
    ///     {
    ///         OperatingSystem = "WINDOWS_2012",
    ///         StorageLocation = new Aws.GameLift.Inputs.BuildStorageLocationArgs
    ///         {
    ///             Bucket = aws_s3_bucket.Test.Id,
    ///             Key = aws_s3_object.Test.Key,
    ///             RoleArn = aws_iam_role.Test.Arn,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GameLift Builds can be imported using the ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:gamelift/build:Build example &lt;build-id&gt;
    /// ```
    /// </summary>
    [AwsResourceType("aws:gamelift/build:Build")]
    public partial class Build : global::Pulumi.CustomResource
    {
        /// <summary>
        /// GameLift Build ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Name of the build
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Operating system that the game server binaries are built to run onE.g., `WINDOWS_2012`, `AMAZON_LINUX` or `AMAZON_LINUX_2`.
        /// </summary>
        [Output("operatingSystem")]
        public Output<string> OperatingSystem { get; private set; } = null!;

        /// <summary>
        /// Information indicating where your game build files are stored. See below.
        /// </summary>
        [Output("storageLocation")]
        public Output<Outputs.BuildStorageLocation> StorageLocation { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Version that is associated with this build.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Build resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Build(string name, BuildArgs args, CustomResourceOptions? options = null)
            : base("aws:gamelift/build:Build", name, args ?? new BuildArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Build(string name, Input<string> id, BuildState? state = null, CustomResourceOptions? options = null)
            : base("aws:gamelift/build:Build", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Build resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Build Get(string name, Input<string> id, BuildState? state = null, CustomResourceOptions? options = null)
        {
            return new Build(name, id, state, options);
        }
    }

    public sealed class BuildArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the build
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Operating system that the game server binaries are built to run onE.g., `WINDOWS_2012`, `AMAZON_LINUX` or `AMAZON_LINUX_2`.
        /// </summary>
        [Input("operatingSystem", required: true)]
        public Input<string> OperatingSystem { get; set; } = null!;

        /// <summary>
        /// Information indicating where your game build files are stored. See below.
        /// </summary>
        [Input("storageLocation", required: true)]
        public Input<Inputs.BuildStorageLocationArgs> StorageLocation { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Version that is associated with this build.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public BuildArgs()
        {
        }
        public static new BuildArgs Empty => new BuildArgs();
    }

    public sealed class BuildState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// GameLift Build ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Name of the build
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Operating system that the game server binaries are built to run onE.g., `WINDOWS_2012`, `AMAZON_LINUX` or `AMAZON_LINUX_2`.
        /// </summary>
        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        /// <summary>
        /// Information indicating where your game build files are stored. See below.
        /// </summary>
        [Input("storageLocation")]
        public Input<Inputs.BuildStorageLocationGetArgs>? StorageLocation { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// Version that is associated with this build.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public BuildState()
        {
        }
        public static new BuildState Empty => new BuildState();
    }
}
