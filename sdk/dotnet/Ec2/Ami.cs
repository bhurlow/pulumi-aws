// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// The AMI resource allows the creation and management of a completely-custom
    /// *Amazon Machine Image* (AMI).
    /// 
    /// If you just want to duplicate an existing AMI, possibly copying it to another
    /// region, it's better to use `aws.ec2.AmiCopy` instead.
    /// 
    /// If you just want to share an existing AMI with another AWS account,
    /// it's better to use `aws.ec2.AmiLaunchPermission` instead.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ami.html.markdown.
    /// </summary>
    public partial class Ami : Pulumi.CustomResource
    {
        /// <summary>
        /// Machine architecture for created instances. Defaults to "x86_64".
        /// </summary>
        [Output("architecture")]
        public Output<string?> Architecture { get; private set; } = null!;

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
        public Output<ImmutableArray<Outputs.AmiEbsBlockDevices>> EbsBlockDevices { get; private set; } = null!;

        /// <summary>
        /// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        /// </summary>
        [Output("enaSupport")]
        public Output<bool?> EnaSupport { get; private set; } = null!;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        [Output("ephemeralBlockDevices")]
        public Output<ImmutableArray<Outputs.AmiEphemeralBlockDevices>> EphemeralBlockDevices { get; private set; } = null!;

        /// <summary>
        /// Path to an S3 object containing an image manifest, e.g. created
        /// by the `ec2-upload-bundle` command in the EC2 command line tools.
        /// </summary>
        [Output("imageLocation")]
        public Output<string> ImageLocation { get; private set; } = null!;

        /// <summary>
        /// The id of the kernel image (AKI) that will be used as the paravirtual
        /// kernel in created instances.
        /// </summary>
        [Output("kernelId")]
        public Output<string?> KernelId { get; private set; } = null!;

        [Output("manageEbsSnapshots")]
        public Output<bool> ManageEbsSnapshots { get; private set; } = null!;

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of an initrd image (ARI) that will be used when booting the
        /// created instances.
        /// </summary>
        [Output("ramdiskId")]
        public Output<string?> RamdiskId { get; private set; } = null!;

        /// <summary>
        /// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
        /// </summary>
        [Output("rootDeviceName")]
        public Output<string?> RootDeviceName { get; private set; } = null!;

        /// <summary>
        /// The Snapshot ID for the root volume (for EBS-backed AMIs)
        /// </summary>
        [Output("rootSnapshotId")]
        public Output<string> RootSnapshotId { get; private set; } = null!;

        /// <summary>
        /// When set to "simple" (the default), enables enhanced networking
        /// for created instances. No other value is supported at this time.
        /// </summary>
        [Output("sriovNetSupport")]
        public Output<string?> SriovNetSupport { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Keyword to choose what virtualization mode created instances
        /// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
        /// changes the set of further arguments that are required, as described below.
        /// </summary>
        [Output("virtualizationType")]
        public Output<string?> VirtualizationType { get; private set; } = null!;


        /// <summary>
        /// Create a Ami resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ami(string name, AmiArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ec2/ami:Ami", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Ami(string name, Input<string> id, AmiState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/ami:Ami", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ami resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ami Get(string name, Input<string> id, AmiState? state = null, CustomResourceOptions? options = null)
        {
            return new Ami(name, id, state, options);
        }
    }

    public sealed class AmiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Machine architecture for created instances. Defaults to "x86_64".
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.AmiEbsBlockDevicesArgs>? _ebsBlockDevices;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiEbsBlockDevicesArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.AmiEbsBlockDevicesArgs>());
            set => _ebsBlockDevices = value;
        }

        /// <summary>
        /// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        /// </summary>
        [Input("enaSupport")]
        public Input<bool>? EnaSupport { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.AmiEphemeralBlockDevicesArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiEphemeralBlockDevicesArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.AmiEphemeralBlockDevicesArgs>());
            set => _ephemeralBlockDevices = value;
        }

        /// <summary>
        /// Path to an S3 object containing an image manifest, e.g. created
        /// by the `ec2-upload-bundle` command in the EC2 command line tools.
        /// </summary>
        [Input("imageLocation")]
        public Input<string>? ImageLocation { get; set; }

        /// <summary>
        /// The id of the kernel image (AKI) that will be used as the paravirtual
        /// kernel in created instances.
        /// </summary>
        [Input("kernelId")]
        public Input<string>? KernelId { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        /// <summary>
        /// When set to "simple" (the default), enables enhanced networking
        /// for created instances. No other value is supported at this time.
        /// </summary>
        [Input("sriovNetSupport")]
        public Input<string>? SriovNetSupport { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Keyword to choose what virtualization mode created instances
        /// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
        /// changes the set of further arguments that are required, as described below.
        /// </summary>
        [Input("virtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        public AmiArgs()
        {
        }
    }

    public sealed class AmiState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Machine architecture for created instances. Defaults to "x86_64".
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.AmiEbsBlockDevicesGetArgs>? _ebsBlockDevices;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiEbsBlockDevicesGetArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.AmiEbsBlockDevicesGetArgs>());
            set => _ebsBlockDevices = value;
        }

        /// <summary>
        /// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        /// </summary>
        [Input("enaSupport")]
        public Input<bool>? EnaSupport { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.AmiEphemeralBlockDevicesGetArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiEphemeralBlockDevicesGetArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.AmiEphemeralBlockDevicesGetArgs>());
            set => _ephemeralBlockDevices = value;
        }

        /// <summary>
        /// Path to an S3 object containing an image manifest, e.g. created
        /// by the `ec2-upload-bundle` command in the EC2 command line tools.
        /// </summary>
        [Input("imageLocation")]
        public Input<string>? ImageLocation { get; set; }

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

        /// <summary>
        /// The Snapshot ID for the root volume (for EBS-backed AMIs)
        /// </summary>
        [Input("rootSnapshotId")]
        public Input<string>? RootSnapshotId { get; set; }

        /// <summary>
        /// When set to "simple" (the default), enables enhanced networking
        /// for created instances. No other value is supported at this time.
        /// </summary>
        [Input("sriovNetSupport")]
        public Input<string>? SriovNetSupport { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Keyword to choose what virtualization mode created instances
        /// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
        /// changes the set of further arguments that are required, as described below.
        /// </summary>
        [Input("virtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        public AmiState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class AmiEbsBlockDevicesArgs : Pulumi.ResourceArgs
    {
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public AmiEbsBlockDevicesArgs()
        {
        }
    }

    public sealed class AmiEbsBlockDevicesGetArgs : Pulumi.ResourceArgs
    {
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public AmiEbsBlockDevicesGetArgs()
        {
        }
    }

    public sealed class AmiEphemeralBlockDevicesArgs : Pulumi.ResourceArgs
    {
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("virtualName", required: true)]
        public Input<string> VirtualName { get; set; } = null!;

        public AmiEphemeralBlockDevicesArgs()
        {
        }
    }

    public sealed class AmiEphemeralBlockDevicesGetArgs : Pulumi.ResourceArgs
    {
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("virtualName", required: true)]
        public Input<string> VirtualName { get; set; } = null!;

        public AmiEphemeralBlockDevicesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AmiEbsBlockDevices
    {
        public readonly bool? DeleteOnTermination;
        public readonly string DeviceName;
        public readonly bool? Encrypted;
        public readonly int? Iops;
        public readonly string? SnapshotId;
        public readonly int VolumeSize;
        public readonly string? VolumeType;

        [OutputConstructor]
        private AmiEbsBlockDevices(
            bool? deleteOnTermination,
            string deviceName,
            bool? encrypted,
            int? iops,
            string? snapshotId,
            int volumeSize,
            string? volumeType)
        {
            DeleteOnTermination = deleteOnTermination;
            DeviceName = deviceName;
            Encrypted = encrypted;
            Iops = iops;
            SnapshotId = snapshotId;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }

    [OutputType]
    public sealed class AmiEphemeralBlockDevices
    {
        public readonly string DeviceName;
        public readonly string VirtualName;

        [OutputConstructor]
        private AmiEphemeralBlockDevices(
            string deviceName,
            string virtualName)
        {
            DeviceName = deviceName;
            VirtualName = virtualName;
        }
    }
    }
}
