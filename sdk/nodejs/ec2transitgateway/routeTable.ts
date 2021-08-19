// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EC2 Transit Gateway Route Table.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2transitgateway.RouteTable("example", {transitGatewayId: aws_ec2_transit_gateway.example.id});
 * ```
 *
 * ## Import
 *
 * `aws_ec2_transit_gateway_route_table` can be imported by using the EC2 Transit Gateway Route Table identifier, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ec2transitgateway/routeTable:RouteTable example tgw-rtb-12345678
 * ```
 */
export class RouteTable extends pulumi.CustomResource {
    /**
     * Get an existing RouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableState, opts?: pulumi.CustomResourceOptions): RouteTable {
        return new RouteTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2transitgateway/routeTable:RouteTable';

    /**
     * Returns true if the given object is an instance of RouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTable.__pulumiType;
    }

    /**
     * EC2 Transit Gateway Route Table Amazon Resource Name (ARN).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Boolean whether this is the default association route table for the EC2 Transit Gateway.
     */
    public /*out*/ readonly defaultAssociationRouteTable!: pulumi.Output<boolean>;
    /**
     * Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
     */
    public /*out*/ readonly defaultPropagationRouteTable!: pulumi.Output<boolean>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    public readonly transitGatewayId!: pulumi.Output<string>;

    /**
     * Create a RouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableArgs | RouteTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteTableState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["defaultAssociationRouteTable"] = state ? state.defaultAssociationRouteTable : undefined;
            inputs["defaultPropagationRouteTable"] = state ? state.defaultPropagationRouteTable : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
        } else {
            const args = argsOrState as RouteTableArgs | undefined;
            if ((!args || args.transitGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["defaultAssociationRouteTable"] = undefined /*out*/;
            inputs["defaultPropagationRouteTable"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTable resources.
 */
export interface RouteTableState {
    /**
     * EC2 Transit Gateway Route Table Amazon Resource Name (ARN).
     */
    arn?: pulumi.Input<string>;
    /**
     * Boolean whether this is the default association route table for the EC2 Transit Gateway.
     */
    defaultAssociationRouteTable?: pulumi.Input<boolean>;
    /**
     * Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
     */
    defaultPropagationRouteTable?: pulumi.Input<boolean>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    transitGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTable resource.
 */
export interface RouteTableArgs {
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    transitGatewayId: pulumi.Input<string>;
}
