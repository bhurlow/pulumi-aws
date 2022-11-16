// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Generates a Core Network policy document in JSON format for use with resources that expect core network policy documents such as `awsccNetworkmanagerCoreNetwork`. It follows the API definition from the [core-network-policy documentation](https://docs.aws.amazon.com/vpc/latest/cloudwan/cloudwan-policies-json.html).
 *
 * Using this data source to generate policy documents is *optional*. It is also valid to use literal JSON strings in your configuration or to use the `file` interpolation function to read a raw JSON policy document from a file.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.networkmanager.getCoreNetworkPolicyDocument({
 *     attachmentPolicies: [
 *         {
 *             action: {
 *                 associationMethod: "constant",
 *                 segment: "shared",
 *             },
 *             conditionLogic: "or",
 *             conditions: [{
 *                 key: "segment",
 *                 operator: "equals",
 *                 type: "tag-value",
 *                 value: "shared",
 *             }],
 *             ruleNumber: 100,
 *         },
 *         {
 *             action: {
 *                 associationMethod: "constant",
 *                 segment: "prod",
 *             },
 *             conditionLogic: "or",
 *             conditions: [{
 *                 key: "segment",
 *                 operator: "equals",
 *                 type: "tag-value",
 *                 value: "prod",
 *             }],
 *             ruleNumber: 200,
 *         },
 *     ],
 *     coreNetworkConfigurations: [{
 *         asnRanges: ["64512-64555"],
 *         edgeLocations: [
 *             {
 *                 asn: "64512",
 *                 location: "us-east-1",
 *             },
 *             {
 *                 asn: "64513",
 *                 location: "eu-central-1",
 *             },
 *         ],
 *         vpnEcmpSupport: false,
 *     }],
 *     segmentActions: [{
 *         action: "share",
 *         mode: "attachment-route",
 *         segment: "shared",
 *         shareWiths: ["*"],
 *     }],
 *     segments: [
 *         {
 *             description: "Segment for shared services",
 *             name: "shared",
 *             requireAttachmentAcceptance: true,
 *         },
 *         {
 *             description: "Segment for prod services",
 *             name: "prod",
 *             requireAttachmentAcceptance: true,
 *         },
 *     ],
 * }));
 * ```
 *
 * `data.aws_networkmanager_core_network_policy_document.test.json` will evaluate to:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 */
export function getCoreNetworkPolicyDocument(args: GetCoreNetworkPolicyDocumentArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreNetworkPolicyDocumentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws:networkmanager/getCoreNetworkPolicyDocument:getCoreNetworkPolicyDocument", {
        "attachmentPolicies": args.attachmentPolicies,
        "coreNetworkConfigurations": args.coreNetworkConfigurations,
        "segmentActions": args.segmentActions,
        "segments": args.segments,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getCoreNetworkPolicyDocument.
 */
export interface GetCoreNetworkPolicyDocumentArgs {
    /**
     * In a core network, all attachments use the block argument `attachmentPolicies` section to map an attachment to a segment. Instead of manually associating a segment to each attachment, attachments use tags, and then the tags are used to associate the attachment to the specified segment. Detailed below.
     */
    attachmentPolicies?: inputs.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicy[];
    /**
     * The core network configuration section defines the Regions where a core network should operate. For AWS Regions that are defined in the policy, the core network creates a Core Network Edge where you can connect attachments. After it's created, each Core Network Edge is peered with every other defined Region and is configured with consistent segment and routing across all Regions. Regions cannot be removed until the associated attachments are deleted. Detailed below.
     */
    coreNetworkConfigurations: inputs.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfiguration[];
    /**
     * A block argument, `segmentActions` define how routing works between segments. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
     */
    segmentActions?: inputs.networkmanager.GetCoreNetworkPolicyDocumentSegmentAction[];
    /**
     * Block argument that defines the different segments in the network. Here you can provide descriptions, change defaults, and provide explicit Regional operational and route filters. The names defined for each segment are used in the `segmentActions` and `attachmentPolicies` section. Each segment is created, and operates, as a completely separated routing domain. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
     */
    segments: inputs.networkmanager.GetCoreNetworkPolicyDocumentSegment[];
    version?: string;
}

/**
 * A collection of values returned by getCoreNetworkPolicyDocument.
 */
export interface GetCoreNetworkPolicyDocumentResult {
    readonly attachmentPolicies?: outputs.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicy[];
    readonly coreNetworkConfigurations: outputs.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfiguration[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Standard JSON policy document rendered based on the arguments above.
     */
    readonly json: string;
    readonly segmentActions?: outputs.networkmanager.GetCoreNetworkPolicyDocumentSegmentAction[];
    readonly segments: outputs.networkmanager.GetCoreNetworkPolicyDocumentSegment[];
    readonly version?: string;
}

export function getCoreNetworkPolicyDocumentOutput(args: GetCoreNetworkPolicyDocumentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCoreNetworkPolicyDocumentResult> {
    return pulumi.output(args).apply(a => getCoreNetworkPolicyDocument(a, opts))
}

/**
 * A collection of arguments for invoking getCoreNetworkPolicyDocument.
 */
export interface GetCoreNetworkPolicyDocumentOutputArgs {
    /**
     * In a core network, all attachments use the block argument `attachmentPolicies` section to map an attachment to a segment. Instead of manually associating a segment to each attachment, attachments use tags, and then the tags are used to associate the attachment to the specified segment. Detailed below.
     */
    attachmentPolicies?: pulumi.Input<pulumi.Input<inputs.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyArgs>[]>;
    /**
     * The core network configuration section defines the Regions where a core network should operate. For AWS Regions that are defined in the policy, the core network creates a Core Network Edge where you can connect attachments. After it's created, each Core Network Edge is peered with every other defined Region and is configured with consistent segment and routing across all Regions. Regions cannot be removed until the associated attachments are deleted. Detailed below.
     */
    coreNetworkConfigurations: pulumi.Input<pulumi.Input<inputs.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs>[]>;
    /**
     * A block argument, `segmentActions` define how routing works between segments. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
     */
    segmentActions?: pulumi.Input<pulumi.Input<inputs.networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArgs>[]>;
    /**
     * Block argument that defines the different segments in the network. Here you can provide descriptions, change defaults, and provide explicit Regional operational and route filters. The names defined for each segment are used in the `segmentActions` and `attachmentPolicies` section. Each segment is created, and operates, as a completely separated routing domain. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
     */
    segments: pulumi.Input<pulumi.Input<inputs.networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs>[]>;
    version?: pulumi.Input<string>;
}
