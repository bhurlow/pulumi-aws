// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS NetworkManager Connect Peer.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * `aws_networkmanager_connect_peer` can be imported using the connect peer ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:networkmanager/connectPeer:ConnectPeer example connect-peer-061f3e96275db1acc
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
    public static readonly __pulumiType = 'aws:networkmanager/connectPeer:ConnectPeer';

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
     * The ARN of the attachment.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Connect peer BGP options.
     */
    public readonly bgpOptions!: pulumi.Output<outputs.networkmanager.ConnectPeerBgpOptions | undefined>;
    /**
     * The configuration of the Connect peer.
     */
    public /*out*/ readonly configurations!: pulumi.Output<outputs.networkmanager.ConnectPeerConfiguration[]>;
    /**
     * The ID of the connection attachment.
     */
    public readonly connectAttachmentId!: pulumi.Output<string>;
    public /*out*/ readonly connectPeerId!: pulumi.Output<string>;
    /**
     * A Connect peer core network address.
     */
    public readonly coreNetworkAddress!: pulumi.Output<string | undefined>;
    /**
     * The ID of a core network.
     */
    public /*out*/ readonly coreNetworkId!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The Region where the peer is located.
     */
    public /*out*/ readonly edgeLocation!: pulumi.Output<string>;
    /**
     * The inside IP addresses used for BGP peering.
     */
    public readonly insideCidrBlocks!: pulumi.Output<string[]>;
    /**
     * The Connect peer address.
     */
    public readonly peerAddress!: pulumi.Output<string>;
    /**
     * The state of the Connect peer.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Key-value tags for the attachment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

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
            resourceInputs["bgpOptions"] = state ? state.bgpOptions : undefined;
            resourceInputs["configurations"] = state ? state.configurations : undefined;
            resourceInputs["connectAttachmentId"] = state ? state.connectAttachmentId : undefined;
            resourceInputs["connectPeerId"] = state ? state.connectPeerId : undefined;
            resourceInputs["coreNetworkAddress"] = state ? state.coreNetworkAddress : undefined;
            resourceInputs["coreNetworkId"] = state ? state.coreNetworkId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["edgeLocation"] = state ? state.edgeLocation : undefined;
            resourceInputs["insideCidrBlocks"] = state ? state.insideCidrBlocks : undefined;
            resourceInputs["peerAddress"] = state ? state.peerAddress : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as ConnectPeerArgs | undefined;
            if ((!args || args.connectAttachmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectAttachmentId'");
            }
            if ((!args || args.insideCidrBlocks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'insideCidrBlocks'");
            }
            if ((!args || args.peerAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerAddress'");
            }
            resourceInputs["bgpOptions"] = args ? args.bgpOptions : undefined;
            resourceInputs["connectAttachmentId"] = args ? args.connectAttachmentId : undefined;
            resourceInputs["coreNetworkAddress"] = args ? args.coreNetworkAddress : undefined;
            resourceInputs["insideCidrBlocks"] = args ? args.insideCidrBlocks : undefined;
            resourceInputs["peerAddress"] = args ? args.peerAddress : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["configurations"] = undefined /*out*/;
            resourceInputs["connectPeerId"] = undefined /*out*/;
            resourceInputs["coreNetworkId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["edgeLocation"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConnectPeer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectPeer resources.
 */
export interface ConnectPeerState {
    /**
     * The ARN of the attachment.
     */
    arn?: pulumi.Input<string>;
    /**
     * The Connect peer BGP options.
     */
    bgpOptions?: pulumi.Input<inputs.networkmanager.ConnectPeerBgpOptions>;
    /**
     * The configuration of the Connect peer.
     */
    configurations?: pulumi.Input<pulumi.Input<inputs.networkmanager.ConnectPeerConfiguration>[]>;
    /**
     * The ID of the connection attachment.
     */
    connectAttachmentId?: pulumi.Input<string>;
    connectPeerId?: pulumi.Input<string>;
    /**
     * A Connect peer core network address.
     */
    coreNetworkAddress?: pulumi.Input<string>;
    /**
     * The ID of a core network.
     */
    coreNetworkId?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    /**
     * The Region where the peer is located.
     */
    edgeLocation?: pulumi.Input<string>;
    /**
     * The inside IP addresses used for BGP peering.
     */
    insideCidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Connect peer address.
     */
    peerAddress?: pulumi.Input<string>;
    /**
     * The state of the Connect peer.
     */
    state?: pulumi.Input<string>;
    /**
     * Key-value tags for the attachment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ConnectPeer resource.
 */
export interface ConnectPeerArgs {
    /**
     * The Connect peer BGP options.
     */
    bgpOptions?: pulumi.Input<inputs.networkmanager.ConnectPeerBgpOptions>;
    /**
     * The ID of the connection attachment.
     */
    connectAttachmentId: pulumi.Input<string>;
    /**
     * A Connect peer core network address.
     */
    coreNetworkAddress?: pulumi.Input<string>;
    /**
     * The inside IP addresses used for BGP peering.
     */
    insideCidrBlocks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Connect peer address.
     */
    peerAddress: pulumi.Input<string>;
    /**
     * Key-value tags for the attachment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
