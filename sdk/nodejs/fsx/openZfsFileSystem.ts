// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an Amazon FSx for OpenZFS file system.
 * See the [FSx OpenZFS User Guide](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/what-is-fsx.html) for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.fsx.OpenZfsFileSystem("test", {
 *     storageCapacity: 64,
 *     subnetIds: [aws_subnet.test1.id],
 *     deploymentType: "SINGLE_AZ_1",
 *     throughputCapacity: 64,
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import FSx File Systems using the `id`. For example:
 *
 * ```sh
 *  $ pulumi import aws:fsx/openZfsFileSystem:OpenZfsFileSystem example fs-543ab12b1ca672f33
 * ```
 *  Certain resource arguments, like `security_group_ids`, do not have a FSx API method for reading the information after creation. If the argument is set in the TODO configuration on an imported resource, TODO will always show a difference. To workaround this behavior, either omit the argument from the TODO configuration or use `ignore_changes` to hide the difference. For example:
 */
export class OpenZfsFileSystem extends pulumi.CustomResource {
    /**
     * Get an existing OpenZfsFileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OpenZfsFileSystemState, opts?: pulumi.CustomResourceOptions): OpenZfsFileSystem {
        return new OpenZfsFileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:fsx/openZfsFileSystem:OpenZfsFileSystem';

    /**
     * Returns true if the given object is an instance of OpenZfsFileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenZfsFileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenZfsFileSystem.__pulumiType;
    }

    /**
     * Amazon Resource Name of the file system.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
     */
    public readonly automaticBackupRetentionDays!: pulumi.Output<number | undefined>;
    /**
     * The ID of the source backup to create the filesystem from.
     */
    public readonly backupId!: pulumi.Output<string | undefined>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to backups. The default value is false.
     */
    public readonly copyTagsToBackups!: pulumi.Output<boolean | undefined>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
     */
    public readonly copyTagsToVolumes!: pulumi.Output<boolean | undefined>;
    /**
     * A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
     */
    public readonly dailyAutomaticBackupStartTime!: pulumi.Output<string>;
    /**
     * The filesystem deployment type. Valid values: `SINGLE_AZ_1`, `SINGLE_AZ_2` and `MULTI_AZ_1`.
     */
    public readonly deploymentType!: pulumi.Output<string>;
    /**
     * The SSD IOPS configuration for the Amazon FSx for OpenZFS file system. See Disk Iops Configuration below.
     */
    public readonly diskIopsConfiguration!: pulumi.Output<outputs.fsx.OpenZfsFileSystemDiskIopsConfiguration>;
    /**
     * DNS name for the file system, e.g., `fs-12345678.fsx.us-west-2.amazonaws.com`
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * (Multi-AZ only) Specifies the IP address range in which the endpoints to access your file system will be created.
     */
    public readonly endpointIpAddressRange!: pulumi.Output<string>;
    /**
     * ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Set of Elastic Network Interface identifiers from which the file system is accessible The first network interface returned is the primary network interface.
     */
    public /*out*/ readonly networkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * AWS account identifier that created the file system.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * (Multi-AZ only) Required when `deploymentType` is set to `MULTI_AZ_1`. This specifies the subnet in which you want the preferred file server to be located.
     */
    public readonly preferredSubnetId!: pulumi.Output<string | undefined>;
    /**
     * The configuration for the root volume of the file system. All other volumes are children or the root volume. See Root Volume Configuration below.
     */
    public readonly rootVolumeConfiguration!: pulumi.Output<outputs.fsx.OpenZfsFileSystemRootVolumeConfiguration>;
    /**
     * Identifier of the root volume, e.g., `fsvol-12345678`
     */
    public /*out*/ readonly rootVolumeId!: pulumi.Output<string>;
    /**
     * (Multi-AZ only) Specifies the route tables in which Amazon FSx creates the rules for routing traffic to the correct file server. You should specify all virtual private cloud (VPC) route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
     */
    public readonly routeTableIds!: pulumi.Output<string[]>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The storage capacity (GiB) of the file system. Valid values between `64` and `524288`.
     */
    public readonly storageCapacity!: pulumi.Output<number | undefined>;
    /**
     * The filesystem storage type. Only `SSD` is supported.
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from.
     */
    public readonly subnetIds!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Throughput (MB/s) of the file system. Valid values depend on `deploymentType`. Must be one of `64`, `128`, `256`, `512`, `1024`, `2048`, `3072`, `4096` for `SINGLE_AZ_1`. Must be one of `160`, `320`, `640`, `1280`, `2560`, `3840`, `5120`, `7680`, `10240` for `SINGLE_AZ_2`.
     */
    public readonly throughputCapacity!: pulumi.Output<number>;
    /**
     * Identifier of the Virtual Private Cloud for the file system.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    public readonly weeklyMaintenanceStartTime!: pulumi.Output<string>;

    /**
     * Create a OpenZfsFileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenZfsFileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OpenZfsFileSystemArgs | OpenZfsFileSystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OpenZfsFileSystemState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["automaticBackupRetentionDays"] = state ? state.automaticBackupRetentionDays : undefined;
            resourceInputs["backupId"] = state ? state.backupId : undefined;
            resourceInputs["copyTagsToBackups"] = state ? state.copyTagsToBackups : undefined;
            resourceInputs["copyTagsToVolumes"] = state ? state.copyTagsToVolumes : undefined;
            resourceInputs["dailyAutomaticBackupStartTime"] = state ? state.dailyAutomaticBackupStartTime : undefined;
            resourceInputs["deploymentType"] = state ? state.deploymentType : undefined;
            resourceInputs["diskIopsConfiguration"] = state ? state.diskIopsConfiguration : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["endpointIpAddressRange"] = state ? state.endpointIpAddressRange : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["networkInterfaceIds"] = state ? state.networkInterfaceIds : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["preferredSubnetId"] = state ? state.preferredSubnetId : undefined;
            resourceInputs["rootVolumeConfiguration"] = state ? state.rootVolumeConfiguration : undefined;
            resourceInputs["rootVolumeId"] = state ? state.rootVolumeId : undefined;
            resourceInputs["routeTableIds"] = state ? state.routeTableIds : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["storageCapacity"] = state ? state.storageCapacity : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["throughputCapacity"] = state ? state.throughputCapacity : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["weeklyMaintenanceStartTime"] = state ? state.weeklyMaintenanceStartTime : undefined;
        } else {
            const args = argsOrState as OpenZfsFileSystemArgs | undefined;
            if ((!args || args.deploymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentType'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if ((!args || args.throughputCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'throughputCapacity'");
            }
            resourceInputs["automaticBackupRetentionDays"] = args ? args.automaticBackupRetentionDays : undefined;
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["copyTagsToBackups"] = args ? args.copyTagsToBackups : undefined;
            resourceInputs["copyTagsToVolumes"] = args ? args.copyTagsToVolumes : undefined;
            resourceInputs["dailyAutomaticBackupStartTime"] = args ? args.dailyAutomaticBackupStartTime : undefined;
            resourceInputs["deploymentType"] = args ? args.deploymentType : undefined;
            resourceInputs["diskIopsConfiguration"] = args ? args.diskIopsConfiguration : undefined;
            resourceInputs["endpointIpAddressRange"] = args ? args.endpointIpAddressRange : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["preferredSubnetId"] = args ? args.preferredSubnetId : undefined;
            resourceInputs["rootVolumeConfiguration"] = args ? args.rootVolumeConfiguration : undefined;
            resourceInputs["routeTableIds"] = args ? args.routeTableIds : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["storageCapacity"] = args ? args.storageCapacity : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["throughputCapacity"] = args ? args.throughputCapacity : undefined;
            resourceInputs["weeklyMaintenanceStartTime"] = args ? args.weeklyMaintenanceStartTime : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["networkInterfaceIds"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["rootVolumeId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(OpenZfsFileSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OpenZfsFileSystem resources.
 */
export interface OpenZfsFileSystemState {
    /**
     * Amazon Resource Name of the file system.
     */
    arn?: pulumi.Input<string>;
    /**
     * The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
     */
    automaticBackupRetentionDays?: pulumi.Input<number>;
    /**
     * The ID of the source backup to create the filesystem from.
     */
    backupId?: pulumi.Input<string>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to backups. The default value is false.
     */
    copyTagsToBackups?: pulumi.Input<boolean>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
     */
    copyTagsToVolumes?: pulumi.Input<boolean>;
    /**
     * A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
     */
    dailyAutomaticBackupStartTime?: pulumi.Input<string>;
    /**
     * The filesystem deployment type. Valid values: `SINGLE_AZ_1`, `SINGLE_AZ_2` and `MULTI_AZ_1`.
     */
    deploymentType?: pulumi.Input<string>;
    /**
     * The SSD IOPS configuration for the Amazon FSx for OpenZFS file system. See Disk Iops Configuration below.
     */
    diskIopsConfiguration?: pulumi.Input<inputs.fsx.OpenZfsFileSystemDiskIopsConfiguration>;
    /**
     * DNS name for the file system, e.g., `fs-12345678.fsx.us-west-2.amazonaws.com`
     */
    dnsName?: pulumi.Input<string>;
    /**
     * (Multi-AZ only) Specifies the IP address range in which the endpoints to access your file system will be created.
     */
    endpointIpAddressRange?: pulumi.Input<string>;
    /**
     * ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Set of Elastic Network Interface identifiers from which the file system is accessible The first network interface returned is the primary network interface.
     */
    networkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * AWS account identifier that created the file system.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * (Multi-AZ only) Required when `deploymentType` is set to `MULTI_AZ_1`. This specifies the subnet in which you want the preferred file server to be located.
     */
    preferredSubnetId?: pulumi.Input<string>;
    /**
     * The configuration for the root volume of the file system. All other volumes are children or the root volume. See Root Volume Configuration below.
     */
    rootVolumeConfiguration?: pulumi.Input<inputs.fsx.OpenZfsFileSystemRootVolumeConfiguration>;
    /**
     * Identifier of the root volume, e.g., `fsvol-12345678`
     */
    rootVolumeId?: pulumi.Input<string>;
    /**
     * (Multi-AZ only) Specifies the route tables in which Amazon FSx creates the rules for routing traffic to the correct file server. You should specify all virtual private cloud (VPC) route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
     */
    routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage capacity (GiB) of the file system. Valid values between `64` and `524288`.
     */
    storageCapacity?: pulumi.Input<number>;
    /**
     * The filesystem storage type. Only `SSD` is supported.
     */
    storageType?: pulumi.Input<string>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from.
     */
    subnetIds?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Throughput (MB/s) of the file system. Valid values depend on `deploymentType`. Must be one of `64`, `128`, `256`, `512`, `1024`, `2048`, `3072`, `4096` for `SINGLE_AZ_1`. Must be one of `160`, `320`, `640`, `1280`, `2560`, `3840`, `5120`, `7680`, `10240` for `SINGLE_AZ_2`.
     */
    throughputCapacity?: pulumi.Input<number>;
    /**
     * Identifier of the Virtual Private Cloud for the file system.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    weeklyMaintenanceStartTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OpenZfsFileSystem resource.
 */
export interface OpenZfsFileSystemArgs {
    /**
     * The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
     */
    automaticBackupRetentionDays?: pulumi.Input<number>;
    /**
     * The ID of the source backup to create the filesystem from.
     */
    backupId?: pulumi.Input<string>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to backups. The default value is false.
     */
    copyTagsToBackups?: pulumi.Input<boolean>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
     */
    copyTagsToVolumes?: pulumi.Input<boolean>;
    /**
     * A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
     */
    dailyAutomaticBackupStartTime?: pulumi.Input<string>;
    /**
     * The filesystem deployment type. Valid values: `SINGLE_AZ_1`, `SINGLE_AZ_2` and `MULTI_AZ_1`.
     */
    deploymentType: pulumi.Input<string>;
    /**
     * The SSD IOPS configuration for the Amazon FSx for OpenZFS file system. See Disk Iops Configuration below.
     */
    diskIopsConfiguration?: pulumi.Input<inputs.fsx.OpenZfsFileSystemDiskIopsConfiguration>;
    /**
     * (Multi-AZ only) Specifies the IP address range in which the endpoints to access your file system will be created.
     */
    endpointIpAddressRange?: pulumi.Input<string>;
    /**
     * ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * (Multi-AZ only) Required when `deploymentType` is set to `MULTI_AZ_1`. This specifies the subnet in which you want the preferred file server to be located.
     */
    preferredSubnetId?: pulumi.Input<string>;
    /**
     * The configuration for the root volume of the file system. All other volumes are children or the root volume. See Root Volume Configuration below.
     */
    rootVolumeConfiguration?: pulumi.Input<inputs.fsx.OpenZfsFileSystemRootVolumeConfiguration>;
    /**
     * (Multi-AZ only) Specifies the route tables in which Amazon FSx creates the rules for routing traffic to the correct file server. You should specify all virtual private cloud (VPC) route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
     */
    routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage capacity (GiB) of the file system. Valid values between `64` and `524288`.
     */
    storageCapacity?: pulumi.Input<number>;
    /**
     * The filesystem storage type. Only `SSD` is supported.
     */
    storageType?: pulumi.Input<string>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from.
     */
    subnetIds: pulumi.Input<string>;
    /**
     * A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Throughput (MB/s) of the file system. Valid values depend on `deploymentType`. Must be one of `64`, `128`, `256`, `512`, `1024`, `2048`, `3072`, `4096` for `SINGLE_AZ_1`. Must be one of `160`, `320`, `640`, `1280`, `2560`, `3840`, `5120`, `7680`, `10240` for `SINGLE_AZ_2`.
     */
    throughputCapacity: pulumi.Input<number>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    weeklyMaintenanceStartTime?: pulumi.Input<string>;
}
