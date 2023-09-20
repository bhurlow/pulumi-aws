// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a Glue Data Quality Ruleset Resource. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/glue-data-quality.html) for a full explanation of the Glue Data Quality Ruleset functionality
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.DataQualityRuleset("example", {ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]"});
 * ```
 * ### With description
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.DataQualityRuleset("example", {
 *     description: "example",
 *     ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
 * });
 * ```
 * ### With tags
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.DataQualityRuleset("example", {
 *     ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
 *     tags: {
 *         hello: "world",
 *     },
 * });
 * ```
 * ### With targetTable
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.DataQualityRuleset("example", {
 *     ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
 *     targetTable: {
 *         databaseName: aws_glue_catalog_database.example.name,
 *         tableName: aws_glue_catalog_table.example.name,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Glue Data Quality Ruleset using the `name`. For example:
 *
 * ```sh
 *  $ pulumi import aws:glue/dataQualityRuleset:DataQualityRuleset example exampleName
 * ```
 */
export class DataQualityRuleset extends pulumi.CustomResource {
    /**
     * Get an existing DataQualityRuleset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataQualityRulesetState, opts?: pulumi.CustomResourceOptions): DataQualityRuleset {
        return new DataQualityRuleset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/dataQualityRuleset:DataQualityRuleset';

    /**
     * Returns true if the given object is an instance of DataQualityRuleset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataQualityRuleset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataQualityRuleset.__pulumiType;
    }

    /**
     * ARN of the Glue Data Quality Ruleset.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The time and date that this data quality ruleset was created.
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * Description of the data quality ruleset.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The time and date that this data quality ruleset was created.
     */
    public /*out*/ readonly lastModifiedOn!: pulumi.Output<string>;
    /**
     * Name of the data quality ruleset.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When a ruleset was created from a recommendation run, this run ID is generated to link the two together.
     */
    public /*out*/ readonly recommendationRunId!: pulumi.Output<string>;
    /**
     * A Data Quality Definition Language (DQDL) ruleset. For more information, see the AWS Glue developer guide.
     */
    public readonly ruleset!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * A Configuration block specifying a target table associated with the data quality ruleset. See `targetTable` below.
     */
    public readonly targetTable!: pulumi.Output<outputs.glue.DataQualityRulesetTargetTable | undefined>;

    /**
     * Create a DataQualityRuleset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataQualityRulesetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataQualityRulesetArgs | DataQualityRulesetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataQualityRulesetState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["createdOn"] = state ? state.createdOn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lastModifiedOn"] = state ? state.lastModifiedOn : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["recommendationRunId"] = state ? state.recommendationRunId : undefined;
            resourceInputs["ruleset"] = state ? state.ruleset : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["targetTable"] = state ? state.targetTable : undefined;
        } else {
            const args = argsOrState as DataQualityRulesetArgs | undefined;
            if ((!args || args.ruleset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleset'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ruleset"] = args ? args.ruleset : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetTable"] = args ? args.targetTable : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdOn"] = undefined /*out*/;
            resourceInputs["lastModifiedOn"] = undefined /*out*/;
            resourceInputs["recommendationRunId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DataQualityRuleset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataQualityRuleset resources.
 */
export interface DataQualityRulesetState {
    /**
     * ARN of the Glue Data Quality Ruleset.
     */
    arn?: pulumi.Input<string>;
    /**
     * The time and date that this data quality ruleset was created.
     */
    createdOn?: pulumi.Input<string>;
    /**
     * Description of the data quality ruleset.
     */
    description?: pulumi.Input<string>;
    /**
     * The time and date that this data quality ruleset was created.
     */
    lastModifiedOn?: pulumi.Input<string>;
    /**
     * Name of the data quality ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * When a ruleset was created from a recommendation run, this run ID is generated to link the two together.
     */
    recommendationRunId?: pulumi.Input<string>;
    /**
     * A Data Quality Definition Language (DQDL) ruleset. For more information, see the AWS Glue developer guide.
     */
    ruleset?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A Configuration block specifying a target table associated with the data quality ruleset. See `targetTable` below.
     */
    targetTable?: pulumi.Input<inputs.glue.DataQualityRulesetTargetTable>;
}

/**
 * The set of arguments for constructing a DataQualityRuleset resource.
 */
export interface DataQualityRulesetArgs {
    /**
     * Description of the data quality ruleset.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the data quality ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * A Data Quality Definition Language (DQDL) ruleset. For more information, see the AWS Glue developer guide.
     */
    ruleset: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A Configuration block specifying a target table associated with the data quality ruleset. See `targetTable` below.
     */
    targetTable?: pulumi.Input<inputs.glue.DataQualityRulesetTargetTable>;
}
