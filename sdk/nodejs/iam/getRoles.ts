// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ARNs and Names of IAM Roles.
 *
 * ## Example Usage
 * ### All roles in an account
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const roles = pulumi.output(aws.iam.getRoles());
 * ```
 * ### Roles filtered by name regex
 *
 * Roles whose role-name contains `project`
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const roles = pulumi.output(aws.iam.getRoles({
 *     nameRegex: ".*project.*",
 * }));
 * ```
 * ### Roles filtered by path prefix
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const roles = pulumi.output(aws.iam.getRoles({
 *     pathPrefix: "/custom-path",
 * }));
 * ```
 * ### Roles provisioned by AWS SSO
 *
 * Roles in the account filtered by path prefix
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const roles = pulumi.output(aws.iam.getRoles({
 *     pathPrefix: "/aws-reserved/sso.amazonaws.com/",
 * }));
 * ```
 *
 * Specific role in the account filtered by name regex and path prefix
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const roles = pulumi.output(aws.iam.getRoles({
 *     nameRegex: "AWSReservedSSO_permission_set_name_.*",
 *     pathPrefix: "/aws-reserved/sso.amazonaws.com/",
 * }));
 * ```
 */
export function getRoles(args?: GetRolesArgs, opts?: pulumi.InvokeOptions): Promise<GetRolesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getRoles:getRoles", {
        "nameRegex": args.nameRegex,
        "pathPrefix": args.pathPrefix,
    }, opts);
}

/**
 * A collection of arguments for invoking getRoles.
 */
export interface GetRolesArgs {
    /**
     * A regex string to apply to the IAM roles list returned by AWS. This allows more advanced filtering not supported from the AWS API.
     * This filtering is done locally on what AWS returns, and could have a performance impact if the result is large. It is recommended to combine this with other
     * options to narrow down the list AWS returns.
     */
    nameRegex?: string;
    /**
     * The path prefix for filtering the results. For example, the prefix `/application_abc/component_xyz/` gets all roles whose path starts with `/application_abc/component_xyz/`. If it is not included, it defaults to a slash (`/`), listing all roles. For more details, check out [list-roles in the AWS CLI reference][1].
     */
    pathPrefix?: string;
}

/**
 * A collection of values returned by getRoles.
 */
export interface GetRolesResult {
    /**
     * Set of ARNs of the matched IAM roles.
     */
    readonly arns: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    /**
     * Set of Names of the matched IAM roles.
     */
    readonly names: string[];
    readonly pathPrefix?: string;
}
