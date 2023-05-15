// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a WAF Regional Rule Group Resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleRule = new aws.wafregional.Rule("exampleRule", {metricName: "example"});
 * const exampleRuleGroup = new aws.wafregional.RuleGroup("exampleRuleGroup", {
 *     metricName: "example",
 *     activatedRules: [{
 *         action: {
 *             type: "COUNT",
 *         },
 *         priority: 50,
 *         ruleId: exampleRule.id,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * WAF Regional Rule Group can be imported using the id, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:wafregional/ruleGroup:RuleGroup example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
 * ```
 */
export class RuleGroup extends pulumi.CustomResource {
    /**
     * Get an existing RuleGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleGroupState, opts?: pulumi.CustomResourceOptions): RuleGroup {
        return new RuleGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafregional/ruleGroup:RuleGroup';

    /**
     * Returns true if the given object is an instance of RuleGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleGroup.__pulumiType;
    }

    /**
     * A list of activated rules, see below
     */
    public readonly activatedRules!: pulumi.Output<outputs.wafregional.RuleGroupActivatedRule[] | undefined>;
    /**
     * The ARN of the WAF Regional Rule Group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A friendly name for the metrics from the rule group
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * A friendly name of the rule group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a RuleGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleGroupArgs | RuleGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleGroupState | undefined;
            resourceInputs["activatedRules"] = state ? state.activatedRules : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["metricName"] = state ? state.metricName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as RuleGroupArgs | undefined;
            if ((!args || args.metricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricName'");
            }
            resourceInputs["activatedRules"] = args ? args.activatedRules : undefined;
            resourceInputs["metricName"] = args ? args.metricName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleGroup resources.
 */
export interface RuleGroupState {
    /**
     * A list of activated rules, see below
     */
    activatedRules?: pulumi.Input<pulumi.Input<inputs.wafregional.RuleGroupActivatedRule>[]>;
    /**
     * The ARN of the WAF Regional Rule Group.
     */
    arn?: pulumi.Input<string>;
    /**
     * A friendly name for the metrics from the rule group
     */
    metricName?: pulumi.Input<string>;
    /**
     * A friendly name of the rule group
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a RuleGroup resource.
 */
export interface RuleGroupArgs {
    /**
     * A list of activated rules, see below
     */
    activatedRules?: pulumi.Input<pulumi.Input<inputs.wafregional.RuleGroupActivatedRule>[]>;
    /**
     * A friendly name for the metrics from the rule group
     */
    metricName: pulumi.Input<string>;
    /**
     * A friendly name of the rule group
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
