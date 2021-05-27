// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an existing backup plan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.backup.getPlan({
 *     planId: "tf_example_backup_plan_id",
 * }));
 * ```
 */
export function getPlan(args: GetPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetPlanResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:backup/getPlan:getPlan", {
        "planId": args.planId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlan.
 */
export interface GetPlanArgs {
    /**
     * The backup plan ID.
     */
    planId: string;
    /**
     * Metadata that you can assign to help organize the plans you create.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getPlan.
 */
export interface GetPlanResult {
    /**
     * The ARN of the backup plan.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The display name of a backup plan.
     */
    readonly name: string;
    readonly planId: string;
    /**
     * Metadata that you can assign to help organize the plans you create.
     */
    readonly tags: {[key: string]: string};
    /**
     * Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
     */
    readonly version: string;
}
