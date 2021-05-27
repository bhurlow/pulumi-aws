// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Retrieves the summary of a WAFv2 IP Set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.wafv2.getIpSet({
 *     name: "some-ip-set",
 *     scope: "REGIONAL",
 * }));
 * ```
 */
export function getIpSet(args: GetIpSetArgs, opts?: pulumi.InvokeOptions): Promise<GetIpSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:wafv2/getIpSet:getIpSet", {
        "name": args.name,
        "scope": args.scope,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpSet.
 */
export interface GetIpSetArgs {
    /**
     * The name of the WAFv2 IP Set.
     */
    name: string;
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     */
    scope: string;
}

/**
 * A collection of values returned by getIpSet.
 */
export interface GetIpSetResult {
    /**
     * An array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation.
     */
    readonly addresses: string[];
    /**
     * The Amazon Resource Name (ARN) of the entity.
     */
    readonly arn: string;
    /**
     * The description of the set that helps with identification.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The IP address version of the set.
     */
    readonly ipAddressVersion: string;
    readonly name: string;
    readonly scope: string;
}
