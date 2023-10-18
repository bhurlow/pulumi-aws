// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Core Network Policy Attachment resource. This puts a Core Network Policy to an existing Core Network and executes the change set, which deploys changes globally based on the policy submitted (Sets the policy to `LIVE`).
 *
 * > **NOTE:** Deleting this resource will not delete the current policy defined in this resource. Deleting this resource will also not revert the current `LIVE` policy to the previous version.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleCoreNetwork = new aws.networkmanager.CoreNetwork("exampleCoreNetwork", {globalNetworkId: aws_networkmanager_global_network.example.id});
 * const exampleCoreNetworkPolicyAttachment = new aws.networkmanager.CoreNetworkPolicyAttachment("exampleCoreNetworkPolicyAttachment", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     policyDocument: data.aws_networkmanager_core_network_policy_document.example.json,
 * });
 * ```
 * ### With VPC Attachment (Single Region)
 *
 * The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `createBasePolicy` argument of the `aws.networkmanager.CoreNetwork` resource to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `pulumi up` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `createBasePolicy` argument. There are 2 options to implement this:
 *
 * - Option 1: Use the `basePolicyDocument` argument in the `aws.networkmanager.CoreNetwork` resource that allows the most customizations to a base policy. Use this to customize the `edgeLocations` `asn`. In the example below, `us-west-2` and ASN `65500` are used in the base policy.
 * - Option 2: Use the `createBasePolicy` argument only. This creates a base policy in the region specified in the `provider` block.
 * ### Option 1 - using basePolicyDocument
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGlobalNetwork = new aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", {});
 * const base = aws.networkmanager.getCoreNetworkPolicyDocument({
 *     coreNetworkConfigurations: [{
 *         asnRanges: ["65022-65534"],
 *         edgeLocations: [{
 *             location: "us-west-2",
 *             asn: "65500",
 *         }],
 *     }],
 *     segments: [{
 *         name: "segment",
 *     }],
 * });
 * const exampleCoreNetwork = new aws.networkmanager.CoreNetwork("exampleCoreNetwork", {
 *     globalNetworkId: exampleGlobalNetwork.id,
 *     basePolicyDocument: base.then(base => base.json),
 *     createBasePolicy: true,
 * });
 * const exampleVpcAttachment = new aws.networkmanager.VpcAttachment("exampleVpcAttachment", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     subnetArns: aws_subnet.example.map(__item => __item.arn),
 *     vpcArn: aws_vpc.example.arn,
 * });
 * const exampleCoreNetworkPolicyDocument = aws.networkmanager.getCoreNetworkPolicyDocumentOutput({
 *     coreNetworkConfigurations: [{
 *         asnRanges: ["65022-65534"],
 *         edgeLocations: [{
 *             location: "us-west-2",
 *             asn: "65500",
 *         }],
 *     }],
 *     segments: [{
 *         name: "segment",
 *     }],
 *     segmentActions: [{
 *         action: "create-route",
 *         segment: "segment",
 *         destinationCidrBlocks: ["0.0.0.0/0"],
 *         destinations: [exampleVpcAttachment.id],
 *     }],
 * });
 * const exampleCoreNetworkPolicyAttachment = new aws.networkmanager.CoreNetworkPolicyAttachment("exampleCoreNetworkPolicyAttachment", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     policyDocument: exampleCoreNetworkPolicyDocument.apply(exampleCoreNetworkPolicyDocument => exampleCoreNetworkPolicyDocument.json),
 * });
 * ```
 * ### Option 2 - createBasePolicy only
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGlobalNetwork = new aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", {});
 * const exampleCoreNetwork = new aws.networkmanager.CoreNetwork("exampleCoreNetwork", {
 *     globalNetworkId: exampleGlobalNetwork.id,
 *     createBasePolicy: true,
 * });
 * const exampleVpcAttachment = new aws.networkmanager.VpcAttachment("exampleVpcAttachment", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     subnetArns: aws_subnet.example.map(__item => __item.arn),
 *     vpcArn: aws_vpc.example.arn,
 * });
 * const exampleCoreNetworkPolicyDocument = aws.networkmanager.getCoreNetworkPolicyDocumentOutput({
 *     coreNetworkConfigurations: [{
 *         asnRanges: ["65022-65534"],
 *         edgeLocations: [{
 *             location: "us-west-2",
 *         }],
 *     }],
 *     segments: [{
 *         name: "segment",
 *     }],
 *     segmentActions: [{
 *         action: "create-route",
 *         segment: "segment",
 *         destinationCidrBlocks: ["0.0.0.0/0"],
 *         destinations: [exampleVpcAttachment.id],
 *     }],
 * });
 * const exampleCoreNetworkPolicyAttachment = new aws.networkmanager.CoreNetworkPolicyAttachment("exampleCoreNetworkPolicyAttachment", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     policyDocument: exampleCoreNetworkPolicyDocument.apply(exampleCoreNetworkPolicyDocument => exampleCoreNetworkPolicyDocument.json),
 * });
 * ```
 * ### With VPC Attachment (Multi-Region)
 *
 * The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `createBasePolicy` argument of the `aws.networkmanager.CoreNetwork` resource to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `pulumi up` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `createBasePolicy` argument. For multi-region in a core network that does not yet have a `LIVE` policy, there are 2 options:
 *
 * - Option 1: Use the `basePolicyDocument` argument that allows the most customizations to a base policy. Use this to customize the `edgeLocations` `asn`. In the example below, `us-west-2`, `us-east-1` and specific ASNs are used in the base policy.
 * - Option 2: Pass a list of regions to the `aws.networkmanager.CoreNetwork` resource `basePolicyRegions` argument. In the example below, `us-west-2` and `us-east-1` are specified in the base policy.
 * ### Option 1 - using basePolicyDocument
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGlobalNetwork = new aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", {});
 * const base = aws.networkmanager.getCoreNetworkPolicyDocument({
 *     coreNetworkConfigurations: [{
 *         asnRanges: ["65022-65534"],
 *         edgeLocations: [
 *             {
 *                 location: "us-west-2",
 *                 asn: "65500",
 *             },
 *             {
 *                 location: "us-east-1",
 *                 asn: "65501",
 *             },
 *         ],
 *     }],
 *     segments: [{
 *         name: "segment",
 *     }],
 * });
 * const exampleCoreNetwork = new aws.networkmanager.CoreNetwork("exampleCoreNetwork", {
 *     globalNetworkId: exampleGlobalNetwork.id,
 *     basePolicyDocument: base.then(base => base.json),
 *     createBasePolicy: true,
 * });
 * const exampleUsWest2 = new aws.networkmanager.VpcAttachment("exampleUsWest2", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     subnetArns: aws_subnet.example_us_west_2.map(__item => __item.arn),
 *     vpcArn: aws_vpc.example_us_west_2.arn,
 * });
 * const exampleUsEast1 = new aws.networkmanager.VpcAttachment("exampleUsEast1", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     subnetArns: aws_subnet.example_us_east_1.map(__item => __item.arn),
 *     vpcArn: aws_vpc.example_us_east_1.arn,
 * }, {
 *     provider: "alternate",
 * });
 * const exampleCoreNetworkPolicyDocument = aws.networkmanager.getCoreNetworkPolicyDocumentOutput({
 *     coreNetworkConfigurations: [{
 *         asnRanges: ["65022-65534"],
 *         edgeLocations: [
 *             {
 *                 location: "us-west-2",
 *                 asn: "65500",
 *             },
 *             {
 *                 location: "us-east-1",
 *                 asn: "65501",
 *             },
 *         ],
 *     }],
 *     segments: [
 *         {
 *             name: "segment",
 *         },
 *         {
 *             name: "segment2",
 *         },
 *     ],
 *     segmentActions: [
 *         {
 *             action: "create-route",
 *             segment: "segment",
 *             destinationCidrBlocks: ["10.0.0.0/16"],
 *             destinations: [exampleUsWest2.id],
 *         },
 *         {
 *             action: "create-route",
 *             segment: "segment",
 *             destinationCidrBlocks: ["10.1.0.0/16"],
 *             destinations: [exampleUsEast1.id],
 *         },
 *     ],
 * });
 * const exampleCoreNetworkPolicyAttachment = new aws.networkmanager.CoreNetworkPolicyAttachment("exampleCoreNetworkPolicyAttachment", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     policyDocument: exampleCoreNetworkPolicyDocument.apply(exampleCoreNetworkPolicyDocument => exampleCoreNetworkPolicyDocument.json),
 * });
 * ```
 * ### Option 2 - using basePolicyRegions
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGlobalNetwork = new aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", {});
 * const exampleCoreNetwork = new aws.networkmanager.CoreNetwork("exampleCoreNetwork", {
 *     globalNetworkId: exampleGlobalNetwork.id,
 *     basePolicyRegions: [
 *         "us-west-2",
 *         "us-east-1",
 *     ],
 *     createBasePolicy: true,
 * });
 * const exampleUsWest2 = new aws.networkmanager.VpcAttachment("exampleUsWest2", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     subnetArns: aws_subnet.example_us_west_2.map(__item => __item.arn),
 *     vpcArn: aws_vpc.example_us_west_2.arn,
 * });
 * const exampleUsEast1 = new aws.networkmanager.VpcAttachment("exampleUsEast1", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     subnetArns: aws_subnet.example_us_east_1.map(__item => __item.arn),
 *     vpcArn: aws_vpc.example_us_east_1.arn,
 * }, {
 *     provider: "alternate",
 * });
 * const exampleCoreNetworkPolicyDocument = aws.networkmanager.getCoreNetworkPolicyDocumentOutput({
 *     coreNetworkConfigurations: [{
 *         asnRanges: ["65022-65534"],
 *         edgeLocations: [
 *             {
 *                 location: "us-west-2",
 *             },
 *             {
 *                 location: "us-east-1",
 *             },
 *         ],
 *     }],
 *     segments: [
 *         {
 *             name: "segment",
 *         },
 *         {
 *             name: "segment2",
 *         },
 *     ],
 *     segmentActions: [
 *         {
 *             action: "create-route",
 *             segment: "segment",
 *             destinationCidrBlocks: ["10.0.0.0/16"],
 *             destinations: [exampleUsWest2.id],
 *         },
 *         {
 *             action: "create-route",
 *             segment: "segment",
 *             destinationCidrBlocks: ["10.1.0.0/16"],
 *             destinations: [exampleUsEast1.id],
 *         },
 *     ],
 * });
 * const exampleCoreNetworkPolicyAttachment = new aws.networkmanager.CoreNetworkPolicyAttachment("exampleCoreNetworkPolicyAttachment", {
 *     coreNetworkId: exampleCoreNetwork.id,
 *     policyDocument: exampleCoreNetworkPolicyDocument.apply(exampleCoreNetworkPolicyDocument => exampleCoreNetworkPolicyDocument.json),
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import `aws_networkmanager_core_network_policy_attachment` using the core network ID. For example:
 *
 * ```sh
 *  $ pulumi import aws:networkmanager/coreNetworkPolicyAttachment:CoreNetworkPolicyAttachment example core-network-0d47f6t230mz46dy4
 * ```
 */
export class CoreNetworkPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing CoreNetworkPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CoreNetworkPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): CoreNetworkPolicyAttachment {
        return new CoreNetworkPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:networkmanager/coreNetworkPolicyAttachment:CoreNetworkPolicyAttachment';

    /**
     * Returns true if the given object is an instance of CoreNetworkPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CoreNetworkPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CoreNetworkPolicyAttachment.__pulumiType;
    }

    /**
     * The ID of the core network that a policy will be attached to and made `LIVE`.
     */
    public readonly coreNetworkId!: pulumi.Output<string>;
    /**
     * Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
     */
    public readonly policyDocument!: pulumi.Output<string>;
    /**
     * Current state of a core network.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a CoreNetworkPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CoreNetworkPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CoreNetworkPolicyAttachmentArgs | CoreNetworkPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CoreNetworkPolicyAttachmentState | undefined;
            resourceInputs["coreNetworkId"] = state ? state.coreNetworkId : undefined;
            resourceInputs["policyDocument"] = state ? state.policyDocument : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as CoreNetworkPolicyAttachmentArgs | undefined;
            if ((!args || args.coreNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'coreNetworkId'");
            }
            if ((!args || args.policyDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyDocument'");
            }
            resourceInputs["coreNetworkId"] = args ? args.coreNetworkId : undefined;
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CoreNetworkPolicyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CoreNetworkPolicyAttachment resources.
 */
export interface CoreNetworkPolicyAttachmentState {
    /**
     * The ID of the core network that a policy will be attached to and made `LIVE`.
     */
    coreNetworkId?: pulumi.Input<string>;
    /**
     * Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
     */
    policyDocument?: pulumi.Input<string>;
    /**
     * Current state of a core network.
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CoreNetworkPolicyAttachment resource.
 */
export interface CoreNetworkPolicyAttachmentArgs {
    /**
     * The ID of the core network that a policy will be attached to and made `LIVE`.
     */
    coreNetworkId: pulumi.Input<string>;
    /**
     * Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
     */
    policyDocument: pulumi.Input<string>;
}
