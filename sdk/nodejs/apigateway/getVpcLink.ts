// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the id of a VPC Link in
 * API Gateway. To fetch the VPC Link you must provide a name to match against.
 * As there is no unique name constraint on API Gateway VPC Links this data source will
 * error if there is more than one match.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myApiGatewayVpcLink = pulumi.output(aws.apigateway.getVpcLink({
 *     name: "my-vpc-link",
 * }));
 * ```
 */
export function getVpcLink(args: GetVpcLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcLinkResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:apigateway/getVpcLink:getVpcLink", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcLink.
 */
export interface GetVpcLinkArgs {
    /**
     * The name of the API Gateway VPC Link to look up. If no API Gateway VPC Link is found with this name, an error will be returned.
     * If multiple API Gateway VPC Links are found with this name, an error will be returned.
     */
    name: string;
    /**
     * Key-value map of resource tags
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVpcLink.
 */
export interface GetVpcLinkResult {
    /**
     * The description of the VPC link.
     */
    readonly description: string;
    /**
     * Set to the ID of the found API Gateway VPC Link.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The status of the VPC link.
     */
    readonly status: string;
    /**
     * The status message of the VPC link.
     */
    readonly statusMessage: string;
    /**
     * Key-value map of resource tags
     */
    readonly tags: {[key: string]: string};
    /**
     * The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
     */
    readonly targetArns: string[];
}
