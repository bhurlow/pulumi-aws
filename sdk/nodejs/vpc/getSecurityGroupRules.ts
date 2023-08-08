// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This resource can be useful for getting back a set of security group rule IDs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.vpc.getSecurityGroupRules({
 *     filters: [{
 *         name: "group-id",
 *         values: [_var.security_group_id],
 *     }],
 * });
 * ```
 */
export function getSecurityGroupRules(args?: GetSecurityGroupRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupRulesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:vpc/getSecurityGroupRules:getSecurityGroupRules", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroupRules.
 */
export interface GetSecurityGroupRulesArgs {
    /**
     * Custom filter block as described below.
     */
    filters?: inputs.vpc.GetSecurityGroupRulesFilter[];
    /**
     * Map of tags, each pair of which must exactly match
     * a pair on the desired security group rule.
     *
     * More complex filters can be expressed using one or more `filter` sub-blocks,
     * which take the following arguments:
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getSecurityGroupRules.
 */
export interface GetSecurityGroupRulesResult {
    readonly filters?: outputs.vpc.GetSecurityGroupRulesFilter[];
    readonly id: string;
    /**
     * List of all the security group rule IDs found.
     */
    readonly ids: string[];
    readonly tags?: {[key: string]: string};
}
/**
 * This resource can be useful for getting back a set of security group rule IDs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.vpc.getSecurityGroupRules({
 *     filters: [{
 *         name: "group-id",
 *         values: [_var.security_group_id],
 *     }],
 * });
 * ```
 */
export function getSecurityGroupRulesOutput(args?: GetSecurityGroupRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupRulesResult> {
    return pulumi.output(args).apply((a: any) => getSecurityGroupRules(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityGroupRules.
 */
export interface GetSecurityGroupRulesOutputArgs {
    /**
     * Custom filter block as described below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.vpc.GetSecurityGroupRulesFilterArgs>[]>;
    /**
     * Map of tags, each pair of which must exactly match
     * a pair on the desired security group rule.
     *
     * More complex filters can be expressed using one or more `filter` sub-blocks,
     * which take the following arguments:
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
