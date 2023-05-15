// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an Amazon FSx for NetApp ONTAP file system.
 * See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/what-is-fsx-ontap.html) for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.fsx.OntapFileSystem("test", {
 *     storageCapacity: 1024,
 *     subnetIds: [
 *         aws_subnet.test1.id,
 *         aws_subnet.test2.id,
 *     ],
 *     deploymentType: "MULTI_AZ_1",
 *     throughputCapacity: 512,
 *     preferredSubnetId: aws_subnet.test1.id,
 * });
 * ```
 *
 * ## Import
 *
 * FSx File Systems can be imported using the `id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:fsx/ontapFileSystem:OntapFileSystem example fs-543ab12b1ca672f33
 * ```
 *
 *  Certain resource arguments, like `security_group_ids`, do not have a FSx API method for reading the information after creation. If the argument is set in the the provider configuration on an imported resource, the provider will always show a difference. To workaround this behavior, either omit the argument from the provider configuration or use `ignore_changes` to hide the difference, e.g., terraform resource "aws_fsx_ontap_file_system" "example" {
 *
 * # ... other configuration ...
 *
 *  security_group_ids = [aws_security_group.example.id]
 *
 * # There is no FSx API for reading security_group_ids
 *
 *  lifecycle {
 *
 *  ignore_changes = [security_group_ids]
 *
 *  } }
 */
export class OntapFileSystem extends pulumi.CustomResource {
    /**
     * Get an existing OntapFileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OntapFileSystemState, opts?: pulumi.CustomResourceOptions): OntapFileSystem {
        return new OntapFileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:fsx/ontapFileSystem:OntapFileSystem';

    /**
     * Returns true if the given object is an instance of OntapFileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OntapFileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OntapFileSystem.__pulumiType;
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
     * A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
     */
    public readonly dailyAutomaticBackupStartTime!: pulumi.Output<string>;
    /**
     * The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
     */
    public readonly deploymentType!: pulumi.Output<string>;
    /**
     * The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
     */
    public readonly diskIopsConfiguration!: pulumi.Output<outputs.fsx.OntapFileSystemDiskIopsConfiguration>;
    /**
     * The Domain Name Service (DNS) name for the file system. You can mount your file system using its DNS name.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
     */
    public readonly endpointIpAddressRange!: pulumi.Output<string>;
    /**
     * The endpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.fsx.OntapFileSystemEndpoint[]>;
    /**
     * The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
     */
    public readonly fsxAdminPassword!: pulumi.Output<string | undefined>;
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
     * The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
     */
    public readonly preferredSubnetId!: pulumi.Output<string>;
    /**
     * Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
     */
    public readonly routeTableIds!: pulumi.Output<string[]>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
     */
    public readonly storageCapacity!: pulumi.Output<number | undefined>;
    /**
     * The filesystem storage type. defaults to `SSD`.
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, `2048`, and `4096`.
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
     * Create a OntapFileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OntapFileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OntapFileSystemArgs | OntapFileSystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OntapFileSystemState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["automaticBackupRetentionDays"] = state ? state.automaticBackupRetentionDays : undefined;
            resourceInputs["dailyAutomaticBackupStartTime"] = state ? state.dailyAutomaticBackupStartTime : undefined;
            resourceInputs["deploymentType"] = state ? state.deploymentType : undefined;
            resourceInputs["diskIopsConfiguration"] = state ? state.diskIopsConfiguration : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["endpointIpAddressRange"] = state ? state.endpointIpAddressRange : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["fsxAdminPassword"] = state ? state.fsxAdminPassword : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["networkInterfaceIds"] = state ? state.networkInterfaceIds : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["preferredSubnetId"] = state ? state.preferredSubnetId : undefined;
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
            const args = argsOrState as OntapFileSystemArgs | undefined;
            if ((!args || args.deploymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentType'");
            }
            if ((!args || args.preferredSubnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'preferredSubnetId'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if ((!args || args.throughputCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'throughputCapacity'");
            }
            resourceInputs["automaticBackupRetentionDays"] = args ? args.automaticBackupRetentionDays : undefined;
            resourceInputs["dailyAutomaticBackupStartTime"] = args ? args.dailyAutomaticBackupStartTime : undefined;
            resourceInputs["deploymentType"] = args ? args.deploymentType : undefined;
            resourceInputs["diskIopsConfiguration"] = args ? args.diskIopsConfiguration : undefined;
            resourceInputs["endpointIpAddressRange"] = args ? args.endpointIpAddressRange : undefined;
            resourceInputs["fsxAdminPassword"] = args?.fsxAdminPassword ? pulumi.secret(args.fsxAdminPassword) : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["preferredSubnetId"] = args ? args.preferredSubnetId : undefined;
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
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["networkInterfaceIds"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["fsxAdminPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(OntapFileSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OntapFileSystem resources.
 */
export interface OntapFileSystemState {
    /**
     * Amazon Resource Name of the file system.
     */
    arn?: pulumi.Input<string>;
    /**
     * The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
     */
    automaticBackupRetentionDays?: pulumi.Input<number>;
    /**
     * A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
     */
    dailyAutomaticBackupStartTime?: pulumi.Input<string>;
    /**
     * The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
     */
    deploymentType?: pulumi.Input<string>;
    /**
     * The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
     */
    diskIopsConfiguration?: pulumi.Input<inputs.fsx.OntapFileSystemDiskIopsConfiguration>;
    /**
     * The Domain Name Service (DNS) name for the file system. You can mount your file system using its DNS name.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
     */
    endpointIpAddressRange?: pulumi.Input<string>;
    /**
     * The endpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.fsx.OntapFileSystemEndpoint>[]>;
    /**
     * The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
     */
    fsxAdminPassword?: pulumi.Input<string>;
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
     * The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
     */
    preferredSubnetId?: pulumi.Input<string>;
    /**
     * Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
     */
    routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
     */
    storageCapacity?: pulumi.Input<number>;
    /**
     * The filesystem storage type. defaults to `SSD`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, `2048`, and `4096`.
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
 * The set of arguments for constructing a OntapFileSystem resource.
 */
export interface OntapFileSystemArgs {
    /**
     * The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
     */
    automaticBackupRetentionDays?: pulumi.Input<number>;
    /**
     * A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
     */
    dailyAutomaticBackupStartTime?: pulumi.Input<string>;
    /**
     * The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
     */
    deploymentType: pulumi.Input<string>;
    /**
     * The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
     */
    diskIopsConfiguration?: pulumi.Input<inputs.fsx.OntapFileSystemDiskIopsConfiguration>;
    /**
     * Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
     */
    endpointIpAddressRange?: pulumi.Input<string>;
    /**
     * The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
     */
    fsxAdminPassword?: pulumi.Input<string>;
    /**
     * ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
     */
    preferredSubnetId: pulumi.Input<string>;
    /**
     * Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
     */
    routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
     */
    storageCapacity?: pulumi.Input<number>;
    /**
     * The filesystem storage type. defaults to `SSD`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, `2048`, and `4096`.
     */
    throughputCapacity: pulumi.Input<number>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    weeklyMaintenanceStartTime?: pulumi.Input<string>;
}
