// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for managing an AWS API Gateway V2 VPC Link.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.apigatewayv2.getVpcLink({
 *     vpcLinkId: "example",
 * });
 * ```
 */
export function getVpcLink(args: GetVpcLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcLinkResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:apigatewayv2/getVpcLink:getVpcLink", {
        "tags": args.tags,
        "vpcLinkId": args.vpcLinkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcLink.
 */
export interface GetVpcLinkArgs {
    /**
     * VPC Link Tags.
     */
    tags?: {[key: string]: string};
    /**
     * VPC Link ID
     */
    vpcLinkId: string;
}

/**
 * A collection of values returned by getVpcLink.
 */
export interface GetVpcLinkResult {
    /**
     * ARN of the VPC Link.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * VPC Link Name.
     */
    readonly name: string;
    /**
     * List of security groups associated with the VPC Link.
     */
    readonly securityGroupIds: string[];
    /**
     * List of subnets attached to the VPC Link.
     */
    readonly subnetIds: string[];
    /**
     * VPC Link Tags.
     */
    readonly tags?: {[key: string]: string};
    readonly vpcLinkId: string;
}
/**
 * Data source for managing an AWS API Gateway V2 VPC Link.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.apigatewayv2.getVpcLink({
 *     vpcLinkId: "example",
 * });
 * ```
 */
export function getVpcLinkOutput(args: GetVpcLinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcLinkResult> {
    return pulumi.output(args).apply((a: any) => getVpcLink(a, opts))
}

/**
 * A collection of arguments for invoking getVpcLink.
 */
export interface GetVpcLinkOutputArgs {
    /**
     * VPC Link Tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * VPC Link ID
     */
    vpcLinkId: pulumi.Input<string>;
}
