// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * `aws.ram.ResourceShare` Retrieve information about a RAM Resource Share.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ram.getResourceShare({
 *     name: "example",
 *     resourceOwner: "SELF",
 * }));
 * ```
 * ## Search by filters
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const tagFilter = pulumi.output(aws.ram.getResourceShare({
 *     filters: [{
 *         name: "NameOfTag",
 *         values: ["exampleNameTagValue"],
 *     }],
 *     name: "MyResourceName",
 *     resourceOwner: "SELF",
 * }));
 * ```
 */
export function getResourceShare(args: GetResourceShareArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceShareResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws:ram/getResourceShare:getResourceShare", {
        "filters": args.filters,
        "name": args.name,
        "resourceOwner": args.resourceOwner,
        "resourceShareStatus": args.resourceShareStatus,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceShare.
 */
export interface GetResourceShareArgs {
    /**
     * A filter used to scope the list e.g., by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
     */
    filters?: inputs.ram.GetResourceShareFilter[];
    /**
     * The name of the tag key to filter on.
     */
    name: string;
    /**
     * The owner of the resource share. Valid values are `SELF` or `OTHER-ACCOUNTS`.
     */
    resourceOwner: string;
    /**
     * Specifies that you want to retrieve details of only those resource shares that have this status. Valid values are `PENDING`, `ACTIVE`, `FAILED`, `DELETING`, and `DELETED`.
     */
    resourceShareStatus?: string;
    /**
     * The Tags attached to the RAM share
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getResourceShare.
 */
export interface GetResourceShareResult {
    /**
     * The Amazon Resource Name (ARN) of the resource share.
     */
    readonly arn: string;
    readonly filters?: outputs.ram.GetResourceShareFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The ID of the AWS account that owns the resource share.
     */
    readonly owningAccountId: string;
    readonly resourceOwner: string;
    readonly resourceShareStatus?: string;
    /**
     * The Status of the RAM share.
     */
    readonly status: string;
    /**
     * The Tags attached to the RAM share
     */
    readonly tags: {[key: string]: string};
}

export function getResourceShareOutput(args: GetResourceShareOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceShareResult> {
    return pulumi.output(args).apply(a => getResourceShare(a, opts))
}

/**
 * A collection of arguments for invoking getResourceShare.
 */
export interface GetResourceShareOutputArgs {
    /**
     * A filter used to scope the list e.g., by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
     */
    filters?: pulumi.Input<pulumi.Input<inputs.ram.GetResourceShareFilterArgs>[]>;
    /**
     * The name of the tag key to filter on.
     */
    name: pulumi.Input<string>;
    /**
     * The owner of the resource share. Valid values are `SELF` or `OTHER-ACCOUNTS`.
     */
    resourceOwner: pulumi.Input<string>;
    /**
     * Specifies that you want to retrieve details of only those resource shares that have this status. Valid values are `PENDING`, `ACTIVE`, `FAILED`, `DELETING`, and `DELETED`.
     */
    resourceShareStatus?: pulumi.Input<string>;
    /**
     * The Tags attached to the RAM share
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
