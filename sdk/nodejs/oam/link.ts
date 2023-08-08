// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS CloudWatch Observability Access Manager Link.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.oam.Link("example", {
 *     labelTemplate: "$AccountName",
 *     resourceTypes: ["AWS::CloudWatch::Metric"],
 *     sinkIdentifier: aws_oam_sink.test.id,
 *     tags: {
 *         Env: "prod",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_oam_link.example
 *
 *  id = "arn:aws:oam:us-west-2:123456789012:link/link-id" } Using `pulumi import`, import CloudWatch Observability Access Manager Link using the `arn`. For exampleconsole % pulumi import aws_oam_link.example arn:aws:oam:us-west-2:123456789012:link/link-id
 */
export class Link extends pulumi.CustomResource {
    /**
     * Get an existing Link resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkState, opts?: pulumi.CustomResourceOptions): Link {
        return new Link(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:oam/link:Link';

    /**
     * Returns true if the given object is an instance of Link.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Link {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Link.__pulumiType;
    }

    /**
     * ARN of the link.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Label that is assigned to this link.
     */
    public /*out*/ readonly label!: pulumi.Output<string>;
    /**
     * Human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
     */
    public readonly labelTemplate!: pulumi.Output<string>;
    /**
     * ID string that AWS generated as part of the link ARN.
     */
    public /*out*/ readonly linkId!: pulumi.Output<string>;
    /**
     * Types of data that the source account shares with the monitoring account.
     */
    public readonly resourceTypes!: pulumi.Output<string[]>;
    /**
     * ARN of the sink that is used for this link.
     */
    public /*out*/ readonly sinkArn!: pulumi.Output<string>;
    /**
     * Identifier of the sink to use to create this link.
     *
     * The following arguments are optional:
     */
    public readonly sinkIdentifier!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Link resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkArgs | LinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LinkState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["labelTemplate"] = state ? state.labelTemplate : undefined;
            resourceInputs["linkId"] = state ? state.linkId : undefined;
            resourceInputs["resourceTypes"] = state ? state.resourceTypes : undefined;
            resourceInputs["sinkArn"] = state ? state.sinkArn : undefined;
            resourceInputs["sinkIdentifier"] = state ? state.sinkIdentifier : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as LinkArgs | undefined;
            if ((!args || args.labelTemplate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'labelTemplate'");
            }
            if ((!args || args.resourceTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceTypes'");
            }
            if ((!args || args.sinkIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sinkIdentifier'");
            }
            resourceInputs["labelTemplate"] = args ? args.labelTemplate : undefined;
            resourceInputs["resourceTypes"] = args ? args.resourceTypes : undefined;
            resourceInputs["sinkIdentifier"] = args ? args.sinkIdentifier : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["label"] = undefined /*out*/;
            resourceInputs["linkId"] = undefined /*out*/;
            resourceInputs["sinkArn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Link.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Link resources.
 */
export interface LinkState {
    /**
     * ARN of the link.
     */
    arn?: pulumi.Input<string>;
    /**
     * Label that is assigned to this link.
     */
    label?: pulumi.Input<string>;
    /**
     * Human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
     */
    labelTemplate?: pulumi.Input<string>;
    /**
     * ID string that AWS generated as part of the link ARN.
     */
    linkId?: pulumi.Input<string>;
    /**
     * Types of data that the source account shares with the monitoring account.
     */
    resourceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ARN of the sink that is used for this link.
     */
    sinkArn?: pulumi.Input<string>;
    /**
     * Identifier of the sink to use to create this link.
     *
     * The following arguments are optional:
     */
    sinkIdentifier?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Link resource.
 */
export interface LinkArgs {
    /**
     * Human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
     */
    labelTemplate: pulumi.Input<string>;
    /**
     * Types of data that the source account shares with the monitoring account.
     */
    resourceTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier of the sink to use to create this link.
     *
     * The following arguments are optional:
     */
    sinkIdentifier: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
