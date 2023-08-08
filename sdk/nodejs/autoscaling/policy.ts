// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an AutoScaling Scaling Policy resource.
 *
 * > **NOTE:** You may want to omit `desiredCapacity` attribute from attached `aws.autoscaling.Group`
 * when using autoscaling policies. It's good practice to pick either
 * [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
 * or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
 * (policy-based) scaling.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bar = new aws.autoscaling.Group("bar", {
 *     availabilityZones: ["us-east-1a"],
 *     maxSize: 5,
 *     minSize: 2,
 *     healthCheckGracePeriod: 300,
 *     healthCheckType: "ELB",
 *     forceDelete: true,
 *     launchConfiguration: aws_launch_configuration.foo.name,
 * });
 * const bat = new aws.autoscaling.Policy("bat", {
 *     scalingAdjustment: 4,
 *     adjustmentType: "ChangeInCapacity",
 *     cooldown: 300,
 *     autoscalingGroupName: bar.name,
 * });
 * ```
 * ### Create target tracking scaling policy using metric math
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.autoscaling.Policy("example", {
 *     autoscalingGroupName: "my-test-asg",
 *     policyType: "TargetTrackingScaling",
 *     targetTrackingConfiguration: {
 *         customizedMetricSpecification: {
 *             metrics: [
 *                 {
 *                     id: "m1",
 *                     label: "Get the queue size (the number of messages waiting to be processed)",
 *                     metricStat: {
 *                         metric: {
 *                             dimensions: [{
 *                                 name: "QueueName",
 *                                 value: "my-queue",
 *                             }],
 *                             metricName: "ApproximateNumberOfMessagesVisible",
 *                             namespace: "AWS/SQS",
 *                         },
 *                         stat: "Sum",
 *                     },
 *                     returnData: false,
 *                 },
 *                 {
 *                     id: "m2",
 *                     label: "Get the group size (the number of InService instances)",
 *                     metricStat: {
 *                         metric: {
 *                             dimensions: [{
 *                                 name: "AutoScalingGroupName",
 *                                 value: "my-asg",
 *                             }],
 *                             metricName: "GroupInServiceInstances",
 *                             namespace: "AWS/AutoScaling",
 *                         },
 *                         stat: "Average",
 *                     },
 *                     returnData: false,
 *                 },
 *                 {
 *                     expression: "m1 / m2",
 *                     id: "e1",
 *                     label: "Calculate the backlog per instance",
 *                     returnData: true,
 *                 },
 *             ],
 *         },
 *         targetValue: 100,
 *     },
 * });
 * ```
 * ### Create predictive scaling policy using customized metrics
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.autoscaling.Policy("example", {
 *     autoscalingGroupName: "my-test-asg",
 *     policyType: "PredictiveScaling",
 *     predictiveScalingConfiguration: {
 *         metricSpecification: {
 *             customizedCapacityMetricSpecification: {
 *                 metricDataQueries: [{
 *                     expression: "SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
 *                     id: "capacity_sum",
 *                 }],
 *             },
 *             customizedLoadMetricSpecification: {
 *                 metricDataQueries: [{
 *                     expression: "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 3600))",
 *                     id: "load_sum",
 *                 }],
 *             },
 *             customizedScalingMetricSpecification: {
 *                 metricDataQueries: [
 *                     {
 *                         expression: "SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
 *                         id: "capacity_sum",
 *                         returnData: false,
 *                     },
 *                     {
 *                         expression: "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 300))",
 *                         id: "load_sum",
 *                         returnData: false,
 *                     },
 *                     {
 *                         expression: "load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)",
 *                         id: "weighted_average",
 *                     },
 *                 ],
 *             },
 *             targetValue: 10,
 *         },
 *     },
 * });
 * ```
 * ### Create predictive scaling policy using customized scaling and predefined load metric
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.autoscaling.Policy("example", {
 *     autoscalingGroupName: "my-test-asg",
 *     policyType: "PredictiveScaling",
 *     predictiveScalingConfiguration: {
 *         metricSpecification: {
 *             customizedScalingMetricSpecification: {
 *                 metricDataQueries: [{
 *                     id: "scaling",
 *                     metricStat: {
 *                         metric: {
 *                             dimensions: [{
 *                                 name: "AutoScalingGroupName",
 *                                 value: "my-test-asg",
 *                             }],
 *                             metricName: "CPUUtilization",
 *                             namespace: "AWS/EC2",
 *                         },
 *                         stat: "Average",
 *                     },
 *                 }],
 *             },
 *             predefinedLoadMetricSpecification: {
 *                 predefinedMetricType: "ASGTotalCPUUtilization",
 *                 resourceLabel: "testLabel",
 *             },
 *             targetValue: 10,
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_autoscaling_policy.test-policy
 *
 *  id = "asg-name/policy-name" } Using `pulumi import`, import AutoScaling scaling policy using the role autoscaling_group_name and name separated by `/`. For exampleconsole % pulumi import aws_autoscaling_policy.test-policy asg-name/policy-name
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:autoscaling/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     */
    public readonly adjustmentType!: pulumi.Output<string | undefined>;
    /**
     * ARN assigned by AWS to the scaling policy.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Name of the autoscaling group.
     */
    public readonly autoscalingGroupName!: pulumi.Output<string>;
    /**
     * Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     */
    public readonly cooldown!: pulumi.Output<number | undefined>;
    /**
     * Whether the scaling policy is enabled or disabled. Default: `true`.
     *
     * The following argument is only available to "SimpleScaling" and "StepScaling" type policies:
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
     */
    public readonly estimatedInstanceWarmup!: pulumi.Output<number | undefined>;
    /**
     * Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
     */
    public readonly metricAggregationType!: pulumi.Output<string>;
    /**
     * Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
     *
     * The following arguments are only available to "SimpleScaling" type policies:
     */
    public readonly minAdjustmentMagnitude!: pulumi.Output<number | undefined>;
    /**
     * Name of the policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
     */
    public readonly policyType!: pulumi.Output<string | undefined>;
    /**
     * Predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
     */
    public readonly predictiveScalingConfiguration!: pulumi.Output<outputs.autoscaling.PolicyPredictiveScalingConfiguration | undefined>;
    /**
     * Number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     */
    public readonly scalingAdjustment!: pulumi.Output<number | undefined>;
    /**
     * Set of adjustments that manage
     * group scaling. These have the following structure:
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as aws from "@pulumi/aws";
     *
     * const example = new aws.autoscaling.Policy("example", {stepAdjustments: [
     *     {
     *         metricIntervalLowerBound: "1",
     *         metricIntervalUpperBound: "2",
     *         scalingAdjustment: -1,
     *     },
     *     {
     *         metricIntervalLowerBound: "2",
     *         metricIntervalUpperBound: "3",
     *         scalingAdjustment: 1,
     *     },
     * ]});
     * ```
     *
     * The following fields are available in step adjustments:
     */
    public readonly stepAdjustments!: pulumi.Output<outputs.autoscaling.PolicyStepAdjustment[] | undefined>;
    /**
     * Target tracking policy. These have the following structure:
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as aws from "@pulumi/aws";
     *
     * const example = new aws.autoscaling.Policy("example", {targetTrackingConfiguration: {
     *     predefinedMetricSpecification: {
     *         predefinedMetricType: "ASGAverageCPUUtilization",
     *     },
     *     targetValue: 40,
     * }});
     * ```
     *
     * The following fields are available in target tracking configuration:
     */
    public readonly targetTrackingConfiguration!: pulumi.Output<outputs.autoscaling.PolicyTargetTrackingConfiguration | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            resourceInputs["adjustmentType"] = state ? state.adjustmentType : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["autoscalingGroupName"] = state ? state.autoscalingGroupName : undefined;
            resourceInputs["cooldown"] = state ? state.cooldown : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["estimatedInstanceWarmup"] = state ? state.estimatedInstanceWarmup : undefined;
            resourceInputs["metricAggregationType"] = state ? state.metricAggregationType : undefined;
            resourceInputs["minAdjustmentMagnitude"] = state ? state.minAdjustmentMagnitude : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["predictiveScalingConfiguration"] = state ? state.predictiveScalingConfiguration : undefined;
            resourceInputs["scalingAdjustment"] = state ? state.scalingAdjustment : undefined;
            resourceInputs["stepAdjustments"] = state ? state.stepAdjustments : undefined;
            resourceInputs["targetTrackingConfiguration"] = state ? state.targetTrackingConfiguration : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.autoscalingGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscalingGroupName'");
            }
            resourceInputs["adjustmentType"] = args ? args.adjustmentType : undefined;
            resourceInputs["autoscalingGroupName"] = args ? args.autoscalingGroupName : undefined;
            resourceInputs["cooldown"] = args ? args.cooldown : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["estimatedInstanceWarmup"] = args ? args.estimatedInstanceWarmup : undefined;
            resourceInputs["metricAggregationType"] = args ? args.metricAggregationType : undefined;
            resourceInputs["minAdjustmentMagnitude"] = args ? args.minAdjustmentMagnitude : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policyType"] = args ? args.policyType : undefined;
            resourceInputs["predictiveScalingConfiguration"] = args ? args.predictiveScalingConfiguration : undefined;
            resourceInputs["scalingAdjustment"] = args ? args.scalingAdjustment : undefined;
            resourceInputs["stepAdjustments"] = args ? args.stepAdjustments : undefined;
            resourceInputs["targetTrackingConfiguration"] = args ? args.targetTrackingConfiguration : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * ARN assigned by AWS to the scaling policy.
     */
    arn?: pulumi.Input<string>;
    /**
     * Name of the autoscaling group.
     */
    autoscalingGroupName?: pulumi.Input<string>;
    /**
     * Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * Whether the scaling policy is enabled or disabled. Default: `true`.
     *
     * The following argument is only available to "SimpleScaling" and "StepScaling" type policies:
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
     */
    estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
     */
    metricAggregationType?: pulumi.Input<string>;
    /**
     * Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
     *
     * The following arguments are only available to "SimpleScaling" type policies:
     */
    minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * Name of the policy.
     */
    name?: pulumi.Input<string>;
    /**
     * Policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
     */
    policyType?: pulumi.Input<string>;
    /**
     * Predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
     */
    predictiveScalingConfiguration?: pulumi.Input<inputs.autoscaling.PolicyPredictiveScalingConfiguration>;
    /**
     * Number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     */
    scalingAdjustment?: pulumi.Input<number>;
    /**
     * Set of adjustments that manage
     * group scaling. These have the following structure:
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as aws from "@pulumi/aws";
     *
     * const example = new aws.autoscaling.Policy("example", {stepAdjustments: [
     *     {
     *         metricIntervalLowerBound: "1",
     *         metricIntervalUpperBound: "2",
     *         scalingAdjustment: -1,
     *     },
     *     {
     *         metricIntervalLowerBound: "2",
     *         metricIntervalUpperBound: "3",
     *         scalingAdjustment: 1,
     *     },
     * ]});
     * ```
     *
     * The following fields are available in step adjustments:
     */
    stepAdjustments?: pulumi.Input<pulumi.Input<inputs.autoscaling.PolicyStepAdjustment>[]>;
    /**
     * Target tracking policy. These have the following structure:
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as aws from "@pulumi/aws";
     *
     * const example = new aws.autoscaling.Policy("example", {targetTrackingConfiguration: {
     *     predefinedMetricSpecification: {
     *         predefinedMetricType: "ASGAverageCPUUtilization",
     *     },
     *     targetValue: 40,
     * }});
     * ```
     *
     * The following fields are available in target tracking configuration:
     */
    targetTrackingConfiguration?: pulumi.Input<inputs.autoscaling.PolicyTargetTrackingConfiguration>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * Name of the autoscaling group.
     */
    autoscalingGroupName: pulumi.Input<string>;
    /**
     * Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * Whether the scaling policy is enabled or disabled. Default: `true`.
     *
     * The following argument is only available to "SimpleScaling" and "StepScaling" type policies:
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
     */
    estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
     */
    metricAggregationType?: pulumi.Input<string>;
    /**
     * Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
     *
     * The following arguments are only available to "SimpleScaling" type policies:
     */
    minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * Name of the policy.
     */
    name?: pulumi.Input<string>;
    /**
     * Policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
     */
    policyType?: pulumi.Input<string>;
    /**
     * Predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
     */
    predictiveScalingConfiguration?: pulumi.Input<inputs.autoscaling.PolicyPredictiveScalingConfiguration>;
    /**
     * Number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     */
    scalingAdjustment?: pulumi.Input<number>;
    /**
     * Set of adjustments that manage
     * group scaling. These have the following structure:
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as aws from "@pulumi/aws";
     *
     * const example = new aws.autoscaling.Policy("example", {stepAdjustments: [
     *     {
     *         metricIntervalLowerBound: "1",
     *         metricIntervalUpperBound: "2",
     *         scalingAdjustment: -1,
     *     },
     *     {
     *         metricIntervalLowerBound: "2",
     *         metricIntervalUpperBound: "3",
     *         scalingAdjustment: 1,
     *     },
     * ]});
     * ```
     *
     * The following fields are available in step adjustments:
     */
    stepAdjustments?: pulumi.Input<pulumi.Input<inputs.autoscaling.PolicyStepAdjustment>[]>;
    /**
     * Target tracking policy. These have the following structure:
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as aws from "@pulumi/aws";
     *
     * const example = new aws.autoscaling.Policy("example", {targetTrackingConfiguration: {
     *     predefinedMetricSpecification: {
     *         predefinedMetricType: "ASGAverageCPUUtilization",
     *     },
     *     targetValue: 40,
     * }});
     * ```
     *
     * The following fields are available in target tracking configuration:
     */
    targetTrackingConfiguration?: pulumi.Input<inputs.autoscaling.PolicyTargetTrackingConfiguration>;
}
