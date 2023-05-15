// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CE Anomaly Subscription.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testAnomalyMonitor = new aws.costexplorer.AnomalyMonitor("testAnomalyMonitor", {
 *     monitorType: "DIMENSIONAL",
 *     monitorDimension: "SERVICE",
 * });
 * const testAnomalySubscription = new aws.costexplorer.AnomalySubscription("testAnomalySubscription", {
 *     threshold: 100,
 *     frequency: "DAILY",
 *     monitorArnLists: [testAnomalyMonitor.arn],
 *     subscribers: [{
 *         type: "EMAIL",
 *         address: "abc@example.com",
 *     }],
 * });
 * ```
 * ### Threshold Expression
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.costexplorer.AnomalySubscription("test", {
 *     frequency: "DAILY",
 *     monitorArnLists: [aws_ce_anomaly_monitor.test.arn],
 *     subscribers: [{
 *         type: "EMAIL",
 *         address: "abc@example.com",
 *     }],
 *     thresholdExpression: {
 *         dimension: {
 *             key: "ANOMALY_TOTAL_IMPACT_ABSOLUTE",
 *             values: ["100.0"],
 *             matchOptions: ["GREATER_THAN_OR_EQUAL"],
 *         },
 *     },
 * });
 * ```
 * ### SNS Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const costAnomalyUpdates = new aws.sns.Topic("costAnomalyUpdates", {});
 * const snsTopicPolicy = pulumi.all([costAnomalyUpdates.arn, costAnomalyUpdates.arn]).apply(([costAnomalyUpdatesArn, costAnomalyUpdatesArn1]) => aws.iam.getPolicyDocumentOutput({
 *     policyId: "__default_policy_ID",
 *     statements: [
 *         {
 *             sid: "AWSAnomalyDetectionSNSPublishingPermissions",
 *             actions: ["SNS:Publish"],
 *             effect: "Allow",
 *             principals: [{
 *                 type: "Service",
 *                 identifiers: ["costalerts.amazonaws.com"],
 *             }],
 *             resources: [costAnomalyUpdatesArn],
 *         },
 *         {
 *             sid: "__default_statement_ID",
 *             actions: [
 *                 "SNS:Subscribe",
 *                 "SNS:SetTopicAttributes",
 *                 "SNS:RemovePermission",
 *                 "SNS:Receive",
 *                 "SNS:Publish",
 *                 "SNS:ListSubscriptionsByTopic",
 *                 "SNS:GetTopicAttributes",
 *                 "SNS:DeleteTopic",
 *                 "SNS:AddPermission",
 *             ],
 *             conditions: [{
 *                 test: "StringEquals",
 *                 variable: "AWS:SourceOwner",
 *                 values: [_var["account-id"]],
 *             }],
 *             effect: "Allow",
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: ["*"],
 *             }],
 *             resources: [costAnomalyUpdatesArn1],
 *         },
 *     ],
 * }));
 * const _default = new aws.sns.TopicPolicy("default", {
 *     arn: costAnomalyUpdates.arn,
 *     policy: snsTopicPolicy.apply(snsTopicPolicy => snsTopicPolicy.json),
 * });
 * const anomalyMonitor = new aws.costexplorer.AnomalyMonitor("anomalyMonitor", {
 *     monitorType: "DIMENSIONAL",
 *     monitorDimension: "SERVICE",
 * });
 * const realtimeSubscription = new aws.costexplorer.AnomalySubscription("realtimeSubscription", {
 *     threshold: 0,
 *     frequency: "IMMEDIATE",
 *     monitorArnLists: [anomalyMonitor.arn],
 *     subscribers: [{
 *         type: "SNS",
 *         address: costAnomalyUpdates.arn,
 *     }],
 * }, {
 *     dependsOn: [_default],
 * });
 * ```
 *
 * ## Import
 *
 * `aws_ce_anomaly_subscription` can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:costexplorer/anomalySubscription:AnomalySubscription example AnomalySubscriptionARN
 * ```
 */
export class AnomalySubscription extends pulumi.CustomResource {
    /**
     * Get an existing AnomalySubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnomalySubscriptionState, opts?: pulumi.CustomResourceOptions): AnomalySubscription {
        return new AnomalySubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:costexplorer/anomalySubscription:AnomalySubscription';

    /**
     * Returns true if the given object is an instance of AnomalySubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnomalySubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnomalySubscription.__pulumiType;
    }

    /**
     * The unique identifier for the AWS account in which the anomaly subscription ought to be created.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * ARN of the anomaly subscription.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
     */
    public readonly frequency!: pulumi.Output<string>;
    /**
     * A list of cost anomaly monitors.
     */
    public readonly monitorArnLists!: pulumi.Output<string[]>;
    /**
     * The name for the subscription.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A subscriber configuration. Multiple subscribers can be defined.
     */
    public readonly subscribers!: pulumi.Output<outputs.costexplorer.AnomalySubscriptionSubscriber[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The dollar value that triggers a notification if the threshold is exceeded. Depracated, use `thresholdExpression` instead.
     *
     * @deprecated use threshold_expression instead
     */
    public readonly threshold!: pulumi.Output<number>;
    /**
     * An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
     */
    public readonly thresholdExpression!: pulumi.Output<outputs.costexplorer.AnomalySubscriptionThresholdExpression>;

    /**
     * Create a AnomalySubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnomalySubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnomalySubscriptionArgs | AnomalySubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AnomalySubscriptionState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["monitorArnLists"] = state ? state.monitorArnLists : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["subscribers"] = state ? state.subscribers : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["threshold"] = state ? state.threshold : undefined;
            resourceInputs["thresholdExpression"] = state ? state.thresholdExpression : undefined;
        } else {
            const args = argsOrState as AnomalySubscriptionArgs | undefined;
            if ((!args || args.frequency === undefined) && !opts.urn) {
                throw new Error("Missing required property 'frequency'");
            }
            if ((!args || args.monitorArnLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorArnLists'");
            }
            if ((!args || args.subscribers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscribers'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["frequency"] = args ? args.frequency : undefined;
            resourceInputs["monitorArnLists"] = args ? args.monitorArnLists : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["subscribers"] = args ? args.subscribers : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["threshold"] = args ? args.threshold : undefined;
            resourceInputs["thresholdExpression"] = args ? args.thresholdExpression : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AnomalySubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnomalySubscription resources.
 */
export interface AnomalySubscriptionState {
    /**
     * The unique identifier for the AWS account in which the anomaly subscription ought to be created.
     */
    accountId?: pulumi.Input<string>;
    /**
     * ARN of the anomaly subscription.
     */
    arn?: pulumi.Input<string>;
    /**
     * The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
     */
    frequency?: pulumi.Input<string>;
    /**
     * A list of cost anomaly monitors.
     */
    monitorArnLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name for the subscription.
     */
    name?: pulumi.Input<string>;
    /**
     * A subscriber configuration. Multiple subscribers can be defined.
     */
    subscribers?: pulumi.Input<pulumi.Input<inputs.costexplorer.AnomalySubscriptionSubscriber>[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The dollar value that triggers a notification if the threshold is exceeded. Depracated, use `thresholdExpression` instead.
     *
     * @deprecated use threshold_expression instead
     */
    threshold?: pulumi.Input<number>;
    /**
     * An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
     */
    thresholdExpression?: pulumi.Input<inputs.costexplorer.AnomalySubscriptionThresholdExpression>;
}

/**
 * The set of arguments for constructing a AnomalySubscription resource.
 */
export interface AnomalySubscriptionArgs {
    /**
     * The unique identifier for the AWS account in which the anomaly subscription ought to be created.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
     */
    frequency: pulumi.Input<string>;
    /**
     * A list of cost anomaly monitors.
     */
    monitorArnLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name for the subscription.
     */
    name?: pulumi.Input<string>;
    /**
     * A subscriber configuration. Multiple subscribers can be defined.
     */
    subscribers: pulumi.Input<pulumi.Input<inputs.costexplorer.AnomalySubscriptionSubscriber>[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The dollar value that triggers a notification if the threshold is exceeded. Depracated, use `thresholdExpression` instead.
     *
     * @deprecated use threshold_expression instead
     */
    threshold?: pulumi.Input<number>;
    /**
     * An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
     */
    thresholdExpression?: pulumi.Input<inputs.costexplorer.AnomalySubscriptionThresholdExpression>;
}
