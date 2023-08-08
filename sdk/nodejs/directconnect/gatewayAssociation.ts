// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates a Direct Connect Gateway with a VGW or transit gateway.
 *
 * To create a cross-account association, create an `aws.directconnect.GatewayAssociationProposal` resource
 * in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
 * by creating an `aws.directconnect.GatewayAssociation` resource with the `proposalId` and `associatedGatewayOwnerAccountId` attributes set.
 *
 * ## Example Usage
 * ### VPN Gateway Association
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
 * const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.255.255.0/28"});
 * const exampleVpnGateway = new aws.ec2.VpnGateway("exampleVpnGateway", {vpcId: exampleVpc.id});
 * const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("exampleGatewayAssociation", {
 *     dxGatewayId: exampleGateway.id,
 *     associatedGatewayId: exampleVpnGateway.id,
 * });
 * ```
 * ### Transit Gateway Association
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
 * const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {});
 * const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("exampleGatewayAssociation", {
 *     dxGatewayId: exampleGateway.id,
 *     associatedGatewayId: exampleTransitGateway.id,
 *     allowedPrefixes: [
 *         "10.255.255.0/30",
 *         "10.255.255.8/30",
 *     ],
 * });
 * ```
 * ### Allowed Prefixes
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
 * const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.255.255.0/28"});
 * const exampleVpnGateway = new aws.ec2.VpnGateway("exampleVpnGateway", {vpcId: exampleVpc.id});
 * const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("exampleGatewayAssociation", {
 *     dxGatewayId: exampleGateway.id,
 *     associatedGatewayId: exampleVpnGateway.id,
 *     allowedPrefixes: [
 *         "210.52.109.0/24",
 *         "175.45.176.0/22",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_dx_gateway_association.example
 *
 *  id = "345508c3-7215-4aef-9832-07c125d5bd0f/vgw-98765432" } Using `pulumi import`, import Direct Connect gateway associations using `dx_gateway_id` together with `associated_gateway_id`. For exampleconsole % pulumi import aws_dx_gateway_association.example 345508c3-7215-4aef-9832-07c125d5bd0f/vgw-98765432
 */
export class GatewayAssociation extends pulumi.CustomResource {
    /**
     * Get an existing GatewayAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayAssociationState, opts?: pulumi.CustomResourceOptions): GatewayAssociation {
        return new GatewayAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/gatewayAssociation:GatewayAssociation';

    /**
     * Returns true if the given object is an instance of GatewayAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayAssociation.__pulumiType;
    }

    /**
     * VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     */
    public readonly allowedPrefixes!: pulumi.Output<string[]>;
    /**
     * The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for single account Direct Connect gateway associations.
     */
    public readonly associatedGatewayId!: pulumi.Output<string>;
    /**
     * The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for cross-account Direct Connect gateway associations.
     */
    public readonly associatedGatewayOwnerAccountId!: pulumi.Output<string>;
    /**
     * The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
     */
    public /*out*/ readonly associatedGatewayType!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect gateway association.
     */
    public /*out*/ readonly dxGatewayAssociationId!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect gateway.
     */
    public readonly dxGatewayId!: pulumi.Output<string>;
    /**
     * The ID of the AWS account that owns the Direct Connect gateway.
     */
    public /*out*/ readonly dxGatewayOwnerAccountId!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect gateway association proposal.
     * Used for cross-account Direct Connect gateway associations.
     */
    public readonly proposalId!: pulumi.Output<string | undefined>;
    /**
     * @deprecated use 'associated_gateway_id' argument instead
     */
    public readonly vpnGatewayId!: pulumi.Output<string | undefined>;

    /**
     * Create a GatewayAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayAssociationArgs | GatewayAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayAssociationState | undefined;
            resourceInputs["allowedPrefixes"] = state ? state.allowedPrefixes : undefined;
            resourceInputs["associatedGatewayId"] = state ? state.associatedGatewayId : undefined;
            resourceInputs["associatedGatewayOwnerAccountId"] = state ? state.associatedGatewayOwnerAccountId : undefined;
            resourceInputs["associatedGatewayType"] = state ? state.associatedGatewayType : undefined;
            resourceInputs["dxGatewayAssociationId"] = state ? state.dxGatewayAssociationId : undefined;
            resourceInputs["dxGatewayId"] = state ? state.dxGatewayId : undefined;
            resourceInputs["dxGatewayOwnerAccountId"] = state ? state.dxGatewayOwnerAccountId : undefined;
            resourceInputs["proposalId"] = state ? state.proposalId : undefined;
            resourceInputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
        } else {
            const args = argsOrState as GatewayAssociationArgs | undefined;
            if ((!args || args.dxGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dxGatewayId'");
            }
            resourceInputs["allowedPrefixes"] = args ? args.allowedPrefixes : undefined;
            resourceInputs["associatedGatewayId"] = args ? args.associatedGatewayId : undefined;
            resourceInputs["associatedGatewayOwnerAccountId"] = args ? args.associatedGatewayOwnerAccountId : undefined;
            resourceInputs["dxGatewayId"] = args ? args.dxGatewayId : undefined;
            resourceInputs["proposalId"] = args ? args.proposalId : undefined;
            resourceInputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
            resourceInputs["associatedGatewayType"] = undefined /*out*/;
            resourceInputs["dxGatewayAssociationId"] = undefined /*out*/;
            resourceInputs["dxGatewayOwnerAccountId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GatewayAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayAssociation resources.
 */
export interface GatewayAssociationState {
    /**
     * VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     */
    allowedPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for single account Direct Connect gateway associations.
     */
    associatedGatewayId?: pulumi.Input<string>;
    /**
     * The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for cross-account Direct Connect gateway associations.
     */
    associatedGatewayOwnerAccountId?: pulumi.Input<string>;
    /**
     * The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
     */
    associatedGatewayType?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway association.
     */
    dxGatewayAssociationId?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway.
     */
    dxGatewayId?: pulumi.Input<string>;
    /**
     * The ID of the AWS account that owns the Direct Connect gateway.
     */
    dxGatewayOwnerAccountId?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway association proposal.
     * Used for cross-account Direct Connect gateway associations.
     */
    proposalId?: pulumi.Input<string>;
    /**
     * @deprecated use 'associated_gateway_id' argument instead
     */
    vpnGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GatewayAssociation resource.
 */
export interface GatewayAssociationArgs {
    /**
     * VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     */
    allowedPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for single account Direct Connect gateway associations.
     */
    associatedGatewayId?: pulumi.Input<string>;
    /**
     * The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for cross-account Direct Connect gateway associations.
     */
    associatedGatewayOwnerAccountId?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway.
     */
    dxGatewayId: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway association proposal.
     * Used for cross-account Direct Connect gateway associations.
     */
    proposalId?: pulumi.Input<string>;
    /**
     * @deprecated use 'associated_gateway_id' argument instead
     */
    vpnGatewayId?: pulumi.Input<string>;
}
