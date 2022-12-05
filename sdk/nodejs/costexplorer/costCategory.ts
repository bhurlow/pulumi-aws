// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CE Cost Category.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.costexplorer.CostCategory("test", {
 *     rules: [
 *         {
 *             rule: {
 *                 dimension: {
 *                     key: "LINKED_ACCOUNT_NAME",
 *                     matchOptions: ["ENDS_WITH"],
 *                     values: ["-prod"],
 *                 },
 *             },
 *             value: "production",
 *         },
 *         {
 *             rule: {
 *                 dimension: {
 *                     key: "LINKED_ACCOUNT_NAME",
 *                     matchOptions: ["ENDS_WITH"],
 *                     values: ["-stg"],
 *                 },
 *             },
 *             value: "staging",
 *         },
 *         {
 *             rule: {
 *                 dimension: {
 *                     key: "LINKED_ACCOUNT_NAME",
 *                     matchOptions: ["ENDS_WITH"],
 *                     values: ["-dev"],
 *                 },
 *             },
 *             value: "testing",
 *         },
 *     ],
 *     ruleVersion: "CostCategoryExpression.v1",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_ce_cost_category` can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import aws:costexplorer/costCategory:CostCategory example costCategoryARN
 * ```
 */
export class CostCategory extends pulumi.CustomResource {
    /**
     * Get an existing CostCategory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CostCategoryState, opts?: pulumi.CustomResourceOptions): CostCategory {
        return new CostCategory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:costexplorer/costCategory:CostCategory';

    /**
     * Returns true if the given object is an instance of CostCategory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CostCategory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CostCategory.__pulumiType;
    }

    /**
     * ARN of the cost category.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Default value for the cost category.
     */
    public readonly defaultValue!: pulumi.Output<string | undefined>;
    /**
     * Effective end data of your Cost Category.
     */
    public /*out*/ readonly effectiveEnd!: pulumi.Output<string>;
    /**
     * The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.
     */
    public readonly effectiveStart!: pulumi.Output<string>;
    /**
     * Unique name for the Cost Category.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Rule schema version in this particular Cost Category.
     */
    public readonly ruleVersion!: pulumi.Output<string>;
    /**
     * Configuration block for the `Expression` object used to categorize costs. See below.
     */
    public readonly rules!: pulumi.Output<outputs.costexplorer.CostCategoryRule[]>;
    /**
     * Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
     */
    public readonly splitChargeRules!: pulumi.Output<outputs.costexplorer.CostCategorySplitChargeRule[] | undefined>;
    /**
     * Configuration block for the specific `Tag` to use for `Expression`. See below.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a CostCategory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CostCategoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CostCategoryArgs | CostCategoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CostCategoryState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["defaultValue"] = state ? state.defaultValue : undefined;
            resourceInputs["effectiveEnd"] = state ? state.effectiveEnd : undefined;
            resourceInputs["effectiveStart"] = state ? state.effectiveStart : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ruleVersion"] = state ? state.ruleVersion : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["splitChargeRules"] = state ? state.splitChargeRules : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as CostCategoryArgs | undefined;
            if ((!args || args.ruleVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleVersion'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            resourceInputs["defaultValue"] = args ? args.defaultValue : undefined;
            resourceInputs["effectiveStart"] = args ? args.effectiveStart : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ruleVersion"] = args ? args.ruleVersion : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["splitChargeRules"] = args ? args.splitChargeRules : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["effectiveEnd"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CostCategory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CostCategory resources.
 */
export interface CostCategoryState {
    /**
     * ARN of the cost category.
     */
    arn?: pulumi.Input<string>;
    /**
     * Default value for the cost category.
     */
    defaultValue?: pulumi.Input<string>;
    /**
     * Effective end data of your Cost Category.
     */
    effectiveEnd?: pulumi.Input<string>;
    /**
     * The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.
     */
    effectiveStart?: pulumi.Input<string>;
    /**
     * Unique name for the Cost Category.
     */
    name?: pulumi.Input<string>;
    /**
     * Rule schema version in this particular Cost Category.
     */
    ruleVersion?: pulumi.Input<string>;
    /**
     * Configuration block for the `Expression` object used to categorize costs. See below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.costexplorer.CostCategoryRule>[]>;
    /**
     * Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
     */
    splitChargeRules?: pulumi.Input<pulumi.Input<inputs.costexplorer.CostCategorySplitChargeRule>[]>;
    /**
     * Configuration block for the specific `Tag` to use for `Expression`. See below.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a CostCategory resource.
 */
export interface CostCategoryArgs {
    /**
     * Default value for the cost category.
     */
    defaultValue?: pulumi.Input<string>;
    /**
     * The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.
     */
    effectiveStart?: pulumi.Input<string>;
    /**
     * Unique name for the Cost Category.
     */
    name?: pulumi.Input<string>;
    /**
     * Rule schema version in this particular Cost Category.
     */
    ruleVersion: pulumi.Input<string>;
    /**
     * Configuration block for the `Expression` object used to categorize costs. See below.
     */
    rules: pulumi.Input<pulumi.Input<inputs.costexplorer.CostCategoryRule>[]>;
    /**
     * Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
     */
    splitChargeRules?: pulumi.Input<pulumi.Input<inputs.costexplorer.CostCategorySplitChargeRule>[]>;
    /**
     * Configuration block for the specific `Tag` to use for `Expression`. See below.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
