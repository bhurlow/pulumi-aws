// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ebs
{
    /// <summary>
    /// Manages a single EBS volume.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Ebs.Volume("example", new()
    ///     {
    ///         AvailabilityZone = "us-west-2a",
    ///         Size = 40,
    ///         Tags = 
    ///         {
    ///             { "Name", "HelloWorld" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; **NOTE:** At least one of `size` or `snapshot_id` is required when specifying an EBS volume
    /// 
    /// ## Import
    /// 
    /// EBS Volumes can be imported using the `id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ebs/volume:Volume id vol-049df61146c4d7901
    /// ```
    /// </summary>
    [AwsResourceType("aws:ebs/volume:Volume")]
    public partial class Volume : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The volume ARN (e.g., arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The AZ where the EBS volume will exist.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Output("encrypted")]
        public Output<bool> Encrypted { get; private set; } = null!;

        /// <summary>
        /// If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
        /// </summary>
        [Output("finalSnapshot")]
        public Output<bool?> FinalSnapshot { get; private set; } = null!;

        /// <summary>
        /// The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
        /// </summary>
        [Output("iops")]
        public Output<int> Iops { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
        /// </summary>
        [Output("multiAttachEnabled")]
        public Output<bool?> MultiAttachEnabled { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Output("outpostArn")]
        public Output<string?> OutpostArn { get; private set; } = null!;

        /// <summary>
        /// The size of the drive in GiBs.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// A snapshot to base the EBS volume off of.
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;

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
        /// The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
        /// </summary>
        [Output("throughput")]
        public Output<int> Throughput { get; private set; } = null!;

        /// <summary>
        /// The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Volume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Volume(string name, VolumeArgs args, CustomResourceOptions? options = null)
            : base("aws:ebs/volume:Volume", name, args ?? new VolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Volume(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
            : base("aws:ebs/volume:Volume", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Volume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Volume Get(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new Volume(name, id, state, options);
        }
    }

    public sealed class VolumeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AZ where the EBS volume will exist.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
        /// </summary>
        [Input("finalSnapshot")]
        public Input<bool>? FinalSnapshot { get; set; }

        /// <summary>
        /// The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
        /// </summary>
        [Input("multiAttachEnabled")]
        public Input<bool>? MultiAttachEnabled { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The size of the drive in GiBs.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// A snapshot to base the EBS volume off of.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

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
        /// The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
        /// </summary>
        [Input("throughput")]
        public Input<int>? Throughput { get; set; }

        /// <summary>
        /// The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VolumeArgs()
        {
        }
        public static new VolumeArgs Empty => new VolumeArgs();
    }

    public sealed class VolumeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The volume ARN (e.g., arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The AZ where the EBS volume will exist.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
        /// </summary>
        [Input("finalSnapshot")]
        public Input<bool>? FinalSnapshot { get; set; }

        /// <summary>
        /// The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
        /// </summary>
        [Input("multiAttachEnabled")]
        public Input<bool>? MultiAttachEnabled { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The size of the drive in GiBs.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// A snapshot to base the EBS volume off of.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

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

        /// <summary>
        /// The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
        /// </summary>
        [Input("throughput")]
        public Input<int>? Throughput { get; set; }

        /// <summary>
        /// The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VolumeState()
        {
        }
        public static new VolumeState Empty => new VolumeState();
    }
}
