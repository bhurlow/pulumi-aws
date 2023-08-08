// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for managing an AWS OpenSearch Serverless VPC Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.opensearch.getServerlessVpcEndpoint({
 *     vpcEndpointId: "vpce-829a4487959e2a839",
 * });
 * ```
 */
export function getServerlessVpcEndpoint(args: GetServerlessVpcEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetServerlessVpcEndpointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:opensearch/getServerlessVpcEndpoint:getServerlessVpcEndpoint", {
        "vpcEndpointId": args.vpcEndpointId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerlessVpcEndpoint.
 */
export interface GetServerlessVpcEndpointArgs {
    /**
     * The unique identifier of the endpoint.
     */
    vpcEndpointId: string;
}

/**
 * A collection of values returned by getServerlessVpcEndpoint.
 */
export interface GetServerlessVpcEndpointResult {
    /**
     * The date the endpoint was created.
     */
    readonly createdDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the endpoint.
     */
    readonly name: string;
    /**
     * The IDs of the security groups that define the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
     */
    readonly securityGroupIds: string[];
    /**
     * The IDs of the subnets from which you access OpenSearch Serverless.
     */
    readonly subnetIds: string[];
    readonly vpcEndpointId: string;
    /**
     * The ID of the VPC from which you access OpenSearch Serverless.
     */
    readonly vpcId: string;
}
/**
 * Data source for managing an AWS OpenSearch Serverless VPC Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.opensearch.getServerlessVpcEndpoint({
 *     vpcEndpointId: "vpce-829a4487959e2a839",
 * });
 * ```
 */
export function getServerlessVpcEndpointOutput(args: GetServerlessVpcEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServerlessVpcEndpointResult> {
    return pulumi.output(args).apply((a: any) => getServerlessVpcEndpoint(a, opts))
}

/**
 * A collection of arguments for invoking getServerlessVpcEndpoint.
 */
export interface GetServerlessVpcEndpointOutputArgs {
    /**
     * The unique identifier of the endpoint.
     */
    vpcEndpointId: pulumi.Input<string>;
}
