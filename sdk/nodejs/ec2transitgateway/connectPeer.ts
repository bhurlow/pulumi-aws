// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EC2 Transit Gateway Connect Peer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleConnect = new aws.ec2transitgateway.Connect("exampleConnect", {
 *     transportAttachmentId: aws_ec2_transit_gateway_vpc_attachment.example.id,
 *     transitGatewayId: aws_ec2_transit_gateway.example.id,
 * });
 * const exampleConnectPeer = new aws.ec2transitgateway.ConnectPeer("exampleConnectPeer", {
 *     peerAddress: "10.1.2.3",
 *     insideCidrBlocks: ["169.254.100.0/29"],
 *     transitGatewayAttachmentId: exampleConnect.id,
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import `aws_ec2_transit_gateway_connect_peer` using the EC2 Transit Gateway Connect Peer identifier. For example:
 *
 * ```sh
 *  $ pulumi import aws:ec2transitgateway/connectPeer:ConnectPeer example tgw-connect-peer-12345678
 * ```
 */
export class ConnectPeer extends pulumi.CustomResource {
    /**
     * Get an existing ConnectPeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectPeerState, opts?: pulumi.CustomResourceOptions): ConnectPeer {
        return new ConnectPeer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2transitgateway/connectPeer:ConnectPeer';

    /**
     * Returns true if the given object is an instance of ConnectPeer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectPeer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectPeer.__pulumiType;
    }

    /**
     * EC2 Transit Gateway Connect Peer ARN
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
     */
    public readonly bgpAsn!: pulumi.Output<string>;
    /**
     * The IP address assigned to customer device, which is used as BGP IP address.
     */
    public /*out*/ readonly bgpPeerAddress!: pulumi.Output<string>;
    /**
     * The IP addresses assigned to Transit Gateway, which are used as BGP IP addresses.
     */
    public /*out*/ readonly bgpTransitGatewayAddresses!: pulumi.Output<string[]>;
    /**
     * The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
     */
    public readonly insideCidrBlocks!: pulumi.Output<string[]>;
    /**
     * The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
     */
    public readonly peerAddress!: pulumi.Output<string>;
    /**
     * Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
     */
    public readonly transitGatewayAddress!: pulumi.Output<string>;
    /**
     * The Transit Gateway Connect
     */
    public readonly transitGatewayAttachmentId!: pulumi.Output<string>;

    /**
     * Create a ConnectPeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectPeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectPeerArgs | ConnectPeerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectPeerState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            resourceInputs["bgpPeerAddress"] = state ? state.bgpPeerAddress : undefined;
            resourceInputs["bgpTransitGatewayAddresses"] = state ? state.bgpTransitGatewayAddresses : undefined;
            resourceInputs["insideCidrBlocks"] = state ? state.insideCidrBlocks : undefined;
            resourceInputs["peerAddress"] = state ? state.peerAddress : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["transitGatewayAddress"] = state ? state.transitGatewayAddress : undefined;
            resourceInputs["transitGatewayAttachmentId"] = state ? state.transitGatewayAttachmentId : undefined;
        } else {
            const args = argsOrState as ConnectPeerArgs | undefined;
            if ((!args || args.insideCidrBlocks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'insideCidrBlocks'");
            }
            if ((!args || args.peerAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerAddress'");
            }
            if ((!args || args.transitGatewayAttachmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayAttachmentId'");
            }
            resourceInputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            resourceInputs["insideCidrBlocks"] = args ? args.insideCidrBlocks : undefined;
            resourceInputs["peerAddress"] = args ? args.peerAddress : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitGatewayAddress"] = args ? args.transitGatewayAddress : undefined;
            resourceInputs["transitGatewayAttachmentId"] = args ? args.transitGatewayAttachmentId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["bgpPeerAddress"] = undefined /*out*/;
            resourceInputs["bgpTransitGatewayAddresses"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ConnectPeer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectPeer resources.
 */
export interface ConnectPeerState {
    /**
     * EC2 Transit Gateway Connect Peer ARN
     */
    arn?: pulumi.Input<string>;
    /**
     * The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
     */
    bgpAsn?: pulumi.Input<string>;
    /**
     * The IP address assigned to customer device, which is used as BGP IP address.
     */
    bgpPeerAddress?: pulumi.Input<string>;
    /**
     * The IP addresses assigned to Transit Gateway, which are used as BGP IP addresses.
     */
    bgpTransitGatewayAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
     */
    insideCidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
     */
    peerAddress?: pulumi.Input<string>;
    /**
     * Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
     */
    transitGatewayAddress?: pulumi.Input<string>;
    /**
     * The Transit Gateway Connect
     */
    transitGatewayAttachmentId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConnectPeer resource.
 */
export interface ConnectPeerArgs {
    /**
     * The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
     */
    bgpAsn?: pulumi.Input<string>;
    /**
     * The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
     */
    insideCidrBlocks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
     */
    peerAddress: pulumi.Input<string>;
    /**
     * Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
     */
    transitGatewayAddress?: pulumi.Input<string>;
    /**
     * The Transit Gateway Connect
     */
    transitGatewayAttachmentId: pulumi.Input<string>;
}
