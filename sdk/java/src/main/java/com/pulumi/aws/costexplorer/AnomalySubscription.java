// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.costexplorer;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.costexplorer.AnomalySubscriptionArgs;
import com.pulumi.aws.costexplorer.inputs.AnomalySubscriptionState;
import com.pulumi.aws.costexplorer.outputs.AnomalySubscriptionSubscriber;
import com.pulumi.aws.costexplorer.outputs.AnomalySubscriptionThresholdExpression;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CE Anomaly Subscription.
 * 
 * ## Example Usage
 * ### Basic Example
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.costexplorer.AnomalyMonitor;
 * import com.pulumi.aws.costexplorer.AnomalyMonitorArgs;
 * import com.pulumi.aws.costexplorer.AnomalySubscription;
 * import com.pulumi.aws.costexplorer.AnomalySubscriptionArgs;
 * import com.pulumi.aws.costexplorer.inputs.AnomalySubscriptionSubscriberArgs;
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
 *         var testAnomalyMonitor = new AnomalyMonitor(&#34;testAnomalyMonitor&#34;, AnomalyMonitorArgs.builder()        
 *             .monitorType(&#34;DIMENSIONAL&#34;)
 *             .monitorDimension(&#34;SERVICE&#34;)
 *             .build());
 * 
 *         var testAnomalySubscription = new AnomalySubscription(&#34;testAnomalySubscription&#34;, AnomalySubscriptionArgs.builder()        
 *             .threshold(100)
 *             .frequency(&#34;DAILY&#34;)
 *             .monitorArnLists(testAnomalyMonitor.arn())
 *             .subscribers(AnomalySubscriptionSubscriberArgs.builder()
 *                 .type(&#34;EMAIL&#34;)
 *                 .address(&#34;abc@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Threshold Expression
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.costexplorer.AnomalySubscription;
 * import com.pulumi.aws.costexplorer.AnomalySubscriptionArgs;
 * import com.pulumi.aws.costexplorer.inputs.AnomalySubscriptionSubscriberArgs;
 * import com.pulumi.aws.costexplorer.inputs.AnomalySubscriptionThresholdExpressionArgs;
 * import com.pulumi.aws.costexplorer.inputs.AnomalySubscriptionThresholdExpressionDimensionArgs;
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
 *         var test = new AnomalySubscription(&#34;test&#34;, AnomalySubscriptionArgs.builder()        
 *             .frequency(&#34;DAILY&#34;)
 *             .monitorArnLists(aws_ce_anomaly_monitor.test().arn())
 *             .subscribers(AnomalySubscriptionSubscriberArgs.builder()
 *                 .type(&#34;EMAIL&#34;)
 *                 .address(&#34;abc@example.com&#34;)
 *                 .build())
 *             .thresholdExpression(AnomalySubscriptionThresholdExpressionArgs.builder()
 *                 .dimension(AnomalySubscriptionThresholdExpressionDimensionArgs.builder()
 *                     .key(&#34;ANOMALY_TOTAL_IMPACT_ABSOLUTE&#34;)
 *                     .values(&#34;100.0&#34;)
 *                     .matchOptions(&#34;GREATER_THAN_OR_EQUAL&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### SNS Example
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sns.Topic;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.sns.TopicPolicy;
 * import com.pulumi.aws.sns.TopicPolicyArgs;
 * import com.pulumi.aws.costexplorer.AnomalyMonitor;
 * import com.pulumi.aws.costexplorer.AnomalyMonitorArgs;
 * import com.pulumi.aws.costexplorer.AnomalySubscription;
 * import com.pulumi.aws.costexplorer.AnomalySubscriptionArgs;
 * import com.pulumi.aws.costexplorer.inputs.AnomalySubscriptionSubscriberArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var costAnomalyUpdates = new Topic(&#34;costAnomalyUpdates&#34;);
 * 
 *         final var snsTopicPolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .policyId(&#34;__default_policy_ID&#34;)
 *             .statements(            
 *                 GetPolicyDocumentStatementArgs.builder()
 *                     .sid(&#34;AWSAnomalyDetectionSNSPublishingPermissions&#34;)
 *                     .actions(&#34;SNS:Publish&#34;)
 *                     .effect(&#34;Allow&#34;)
 *                     .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                         .type(&#34;Service&#34;)
 *                         .identifiers(&#34;costalerts.amazonaws.com&#34;)
 *                         .build())
 *                     .resources(costAnomalyUpdates.arn())
 *                     .build(),
 *                 GetPolicyDocumentStatementArgs.builder()
 *                     .sid(&#34;__default_statement_ID&#34;)
 *                     .actions(                    
 *                         &#34;SNS:Subscribe&#34;,
 *                         &#34;SNS:SetTopicAttributes&#34;,
 *                         &#34;SNS:RemovePermission&#34;,
 *                         &#34;SNS:Receive&#34;,
 *                         &#34;SNS:Publish&#34;,
 *                         &#34;SNS:ListSubscriptionsByTopic&#34;,
 *                         &#34;SNS:GetTopicAttributes&#34;,
 *                         &#34;SNS:DeleteTopic&#34;,
 *                         &#34;SNS:AddPermission&#34;)
 *                     .conditions(GetPolicyDocumentStatementConditionArgs.builder()
 *                         .test(&#34;StringEquals&#34;)
 *                         .variable(&#34;AWS:SourceOwner&#34;)
 *                         .values(var_.account-id())
 *                         .build())
 *                     .effect(&#34;Allow&#34;)
 *                     .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                         .type(&#34;AWS&#34;)
 *                         .identifiers(&#34;*&#34;)
 *                         .build())
 *                     .resources(costAnomalyUpdates.arn())
 *                     .build())
 *             .build());
 * 
 *         var default_ = new TopicPolicy(&#34;default&#34;, TopicPolicyArgs.builder()        
 *             .arn(costAnomalyUpdates.arn())
 *             .policy(snsTopicPolicy.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult).applyValue(snsTopicPolicy -&gt; snsTopicPolicy.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json())))
 *             .build());
 * 
 *         var anomalyMonitor = new AnomalyMonitor(&#34;anomalyMonitor&#34;, AnomalyMonitorArgs.builder()        
 *             .monitorType(&#34;DIMENSIONAL&#34;)
 *             .monitorDimension(&#34;SERVICE&#34;)
 *             .build());
 * 
 *         var realtimeSubscription = new AnomalySubscription(&#34;realtimeSubscription&#34;, AnomalySubscriptionArgs.builder()        
 *             .threshold(0)
 *             .frequency(&#34;IMMEDIATE&#34;)
 *             .monitorArnLists(anomalyMonitor.arn())
 *             .subscribers(AnomalySubscriptionSubscriberArgs.builder()
 *                 .type(&#34;SNS&#34;)
 *                 .address(costAnomalyUpdates.arn())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(default_)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_ce_anomaly_subscription.example
 * 
 *  id = &#34;AnomalySubscriptionARN&#34; } Using `pulumi import`, import `aws_ce_anomaly_subscription` using the `id`. For exampleconsole % pulumi import aws_ce_anomaly_subscription.example AnomalySubscriptionARN
 * 
 */
@ResourceType(type="aws:costexplorer/anomalySubscription:AnomalySubscription")
public class AnomalySubscription extends com.pulumi.resources.CustomResource {
    /**
     * The unique identifier for the AWS account in which the anomaly subscription ought to be created.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return The unique identifier for the AWS account in which the anomaly subscription ought to be created.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * ARN of the anomaly subscription.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the anomaly subscription.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
     * 
     */
    @Export(name="frequency", refs={String.class}, tree="[0]")
    private Output<String> frequency;

    /**
     * @return The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
     * 
     */
    public Output<String> frequency() {
        return this.frequency;
    }
    /**
     * A list of cost anomaly monitors.
     * 
     */
    @Export(name="monitorArnLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> monitorArnLists;

    /**
     * @return A list of cost anomaly monitors.
     * 
     */
    public Output<List<String>> monitorArnLists() {
        return this.monitorArnLists;
    }
    /**
     * The name for the subscription.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name for the subscription.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A subscriber configuration. Multiple subscribers can be defined.
     * 
     */
    @Export(name="subscribers", refs={List.class,AnomalySubscriptionSubscriber.class}, tree="[0,1]")
    private Output<List<AnomalySubscriptionSubscriber>> subscribers;

    /**
     * @return A subscriber configuration. Multiple subscribers can be defined.
     * 
     */
    public Output<List<AnomalySubscriptionSubscriber>> subscribers() {
        return this.subscribers;
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
     * 
     */
    @Export(name="thresholdExpression", refs={AnomalySubscriptionThresholdExpression.class}, tree="[0]")
    private Output<AnomalySubscriptionThresholdExpression> thresholdExpression;

    /**
     * @return An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
     * 
     */
    public Output<AnomalySubscriptionThresholdExpression> thresholdExpression() {
        return this.thresholdExpression;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AnomalySubscription(String name) {
        this(name, AnomalySubscriptionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AnomalySubscription(String name, AnomalySubscriptionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AnomalySubscription(String name, AnomalySubscriptionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:costexplorer/anomalySubscription:AnomalySubscription", name, args == null ? AnomalySubscriptionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AnomalySubscription(String name, Output<String> id, @Nullable AnomalySubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:costexplorer/anomalySubscription:AnomalySubscription", name, state, makeResourceOptions(options, id));
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
    public static AnomalySubscription get(String name, Output<String> id, @Nullable AnomalySubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AnomalySubscription(name, id, state, options);
    }
}
