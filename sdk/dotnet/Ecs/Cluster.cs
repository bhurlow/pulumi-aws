// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs
{
    /// <summary>
    /// Provides an ECS cluster.
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
    ///     var foo = new Aws.Ecs.Cluster("foo", new()
    ///     {
    ///         Settings = new[]
    ///         {
    ///             new Aws.Ecs.Inputs.ClusterSettingArgs
    ///             {
    ///                 Name = "containerInsights",
    ///                 Value = "enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Example with Log Configuration
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleKey = new Aws.Kms.Key("exampleKey", new()
    ///     {
    ///         Description = "example",
    ///         DeletionWindowInDays = 7,
    ///     });
    /// 
    ///     var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");
    /// 
    ///     var test = new Aws.Ecs.Cluster("test", new()
    ///     {
    ///         Configuration = new Aws.Ecs.Inputs.ClusterConfigurationArgs
    ///         {
    ///             ExecuteCommandConfiguration = new Aws.Ecs.Inputs.ClusterConfigurationExecuteCommandConfigurationArgs
    ///             {
    ///                 KmsKeyId = exampleKey.Arn,
    ///                 Logging = "OVERRIDE",
    ///                 LogConfiguration = new Aws.Ecs.Inputs.ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs
    ///                 {
    ///                     CloudWatchEncryptionEnabled = true,
    ///                     CloudWatchLogGroupName = exampleLogGroup.Name,
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
    /// Using `pulumi import`, import ECS clusters using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecs/cluster:Cluster stateless stateless-app
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecs/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN that identifies the cluster.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The execute command configuration for the cluster. Detailed below.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.ClusterConfiguration?> Configuration { get; private set; } = null!;

        /// <summary>
        /// Name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configures a default Service Connect namespace. Detailed below.
        /// </summary>
        [Output("serviceConnectDefaults")]
        public Output<Outputs.ClusterServiceConnectDefaults?> ServiceConnectDefaults { get; private set; } = null!;

        /// <summary>
        /// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        /// </summary>
        [Output("settings")]
        public Output<ImmutableArray<Outputs.ClusterSetting>> Settings { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ecs/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecs/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The execute command configuration for the cluster. Detailed below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.ClusterConfigurationArgs>? Configuration { get; set; }

        /// <summary>
        /// Name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configures a default Service Connect namespace. Detailed below.
        /// </summary>
        [Input("serviceConnectDefaults")]
        public Input<Inputs.ClusterServiceConnectDefaultsArgs>? ServiceConnectDefaults { get; set; }

        [Input("settings")]
        private InputList<Inputs.ClusterSettingArgs>? _settings;

        /// <summary>
        /// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        /// </summary>
        public InputList<Inputs.ClusterSettingArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.ClusterSettingArgs>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN that identifies the cluster.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The execute command configuration for the cluster. Detailed below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.ClusterConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// Name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configures a default Service Connect namespace. Detailed below.
        /// </summary>
        [Input("serviceConnectDefaults")]
        public Input<Inputs.ClusterServiceConnectDefaultsGetArgs>? ServiceConnectDefaults { get; set; }

        [Input("settings")]
        private InputList<Inputs.ClusterSettingGetArgs>? _settings;

        /// <summary>
        /// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        /// </summary>
        public InputList<Inputs.ClusterSettingGetArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.ClusterSettingGetArgs>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
