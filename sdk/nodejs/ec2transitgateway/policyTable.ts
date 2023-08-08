// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EC2 Transit Gateway Policy Table.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2transitgateway.PolicyTable("example", {
 *     transitGatewayId: aws_ec2_transit_gateway.example.id,
 *     tags: {
 *         Name: "Example Policy Table",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ec2_transit_gateway_policy_table.example
 *
 *  id = "tgw-rtb-12345678" } Using `pulumi import`, import `aws_ec2_transit_gateway_policy_table` using the EC2 Transit Gateway Policy Table identifier. For exampleconsole % pulumi import aws_ec2_transit_gateway_policy_table.example tgw-rtb-12345678
 */
export class PolicyTable extends pulumi.CustomResource {
    /**
     * Get an existing PolicyTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyTableState, opts?: pulumi.CustomResourceOptions): PolicyTable {
        return new PolicyTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2transitgateway/policyTable:PolicyTable';

    /**
     * Returns true if the given object is an instance of PolicyTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyTable.__pulumiType;
    }

    /**
     * EC2 Transit Gateway Policy Table Amazon Resource Name (ARN).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The state of the EC2 Transit Gateway Policy Table.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * EC2 Transit Gateway identifier.
     */
    public readonly transitGatewayId!: pulumi.Output<string>;

    /**
     * Create a PolicyTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyTableArgs | PolicyTableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyTableState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
        } else {
            const args = argsOrState as PolicyTableArgs | undefined;
            if ((!args || args.transitGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayId'");
            }
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicyTable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyTable resources.
 */
export interface PolicyTableState {
    /**
     * EC2 Transit Gateway Policy Table Amazon Resource Name (ARN).
     */
    arn?: pulumi.Input<string>;
    /**
     * The state of the EC2 Transit Gateway Policy Table.
     */
    state?: pulumi.Input<string>;
    /**
     * Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * EC2 Transit Gateway identifier.
     */
    transitGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyTable resource.
 */
export interface PolicyTableArgs {
    /**
     * Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * EC2 Transit Gateway identifier.
     */
    transitGatewayId: pulumi.Input<string>;
}
