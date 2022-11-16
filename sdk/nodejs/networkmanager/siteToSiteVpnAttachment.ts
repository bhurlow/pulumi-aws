// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS NetworkManager SiteToSiteAttachment.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.networkmanager.SiteToSiteVpnAttachment("example", {
 *     coreNetworkId: awscc_networkmanager_core_network.example.id,
 *     vpnConnectionArn: aws_vpn_connection.example.arn,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_networkmanager_site_to_site_vpn_attachment` can be imported using the attachment ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:networkmanager/siteToSiteVpnAttachment:SiteToSiteVpnAttachment example attachment-0f8fa60d2238d1bd8
 * ```
 */
export class SiteToSiteVpnAttachment extends pulumi.CustomResource {
    /**
     * Get an existing SiteToSiteVpnAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SiteToSiteVpnAttachmentState, opts?: pulumi.CustomResourceOptions): SiteToSiteVpnAttachment {
        return new SiteToSiteVpnAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:networkmanager/siteToSiteVpnAttachment:SiteToSiteVpnAttachment';

    /**
     * Returns true if the given object is an instance of SiteToSiteVpnAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SiteToSiteVpnAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SiteToSiteVpnAttachment.__pulumiType;
    }

    /**
     * The ARN of the attachment.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The policy rule number associated with the attachment.
     */
    public /*out*/ readonly attachmentPolicyRuleNumber!: pulumi.Output<number>;
    /**
     * The type of attachment.
     */
    public /*out*/ readonly attachmentType!: pulumi.Output<string>;
    /**
     * The ARN of a core network.
     */
    public /*out*/ readonly coreNetworkArn!: pulumi.Output<string>;
    /**
     * The ID of a core network for the VPN attachment.
     */
    public readonly coreNetworkId!: pulumi.Output<string>;
    /**
     * The Region where the edge is located.
     */
    public /*out*/ readonly edgeLocation!: pulumi.Output<string>;
    /**
     * The ID of the attachment account owner.
     */
    public /*out*/ readonly ownerAccountId!: pulumi.Output<string>;
    /**
     * The attachment resource ARN.
     */
    public /*out*/ readonly resourceArn!: pulumi.Output<string>;
    /**
     * The name of the segment attachment.
     */
    public /*out*/ readonly segmentName!: pulumi.Output<string>;
    /**
     * The state of the attachment.
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
     * The ARN of the site-to-site VPN connection.
     */
    public readonly vpnConnectionArn!: pulumi.Output<string>;

    /**
     * Create a SiteToSiteVpnAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteToSiteVpnAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SiteToSiteVpnAttachmentArgs | SiteToSiteVpnAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SiteToSiteVpnAttachmentState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["attachmentPolicyRuleNumber"] = state ? state.attachmentPolicyRuleNumber : undefined;
            resourceInputs["attachmentType"] = state ? state.attachmentType : undefined;
            resourceInputs["coreNetworkArn"] = state ? state.coreNetworkArn : undefined;
            resourceInputs["coreNetworkId"] = state ? state.coreNetworkId : undefined;
            resourceInputs["edgeLocation"] = state ? state.edgeLocation : undefined;
            resourceInputs["ownerAccountId"] = state ? state.ownerAccountId : undefined;
            resourceInputs["resourceArn"] = state ? state.resourceArn : undefined;
            resourceInputs["segmentName"] = state ? state.segmentName : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpnConnectionArn"] = state ? state.vpnConnectionArn : undefined;
        } else {
            const args = argsOrState as SiteToSiteVpnAttachmentArgs | undefined;
            if ((!args || args.coreNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'coreNetworkId'");
            }
            if ((!args || args.vpnConnectionArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnConnectionArn'");
            }
            resourceInputs["coreNetworkId"] = args ? args.coreNetworkId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpnConnectionArn"] = args ? args.vpnConnectionArn : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["attachmentPolicyRuleNumber"] = undefined /*out*/;
            resourceInputs["attachmentType"] = undefined /*out*/;
            resourceInputs["coreNetworkArn"] = undefined /*out*/;
            resourceInputs["edgeLocation"] = undefined /*out*/;
            resourceInputs["ownerAccountId"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["segmentName"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SiteToSiteVpnAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SiteToSiteVpnAttachment resources.
 */
export interface SiteToSiteVpnAttachmentState {
    /**
     * The ARN of the attachment.
     */
    arn?: pulumi.Input<string>;
    /**
     * The policy rule number associated with the attachment.
     */
    attachmentPolicyRuleNumber?: pulumi.Input<number>;
    /**
     * The type of attachment.
     */
    attachmentType?: pulumi.Input<string>;
    /**
     * The ARN of a core network.
     */
    coreNetworkArn?: pulumi.Input<string>;
    /**
     * The ID of a core network for the VPN attachment.
     */
    coreNetworkId?: pulumi.Input<string>;
    /**
     * The Region where the edge is located.
     */
    edgeLocation?: pulumi.Input<string>;
    /**
     * The ID of the attachment account owner.
     */
    ownerAccountId?: pulumi.Input<string>;
    /**
     * The attachment resource ARN.
     */
    resourceArn?: pulumi.Input<string>;
    /**
     * The name of the segment attachment.
     */
    segmentName?: pulumi.Input<string>;
    /**
     * The state of the attachment.
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
    /**
     * The ARN of the site-to-site VPN connection.
     */
    vpnConnectionArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SiteToSiteVpnAttachment resource.
 */
export interface SiteToSiteVpnAttachmentArgs {
    /**
     * The ID of a core network for the VPN attachment.
     */
    coreNetworkId: pulumi.Input<string>;
    /**
     * Key-value tags for the attachment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ARN of the site-to-site VPN connection.
     */
    vpnConnectionArn: pulumi.Input<string>;
}
