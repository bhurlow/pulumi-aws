// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fsx
{
    /// <summary>
    /// Manages a FSx Lustre File System. See the [FSx Lustre Guide](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html) for more information.
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
    ///         var example = new Aws.Fsx.LustreFileSystem("example", new Aws.Fsx.LustreFileSystemArgs
    ///         {
    ///             ImportPath = $"s3://{aws_s3_bucket.Example.Bucket}",
    ///             StorageCapacity = 1200,
    ///             SubnetIds = 
    ///             {
    ///                 aws_subnet.Example.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// FSx File Systems can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:fsx/lustreFileSystem:LustreFileSystem example fs-543ab12b1ca672f33
    /// ```
    /// 
    ///  Certain resource arguments, like `security_group_ids`, do not have a FSx API method for reading the information after creation. If the argument is set in the provider configuration on an imported resource, this provider will always show a difference. To workaround this behavior, either omit the argument from the provider configuration or use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to hide the difference, e.g. terraform resource "aws_fsx_lustre_file_system" "example" {
    /// 
    /// # ... other configuration ...
    /// 
    ///  security_group_ids = [aws_security_group.example.id]
    /// 
    /// # There is no FSx API for reading security_group_ids
    /// 
    ///  lifecycle {
    /// 
    ///  ignore_changes = [security_group_ids]
    /// 
    ///  } }
    /// </summary>
    [AwsResourceType("aws:fsx/lustreFileSystem:LustreFileSystem")]
    public partial class LustreFileSystem : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// How Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. see [Auto Import Data Repo](https://docs.aws.amazon.com/fsx/latest/LustreGuide/autoimport-data-repo.html) for more details.
        /// </summary>
        [Output("autoImportPolicy")]
        public Output<string> AutoImportPolicy { get; private set; } = null!;

        /// <summary>
        /// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. only valid for `PERSISTENT_1` deployment_type.
        /// </summary>
        [Output("automaticBackupRetentionDays")]
        public Output<int> AutomaticBackupRetentionDays { get; private set; } = null!;

        /// <summary>
        /// A boolean flag indicating whether tags for the file system should be copied to backups. Applicable for `PERSISTENT_1` deployment_type. The default value is false.
        /// </summary>
        [Output("copyTagsToBackups")]
        public Output<bool?> CopyTagsToBackups { get; private set; } = null!;

        /// <summary>
        /// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. only valid for `PERSISTENT_1` deployment_type. Requires `automatic_backup_retention_days` to be set.
        /// </summary>
        [Output("dailyAutomaticBackupStartTime")]
        public Output<string> DailyAutomaticBackupStartTime { get; private set; } = null!;

        /// <summary>
        /// - The filesystem deployment type. One of: `SCRATCH_1`, `SCRATCH_2`, `PERSISTENT_1`. `SCRATCH_1` deployment types cannot have `storage_capacity` increased.
        /// </summary>
        [Output("deploymentType")]
        public Output<string?> DeploymentType { get; private set; } = null!;

        /// <summary>
        /// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// - The type of drive cache used by `PERSISTENT_1` filesystems that are provisioned with `HDD` storage_type. Required for `HDD` storage_type, set to either `READ` or `NONE`.
        /// </summary>
        [Output("driveCacheType")]
        public Output<string?> DriveCacheType { get; private set; } = null!;

        /// <summary>
        /// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        /// </summary>
        [Output("exportPath")]
        public Output<string> ExportPath { get; private set; } = null!;

        /// <summary>
        /// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        /// </summary>
        [Output("importPath")]
        public Output<string?> ImportPath { get; private set; } = null!;

        /// <summary>
        /// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        /// </summary>
        [Output("importedFileChunkSize")]
        public Output<int> ImportedFileChunkSize { get; private set; } = null!;

        /// <summary>
        /// ARN for the KMS Key to encrypt the file system at rest, applicable for `PERSISTENT_1` deployment_type. Defaults to an AWS managed KMS Key.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The value to be used when mounting the filesystem.
        /// </summary>
        [Output("mountName")]
        public Output<string> MountName { get; private set; } = null!;

        /// <summary>
        /// Set of Elastic Network Interface identifiers from which the file system is accessible. As explained in the [documentation](https://docs.aws.amazon.com/fsx/latest/LustreGuide/mounting-on-premises.html), the first network interface returned is the primary network interface.
        /// </summary>
        [Output("networkInterfaceIds")]
        public Output<ImmutableArray<string>> NetworkInterfaceIds { get; private set; } = null!;

        /// <summary>
        /// AWS account identifier that created the file system.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the `PERSISTENT_1` deployment_type. Valid values for `SSD` storage_type are 50, 100, 200. Valid values for `HDD` storage_type are 12, 40.
        /// </summary>
        [Output("perUnitStorageThroughput")]
        public Output<int?> PerUnitStorageThroughput { get; private set; } = null!;

        /// <summary>
        /// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
        /// </summary>
        [Output("storageCapacity")]
        public Output<int> StorageCapacity { get; private set; } = null!;

        /// <summary>
        /// - The filesystem storage type. Either `SSD` or `HDD`, defaults to `SSD`. `HDD` is only supported on `PERSISTENT_1` deployment types.
        /// </summary>
        [Output("storageType")]
        public Output<string?> StorageType { get; private set; } = null!;

        /// <summary>
        /// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        /// </summary>
        [Output("subnetIds")]
        public Output<string> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the file system.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Identifier of the Virtual Private Cloud for the file system.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        /// </summary>
        [Output("weeklyMaintenanceStartTime")]
        public Output<string> WeeklyMaintenanceStartTime { get; private set; } = null!;


        /// <summary>
        /// Create a LustreFileSystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LustreFileSystem(string name, LustreFileSystemArgs args, CustomResourceOptions? options = null)
            : base("aws:fsx/lustreFileSystem:LustreFileSystem", name, args ?? new LustreFileSystemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LustreFileSystem(string name, Input<string> id, LustreFileSystemState? state = null, CustomResourceOptions? options = null)
            : base("aws:fsx/lustreFileSystem:LustreFileSystem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LustreFileSystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LustreFileSystem Get(string name, Input<string> id, LustreFileSystemState? state = null, CustomResourceOptions? options = null)
        {
            return new LustreFileSystem(name, id, state, options);
        }
    }

    public sealed class LustreFileSystemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. see [Auto Import Data Repo](https://docs.aws.amazon.com/fsx/latest/LustreGuide/autoimport-data-repo.html) for more details.
        /// </summary>
        [Input("autoImportPolicy")]
        public Input<string>? AutoImportPolicy { get; set; }

        /// <summary>
        /// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. only valid for `PERSISTENT_1` deployment_type.
        /// </summary>
        [Input("automaticBackupRetentionDays")]
        public Input<int>? AutomaticBackupRetentionDays { get; set; }

        /// <summary>
        /// A boolean flag indicating whether tags for the file system should be copied to backups. Applicable for `PERSISTENT_1` deployment_type. The default value is false.
        /// </summary>
        [Input("copyTagsToBackups")]
        public Input<bool>? CopyTagsToBackups { get; set; }

        /// <summary>
        /// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. only valid for `PERSISTENT_1` deployment_type. Requires `automatic_backup_retention_days` to be set.
        /// </summary>
        [Input("dailyAutomaticBackupStartTime")]
        public Input<string>? DailyAutomaticBackupStartTime { get; set; }

        /// <summary>
        /// - The filesystem deployment type. One of: `SCRATCH_1`, `SCRATCH_2`, `PERSISTENT_1`. `SCRATCH_1` deployment types cannot have `storage_capacity` increased.
        /// </summary>
        [Input("deploymentType")]
        public Input<string>? DeploymentType { get; set; }

        /// <summary>
        /// - The type of drive cache used by `PERSISTENT_1` filesystems that are provisioned with `HDD` storage_type. Required for `HDD` storage_type, set to either `READ` or `NONE`.
        /// </summary>
        [Input("driveCacheType")]
        public Input<string>? DriveCacheType { get; set; }

        /// <summary>
        /// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        /// </summary>
        [Input("exportPath")]
        public Input<string>? ExportPath { get; set; }

        /// <summary>
        /// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        /// </summary>
        [Input("importPath")]
        public Input<string>? ImportPath { get; set; }

        /// <summary>
        /// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        /// </summary>
        [Input("importedFileChunkSize")]
        public Input<int>? ImportedFileChunkSize { get; set; }

        /// <summary>
        /// ARN for the KMS Key to encrypt the file system at rest, applicable for `PERSISTENT_1` deployment_type. Defaults to an AWS managed KMS Key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the `PERSISTENT_1` deployment_type. Valid values for `SSD` storage_type are 50, 100, 200. Valid values for `HDD` storage_type are 12, 40.
        /// </summary>
        [Input("perUnitStorageThroughput")]
        public Input<int>? PerUnitStorageThroughput { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
        /// </summary>
        [Input("storageCapacity", required: true)]
        public Input<int> StorageCapacity { get; set; } = null!;

        /// <summary>
        /// - The filesystem storage type. Either `SSD` or `HDD`, defaults to `SSD`. `HDD` is only supported on `PERSISTENT_1` deployment types.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        /// <summary>
        /// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        /// </summary>
        [Input("subnetIds", required: true)]
        public Input<string> SubnetIds { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the file system.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        /// </summary>
        [Input("weeklyMaintenanceStartTime")]
        public Input<string>? WeeklyMaintenanceStartTime { get; set; }

        public LustreFileSystemArgs()
        {
        }
    }

    public sealed class LustreFileSystemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// How Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. see [Auto Import Data Repo](https://docs.aws.amazon.com/fsx/latest/LustreGuide/autoimport-data-repo.html) for more details.
        /// </summary>
        [Input("autoImportPolicy")]
        public Input<string>? AutoImportPolicy { get; set; }

        /// <summary>
        /// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. only valid for `PERSISTENT_1` deployment_type.
        /// </summary>
        [Input("automaticBackupRetentionDays")]
        public Input<int>? AutomaticBackupRetentionDays { get; set; }

        /// <summary>
        /// A boolean flag indicating whether tags for the file system should be copied to backups. Applicable for `PERSISTENT_1` deployment_type. The default value is false.
        /// </summary>
        [Input("copyTagsToBackups")]
        public Input<bool>? CopyTagsToBackups { get; set; }

        /// <summary>
        /// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. only valid for `PERSISTENT_1` deployment_type. Requires `automatic_backup_retention_days` to be set.
        /// </summary>
        [Input("dailyAutomaticBackupStartTime")]
        public Input<string>? DailyAutomaticBackupStartTime { get; set; }

        /// <summary>
        /// - The filesystem deployment type. One of: `SCRATCH_1`, `SCRATCH_2`, `PERSISTENT_1`. `SCRATCH_1` deployment types cannot have `storage_capacity` increased.
        /// </summary>
        [Input("deploymentType")]
        public Input<string>? DeploymentType { get; set; }

        /// <summary>
        /// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// - The type of drive cache used by `PERSISTENT_1` filesystems that are provisioned with `HDD` storage_type. Required for `HDD` storage_type, set to either `READ` or `NONE`.
        /// </summary>
        [Input("driveCacheType")]
        public Input<string>? DriveCacheType { get; set; }

        /// <summary>
        /// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        /// </summary>
        [Input("exportPath")]
        public Input<string>? ExportPath { get; set; }

        /// <summary>
        /// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        /// </summary>
        [Input("importPath")]
        public Input<string>? ImportPath { get; set; }

        /// <summary>
        /// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        /// </summary>
        [Input("importedFileChunkSize")]
        public Input<int>? ImportedFileChunkSize { get; set; }

        /// <summary>
        /// ARN for the KMS Key to encrypt the file system at rest, applicable for `PERSISTENT_1` deployment_type. Defaults to an AWS managed KMS Key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The value to be used when mounting the filesystem.
        /// </summary>
        [Input("mountName")]
        public Input<string>? MountName { get; set; }

        [Input("networkInterfaceIds")]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// Set of Elastic Network Interface identifiers from which the file system is accessible. As explained in the [documentation](https://docs.aws.amazon.com/fsx/latest/LustreGuide/mounting-on-premises.html), the first network interface returned is the primary network interface.
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// AWS account identifier that created the file system.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the `PERSISTENT_1` deployment_type. Valid values for `SSD` storage_type are 50, 100, 200. Valid values for `HDD` storage_type are 12, 40.
        /// </summary>
        [Input("perUnitStorageThroughput")]
        public Input<int>? PerUnitStorageThroughput { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
        /// </summary>
        [Input("storageCapacity")]
        public Input<int>? StorageCapacity { get; set; }

        /// <summary>
        /// - The filesystem storage type. Either `SSD` or `HDD`, defaults to `SSD`. `HDD` is only supported on `PERSISTENT_1` deployment types.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        /// <summary>
        /// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        /// </summary>
        [Input("subnetIds")]
        public Input<string>? SubnetIds { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the file system.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier of the Virtual Private Cloud for the file system.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        /// </summary>
        [Input("weeklyMaintenanceStartTime")]
        public Input<string>? WeeklyMaintenanceStartTime { get; set; }

        public LustreFileSystemState()
        {
        }
    }
}
