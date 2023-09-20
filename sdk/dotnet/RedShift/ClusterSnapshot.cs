// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift
{
    /// <summary>
    /// Creates a Redshift cluster snapshot
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Redshift Cluster Snapshots using `snapshot_identifier`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:redshift/clusterSnapshot:ClusterSnapshot test example
    /// ```
    /// </summary>
    [AwsResourceType("aws:redshift/clusterSnapshot:ClusterSnapshot")]
    public partial class ClusterSnapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the snapshot.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The cluster identifier for which you want a snapshot.
        /// </summary>
        [Output("clusterIdentifier")]
        public Output<string> ClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// The Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
        /// </summary>
        [Output("manualSnapshotRetentionPeriod")]
        public Output<int?> ManualSnapshotRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// For manual snapshots, the Amazon Web Services account used to create or copy the snapshot. For automatic snapshots, the owner of the cluster. The owner can perform all snapshot actions, such as sharing a manual snapshot.
        /// </summary>
        [Output("ownerAccount")]
        public Output<string> OwnerAccount { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
        /// </summary>
        [Output("snapshotIdentifier")]
        public Output<string> SnapshotIdentifier { get; private set; } = null!;

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
        /// Create a ClusterSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterSnapshot(string name, ClusterSnapshotArgs args, CustomResourceOptions? options = null)
            : base("aws:redshift/clusterSnapshot:ClusterSnapshot", name, args ?? new ClusterSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterSnapshot(string name, Input<string> id, ClusterSnapshotState? state = null, CustomResourceOptions? options = null)
            : base("aws:redshift/clusterSnapshot:ClusterSnapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterSnapshot Get(string name, Input<string> id, ClusterSnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterSnapshot(name, id, state, options);
        }
    }

    public sealed class ClusterSnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster identifier for which you want a snapshot.
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public Input<string> ClusterIdentifier { get; set; } = null!;

        /// <summary>
        /// The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
        /// </summary>
        [Input("manualSnapshotRetentionPeriod")]
        public Input<int>? ManualSnapshotRetentionPeriod { get; set; }

        /// <summary>
        /// A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
        /// </summary>
        [Input("snapshotIdentifier", required: true)]
        public Input<string> SnapshotIdentifier { get; set; } = null!;

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

        public ClusterSnapshotArgs()
        {
        }
        public static new ClusterSnapshotArgs Empty => new ClusterSnapshotArgs();
    }

    public sealed class ClusterSnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the snapshot.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The cluster identifier for which you want a snapshot.
        /// </summary>
        [Input("clusterIdentifier")]
        public Input<string>? ClusterIdentifier { get; set; }

        /// <summary>
        /// The Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
        /// </summary>
        [Input("manualSnapshotRetentionPeriod")]
        public Input<int>? ManualSnapshotRetentionPeriod { get; set; }

        /// <summary>
        /// For manual snapshots, the Amazon Web Services account used to create or copy the snapshot. For automatic snapshots, the owner of the cluster. The owner can perform all snapshot actions, such as sharing a manual snapshot.
        /// </summary>
        [Input("ownerAccount")]
        public Input<string>? OwnerAccount { get; set; }

        /// <summary>
        /// A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
        /// </summary>
        [Input("snapshotIdentifier")]
        public Input<string>? SnapshotIdentifier { get; set; }

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

        public ClusterSnapshotState()
        {
        }
        public static new ClusterSnapshotState Empty => new ClusterSnapshotState();
    }
}
