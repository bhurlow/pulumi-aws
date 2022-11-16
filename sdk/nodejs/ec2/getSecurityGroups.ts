// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to get IDs and VPC membership of Security Groups that are created outside this provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.ec2.getSecurityGroups({
 *     tags: {
 *         Application: "k8s",
 *         Environment: "dev",
 *     },
 * }));
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.ec2.getSecurityGroups({
 *     filters: [
 *         {
 *             name: "group-name",
 *             values: ["*nodes*"],
 *         },
 *         {
 *             name: "vpc-id",
 *             values: [_var.vpc_id],
 *         },
 *     ],
 * });
 * ```
 */
export function getSecurityGroups(args?: GetSecurityGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws:ec2/getSecurityGroups:getSecurityGroups", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroups.
 */
export interface GetSecurityGroupsArgs {
    /**
     * One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out [describe-security-groups in the AWS CLI reference][1].
     */
    filters?: inputs.ec2.GetSecurityGroupsFilter[];
    /**
     * Map of tags, each pair of which must exactly match for desired security groups.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getSecurityGroups.
 */
export interface GetSecurityGroupsResult {
    /**
     * ARNs of the matched security groups.
     */
    readonly arns: string[];
    readonly filters?: outputs.ec2.GetSecurityGroupsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IDs of the matches security groups.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: string};
    /**
     * VPC IDs of the matched security groups. The data source's tag or filter *will span VPCs* unless the `vpc-id` filter is also used.
     */
    readonly vpcIds: string[];
}

export function getSecurityGroupsOutput(args?: GetSecurityGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupsResult> {
    return pulumi.output(args).apply(a => getSecurityGroups(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityGroups.
 */
export interface GetSecurityGroupsOutputArgs {
    /**
     * One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out [describe-security-groups in the AWS CLI reference][1].
     */
    filters?: pulumi.Input<pulumi.Input<inputs.ec2.GetSecurityGroupsFilterArgs>[]>;
    /**
     * Map of tags, each pair of which must exactly match for desired security groups.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
