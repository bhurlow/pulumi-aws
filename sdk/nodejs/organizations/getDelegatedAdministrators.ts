// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Get a list the AWS accounts that are designated as delegated administrators in this organization
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.organizations.getDelegatedAdministrators({
 *     servicePrincipal: "SERVICE PRINCIPAL",
 * }));
 * ```
 */
export function getDelegatedAdministrators(args?: GetDelegatedAdministratorsArgs, opts?: pulumi.InvokeOptions): Promise<GetDelegatedAdministratorsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:organizations/getDelegatedAdministrators:getDelegatedAdministrators", {
        "servicePrincipal": args.servicePrincipal,
    }, opts);
}

/**
 * A collection of arguments for invoking getDelegatedAdministrators.
 */
export interface GetDelegatedAdministratorsArgs {
    /**
     * Specifies a service principal name. If specified, then the operation lists the delegated administrators only for the specified service. If you don't specify a service principal, the operation lists all delegated administrators for all services in your organization.
     */
    servicePrincipal?: string;
}

/**
 * A collection of values returned by getDelegatedAdministrators.
 */
export interface GetDelegatedAdministratorsResult {
    /**
     * The list of delegated administrators in your organization, which have the following attributes:
     */
    readonly delegatedAdministrators: outputs.organizations.GetDelegatedAdministratorsDelegatedAdministrator[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly servicePrincipal?: string;
}
