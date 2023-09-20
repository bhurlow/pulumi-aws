// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an Amazon FSx for OpenZFS volume.
 * See the [FSx OpenZFS User Guide](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/what-is-fsx.html) for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.fsx.OpenZfsVolume("test", {parentVolumeId: aws_fsx_openzfs_file_system.test.root_volume_id});
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import FSx Volumes using the `id`. For example:
 *
 * ```sh
 *  $ pulumi import aws:fsx/openZfsVolume:OpenZfsVolume example fsvol-543ab12b1ca672f33
 * ```
 */
export class OpenZfsVolume extends pulumi.CustomResource {
    /**
     * Get an existing OpenZfsVolume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OpenZfsVolumeState, opts?: pulumi.CustomResourceOptions): OpenZfsVolume {
        return new OpenZfsVolume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:fsx/openZfsVolume:OpenZfsVolume';

    /**
     * Returns true if the given object is an instance of OpenZfsVolume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenZfsVolume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenZfsVolume.__pulumiType;
    }

    /**
     * Amazon Resource Name of the file system.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
     */
    public readonly copyTagsToSnapshots!: pulumi.Output<boolean | undefined>;
    /**
     * Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
     */
    public readonly dataCompressionType!: pulumi.Output<string | undefined>;
    /**
     * The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
     */
    public readonly nfsExports!: pulumi.Output<outputs.fsx.OpenZfsVolumeNfsExports | undefined>;
    /**
     * The ARN of the source snapshot to create the volume from.
     */
    public readonly originSnapshot!: pulumi.Output<outputs.fsx.OpenZfsVolumeOriginSnapshot | undefined>;
    /**
     * The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `aws.fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `aws.fsx.OpenZfsVolume`.
     */
    public readonly parentVolumeId!: pulumi.Output<string>;
    /**
     * specifies whether the volume is read-only. Default is false.
     */
    public readonly readOnly!: pulumi.Output<boolean>;
    /**
     * The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
     */
    public readonly recordSizeKib!: pulumi.Output<number | undefined>;
    /**
     * The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
     */
    public readonly storageCapacityQuotaGib!: pulumi.Output<number>;
    /**
     * The amount of storage in gibibytes (GiB) to reserve from the parent volume.
     */
    public readonly storageCapacityReservationGib!: pulumi.Output<number>;
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
     * Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
     */
    public readonly userAndGroupQuotas!: pulumi.Output<outputs.fsx.OpenZfsVolumeUserAndGroupQuota[]>;
    public readonly volumeType!: pulumi.Output<string | undefined>;

    /**
     * Create a OpenZfsVolume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenZfsVolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OpenZfsVolumeArgs | OpenZfsVolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OpenZfsVolumeState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["copyTagsToSnapshots"] = state ? state.copyTagsToSnapshots : undefined;
            resourceInputs["dataCompressionType"] = state ? state.dataCompressionType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nfsExports"] = state ? state.nfsExports : undefined;
            resourceInputs["originSnapshot"] = state ? state.originSnapshot : undefined;
            resourceInputs["parentVolumeId"] = state ? state.parentVolumeId : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["recordSizeKib"] = state ? state.recordSizeKib : undefined;
            resourceInputs["storageCapacityQuotaGib"] = state ? state.storageCapacityQuotaGib : undefined;
            resourceInputs["storageCapacityReservationGib"] = state ? state.storageCapacityReservationGib : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["userAndGroupQuotas"] = state ? state.userAndGroupQuotas : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as OpenZfsVolumeArgs | undefined;
            if ((!args || args.parentVolumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parentVolumeId'");
            }
            resourceInputs["copyTagsToSnapshots"] = args ? args.copyTagsToSnapshots : undefined;
            resourceInputs["dataCompressionType"] = args ? args.dataCompressionType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nfsExports"] = args ? args.nfsExports : undefined;
            resourceInputs["originSnapshot"] = args ? args.originSnapshot : undefined;
            resourceInputs["parentVolumeId"] = args ? args.parentVolumeId : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["recordSizeKib"] = args ? args.recordSizeKib : undefined;
            resourceInputs["storageCapacityQuotaGib"] = args ? args.storageCapacityQuotaGib : undefined;
            resourceInputs["storageCapacityReservationGib"] = args ? args.storageCapacityReservationGib : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userAndGroupQuotas"] = args ? args.userAndGroupQuotas : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(OpenZfsVolume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OpenZfsVolume resources.
 */
export interface OpenZfsVolumeState {
    /**
     * Amazon Resource Name of the file system.
     */
    arn?: pulumi.Input<string>;
    /**
     * A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
     */
    copyTagsToSnapshots?: pulumi.Input<boolean>;
    /**
     * Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
     */
    dataCompressionType?: pulumi.Input<string>;
    /**
     * The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
     */
    name?: pulumi.Input<string>;
    /**
     * NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
     */
    nfsExports?: pulumi.Input<inputs.fsx.OpenZfsVolumeNfsExports>;
    /**
     * The ARN of the source snapshot to create the volume from.
     */
    originSnapshot?: pulumi.Input<inputs.fsx.OpenZfsVolumeOriginSnapshot>;
    /**
     * The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `aws.fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `aws.fsx.OpenZfsVolume`.
     */
    parentVolumeId?: pulumi.Input<string>;
    /**
     * specifies whether the volume is read-only. Default is false.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
     */
    recordSizeKib?: pulumi.Input<number>;
    /**
     * The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
     */
    storageCapacityQuotaGib?: pulumi.Input<number>;
    /**
     * The amount of storage in gibibytes (GiB) to reserve from the parent volume.
     */
    storageCapacityReservationGib?: pulumi.Input<number>;
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
     * Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
     */
    userAndGroupQuotas?: pulumi.Input<pulumi.Input<inputs.fsx.OpenZfsVolumeUserAndGroupQuota>[]>;
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OpenZfsVolume resource.
 */
export interface OpenZfsVolumeArgs {
    /**
     * A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
     */
    copyTagsToSnapshots?: pulumi.Input<boolean>;
    /**
     * Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
     */
    dataCompressionType?: pulumi.Input<string>;
    /**
     * The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
     */
    name?: pulumi.Input<string>;
    /**
     * NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
     */
    nfsExports?: pulumi.Input<inputs.fsx.OpenZfsVolumeNfsExports>;
    /**
     * The ARN of the source snapshot to create the volume from.
     */
    originSnapshot?: pulumi.Input<inputs.fsx.OpenZfsVolumeOriginSnapshot>;
    /**
     * The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `aws.fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `aws.fsx.OpenZfsVolume`.
     */
    parentVolumeId: pulumi.Input<string>;
    /**
     * specifies whether the volume is read-only. Default is false.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
     */
    recordSizeKib?: pulumi.Input<number>;
    /**
     * The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
     */
    storageCapacityQuotaGib?: pulumi.Input<number>;
    /**
     * The amount of storage in gibibytes (GiB) to reserve from the parent volume.
     */
    storageCapacityReservationGib?: pulumi.Input<number>;
    /**
     * A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
     */
    userAndGroupQuotas?: pulumi.Input<pulumi.Input<inputs.fsx.OpenZfsVolumeUserAndGroupQuota>[]>;
    volumeType?: pulumi.Input<string>;
}
