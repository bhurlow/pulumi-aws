// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The VPC Endpoint data source provides details about
 * a specific VPC endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const s3 = aws.ec2.getVpcEndpoint({
 *     vpcId: aws_vpc.foo.id,
 *     serviceName: "com.amazonaws.us-west-2.s3",
 * });
 * const privateS3 = new aws.ec2.VpcEndpointRouteTableAssociation("privateS3", {
 *     vpcEndpointId: s3.then(s3 => s3.id),
 *     routeTableId: aws_route_table["private"].id,
 * });
 * ```
 */
export function getVpcEndpoint(args?: GetVpcEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcEndpointResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:ec2/getVpcEndpoint:getVpcEndpoint", {
        "filters": args.filters,
        "id": args.id,
        "serviceName": args.serviceName,
        "state": args.state,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcEndpoint.
 */
export interface GetVpcEndpointArgs {
    /**
     * Custom filter block as described below.
     */
    filters?: inputs.ec2.GetVpcEndpointFilter[];
    /**
     * ID of the specific VPC Endpoint to retrieve.
     */
    id?: string;
    /**
     * Service name of the specific VPC Endpoint to retrieve. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
     */
    serviceName?: string;
    /**
     * State of the specific VPC Endpoint to retrieve.
     */
    state?: string;
    /**
     * Map of tags, each pair of which must exactly match
     * a pair on the specific VPC Endpoint to retrieve.
     */
    tags?: {[key: string]: string};
    /**
     * ID of the VPC in which the specific VPC Endpoint is used.
     *
     * More complex filters can be expressed using one or more `filter` sub-blocks,
     * which take the following arguments:
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getVpcEndpoint.
 */
export interface GetVpcEndpointResult {
    /**
     * ARN of the VPC endpoint.
     */
    readonly arn: string;
    /**
     * List of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
     */
    readonly cidrBlocks: string[];
    /**
     * DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS entry blocks are documented below.
     */
    readonly dnsEntries: outputs.ec2.GetVpcEndpointDnsEntry[];
    /**
     * DNS options for the VPC Endpoint. DNS options blocks are documented below.
     */
    readonly dnsOptions: outputs.ec2.GetVpcEndpointDnsOption[];
    readonly filters?: outputs.ec2.GetVpcEndpointFilter[];
    readonly id: string;
    readonly ipAddressType: string;
    /**
     * One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
     */
    readonly networkInterfaceIds: string[];
    /**
     * ID of the AWS account that owns the VPC endpoint.
     */
    readonly ownerId: string;
    /**
     * Policy document associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
     */
    readonly policy: string;
    /**
     * Prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
     */
    readonly prefixListId: string;
    /**
     * Whether or not the VPC is associated with a private hosted zone - `true` or `false`. Applicable for endpoints of type `Interface`.
     */
    readonly privateDnsEnabled: boolean;
    /**
     * Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
     */
    readonly requesterManaged: boolean;
    /**
     * One or more route tables associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
     */
    readonly routeTableIds: string[];
    /**
     * One or more security groups associated with the network interfaces. Applicable for endpoints of type `Interface`.
     */
    readonly securityGroupIds: string[];
    readonly serviceName: string;
    readonly state: string;
    /**
     * One or more subnets in which the VPC Endpoint is located. Applicable for endpoints of type `Interface`.
     */
    readonly subnetIds: string[];
    readonly tags: {[key: string]: string};
    /**
     * VPC Endpoint type, `Gateway` or `Interface`.
     */
    readonly vpcEndpointType: string;
    readonly vpcId: string;
}
/**
 * The VPC Endpoint data source provides details about
 * a specific VPC endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const s3 = aws.ec2.getVpcEndpoint({
 *     vpcId: aws_vpc.foo.id,
 *     serviceName: "com.amazonaws.us-west-2.s3",
 * });
 * const privateS3 = new aws.ec2.VpcEndpointRouteTableAssociation("privateS3", {
 *     vpcEndpointId: s3.then(s3 => s3.id),
 *     routeTableId: aws_route_table["private"].id,
 * });
 * ```
 */
export function getVpcEndpointOutput(args?: GetVpcEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcEndpointResult> {
    return pulumi.output(args).apply((a: any) => getVpcEndpoint(a, opts))
}

/**
 * A collection of arguments for invoking getVpcEndpoint.
 */
export interface GetVpcEndpointOutputArgs {
    /**
     * Custom filter block as described below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.ec2.GetVpcEndpointFilterArgs>[]>;
    /**
     * ID of the specific VPC Endpoint to retrieve.
     */
    id?: pulumi.Input<string>;
    /**
     * Service name of the specific VPC Endpoint to retrieve. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
     */
    serviceName?: pulumi.Input<string>;
    /**
     * State of the specific VPC Endpoint to retrieve.
     */
    state?: pulumi.Input<string>;
    /**
     * Map of tags, each pair of which must exactly match
     * a pair on the specific VPC Endpoint to retrieve.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ID of the VPC in which the specific VPC Endpoint is used.
     *
     * More complex filters can be expressed using one or more `filter` sub-blocks,
     * which take the following arguments:
     */
    vpcId?: pulumi.Input<string>;
}
