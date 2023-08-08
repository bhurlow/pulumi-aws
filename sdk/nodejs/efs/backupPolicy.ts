// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an Elastic File System (EFS) Backup Policy resource.
 * Backup policies turn automatic backups on or off for an existing file system.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const fs = new aws.efs.FileSystem("fs", {});
 * const policy = new aws.efs.BackupPolicy("policy", {
 *     fileSystemId: fs.id,
 *     backupPolicy: {
 *         status: "ENABLED",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_efs_backup_policy.example
 *
 *  id = "fs-6fa144c6" } Using `pulumi import`, import the EFS backup policies using the `id`. For exampleconsole % pulumi import aws_efs_backup_policy.example fs-6fa144c6
 */
export class BackupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupPolicyState, opts?: pulumi.CustomResourceOptions): BackupPolicy {
        return new BackupPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:efs/backupPolicy:BackupPolicy';

    /**
     * Returns true if the given object is an instance of BackupPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPolicy.__pulumiType;
    }

    /**
     * A backupPolicy object (documented below).
     */
    public readonly backupPolicy!: pulumi.Output<outputs.efs.BackupPolicyBackupPolicy>;
    /**
     * The ID of the EFS file system.
     */
    public readonly fileSystemId!: pulumi.Output<string>;

    /**
     * Create a BackupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupPolicyArgs | BackupPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupPolicyState | undefined;
            resourceInputs["backupPolicy"] = state ? state.backupPolicy : undefined;
            resourceInputs["fileSystemId"] = state ? state.fileSystemId : undefined;
        } else {
            const args = argsOrState as BackupPolicyArgs | undefined;
            if ((!args || args.backupPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupPolicy'");
            }
            if ((!args || args.fileSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemId'");
            }
            resourceInputs["backupPolicy"] = args ? args.backupPolicy : undefined;
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupPolicy resources.
 */
export interface BackupPolicyState {
    /**
     * A backupPolicy object (documented below).
     */
    backupPolicy?: pulumi.Input<inputs.efs.BackupPolicyBackupPolicy>;
    /**
     * The ID of the EFS file system.
     */
    fileSystemId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupPolicy resource.
 */
export interface BackupPolicyArgs {
    /**
     * A backupPolicy object (documented below).
     */
    backupPolicy: pulumi.Input<inputs.efs.BackupPolicyBackupPolicy>;
    /**
     * The ID of the EFS file system.
     */
    fileSystemId: pulumi.Input<string>;
}
