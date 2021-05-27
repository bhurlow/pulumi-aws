// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an existing backup vault.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.backup.getVault({
 *     name: "example_backup_vault",
 * }));
 * ```
 */
export function getVault(args: GetVaultArgs, opts?: pulumi.InvokeOptions): Promise<GetVaultResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:backup/getVault:getVault", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVault.
 */
export interface GetVaultArgs {
    /**
     * The name of the backup vault.
     */
    name: string;
    /**
     * Metadata that you can assign to help organize the resources that you create.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVault.
 */
export interface GetVaultResult {
    /**
     * The ARN of the vault.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The server-side encryption key that is used to protect your backups.
     */
    readonly kmsKeyArn: string;
    readonly name: string;
    /**
     * The number of recovery points that are stored in a backup vault.
     */
    readonly recoveryPoints: number;
    /**
     * Metadata that you can assign to help organize the resources that you create.
     */
    readonly tags: {[key: string]: string};
}
