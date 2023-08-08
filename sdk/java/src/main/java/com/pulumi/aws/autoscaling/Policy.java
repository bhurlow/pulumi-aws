// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.autoscaling;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.autoscaling.PolicyArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyState;
import com.pulumi.aws.autoscaling.outputs.PolicyPredictiveScalingConfiguration;
import com.pulumi.aws.autoscaling.outputs.PolicyStepAdjustment;
import com.pulumi.aws.autoscaling.outputs.PolicyTargetTrackingConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an AutoScaling Scaling Policy resource.
 * 
 * &gt; **NOTE:** You may want to omit `desired_capacity` attribute from attached `aws.autoscaling.Group`
 * when using autoscaling policies. It&#39;s good practice to pick either
 * [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
 * or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
 * (policy-based) scaling.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.autoscaling.Group;
 * import com.pulumi.aws.autoscaling.GroupArgs;
 * import com.pulumi.aws.autoscaling.Policy;
 * import com.pulumi.aws.autoscaling.PolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var bar = new Group(&#34;bar&#34;, GroupArgs.builder()        
 *             .availabilityZones(&#34;us-east-1a&#34;)
 *             .maxSize(5)
 *             .minSize(2)
 *             .healthCheckGracePeriod(300)
 *             .healthCheckType(&#34;ELB&#34;)
 *             .forceDelete(true)
 *             .launchConfiguration(aws_launch_configuration.foo().name())
 *             .build());
 * 
 *         var bat = new Policy(&#34;bat&#34;, PolicyArgs.builder()        
 *             .scalingAdjustment(4)
 *             .adjustmentType(&#34;ChangeInCapacity&#34;)
 *             .cooldown(300)
 *             .autoscalingGroupName(bar.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create target tracking scaling policy using metric math
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.autoscaling.Policy;
 * import com.pulumi.aws.autoscaling.PolicyArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Policy(&#34;example&#34;, PolicyArgs.builder()        
 *             .autoscalingGroupName(&#34;my-test-asg&#34;)
 *             .policyType(&#34;TargetTrackingScaling&#34;)
 *             .targetTrackingConfiguration(PolicyTargetTrackingConfigurationArgs.builder()
 *                 .customizedMetricSpecification(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs.builder()
 *                     .metrics(                    
 *                         PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs.builder()
 *                             .id(&#34;m1&#34;)
 *                             .label(&#34;Get the queue size (the number of messages waiting to be processed)&#34;)
 *                             .metricStat(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs.builder()
 *                                 .metric(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs.builder()
 *                                     .dimensions(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs.builder()
 *                                         .name(&#34;QueueName&#34;)
 *                                         .value(&#34;my-queue&#34;)
 *                                         .build())
 *                                     .metricName(&#34;ApproximateNumberOfMessagesVisible&#34;)
 *                                     .namespace(&#34;AWS/SQS&#34;)
 *                                     .build())
 *                                 .stat(&#34;Sum&#34;)
 *                                 .build())
 *                             .returnData(false)
 *                             .build(),
 *                         PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs.builder()
 *                             .id(&#34;m2&#34;)
 *                             .label(&#34;Get the group size (the number of InService instances)&#34;)
 *                             .metricStat(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs.builder()
 *                                 .metric(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs.builder()
 *                                     .dimensions(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs.builder()
 *                                         .name(&#34;AutoScalingGroupName&#34;)
 *                                         .value(&#34;my-asg&#34;)
 *                                         .build())
 *                                     .metricName(&#34;GroupInServiceInstances&#34;)
 *                                     .namespace(&#34;AWS/AutoScaling&#34;)
 *                                     .build())
 *                                 .stat(&#34;Average&#34;)
 *                                 .build())
 *                             .returnData(false)
 *                             .build(),
 *                         PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs.builder()
 *                             .expression(&#34;m1 / m2&#34;)
 *                             .id(&#34;e1&#34;)
 *                             .label(&#34;Calculate the backlog per instance&#34;)
 *                             .returnData(true)
 *                             .build())
 *                     .build())
 *                 .targetValue(100)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create predictive scaling policy using customized metrics
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.autoscaling.Policy;
 * import com.pulumi.aws.autoscaling.PolicyArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Policy(&#34;example&#34;, PolicyArgs.builder()        
 *             .autoscalingGroupName(&#34;my-test-asg&#34;)
 *             .policyType(&#34;PredictiveScaling&#34;)
 *             .predictiveScalingConfiguration(PolicyPredictiveScalingConfigurationArgs.builder()
 *                 .metricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationArgs.builder()
 *                     .customizedCapacityMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs.builder()
 *                         .metricDataQueries(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs.builder()
 *                             .expression(&#34;SUM(SEARCH(&#39;{AWS/AutoScaling,AutoScalingGroupName} MetricName=\&#34;GroupInServiceIntances\&#34; my-test-asg&#39;, &#39;Average&#39;, 300))&#34;)
 *                             .id(&#34;capacity_sum&#34;)
 *                             .build())
 *                         .build())
 *                     .customizedLoadMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs.builder()
 *                         .metricDataQueries(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryArgs.builder()
 *                             .expression(&#34;SUM(SEARCH(&#39;{AWS/EC2,AutoScalingGroupName} MetricName=\&#34;CPUUtilization\&#34; my-test-asg&#39;, &#39;Sum&#39;, 3600))&#34;)
 *                             .id(&#34;load_sum&#34;)
 *                             .build())
 *                         .build())
 *                     .customizedScalingMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs.builder()
 *                         .metricDataQueries(                        
 *                             PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
 *                                 .expression(&#34;SUM(SEARCH(&#39;{AWS/AutoScaling,AutoScalingGroupName} MetricName=\&#34;GroupInServiceIntances\&#34; my-test-asg&#39;, &#39;Average&#39;, 300))&#34;)
 *                                 .id(&#34;capacity_sum&#34;)
 *                                 .returnData(false)
 *                                 .build(),
 *                             PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
 *                                 .expression(&#34;SUM(SEARCH(&#39;{AWS/EC2,AutoScalingGroupName} MetricName=\&#34;CPUUtilization\&#34; my-test-asg&#39;, &#39;Sum&#39;, 300))&#34;)
 *                                 .id(&#34;load_sum&#34;)
 *                                 .returnData(false)
 *                                 .build(),
 *                             PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
 *                                 .expression(&#34;load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)&#34;)
 *                                 .id(&#34;weighted_average&#34;)
 *                                 .build())
 *                         .build())
 *                     .targetValue(10)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create predictive scaling policy using customized scaling and predefined load metric
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.autoscaling.Policy;
 * import com.pulumi.aws.autoscaling.PolicyArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs;
 * import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Policy(&#34;example&#34;, PolicyArgs.builder()        
 *             .autoscalingGroupName(&#34;my-test-asg&#34;)
 *             .policyType(&#34;PredictiveScaling&#34;)
 *             .predictiveScalingConfiguration(PolicyPredictiveScalingConfigurationArgs.builder()
 *                 .metricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationArgs.builder()
 *                     .customizedScalingMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs.builder()
 *                         .metricDataQueries(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
 *                             .id(&#34;scaling&#34;)
 *                             .metricStat(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatArgs.builder()
 *                                 .metric(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricArgs.builder()
 *                                     .dimensions(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArgs.builder()
 *                                         .name(&#34;AutoScalingGroupName&#34;)
 *                                         .value(&#34;my-test-asg&#34;)
 *                                         .build())
 *                                     .metricName(&#34;CPUUtilization&#34;)
 *                                     .namespace(&#34;AWS/EC2&#34;)
 *                                     .build())
 *                                 .stat(&#34;Average&#34;)
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .predefinedLoadMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs.builder()
 *                         .predefinedMetricType(&#34;ASGTotalCPUUtilization&#34;)
 *                         .resourceLabel(&#34;testLabel&#34;)
 *                         .build())
 *                     .targetValue(10)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_autoscaling_policy.test-policy
 * 
 *  id = &#34;asg-name/policy-name&#34; } Using `pulumi import`, import AutoScaling scaling policy using the role autoscaling_group_name and name separated by `/`. For exampleconsole % pulumi import aws_autoscaling_policy.test-policy asg-name/policy-name
 * 
 */
@ResourceType(type="aws:autoscaling/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     * 
     */
    @Export(name="adjustmentType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> adjustmentType;

    /**
     * @return Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     * 
     */
    public Output<Optional<String>> adjustmentType() {
        return Codegen.optional(this.adjustmentType);
    }
    /**
     * ARN assigned by AWS to the scaling policy.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN assigned by AWS to the scaling policy.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Name of the autoscaling group.
     * 
     */
    @Export(name="autoscalingGroupName", refs={String.class}, tree="[0]")
    private Output<String> autoscalingGroupName;

    /**
     * @return Name of the autoscaling group.
     * 
     */
    public Output<String> autoscalingGroupName() {
        return this.autoscalingGroupName;
    }
    /**
     * Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     * 
     */
    @Export(name="cooldown", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cooldown;

    /**
     * @return Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     * 
     */
    public Output<Optional<Integer>> cooldown() {
        return Codegen.optional(this.cooldown);
    }
    /**
     * Whether the scaling policy is enabled or disabled. Default: `true`.
     * 
     * The following argument is only available to &#34;SimpleScaling&#34; and &#34;StepScaling&#34; type policies:
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether the scaling policy is enabled or disabled. Default: `true`.
     * 
     * The following argument is only available to &#34;SimpleScaling&#34; and &#34;StepScaling&#34; type policies:
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group&#39;s specified cooldown period.
     * 
     */
    @Export(name="estimatedInstanceWarmup", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> estimatedInstanceWarmup;

    /**
     * @return Estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group&#39;s specified cooldown period.
     * 
     */
    public Output<Optional<Integer>> estimatedInstanceWarmup() {
        return Codegen.optional(this.estimatedInstanceWarmup);
    }
    /**
     * Aggregation type for the policy&#39;s metrics. Valid values are &#34;Minimum&#34;, &#34;Maximum&#34;, and &#34;Average&#34;. Without a value, AWS will treat the aggregation type as &#34;Average&#34;.
     * 
     */
    @Export(name="metricAggregationType", refs={String.class}, tree="[0]")
    private Output<String> metricAggregationType;

    /**
     * @return Aggregation type for the policy&#39;s metrics. Valid values are &#34;Minimum&#34;, &#34;Maximum&#34;, and &#34;Average&#34;. Without a value, AWS will treat the aggregation type as &#34;Average&#34;.
     * 
     */
    public Output<String> metricAggregationType() {
        return this.metricAggregationType;
    }
    /**
     * Minimum value to scale by when `adjustment_type` is set to `PercentChangeInCapacity`.
     * 
     * The following arguments are only available to &#34;SimpleScaling&#34; type policies:
     * 
     */
    @Export(name="minAdjustmentMagnitude", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minAdjustmentMagnitude;

    /**
     * @return Minimum value to scale by when `adjustment_type` is set to `PercentChangeInCapacity`.
     * 
     * The following arguments are only available to &#34;SimpleScaling&#34; type policies:
     * 
     */
    public Output<Optional<Integer>> minAdjustmentMagnitude() {
        return Codegen.optional(this.minAdjustmentMagnitude);
    }
    /**
     * Name of the policy.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Policy type, either &#34;SimpleScaling&#34;, &#34;StepScaling&#34;, &#34;TargetTrackingScaling&#34;, or &#34;PredictiveScaling&#34;. If this value isn&#39;t provided, AWS will default to &#34;SimpleScaling.&#34;
     * 
     */
    @Export(name="policyType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> policyType;

    /**
     * @return Policy type, either &#34;SimpleScaling&#34;, &#34;StepScaling&#34;, &#34;TargetTrackingScaling&#34;, or &#34;PredictiveScaling&#34;. If this value isn&#39;t provided, AWS will default to &#34;SimpleScaling.&#34;
     * 
     */
    public Output<Optional<String>> policyType() {
        return Codegen.optional(this.policyType);
    }
    /**
     * Predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
     * 
     */
    @Export(name="predictiveScalingConfiguration", refs={PolicyPredictiveScalingConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ PolicyPredictiveScalingConfiguration> predictiveScalingConfiguration;

    /**
     * @return Predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
     * 
     */
    public Output<Optional<PolicyPredictiveScalingConfiguration>> predictiveScalingConfiguration() {
        return Codegen.optional(this.predictiveScalingConfiguration);
    }
    /**
     * Number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     * 
     */
    @Export(name="scalingAdjustment", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> scalingAdjustment;

    /**
     * @return Number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     * 
     */
    public Output<Optional<Integer>> scalingAdjustment() {
        return Codegen.optional(this.scalingAdjustment);
    }
    /**
     * Set of adjustments that manage
     * group scaling. These have the following structure:
     * 
     * The following fields are available in step adjustments:
     * 
     */
    @Export(name="stepAdjustments", refs={List.class,PolicyStepAdjustment.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PolicyStepAdjustment>> stepAdjustments;

    /**
     * @return Set of adjustments that manage
     * group scaling. These have the following structure:
     * 
     * The following fields are available in step adjustments:
     * 
     */
    public Output<Optional<List<PolicyStepAdjustment>>> stepAdjustments() {
        return Codegen.optional(this.stepAdjustments);
    }
    /**
     * Target tracking policy. These have the following structure:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.autoscaling.Policy;
     * import com.pulumi.aws.autoscaling.PolicyArgs;
     * import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationArgs;
     * import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var example = new Policy(&#34;example&#34;, PolicyArgs.builder()        
     *             .targetTrackingConfiguration(PolicyTargetTrackingConfigurationArgs.builder()
     *                 .predefinedMetricSpecification(PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs.builder()
     *                     .predefinedMetricType(&#34;ASGAverageCPUUtilization&#34;)
     *                     .build())
     *                 .targetValue(40)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     * The following fields are available in target tracking configuration:
     * 
     */
    @Export(name="targetTrackingConfiguration", refs={PolicyTargetTrackingConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ PolicyTargetTrackingConfiguration> targetTrackingConfiguration;

    /**
     * @return Target tracking policy. These have the following structure:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.autoscaling.Policy;
     * import com.pulumi.aws.autoscaling.PolicyArgs;
     * import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationArgs;
     * import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var example = new Policy(&#34;example&#34;, PolicyArgs.builder()        
     *             .targetTrackingConfiguration(PolicyTargetTrackingConfigurationArgs.builder()
     *                 .predefinedMetricSpecification(PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs.builder()
     *                     .predefinedMetricType(&#34;ASGAverageCPUUtilization&#34;)
     *                     .build())
     *                 .targetValue(40)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     * The following fields are available in target tracking configuration:
     * 
     */
    public Output<Optional<PolicyTargetTrackingConfiguration>> targetTrackingConfiguration() {
        return Codegen.optional(this.targetTrackingConfiguration);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Policy(String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(String name, PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(String name, PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:autoscaling/policy:Policy", name, args == null ? PolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Policy(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:autoscaling/policy:Policy", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Policy get(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}
