// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Get a list the AWS services for which the specified account is a delegated administrator
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.organizations.getDelegatedServices({
 *     accountId: "AWS ACCOUNT ID",
 * }));
 * ```
 */
export function getDelegatedServices(args: GetDelegatedServicesArgs, opts?: pulumi.InvokeOptions): Promise<GetDelegatedServicesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:organizations/getDelegatedServices:getDelegatedServices", {
        "accountId": args.accountId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDelegatedServices.
 */
export interface GetDelegatedServicesArgs {
    /**
     * The account ID number of a delegated administrator account in the organization.
     */
    accountId: string;
}

/**
 * A collection of values returned by getDelegatedServices.
 */
export interface GetDelegatedServicesResult {
    readonly accountId: string;
    /**
     * The services for which the account is a delegated administrator, which have the following attributes:
     */
    readonly delegatedServices: outputs.organizations.GetDelegatedServicesDelegatedService[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
