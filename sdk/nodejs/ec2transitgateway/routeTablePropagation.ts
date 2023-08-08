// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EC2 Transit Gateway Route Table propagation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2transitgateway.RouteTablePropagation("example", {
 *     transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment.example.id,
 *     transitGatewayRouteTableId: aws_ec2_transit_gateway_route_table.example.id,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ec2_transit_gateway_route_table_propagation.example
 *
 *  id = "tgw-rtb-12345678_tgw-attach-87654321" } Using `pulumi import`, import `aws_ec2_transit_gateway_route_table_propagation` using the EC2 Transit Gateway Route Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier. For exampleconsole % pulumi import aws_ec2_transit_gateway_route_table_propagation.example tgw-rtb-12345678_tgw-attach-87654321
 */
export class RouteTablePropagation extends pulumi.CustomResource {
    /**
     * Get an existing RouteTablePropagation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTablePropagationState, opts?: pulumi.CustomResourceOptions): RouteTablePropagation {
        return new RouteTablePropagation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2transitgateway/routeTablePropagation:RouteTablePropagation';

    /**
     * Returns true if the given object is an instance of RouteTablePropagation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTablePropagation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTablePropagation.__pulumiType;
    }

    /**
     * Identifier of the resource
     */
    public /*out*/ readonly resourceId!: pulumi.Output<string>;
    /**
     * Type of the resource
     */
    public /*out*/ readonly resourceType!: pulumi.Output<string>;
    /**
     * Identifier of EC2 Transit Gateway Attachment.
     */
    public readonly transitGatewayAttachmentId!: pulumi.Output<string>;
    /**
     * Identifier of EC2 Transit Gateway Route Table.
     */
    public readonly transitGatewayRouteTableId!: pulumi.Output<string>;

    /**
     * Create a RouteTablePropagation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTablePropagationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTablePropagationArgs | RouteTablePropagationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteTablePropagationState | undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["transitGatewayAttachmentId"] = state ? state.transitGatewayAttachmentId : undefined;
            resourceInputs["transitGatewayRouteTableId"] = state ? state.transitGatewayRouteTableId : undefined;
        } else {
            const args = argsOrState as RouteTablePropagationArgs | undefined;
            if ((!args || args.transitGatewayAttachmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayAttachmentId'");
            }
            if ((!args || args.transitGatewayRouteTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayRouteTableId'");
            }
            resourceInputs["transitGatewayAttachmentId"] = args ? args.transitGatewayAttachmentId : undefined;
            resourceInputs["transitGatewayRouteTableId"] = args ? args.transitGatewayRouteTableId : undefined;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouteTablePropagation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTablePropagation resources.
 */
export interface RouteTablePropagationState {
    /**
     * Identifier of the resource
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Type of the resource
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway Attachment.
     */
    transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway Route Table.
     */
    transitGatewayRouteTableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTablePropagation resource.
 */
export interface RouteTablePropagationArgs {
    /**
     * Identifier of EC2 Transit Gateway Attachment.
     */
    transitGatewayAttachmentId: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway Route Table.
     */
    transitGatewayRouteTableId: pulumi.Input<string>;
}
