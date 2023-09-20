// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage the accepter's side of a VPC Peering Connection.
 *
 * When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
 * VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
 * accepter's account.
 * The requester can use the `aws.ec2.VpcPeeringConnection` resource to manage its side of the connection
 * and the accepter can use the `aws.ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
 * connection into management.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const peer = new aws.Provider("peer", {region: "us-west-2"});
 * // Accepter's credentials.
 * const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
 * const peerVpc = new aws.ec2.Vpc("peerVpc", {cidrBlock: "10.1.0.0/16"}, {
 *     provider: aws.peer,
 * });
 * const peerCallerIdentity = aws.getCallerIdentity({});
 * // Requester's side of the connection.
 * const peerVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection", {
 *     vpcId: main.id,
 *     peerVpcId: peerVpc.id,
 *     peerOwnerId: peerCallerIdentity.then(peerCallerIdentity => peerCallerIdentity.accountId),
 *     peerRegion: "us-west-2",
 *     autoAccept: false,
 *     tags: {
 *         Side: "Requester",
 *     },
 * });
 * // Accepter's side of the connection.
 * const peerVpcPeeringConnectionAccepter = new aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", {
 *     vpcPeeringConnectionId: peerVpcPeeringConnection.id,
 *     autoAccept: true,
 *     tags: {
 *         Side: "Accepter",
 *     },
 * }, {
 *     provider: aws.peer,
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import VPC Peering Connection Accepters using the Peering Connection ID. For example:
 *
 * ```sh
 *  $ pulumi import aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter example pcx-12345678
 * ```
 *  Certain resource arguments, like `auto_accept`, do not have an EC2 API method for reading the information after peering connection creation. If the argument is set in the TODO configuration on an imported resource, TODO will always show a difference. To workaround this behavior, either omit the argument from the TODO configuration or use `ignore_changes` to hide the difference. For example:
 */
export class VpcPeeringConnectionAccepter extends pulumi.CustomResource {
    /**
     * Get an existing VpcPeeringConnectionAccepter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcPeeringConnectionAccepterState, opts?: pulumi.CustomResourceOptions): VpcPeeringConnectionAccepter {
        return new VpcPeeringConnectionAccepter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter';

    /**
     * Returns true if the given object is an instance of VpcPeeringConnectionAccepter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcPeeringConnectionAccepter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcPeeringConnectionAccepter.__pulumiType;
    }

    /**
     * The status of the VPC Peering Connection request.
     */
    public /*out*/ readonly acceptStatus!: pulumi.Output<string>;
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     */
    public readonly accepter!: pulumi.Output<outputs.ec2.VpcPeeringConnectionAccepterAccepter>;
    /**
     * Whether or not to accept the peering request. Defaults to `false`.
     */
    public readonly autoAccept!: pulumi.Output<boolean | undefined>;
    /**
     * The AWS account ID of the owner of the requester VPC.
     */
    public /*out*/ readonly peerOwnerId!: pulumi.Output<string>;
    /**
     * The region of the accepter VPC.
     */
    public /*out*/ readonly peerRegion!: pulumi.Output<string>;
    /**
     * The ID of the requester VPC.
     */
    public /*out*/ readonly peerVpcId!: pulumi.Output<string>;
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     */
    public readonly requester!: pulumi.Output<outputs.ec2.VpcPeeringConnectionAccepterRequester>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of the accepter VPC.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The VPC Peering Connection ID to manage.
     */
    public readonly vpcPeeringConnectionId!: pulumi.Output<string>;

    /**
     * Create a VpcPeeringConnectionAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcPeeringConnectionAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcPeeringConnectionAccepterArgs | VpcPeeringConnectionAccepterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcPeeringConnectionAccepterState | undefined;
            resourceInputs["acceptStatus"] = state ? state.acceptStatus : undefined;
            resourceInputs["accepter"] = state ? state.accepter : undefined;
            resourceInputs["autoAccept"] = state ? state.autoAccept : undefined;
            resourceInputs["peerOwnerId"] = state ? state.peerOwnerId : undefined;
            resourceInputs["peerRegion"] = state ? state.peerRegion : undefined;
            resourceInputs["peerVpcId"] = state ? state.peerVpcId : undefined;
            resourceInputs["requester"] = state ? state.requester : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcPeeringConnectionId"] = state ? state.vpcPeeringConnectionId : undefined;
        } else {
            const args = argsOrState as VpcPeeringConnectionAccepterArgs | undefined;
            if ((!args || args.vpcPeeringConnectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcPeeringConnectionId'");
            }
            resourceInputs["accepter"] = args ? args.accepter : undefined;
            resourceInputs["autoAccept"] = args ? args.autoAccept : undefined;
            resourceInputs["requester"] = args ? args.requester : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcPeeringConnectionId"] = args ? args.vpcPeeringConnectionId : undefined;
            resourceInputs["acceptStatus"] = undefined /*out*/;
            resourceInputs["peerOwnerId"] = undefined /*out*/;
            resourceInputs["peerRegion"] = undefined /*out*/;
            resourceInputs["peerVpcId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(VpcPeeringConnectionAccepter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcPeeringConnectionAccepter resources.
 */
export interface VpcPeeringConnectionAccepterState {
    /**
     * The status of the VPC Peering Connection request.
     */
    acceptStatus?: pulumi.Input<string>;
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     */
    accepter?: pulumi.Input<inputs.ec2.VpcPeeringConnectionAccepterAccepter>;
    /**
     * Whether or not to accept the peering request. Defaults to `false`.
     */
    autoAccept?: pulumi.Input<boolean>;
    /**
     * The AWS account ID of the owner of the requester VPC.
     */
    peerOwnerId?: pulumi.Input<string>;
    /**
     * The region of the accepter VPC.
     */
    peerRegion?: pulumi.Input<string>;
    /**
     * The ID of the requester VPC.
     */
    peerVpcId?: pulumi.Input<string>;
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     */
    requester?: pulumi.Input<inputs.ec2.VpcPeeringConnectionAccepterRequester>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the accepter VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The VPC Peering Connection ID to manage.
     */
    vpcPeeringConnectionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcPeeringConnectionAccepter resource.
 */
export interface VpcPeeringConnectionAccepterArgs {
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     */
    accepter?: pulumi.Input<inputs.ec2.VpcPeeringConnectionAccepterAccepter>;
    /**
     * Whether or not to accept the peering request. Defaults to `false`.
     */
    autoAccept?: pulumi.Input<boolean>;
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     */
    requester?: pulumi.Input<inputs.ec2.VpcPeeringConnectionAccepterRequester>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VPC Peering Connection ID to manage.
     */
    vpcPeeringConnectionId: pulumi.Input<string>;
}
