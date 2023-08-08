// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides network associations for AWS Client VPN endpoints. For more information on usage, please see the
 * [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2clientvpn.NetworkAssociation("example", {
 *     clientVpnEndpointId: aws_ec2_client_vpn_endpoint.example.id,
 *     subnetId: aws_subnet.example.id,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ec2_client_vpn_network_association.example
 *
 *  id = "cvpn-endpoint-0ac3a1abbccddd666,vpn-assoc-0b8db902465d069ad" } Using `pulumi import`, import AWS Client VPN network associations using the endpoint ID and the association ID. Values are separated by a `,`. For exampleconsole % pulumi import aws_ec2_client_vpn_network_association.example cvpn-endpoint-0ac3a1abbccddd666,vpn-assoc-0b8db902465d069ad
 */
export class NetworkAssociation extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAssociationState, opts?: pulumi.CustomResourceOptions): NetworkAssociation {
        return new NetworkAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2clientvpn/networkAssociation:NetworkAssociation';

    /**
     * Returns true if the given object is an instance of NetworkAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAssociation.__pulumiType;
    }

    /**
     * The unique ID of the target network association.
     */
    public /*out*/ readonly associationId!: pulumi.Output<string>;
    /**
     * The ID of the Client VPN endpoint.
     */
    public readonly clientVpnEndpointId!: pulumi.Output<string>;
    /**
     * The ID of the subnet to associate with the Client VPN endpoint.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The ID of the VPC in which the target subnet is located.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a NetworkAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAssociationArgs | NetworkAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkAssociationState | undefined;
            resourceInputs["associationId"] = state ? state.associationId : undefined;
            resourceInputs["clientVpnEndpointId"] = state ? state.clientVpnEndpointId : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as NetworkAssociationArgs | undefined;
            if ((!args || args.clientVpnEndpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientVpnEndpointId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["clientVpnEndpointId"] = args ? args.clientVpnEndpointId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["associationId"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAssociation resources.
 */
export interface NetworkAssociationState {
    /**
     * The unique ID of the target network association.
     */
    associationId?: pulumi.Input<string>;
    /**
     * The ID of the Client VPN endpoint.
     */
    clientVpnEndpointId?: pulumi.Input<string>;
    /**
     * The ID of the subnet to associate with the Client VPN endpoint.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The ID of the VPC in which the target subnet is located.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAssociation resource.
 */
export interface NetworkAssociationArgs {
    /**
     * The ID of the Client VPN endpoint.
     */
    clientVpnEndpointId: pulumi.Input<string>;
    /**
     * The ID of the subnet to associate with the Client VPN endpoint.
     */
    subnetId: pulumi.Input<string>;
}
