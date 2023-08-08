// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS RBin Rule.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.rbin.Rule("example", {
 *     description: "example_rule",
 *     resourceTags: [{
 *         resourceTagKey: "tag_key",
 *         resourceTagValue: "tag_value",
 *     }],
 *     resourceType: "EBS_SNAPSHOT",
 *     retentionPeriod: {
 *         retentionPeriodUnit: "DAYS",
 *         retentionPeriodValue: 10,
 *     },
 *     tags: {
 *         test_tag_key: "test_tag_value",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_rbin_rule.example
 *
 *  id = "examplerule" } Using `pulumi import`, import RBin Rule using the `id`. For exampleconsole % pulumi import aws_rbin_rule.example examplerule
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rbin/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The retention rule description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Information about the retention rule lock configuration. See `lockConfiguration` below.
     */
    public readonly lockConfiguration!: pulumi.Output<outputs.rbin.RuleLockConfiguration | undefined>;
    /**
     * (Timestamp) The date and time at which the unlock delay is set to expire. Only returned for retention rules that have been unlocked and that are still within the unlock delay period.
     */
    public /*out*/ readonly lockEndTime!: pulumi.Output<string>;
    /**
     * (Optional) The lock state of the retention rules to list. Only retention rules with the specified lock state are returned. Valid values are `locked`, `pendingUnlock`, `unlocked`.
     */
    public /*out*/ readonly lockState!: pulumi.Output<string>;
    /**
     * Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
     */
    public readonly resourceTags!: pulumi.Output<outputs.rbin.RuleResourceTag[]>;
    /**
     * The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
     *
     * The following arguments are optional:
     */
    public readonly retentionPeriod!: pulumi.Output<outputs.rbin.RuleRetentionPeriod>;
    /**
     * (String) The state of the retention rule. Only retention rules that are in the `available` state retain resources. Valid values include `pending` and `available`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lockConfiguration"] = state ? state.lockConfiguration : undefined;
            resourceInputs["lockEndTime"] = state ? state.lockEndTime : undefined;
            resourceInputs["lockState"] = state ? state.lockState : undefined;
            resourceInputs["resourceTags"] = state ? state.resourceTags : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.retentionPeriod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retentionPeriod'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["lockConfiguration"] = args ? args.lockConfiguration : undefined;
            resourceInputs["resourceTags"] = args ? args.resourceTags : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["lockEndTime"] = undefined /*out*/;
            resourceInputs["lockState"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    arn?: pulumi.Input<string>;
    /**
     * The retention rule description.
     */
    description?: pulumi.Input<string>;
    /**
     * Information about the retention rule lock configuration. See `lockConfiguration` below.
     */
    lockConfiguration?: pulumi.Input<inputs.rbin.RuleLockConfiguration>;
    /**
     * (Timestamp) The date and time at which the unlock delay is set to expire. Only returned for retention rules that have been unlocked and that are still within the unlock delay period.
     */
    lockEndTime?: pulumi.Input<string>;
    /**
     * (Optional) The lock state of the retention rules to list. Only retention rules with the specified lock state are returned. Valid values are `locked`, `pendingUnlock`, `unlocked`.
     */
    lockState?: pulumi.Input<string>;
    /**
     * Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
     */
    resourceTags?: pulumi.Input<pulumi.Input<inputs.rbin.RuleResourceTag>[]>;
    /**
     * The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
     *
     * The following arguments are optional:
     */
    retentionPeriod?: pulumi.Input<inputs.rbin.RuleRetentionPeriod>;
    /**
     * (String) The state of the retention rule. Only retention rules that are in the `available` state retain resources. Valid values include `pending` and `available`.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * The retention rule description.
     */
    description?: pulumi.Input<string>;
    /**
     * Information about the retention rule lock configuration. See `lockConfiguration` below.
     */
    lockConfiguration?: pulumi.Input<inputs.rbin.RuleLockConfiguration>;
    /**
     * Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
     */
    resourceTags?: pulumi.Input<pulumi.Input<inputs.rbin.RuleResourceTag>[]>;
    /**
     * The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
     *
     * The following arguments are optional:
     */
    retentionPeriod: pulumi.Input<inputs.rbin.RuleRetentionPeriod>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
