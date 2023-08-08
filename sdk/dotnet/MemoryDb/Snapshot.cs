// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MemoryDb
{
    /// <summary>
    /// Provides a MemoryDB Snapshot.
    /// 
    /// More information about snapshot and restore can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/snapshots.html).
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
    ///     var example = new Aws.MemoryDb.Snapshot("example", new()
    ///     {
    ///         ClusterName = aws_memorydb_cluster.Example.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_memorydb_snapshot.example
    /// 
    ///  id = "my-snapshot" } Using `pulumi import`, import a snapshot using the `name`. For exampleconsole % pulumi import aws_memorydb_snapshot.example my-snapshot
    /// </summary>
    [AwsResourceType("aws:memorydb/snapshot:Snapshot")]
    public partial class Snapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the snapshot.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The configuration of the cluster from which the snapshot was taken.
        /// </summary>
        [Output("clusterConfigurations")]
        public Output<ImmutableArray<Outputs.SnapshotClusterConfiguration>> ClusterConfigurations { get; private set; } = null!;

        /// <summary>
        /// Name of the MemoryDB cluster to take a snapshot of.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// ARN of the KMS key used to encrypt the snapshot at rest.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// Name of the snapshot. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the snapshot is from an automatic backup (`automated`) or was created manually (`manual`).
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

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
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("aws:memorydb/snapshot:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
            : base("aws:memorydb/snapshot:Snapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, state, options);
        }
    }

    public sealed class SnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the MemoryDB cluster to take a snapshot of.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// ARN of the KMS key used to encrypt the snapshot at rest.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Name of the snapshot. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

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

        public SnapshotArgs()
        {
        }
        public static new SnapshotArgs Empty => new SnapshotArgs();
    }

    public sealed class SnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the snapshot.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("clusterConfigurations")]
        private InputList<Inputs.SnapshotClusterConfigurationGetArgs>? _clusterConfigurations;

        /// <summary>
        /// The configuration of the cluster from which the snapshot was taken.
        /// </summary>
        public InputList<Inputs.SnapshotClusterConfigurationGetArgs> ClusterConfigurations
        {
            get => _clusterConfigurations ?? (_clusterConfigurations = new InputList<Inputs.SnapshotClusterConfigurationGetArgs>());
            set => _clusterConfigurations = value;
        }

        /// <summary>
        /// Name of the MemoryDB cluster to take a snapshot of.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// ARN of the KMS key used to encrypt the snapshot at rest.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Name of the snapshot. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Indicates whether the snapshot is from an automatic backup (`automated`) or was created manually (`manual`).
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

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
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public SnapshotState()
        {
        }
        public static new SnapshotState Empty => new SnapshotState();
    }
}
