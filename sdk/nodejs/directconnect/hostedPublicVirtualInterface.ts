// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect hosted public virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
 * A hosted virtual interface is a virtual interface that is owned by another AWS account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.directconnect.HostedPublicVirtualInterface("foo", {
 *     addressFamily: "ipv4",
 *     amazonAddress: "175.45.176.2/30",
 *     bgpAsn: 65352,
 *     connectionId: "dxcon-zzzzzzzz",
 *     customerAddress: "175.45.176.1/30",
 *     routeFilterPrefixes: [
 *         "210.52.109.0/24",
 *         "175.45.176.0/22",
 *     ],
 *     vlan: 4094,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_dx_hosted_public_virtual_interface.test
 *
 *  id = "dxvif-33cc44dd" } Using `pulumi import`, import Direct Connect hosted public virtual interfaces using the VIF `id`. For exampleconsole % pulumi import aws_dx_hosted_public_virtual_interface.test dxvif-33cc44dd
 */
export class HostedPublicVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing HostedPublicVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostedPublicVirtualInterfaceState, opts?: pulumi.CustomResourceOptions): HostedPublicVirtualInterface {
        return new HostedPublicVirtualInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface';

    /**
     * Returns true if the given object is an instance of HostedPublicVirtualInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostedPublicVirtualInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostedPublicVirtualInterface.__pulumiType;
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
     * The name for the virtual interface.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    public readonly ownerAccountId!: pulumi.Output<string>;
    /**
     * A list of routes to be advertised to the AWS network in this region.
     */
    public readonly routeFilterPrefixes!: pulumi.Output<string[]>;
    /**
     * The VLAN ID.
     */
    public readonly vlan!: pulumi.Output<number>;

    /**
     * Create a HostedPublicVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostedPublicVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostedPublicVirtualInterfaceArgs | HostedPublicVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostedPublicVirtualInterfaceState | undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            resourceInputs["amazonSideAsn"] = state ? state.amazonSideAsn : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["awsDevice"] = state ? state.awsDevice : undefined;
            resourceInputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            resourceInputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            resourceInputs["connectionId"] = state ? state.connectionId : undefined;
            resourceInputs["customerAddress"] = state ? state.customerAddress : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ownerAccountId"] = state ? state.ownerAccountId : undefined;
            resourceInputs["routeFilterPrefixes"] = state ? state.routeFilterPrefixes : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as HostedPublicVirtualInterfaceArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.bgpAsn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            if ((!args || args.ownerAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ownerAccountId'");
            }
            if ((!args || args.routeFilterPrefixes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeFilterPrefixes'");
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
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerAccountId"] = args ? args.ownerAccountId : undefined;
            resourceInputs["routeFilterPrefixes"] = args ? args.routeFilterPrefixes : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
            resourceInputs["amazonSideAsn"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsDevice"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostedPublicVirtualInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostedPublicVirtualInterface resources.
 */
export interface HostedPublicVirtualInterfaceState {
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
     * The name for the virtual interface.
     */
    name?: pulumi.Input<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    ownerAccountId?: pulumi.Input<string>;
    /**
     * A list of routes to be advertised to the AWS network in this region.
     */
    routeFilterPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The VLAN ID.
     */
    vlan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HostedPublicVirtualInterface resource.
 */
export interface HostedPublicVirtualInterfaceArgs {
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
     * The name for the virtual interface.
     */
    name?: pulumi.Input<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    ownerAccountId: pulumi.Input<string>;
    /**
     * A list of routes to be advertised to the AWS network in this region.
     */
    routeFilterPrefixes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The VLAN ID.
     */
    vlan: pulumi.Input<number>;
}
