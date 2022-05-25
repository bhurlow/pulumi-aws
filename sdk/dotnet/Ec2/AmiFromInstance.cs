// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// The "AMI from instance" resource allows the creation of an Amazon Machine
    /// Image (AMI) modelled after an existing EBS-backed EC2 instance.
    /// 
    /// The created AMI will refer to implicitly-created snapshots of the instance's
    /// EBS volumes and mimick its assigned block device configuration at the time
    /// the resource is created.
    /// 
    /// This resource is best applied to an instance that is stopped when this instance
    /// is created, so that the contents of the created image are predictable. When
    /// applied to an instance that is running, *the instance will be stopped before taking
    /// the snapshots and then started back up again*, resulting in a period of
    /// downtime.
    /// 
    /// Note that the source instance is inspected only at the initial creation of this
    /// resource. Ongoing updates to the referenced instance will not be propagated into
    /// the generated AMI. Users may taint or otherwise recreate the resource in order
    /// to produce a fresh snapshot.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ec2.AmiFromInstance("example", new Aws.Ec2.AmiFromInstanceArgs
    ///         {
    ///             SourceInstanceId = "i-xxxxxxxx",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/amiFromInstance:AmiFromInstance")]
    public partial class AmiFromInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Machine architecture for created instances. Defaults to "x86_64".
        /// </summary>
        [Output("architecture")]
        public Output<string> Architecture { get; private set; } = null!;

        /// <summary>
        /// The ARN of the AMI.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
        /// </summary>
        [Output("bootMode")]
        public Output<string> BootMode { get; private set; } = null!;

        /// <summary>
        /// The date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
        /// </summary>
        [Output("deprecationTime")]
        public Output<string?> DeprecationTime { get; private set; } = null!;

        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        [Output("ebsBlockDevices")]
        public Output<ImmutableArray<Outputs.AmiFromInstanceEbsBlockDevice>> EbsBlockDevices { get; private set; } = null!;

        /// <summary>
        /// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        /// </summary>
        [Output("enaSupport")]
        public Output<bool> EnaSupport { get; private set; } = null!;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        [Output("ephemeralBlockDevices")]
        public Output<ImmutableArray<Outputs.AmiFromInstanceEphemeralBlockDevice>> EphemeralBlockDevices { get; private set; } = null!;

        [Output("hypervisor")]
        public Output<string> Hypervisor { get; private set; } = null!;

        /// <summary>
        /// Path to an S3 object containing an image manifest, e.g., created
        /// by the `ec2-upload-bundle` command in the EC2 command line tools.
        /// </summary>
        [Output("imageLocation")]
        public Output<string> ImageLocation { get; private set; } = null!;

        [Output("imageOwnerAlias")]
        public Output<string> ImageOwnerAlias { get; private set; } = null!;

        [Output("imageType")]
        public Output<string> ImageType { get; private set; } = null!;

        /// <summary>
        /// The id of the kernel image (AKI) that will be used as the paravirtual
        /// kernel in created instances.
        /// </summary>
        [Output("kernelId")]
        public Output<string> KernelId { get; private set; } = null!;

        [Output("manageEbsSnapshots")]
        public Output<bool> ManageEbsSnapshots { get; private set; } = null!;

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        [Output("platformDetails")]
        public Output<string> PlatformDetails { get; private set; } = null!;

        [Output("public")]
        public Output<bool> Public { get; private set; } = null!;

        /// <summary>
        /// The id of an initrd image (ARI) that will be used when booting the
        /// created instances.
        /// </summary>
        [Output("ramdiskId")]
        public Output<string> RamdiskId { get; private set; } = null!;

        /// <summary>
        /// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
        /// </summary>
        [Output("rootDeviceName")]
        public Output<string> RootDeviceName { get; private set; } = null!;

        [Output("rootSnapshotId")]
        public Output<string> RootSnapshotId { get; private set; } = null!;

        /// <summary>
        /// Boolean that overrides the behavior of stopping
        /// the instance before snapshotting. This is risky since it may cause a snapshot of an
        /// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
        /// guarantees that no filesystem writes will be underway at the time of snapshot.
        /// </summary>
        [Output("snapshotWithoutReboot")]
        public Output<bool?> SnapshotWithoutReboot { get; private set; } = null!;

        /// <summary>
        /// The id of the instance to use as the basis of the AMI.
        /// </summary>
        [Output("sourceInstanceId")]
        public Output<string> SourceInstanceId { get; private set; } = null!;

        /// <summary>
        /// When set to "simple" (the default), enables enhanced networking
        /// for created instances. No other value is supported at this time.
        /// </summary>
        [Output("sriovNetSupport")]
        public Output<string> SriovNetSupport { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        [Output("usageOperation")]
        public Output<string> UsageOperation { get; private set; } = null!;

        /// <summary>
        /// Keyword to choose what virtualization mode created instances
        /// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
        /// changes the set of further arguments that are required, as described below.
        /// </summary>
        [Output("virtualizationType")]
        public Output<string> VirtualizationType { get; private set; } = null!;


        /// <summary>
        /// Create a AmiFromInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AmiFromInstance(string name, AmiFromInstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/amiFromInstance:AmiFromInstance", name, args ?? new AmiFromInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AmiFromInstance(string name, Input<string> id, AmiFromInstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/amiFromInstance:AmiFromInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AmiFromInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AmiFromInstance Get(string name, Input<string> id, AmiFromInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new AmiFromInstance(name, id, state, options);
        }
    }

    public sealed class AmiFromInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
        /// </summary>
        [Input("deprecationTime")]
        public Input<string>? DeprecationTime { get; set; }

        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.AmiFromInstanceEbsBlockDeviceArgs>? _ebsBlockDevices;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiFromInstanceEbsBlockDeviceArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.AmiFromInstanceEbsBlockDeviceArgs>());
            set => _ebsBlockDevices = value;
        }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.AmiFromInstanceEphemeralBlockDeviceArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiFromInstanceEphemeralBlockDeviceArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.AmiFromInstanceEphemeralBlockDeviceArgs>());
            set => _ephemeralBlockDevices = value;
        }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Boolean that overrides the behavior of stopping
        /// the instance before snapshotting. This is risky since it may cause a snapshot of an
        /// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
        /// guarantees that no filesystem writes will be underway at the time of snapshot.
        /// </summary>
        [Input("snapshotWithoutReboot")]
        public Input<bool>? SnapshotWithoutReboot { get; set; }

        /// <summary>
        /// The id of the instance to use as the basis of the AMI.
        /// </summary>
        [Input("sourceInstanceId", required: true)]
        public Input<string> SourceInstanceId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AmiFromInstanceArgs()
        {
        }
    }

    public sealed class AmiFromInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Machine architecture for created instances. Defaults to "x86_64".
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The ARN of the AMI.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
        /// </summary>
        [Input("bootMode")]
        public Input<string>? BootMode { get; set; }

        /// <summary>
        /// The date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
        /// </summary>
        [Input("deprecationTime")]
        public Input<string>? DeprecationTime { get; set; }

        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.AmiFromInstanceEbsBlockDeviceGetArgs>? _ebsBlockDevices;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiFromInstanceEbsBlockDeviceGetArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.AmiFromInstanceEbsBlockDeviceGetArgs>());
            set => _ebsBlockDevices = value;
        }

        /// <summary>
        /// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        /// </summary>
        [Input("enaSupport")]
        public Input<bool>? EnaSupport { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.AmiFromInstanceEphemeralBlockDeviceGetArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiFromInstanceEphemeralBlockDeviceGetArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.AmiFromInstanceEphemeralBlockDeviceGetArgs>());
            set => _ephemeralBlockDevices = value;
        }

        [Input("hypervisor")]
        public Input<string>? Hypervisor { get; set; }

        /// <summary>
        /// Path to an S3 object containing an image manifest, e.g., created
        /// by the `ec2-upload-bundle` command in the EC2 command line tools.
        /// </summary>
        [Input("imageLocation")]
        public Input<string>? ImageLocation { get; set; }

        [Input("imageOwnerAlias")]
        public Input<string>? ImageOwnerAlias { get; set; }

        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        /// <summary>
        /// The id of the kernel image (AKI) that will be used as the paravirtual
        /// kernel in created instances.
        /// </summary>
        [Input("kernelId")]
        public Input<string>? KernelId { get; set; }

        [Input("manageEbsSnapshots")]
        public Input<bool>? ManageEbsSnapshots { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("platform")]
        public Input<string>? Platform { get; set; }

        [Input("platformDetails")]
        public Input<string>? PlatformDetails { get; set; }

        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// The id of an initrd image (ARI) that will be used when booting the
        /// created instances.
        /// </summary>
        [Input("ramdiskId")]
        public Input<string>? RamdiskId { get; set; }

        /// <summary>
        /// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
        /// </summary>
        [Input("rootDeviceName")]
        public Input<string>? RootDeviceName { get; set; }

        [Input("rootSnapshotId")]
        public Input<string>? RootSnapshotId { get; set; }

        /// <summary>
        /// Boolean that overrides the behavior of stopping
        /// the instance before snapshotting. This is risky since it may cause a snapshot of an
        /// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
        /// guarantees that no filesystem writes will be underway at the time of snapshot.
        /// </summary>
        [Input("snapshotWithoutReboot")]
        public Input<bool>? SnapshotWithoutReboot { get; set; }

        /// <summary>
        /// The id of the instance to use as the basis of the AMI.
        /// </summary>
        [Input("sourceInstanceId")]
        public Input<string>? SourceInstanceId { get; set; }

        /// <summary>
        /// When set to "simple" (the default), enables enhanced networking
        /// for created instances. No other value is supported at this time.
        /// </summary>
        [Input("sriovNetSupport")]
        public Input<string>? SriovNetSupport { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("usageOperation")]
        public Input<string>? UsageOperation { get; set; }

        /// <summary>
        /// Keyword to choose what virtualization mode created instances
        /// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
        /// changes the set of further arguments that are required, as described below.
        /// </summary>
        [Input("virtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        public AmiFromInstanceState()
        {
        }
    }
}
