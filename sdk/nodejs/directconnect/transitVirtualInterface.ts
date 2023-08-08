// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect transit virtual interface resource.
 * A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one or more transit gateways.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
 * const exampleTransitVirtualInterface = new aws.directconnect.TransitVirtualInterface("exampleTransitVirtualInterface", {
 *     connectionId: aws_dx_connection.example.id,
 *     dxGatewayId: exampleGateway.id,
 *     vlan: 4094,
 *     addressFamily: "ipv4",
 *     bgpAsn: 65352,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_dx_transit_virtual_interface.test
 *
 *  id = "dxvif-33cc44dd" } Using `pulumi import`, import Direct Connect transit virtual interfaces using the VIF `id`. For exampleconsole % pulumi import aws_dx_transit_virtual_interface.test dxvif-33cc44dd
 */
export class TransitVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing TransitVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitVirtualInterfaceState, opts?: pulumi.CustomResourceOptions): TransitVirtualInterface {
        return new TransitVirtualInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/transitVirtualInterface:TransitVirtualInterface';

    /**
     * Returns true if the given object is an instance of TransitVirtualInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitVirtualInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitVirtualInterface.__pulumiType;
    }

    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    public readonly amazonAddress!: pulumi.Output<string>;
    public /*out*/ readonly amazonSideAsn!: pulumi.Output<string>;
    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Direct Connect endpoint on which the virtual interface terminates.
     */
    public /*out*/ readonly awsDevice!: pulumi.Output<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    public readonly bgpAsn!: pulumi.Output<number>;
    /**
     * The authentication key for BGP configuration.
     */
    public readonly bgpAuthKey!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    public readonly customerAddress!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    public readonly dxGatewayId!: pulumi.Output<string>;
    /**
     * Indicates whether jumbo frames (8500 MTU) are supported.
     */
    public /*out*/ readonly jumboFrameCapable!: pulumi.Output<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * The name for the virtual interface.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates whether to enable or disable SiteLink.
     */
    public readonly sitelinkEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The VLAN ID.
     */
    public readonly vlan!: pulumi.Output<number>;

    /**
     * Create a TransitVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitVirtualInterfaceArgs | TransitVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransitVirtualInterfaceState | undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            resourceInputs["amazonSideAsn"] = state ? state.amazonSideAsn : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["awsDevice"] = state ? state.awsDevice : undefined;
            resourceInputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            resourceInputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            resourceInputs["connectionId"] = state ? state.connectionId : undefined;
            resourceInputs["customerAddress"] = state ? state.customerAddress : undefined;
            resourceInputs["dxGatewayId"] = state ? state.dxGatewayId : undefined;
            resourceInputs["jumboFrameCapable"] = state ? state.jumboFrameCapable : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sitelinkEnabled"] = state ? state.sitelinkEnabled : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as TransitVirtualInterfaceArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.bgpAsn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            if ((!args || args.dxGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dxGatewayId'");
            }
            if ((!args || args.vlan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vlan'");
            }
            resourceInputs["addressFamily"] = args ? args.addressFamily : undefined;
            resourceInputs["amazonAddress"] = args ? args.amazonAddress : undefined;
            resourceInputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            resourceInputs["bgpAuthKey"] = args ? args.bgpAuthKey : undefined;
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["customerAddress"] = args ? args.customerAddress : undefined;
            resourceInputs["dxGatewayId"] = args ? args.dxGatewayId : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sitelinkEnabled"] = args ? args.sitelinkEnabled : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
            resourceInputs["amazonSideAsn"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsDevice"] = undefined /*out*/;
            resourceInputs["jumboFrameCapable"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitVirtualInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitVirtualInterface resources.
 */
export interface TransitVirtualInterfaceState {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    addressFamily?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    amazonAddress?: pulumi.Input<string>;
    amazonSideAsn?: pulumi.Input<string>;
    /**
     * The ARN of the virtual interface.
     */
    arn?: pulumi.Input<string>;
    /**
     * The Direct Connect endpoint on which the virtual interface terminates.
     */
    awsDevice?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    bgpAsn?: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    connectionId?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    customerAddress?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    dxGatewayId?: pulumi.Input<string>;
    /**
     * Indicates whether jumbo frames (8500 MTU) are supported.
     */
    jumboFrameCapable?: pulumi.Input<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicates whether to enable or disable SiteLink.
     */
    sitelinkEnabled?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VLAN ID.
     */
    vlan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a TransitVirtualInterface resource.
 */
export interface TransitVirtualInterfaceArgs {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    addressFamily: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    amazonAddress?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    bgpAsn: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    connectionId: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    customerAddress?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    dxGatewayId: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicates whether to enable or disable SiteLink.
     */
    sitelinkEnabled?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VLAN ID.
     */
    vlan: pulumi.Input<number>;
}
