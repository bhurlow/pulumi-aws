// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a single EBS volume.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ebs.Volume("example", {
 *     availabilityZone: "us-west-2a",
 *     size: 40,
 *     tags: {
 *         Name: "HelloWorld",
 *     },
 * });
 * ```
 *
 * > **NOTE:** At least one of `size` or `snapshotId` is required when specifying an EBS volume
 *
 * ## Import
 *
 * EBS Volumes can be imported using the `id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:ebs/volume:Volume id vol-049df61146c4d7901
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ebs/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * The volume ARN (e.g., arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AZ where the EBS volume will exist.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * If true, the disk will be encrypted.
     */
    public readonly encrypted!: pulumi.Output<boolean>;
    /**
     * If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
     */
    public readonly finalSnapshot!: pulumi.Output<boolean | undefined>;
    /**
     * The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
     */
    public readonly iops!: pulumi.Output<number>;
    /**
     * The ARN for the KMS encryption key. When specifying `kmsKeyId`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
     */
    public readonly multiAttachEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    public readonly outpostArn!: pulumi.Output<string | undefined>;
    /**
     * The size of the drive in GiBs.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * A snapshot to base the EBS volume off of.
     */
    public readonly snapshotId!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
     */
    public readonly throughput!: pulumi.Output<number>;
    /**
     * The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["encrypted"] = state ? state.encrypted : undefined;
            resourceInputs["finalSnapshot"] = state ? state.finalSnapshot : undefined;
            resourceInputs["iops"] = state ? state.iops : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["multiAttachEnabled"] = state ? state.multiAttachEnabled : undefined;
            resourceInputs["outpostArn"] = state ? state.outpostArn : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["throughput"] = state ? state.throughput : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["encrypted"] = args ? args.encrypted : undefined;
            resourceInputs["finalSnapshot"] = args ? args.finalSnapshot : undefined;
            resourceInputs["iops"] = args ? args.iops : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["multiAttachEnabled"] = args ? args.multiAttachEnabled : undefined;
            resourceInputs["outpostArn"] = args ? args.outpostArn : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["throughput"] = args ? args.throughput : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * The volume ARN (e.g., arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
     */
    arn?: pulumi.Input<string>;
    /**
     * The AZ where the EBS volume will exist.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted.
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
     */
    finalSnapshot?: pulumi.Input<boolean>;
    /**
     * The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
     */
    iops?: pulumi.Input<number>;
    /**
     * The ARN for the KMS encryption key. When specifying `kmsKeyId`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
     */
    multiAttachEnabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    outpostArn?: pulumi.Input<string>;
    /**
     * The size of the drive in GiBs.
     */
    size?: pulumi.Input<number>;
    /**
     * A snapshot to base the EBS volume off of.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
     */
    throughput?: pulumi.Input<number>;
    /**
     * The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The AZ where the EBS volume will exist.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted.
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
     */
    finalSnapshot?: pulumi.Input<boolean>;
    /**
     * The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
     */
    iops?: pulumi.Input<number>;
    /**
     * The ARN for the KMS encryption key. When specifying `kmsKeyId`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
     */
    multiAttachEnabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    outpostArn?: pulumi.Input<string>;
    /**
     * The size of the drive in GiBs.
     */
    size?: pulumi.Input<number>;
    /**
     * A snapshot to base the EBS volume off of.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
     */
    throughput?: pulumi.Input<number>;
    /**
     * The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
     */
    type?: pulumi.Input<string>;
}
