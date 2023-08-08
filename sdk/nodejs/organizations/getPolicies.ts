// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for managing an AWS Organizations Policies.
 *
 * ## Example Usage
 */
export function getPolicies(args: GetPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetPoliciesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:organizations/getPolicies:getPolicies", {
        "filter": args.filter,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicies.
 */
export interface GetPoliciesArgs {
    /**
     * The type of policies to be returned in the response. Valid values are `SERVICE_CONTROL_POLICY | TAG_POLICY | BACKUP_POLICY | AISERVICES_OPT_OUT_POLICY`
     */
    filter: string;
}

/**
 * A collection of values returned by getPolicies.
 */
export interface GetPoliciesResult {
    readonly filter: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of all the policy ids found.
     */
    readonly ids: string[];
}
/**
 * Data source for managing an AWS Organizations Policies.
 *
 * ## Example Usage
 */
export function getPoliciesOutput(args: GetPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getPolicies.
 */
export interface GetPoliciesOutputArgs {
    /**
     * The type of policies to be returned in the response. Valid values are `SERVICE_CONTROL_POLICY | TAG_POLICY | BACKUP_POLICY | AISERVICES_OPT_OUT_POLICY`
     */
    filter: pulumi.Input<string>;
}
