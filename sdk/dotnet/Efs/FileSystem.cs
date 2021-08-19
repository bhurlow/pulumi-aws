// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Efs
{
    /// <summary>
    /// Provides an Elastic File System (EFS) File System resource.
    /// 
    /// ## Example Usage
    /// ### EFS File System w/ tags
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Aws.Efs.FileSystem("foo", new Aws.Efs.FileSystemArgs
    ///         {
    ///             Tags = 
    ///             {
    ///                 { "Name", "MyProduct" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Using lifecycle policy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooWithLifecylePolicy = new Aws.Efs.FileSystem("fooWithLifecylePolicy", new Aws.Efs.FileSystemArgs
    ///         {
    ///             LifecyclePolicy = new Aws.Efs.Inputs.FileSystemLifecyclePolicyArgs
    ///             {
    ///                 TransitionToIa = "AFTER_30_DAYS",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// The EFS file systems can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:efs/fileSystem:FileSystem foo fs-6fa144c6
    /// ```
    /// </summary>
    [AwsResourceType("aws:efs/fileSystem:FileSystem")]
    public partial class FileSystem : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The identifier of the Availability Zone in which the file system's One Zone storage classes exist.
        /// </summary>
        [Output("availabilityZoneId")]
        public Output<string> AvailabilityZoneId { get; private set; } = null!;

        /// <summary>
        /// the AWS Availability Zone in which to create the file system. Used to create a file system that uses One Zone storage classes. See [user guide](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) for more information.
        /// </summary>
        [Output("availabilityZoneName")]
        public Output<string> AvailabilityZoneName { get; private set; } = null!;

        /// <summary>
        /// A unique name (a maximum of 64 characters are allowed)
        /// used as reference when creating the Elastic File System to ensure idempotent file
        /// system creation. By default generated by this provider. See [Elastic File System]
        /// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
        /// </summary>
        [Output("creationToken")]
        public Output<string> CreationToken { get; private set; } = null!;

        /// <summary>
        /// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Output("encrypted")]
        public Output<bool> Encrypted { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
        /// </summary>
        [Output("lifecyclePolicy")]
        public Output<Outputs.FileSystemLifecyclePolicy?> LifecyclePolicy { get; private set; } = null!;

        /// <summary>
        /// The current number of mount targets that the file system has.
        /// </summary>
        [Output("numberOfMountTargets")]
        public Output<int> NumberOfMountTargets { get; private set; } = null!;

        /// <summary>
        /// The AWS account that created the file system. If the file system was createdby an IAM user, the parent account to which the user belongs is the owner.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
        /// </summary>
        [Output("performanceMode")]
        public Output<string> PerformanceMode { get; private set; } = null!;

        /// <summary>
        /// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughput_mode` set to `provisioned`.
        /// </summary>
        [Output("provisionedThroughputInMibps")]
        public Output<double?> ProvisionedThroughputInMibps { get; private set; } = null!;

        /// <summary>
        /// The latest known metered size (in bytes) of data stored in the file system, the value is not the exact size that the file system was at any point in time. See Size In Bytes.
        /// </summary>
        [Output("sizeInBytes")]
        public Output<ImmutableArray<Outputs.FileSystemSizeInByte>> SizeInBytes { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisioned_throughput_in_mibps`.
        /// </summary>
        [Output("throughputMode")]
        public Output<string?> ThroughputMode { get; private set; } = null!;


        /// <summary>
        /// Create a FileSystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileSystem(string name, FileSystemArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:efs/fileSystem:FileSystem", name, args ?? new FileSystemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileSystem(string name, Input<string> id, FileSystemState? state = null, CustomResourceOptions? options = null)
            : base("aws:efs/fileSystem:FileSystem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FileSystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileSystem Get(string name, Input<string> id, FileSystemState? state = null, CustomResourceOptions? options = null)
        {
            return new FileSystem(name, id, state, options);
        }
    }

    public sealed class FileSystemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the AWS Availability Zone in which to create the file system. Used to create a file system that uses One Zone storage classes. See [user guide](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) for more information.
        /// </summary>
        [Input("availabilityZoneName")]
        public Input<string>? AvailabilityZoneName { get; set; }

        /// <summary>
        /// A unique name (a maximum of 64 characters are allowed)
        /// used as reference when creating the Elastic File System to ensure idempotent file
        /// system creation. By default generated by this provider. See [Elastic File System]
        /// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
        /// </summary>
        [Input("creationToken")]
        public Input<string>? CreationToken { get; set; }

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
        /// </summary>
        [Input("lifecyclePolicy")]
        public Input<Inputs.FileSystemLifecyclePolicyArgs>? LifecyclePolicy { get; set; }

        /// <summary>
        /// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
        /// </summary>
        [Input("performanceMode")]
        public Input<string>? PerformanceMode { get; set; }

        /// <summary>
        /// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughput_mode` set to `provisioned`.
        /// </summary>
        [Input("provisionedThroughputInMibps")]
        public Input<double>? ProvisionedThroughputInMibps { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisioned_throughput_in_mibps`.
        /// </summary>
        [Input("throughputMode")]
        public Input<string>? ThroughputMode { get; set; }

        public FileSystemArgs()
        {
        }
    }

    public sealed class FileSystemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The identifier of the Availability Zone in which the file system's One Zone storage classes exist.
        /// </summary>
        [Input("availabilityZoneId")]
        public Input<string>? AvailabilityZoneId { get; set; }

        /// <summary>
        /// the AWS Availability Zone in which to create the file system. Used to create a file system that uses One Zone storage classes. See [user guide](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) for more information.
        /// </summary>
        [Input("availabilityZoneName")]
        public Input<string>? AvailabilityZoneName { get; set; }

        /// <summary>
        /// A unique name (a maximum of 64 characters are allowed)
        /// used as reference when creating the Elastic File System to ensure idempotent file
        /// system creation. By default generated by this provider. See [Elastic File System]
        /// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
        /// </summary>
        [Input("creationToken")]
        public Input<string>? CreationToken { get; set; }

        /// <summary>
        /// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
        /// </summary>
        [Input("lifecyclePolicy")]
        public Input<Inputs.FileSystemLifecyclePolicyGetArgs>? LifecyclePolicy { get; set; }

        /// <summary>
        /// The current number of mount targets that the file system has.
        /// </summary>
        [Input("numberOfMountTargets")]
        public Input<int>? NumberOfMountTargets { get; set; }

        /// <summary>
        /// The AWS account that created the file system. If the file system was createdby an IAM user, the parent account to which the user belongs is the owner.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
        /// </summary>
        [Input("performanceMode")]
        public Input<string>? PerformanceMode { get; set; }

        /// <summary>
        /// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughput_mode` set to `provisioned`.
        /// </summary>
        [Input("provisionedThroughputInMibps")]
        public Input<double>? ProvisionedThroughputInMibps { get; set; }

        [Input("sizeInBytes")]
        private InputList<Inputs.FileSystemSizeInByteGetArgs>? _sizeInBytes;

        /// <summary>
        /// The latest known metered size (in bytes) of data stored in the file system, the value is not the exact size that the file system was at any point in time. See Size In Bytes.
        /// </summary>
        public InputList<Inputs.FileSystemSizeInByteGetArgs> SizeInBytes
        {
            get => _sizeInBytes ?? (_sizeInBytes = new InputList<Inputs.FileSystemSizeInByteGetArgs>());
            set => _sizeInBytes = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;
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

        /// <summary>
        /// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisioned_throughput_in_mibps`.
        /// </summary>
        [Input("throughputMode")]
        public Input<string>? ThroughputMode { get; set; }

        public FileSystemState()
        {
        }
    }
}
