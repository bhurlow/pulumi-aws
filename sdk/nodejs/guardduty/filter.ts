// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a GuardDuty filter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myFilter = new aws.guardduty.Filter("myFilter", {
 *     action: "ARCHIVE",
 *     detectorId: aws_guardduty_detector.example.id,
 *     rank: 1,
 *     findingCriteria: {
 *         criterions: [
 *             {
 *                 field: "region",
 *                 equals: ["eu-west-1"],
 *             },
 *             {
 *                 field: "service.additionalInfo.threatListName",
 *                 notEquals: [
 *                     "some-threat",
 *                     "another-threat",
 *                 ],
 *             },
 *             {
 *                 field: "updatedAt",
 *                 greaterThan: "2020-01-01T00:00:00Z",
 *                 lessThan: "2020-02-01T00:00:00Z",
 *             },
 *             {
 *                 field: "severity",
 *                 greaterThanOrEqual: "4",
 *             },
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * GuardDuty filters can be imported using the detector ID and filter's name separated by a colon, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:guardduty/filter:Filter MyFilter 00b00fd5aecc0ab60a708659477e9617:MyFilter
 * ```
 */
export class Filter extends pulumi.CustomResource {
    /**
     * Get an existing Filter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FilterState, opts?: pulumi.CustomResourceOptions): Filter {
        return new Filter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:guardduty/filter:Filter';

    /**
     * Returns true if the given object is an instance of Filter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Filter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Filter.__pulumiType;
    }

    /**
     * Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The ARN of the GuardDuty filter.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the filter.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ID of a GuardDuty detector, attached to your account.
     */
    public readonly detectorId!: pulumi.Output<string>;
    /**
     * Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
     */
    public readonly findingCriteria!: pulumi.Output<outputs.guardduty.FilterFindingCriteria>;
    /**
     * The name of your filter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
     */
    public readonly rank!: pulumi.Output<number>;
    /**
     * The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Filter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FilterArgs | FilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FilterState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["detectorId"] = state ? state.detectorId : undefined;
            resourceInputs["findingCriteria"] = state ? state.findingCriteria : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rank"] = state ? state.rank : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as FilterArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.detectorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'detectorId'");
            }
            if ((!args || args.findingCriteria === undefined) && !opts.urn) {
                throw new Error("Missing required property 'findingCriteria'");
            }
            if ((!args || args.rank === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rank'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["detectorId"] = args ? args.detectorId : undefined;
            resourceInputs["findingCriteria"] = args ? args.findingCriteria : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rank"] = args ? args.rank : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Filter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Filter resources.
 */
export interface FilterState {
    /**
     * Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
     */
    action?: pulumi.Input<string>;
    /**
     * The ARN of the GuardDuty filter.
     */
    arn?: pulumi.Input<string>;
    /**
     * Description of the filter.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of a GuardDuty detector, attached to your account.
     */
    detectorId?: pulumi.Input<string>;
    /**
     * Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
     */
    findingCriteria?: pulumi.Input<inputs.guardduty.FilterFindingCriteria>;
    /**
     * The name of your filter.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
     */
    rank?: pulumi.Input<number>;
    /**
     * The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Filter resource.
 */
export interface FilterArgs {
    /**
     * Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
     */
    action: pulumi.Input<string>;
    /**
     * Description of the filter.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of a GuardDuty detector, attached to your account.
     */
    detectorId: pulumi.Input<string>;
    /**
     * Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
     */
    findingCriteria: pulumi.Input<inputs.guardduty.FilterFindingCriteria>;
    /**
     * The name of your filter.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
     */
    rank: pulumi.Input<number>;
    /**
     * The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
