// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sns;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.sns.TopicArgs;
import com.pulumi.aws.sns.inputs.TopicState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an SNS topic resource
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sns.Topic;
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
 *         var userUpdates = new Topic(&#34;userUpdates&#34;);
 * 
 *     }
 * }
 * ```
 * ## Example with Delivery Policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sns.Topic;
 * import com.pulumi.aws.sns.TopicArgs;
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
 *         var userUpdates = new Topic(&#34;userUpdates&#34;, TopicArgs.builder()        
 *             .deliveryPolicy(&#34;&#34;&#34;
 * {
 *   &#34;http&#34;: {
 *     &#34;defaultHealthyRetryPolicy&#34;: {
 *       &#34;minDelayTarget&#34;: 20,
 *       &#34;maxDelayTarget&#34;: 20,
 *       &#34;numRetries&#34;: 3,
 *       &#34;numMaxDelayRetries&#34;: 0,
 *       &#34;numNoDelayRetries&#34;: 0,
 *       &#34;numMinDelayRetries&#34;: 0,
 *       &#34;backoffFunction&#34;: &#34;linear&#34;
 *     },
 *     &#34;disableSubscriptionOverrides&#34;: false,
 *     &#34;defaultThrottlePolicy&#34;: {
 *       &#34;maxReceivesPerSecond&#34;: 1
 *     }
 *   }
 * }
 * 
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Example with Server-side encryption (SSE)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sns.Topic;
 * import com.pulumi.aws.sns.TopicArgs;
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
 *         var userUpdates = new Topic(&#34;userUpdates&#34;, TopicArgs.builder()        
 *             .kmsMasterKeyId(&#34;alias/aws/sns&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Example with First-In-First-Out (FIFO)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sns.Topic;
 * import com.pulumi.aws.sns.TopicArgs;
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
 *         var userUpdates = new Topic(&#34;userUpdates&#34;, TopicArgs.builder()        
 *             .contentBasedDeduplication(true)
 *             .fifoTopic(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Message Delivery Status Arguments
 * 
 * The `&lt;endpoint&gt;_success_feedback_role_arn` and `&lt;endpoint&gt;_failure_feedback_role_arn` arguments are used to give Amazon SNS write access to use CloudWatch Logs on your behalf. The `&lt;endpoint&gt;_success_feedback_sample_rate` argument is for specifying the sample rate percentage (0-100) of successfully delivered messages. After you configure the  `&lt;endpoint&gt;_failure_feedback_role_arn` argument, then all failed message deliveries generate CloudWatch Logs.
 * 
 * ## Import
 * 
 * Using `pulumi import`, import SNS Topics using the topic `arn`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:sns/topic:Topic user_updates arn:aws:sns:us-west-2:0123456789012:my-topic
 * ```
 * 
 */
@ResourceType(type="aws:sns/topic:Topic")
public class Topic extends com.pulumi.resources.CustomResource {
    /**
     * IAM role for failure feedback
     * 
     */
    @Export(name="applicationFailureFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> applicationFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Output<Optional<String>> applicationFailureFeedbackRoleArn() {
        return Codegen.optional(this.applicationFailureFeedbackRoleArn);
    }
    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Export(name="applicationSuccessFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> applicationSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Output<Optional<String>> applicationSuccessFeedbackRoleArn() {
        return Codegen.optional(this.applicationSuccessFeedbackRoleArn);
    }
    /**
     * Percentage of success to sample
     * 
     */
    @Export(name="applicationSuccessFeedbackSampleRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> applicationSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Output<Optional<Integer>> applicationSuccessFeedbackSampleRate() {
        return Codegen.optional(this.applicationSuccessFeedbackSampleRate);
    }
    /**
     * The ARN of the SNS topic, as a more obvious property (clone of id)
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the SNS topic, as a more obvious property (clone of id)
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
     * 
     */
    @Export(name="contentBasedDeduplication", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> contentBasedDeduplication;

    /**
     * @return Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
     * 
     */
    public Output<Optional<Boolean>> contentBasedDeduplication() {
        return Codegen.optional(this.contentBasedDeduplication);
    }
    /**
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     * 
     */
    @Export(name="deliveryPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deliveryPolicy;

    /**
     * @return The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     * 
     */
    public Output<Optional<String>> deliveryPolicy() {
        return Codegen.optional(this.deliveryPolicy);
    }
    /**
     * The display name for the topic
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name for the topic
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
     * 
     */
    @Export(name="fifoTopic", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> fifoTopic;

    /**
     * @return Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
     * 
     */
    public Output<Optional<Boolean>> fifoTopic() {
        return Codegen.optional(this.fifoTopic);
    }
    /**
     * IAM role for failure feedback
     * 
     */
    @Export(name="firehoseFailureFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> firehoseFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Output<Optional<String>> firehoseFailureFeedbackRoleArn() {
        return Codegen.optional(this.firehoseFailureFeedbackRoleArn);
    }
    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Export(name="firehoseSuccessFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> firehoseSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Output<Optional<String>> firehoseSuccessFeedbackRoleArn() {
        return Codegen.optional(this.firehoseSuccessFeedbackRoleArn);
    }
    /**
     * Percentage of success to sample
     * 
     */
    @Export(name="firehoseSuccessFeedbackSampleRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> firehoseSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Output<Optional<Integer>> firehoseSuccessFeedbackSampleRate() {
        return Codegen.optional(this.firehoseSuccessFeedbackSampleRate);
    }
    /**
     * IAM role for failure feedback
     * 
     */
    @Export(name="httpFailureFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> httpFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Output<Optional<String>> httpFailureFeedbackRoleArn() {
        return Codegen.optional(this.httpFailureFeedbackRoleArn);
    }
    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Export(name="httpSuccessFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> httpSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Output<Optional<String>> httpSuccessFeedbackRoleArn() {
        return Codegen.optional(this.httpSuccessFeedbackRoleArn);
    }
    /**
     * Percentage of success to sample
     * 
     */
    @Export(name="httpSuccessFeedbackSampleRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> httpSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Output<Optional<Integer>> httpSuccessFeedbackSampleRate() {
        return Codegen.optional(this.httpSuccessFeedbackSampleRate);
    }
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     * 
     */
    @Export(name="kmsMasterKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsMasterKeyId;

    /**
     * @return The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     * 
     */
    public Output<Optional<String>> kmsMasterKeyId() {
        return Codegen.optional(this.kmsMasterKeyId);
    }
    /**
     * IAM role for failure feedback
     * 
     */
    @Export(name="lambdaFailureFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lambdaFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Output<Optional<String>> lambdaFailureFeedbackRoleArn() {
        return Codegen.optional(this.lambdaFailureFeedbackRoleArn);
    }
    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Export(name="lambdaSuccessFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lambdaSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Output<Optional<String>> lambdaSuccessFeedbackRoleArn() {
        return Codegen.optional(this.lambdaSuccessFeedbackRoleArn);
    }
    /**
     * Percentage of success to sample
     * 
     */
    @Export(name="lambdaSuccessFeedbackSampleRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> lambdaSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Output<Optional<Integer>> lambdaSuccessFeedbackSampleRate() {
        return Codegen.optional(this.lambdaSuccessFeedbackSampleRate);
    }
    /**
     * The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`
     * 
     */
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`
     * 
     */
    public Output<String> namePrefix() {
        return this.namePrefix;
    }
    /**
     * The AWS Account ID of the SNS topic owner
     * 
     */
    @Export(name="owner", refs={String.class}, tree="[0]")
    private Output<String> owner;

    /**
     * @return The AWS Account ID of the SNS topic owner
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * The fully-formed AWS policy as JSON.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return The fully-formed AWS policy as JSON.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
     * 
     */
    @Export(name="signatureVersion", refs={Integer.class}, tree="[0]")
    private Output<Integer> signatureVersion;

    /**
     * @return If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
     * 
     */
    public Output<Integer> signatureVersion() {
        return this.signatureVersion;
    }
    /**
     * IAM role for failure feedback
     * 
     */
    @Export(name="sqsFailureFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sqsFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Output<Optional<String>> sqsFailureFeedbackRoleArn() {
        return Codegen.optional(this.sqsFailureFeedbackRoleArn);
    }
    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Export(name="sqsSuccessFeedbackRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sqsSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Output<Optional<String>> sqsSuccessFeedbackRoleArn() {
        return Codegen.optional(this.sqsSuccessFeedbackRoleArn);
    }
    /**
     * Percentage of success to sample
     * 
     */
    @Export(name="sqsSuccessFeedbackSampleRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sqsSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Output<Optional<Integer>> sqsSuccessFeedbackSampleRate() {
        return Codegen.optional(this.sqsSuccessFeedbackSampleRate);
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
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
     * Tracing mode of an Amazon SNS topic. Valid values: `&#34;PassThrough&#34;`, `&#34;Active&#34;`.
     * 
     */
    @Export(name="tracingConfig", refs={String.class}, tree="[0]")
    private Output<String> tracingConfig;

    /**
     * @return Tracing mode of an Amazon SNS topic. Valid values: `&#34;PassThrough&#34;`, `&#34;Active&#34;`.
     * 
     */
    public Output<String> tracingConfig() {
        return this.tracingConfig;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Topic(String name) {
        this(name, TopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Topic(String name, @Nullable TopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Topic(String name, @Nullable TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sns/topic:Topic", name, args == null ? TopicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Topic(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sns/topic:Topic", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static Topic get(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Topic(name, id, state, options);
    }
}
