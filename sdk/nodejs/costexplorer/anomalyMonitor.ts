// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CE Anomaly Monitor.
 *
 * ## Example Usage
 *
 * There are two main types of a Cost Anomaly Monitor: `DIMENSIONAL` and `CUSTOM`.
 * ### Dimensional Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const serviceMonitor = new aws.costexplorer.AnomalyMonitor("serviceMonitor", {
 *     monitorDimension: "SERVICE",
 *     monitorType: "DIMENSIONAL",
 * });
 * ```
 * ### Custom Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.costexplorer.AnomalyMonitor("test", {
 *     monitorType: "CUSTOM",
 *     monitorSpecification: JSON.stringify({
 *         And: undefined,
 *         CostCategories: undefined,
 *         Dimensions: undefined,
 *         Not: undefined,
 *         Or: undefined,
 *         Tags: {
 *             Key: "CostCenter",
 *             MatchOptions: undefined,
 *             Values: ["10000"],
 *         },
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * `aws_ce_anomaly_monitor` can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:costexplorer/anomalyMonitor:AnomalyMonitor example costAnomalyMonitorARN
 * ```
 */
export class AnomalyMonitor extends pulumi.CustomResource {
    /**
     * Get an existing AnomalyMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnomalyMonitorState, opts?: pulumi.CustomResourceOptions): AnomalyMonitor {
        return new AnomalyMonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:costexplorer/anomalyMonitor:AnomalyMonitor';

    /**
     * Returns true if the given object is an instance of AnomalyMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnomalyMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnomalyMonitor.__pulumiType;
    }

    /**
     * ARN of the anomaly monitor.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The dimensions to evaluate. Valid values: `SERVICE`.
     */
    public readonly monitorDimension!: pulumi.Output<string | undefined>;
    /**
     * A valid JSON representation for the [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object.
     */
    public readonly monitorSpecification!: pulumi.Output<string | undefined>;
    /**
     * The possible type values. Valid values: `DIMENSIONAL` | `CUSTOM`.
     */
    public readonly monitorType!: pulumi.Output<string>;
    /**
     * The name of the monitor.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a AnomalyMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnomalyMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnomalyMonitorArgs | AnomalyMonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AnomalyMonitorState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["monitorDimension"] = state ? state.monitorDimension : undefined;
            resourceInputs["monitorSpecification"] = state ? state.monitorSpecification : undefined;
            resourceInputs["monitorType"] = state ? state.monitorType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as AnomalyMonitorArgs | undefined;
            if ((!args || args.monitorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorType'");
            }
            resourceInputs["monitorDimension"] = args ? args.monitorDimension : undefined;
            resourceInputs["monitorSpecification"] = args ? args.monitorSpecification : undefined;
            resourceInputs["monitorType"] = args ? args.monitorType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AnomalyMonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnomalyMonitor resources.
 */
export interface AnomalyMonitorState {
    /**
     * ARN of the anomaly monitor.
     */
    arn?: pulumi.Input<string>;
    /**
     * The dimensions to evaluate. Valid values: `SERVICE`.
     */
    monitorDimension?: pulumi.Input<string>;
    /**
     * A valid JSON representation for the [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object.
     */
    monitorSpecification?: pulumi.Input<string>;
    /**
     * The possible type values. Valid values: `DIMENSIONAL` | `CUSTOM`.
     */
    monitorType?: pulumi.Input<string>;
    /**
     * The name of the monitor.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a AnomalyMonitor resource.
 */
export interface AnomalyMonitorArgs {
    /**
     * The dimensions to evaluate. Valid values: `SERVICE`.
     */
    monitorDimension?: pulumi.Input<string>;
    /**
     * A valid JSON representation for the [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object.
     */
    monitorSpecification?: pulumi.Input<string>;
    /**
     * The possible type values. Valid values: `DIMENSIONAL` | `CUSTOM`.
     */
    monitorType: pulumi.Input<string>;
    /**
     * The name of the monitor.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
