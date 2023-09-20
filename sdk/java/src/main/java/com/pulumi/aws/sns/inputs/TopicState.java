// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TopicState extends com.pulumi.resources.ResourceArgs {

    public static final TopicState Empty = new TopicState();

    /**
     * IAM role for failure feedback
     * 
     */
    @Import(name="applicationFailureFeedbackRoleArn")
    private @Nullable Output<String> applicationFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Optional<Output<String>> applicationFailureFeedbackRoleArn() {
        return Optional.ofNullable(this.applicationFailureFeedbackRoleArn);
    }

    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Import(name="applicationSuccessFeedbackRoleArn")
    private @Nullable Output<String> applicationSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Optional<Output<String>> applicationSuccessFeedbackRoleArn() {
        return Optional.ofNullable(this.applicationSuccessFeedbackRoleArn);
    }

    /**
     * Percentage of success to sample
     * 
     */
    @Import(name="applicationSuccessFeedbackSampleRate")
    private @Nullable Output<Integer> applicationSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Optional<Output<Integer>> applicationSuccessFeedbackSampleRate() {
        return Optional.ofNullable(this.applicationSuccessFeedbackSampleRate);
    }

    /**
     * The ARN of the SNS topic, as a more obvious property (clone of id)
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the SNS topic, as a more obvious property (clone of id)
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
     * 
     */
    @Import(name="contentBasedDeduplication")
    private @Nullable Output<Boolean> contentBasedDeduplication;

    /**
     * @return Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
     * 
     */
    public Optional<Output<Boolean>> contentBasedDeduplication() {
        return Optional.ofNullable(this.contentBasedDeduplication);
    }

    /**
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     * 
     */
    @Import(name="deliveryPolicy")
    private @Nullable Output<String> deliveryPolicy;

    /**
     * @return The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     * 
     */
    public Optional<Output<String>> deliveryPolicy() {
        return Optional.ofNullable(this.deliveryPolicy);
    }

    /**
     * The display name for the topic
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name for the topic
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
     * 
     */
    @Import(name="fifoTopic")
    private @Nullable Output<Boolean> fifoTopic;

    /**
     * @return Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
     * 
     */
    public Optional<Output<Boolean>> fifoTopic() {
        return Optional.ofNullable(this.fifoTopic);
    }

    /**
     * IAM role for failure feedback
     * 
     */
    @Import(name="firehoseFailureFeedbackRoleArn")
    private @Nullable Output<String> firehoseFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Optional<Output<String>> firehoseFailureFeedbackRoleArn() {
        return Optional.ofNullable(this.firehoseFailureFeedbackRoleArn);
    }

    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Import(name="firehoseSuccessFeedbackRoleArn")
    private @Nullable Output<String> firehoseSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Optional<Output<String>> firehoseSuccessFeedbackRoleArn() {
        return Optional.ofNullable(this.firehoseSuccessFeedbackRoleArn);
    }

    /**
     * Percentage of success to sample
     * 
     */
    @Import(name="firehoseSuccessFeedbackSampleRate")
    private @Nullable Output<Integer> firehoseSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Optional<Output<Integer>> firehoseSuccessFeedbackSampleRate() {
        return Optional.ofNullable(this.firehoseSuccessFeedbackSampleRate);
    }

    /**
     * IAM role for failure feedback
     * 
     */
    @Import(name="httpFailureFeedbackRoleArn")
    private @Nullable Output<String> httpFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Optional<Output<String>> httpFailureFeedbackRoleArn() {
        return Optional.ofNullable(this.httpFailureFeedbackRoleArn);
    }

    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Import(name="httpSuccessFeedbackRoleArn")
    private @Nullable Output<String> httpSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Optional<Output<String>> httpSuccessFeedbackRoleArn() {
        return Optional.ofNullable(this.httpSuccessFeedbackRoleArn);
    }

    /**
     * Percentage of success to sample
     * 
     */
    @Import(name="httpSuccessFeedbackSampleRate")
    private @Nullable Output<Integer> httpSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Optional<Output<Integer>> httpSuccessFeedbackSampleRate() {
        return Optional.ofNullable(this.httpSuccessFeedbackSampleRate);
    }

    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     * 
     */
    @Import(name="kmsMasterKeyId")
    private @Nullable Output<String> kmsMasterKeyId;

    /**
     * @return The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     * 
     */
    public Optional<Output<String>> kmsMasterKeyId() {
        return Optional.ofNullable(this.kmsMasterKeyId);
    }

    /**
     * IAM role for failure feedback
     * 
     */
    @Import(name="lambdaFailureFeedbackRoleArn")
    private @Nullable Output<String> lambdaFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Optional<Output<String>> lambdaFailureFeedbackRoleArn() {
        return Optional.ofNullable(this.lambdaFailureFeedbackRoleArn);
    }

    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Import(name="lambdaSuccessFeedbackRoleArn")
    private @Nullable Output<String> lambdaSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Optional<Output<String>> lambdaSuccessFeedbackRoleArn() {
        return Optional.ofNullable(this.lambdaSuccessFeedbackRoleArn);
    }

    /**
     * Percentage of success to sample
     * 
     */
    @Import(name="lambdaSuccessFeedbackSampleRate")
    private @Nullable Output<Integer> lambdaSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Optional<Output<Integer>> lambdaSuccessFeedbackSampleRate() {
        return Optional.ofNullable(this.lambdaSuccessFeedbackSampleRate);
    }

    /**
     * The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * The AWS Account ID of the SNS topic owner
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return The AWS Account ID of the SNS topic owner
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * The fully-formed AWS policy as JSON.
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return The fully-formed AWS policy as JSON.
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    /**
     * If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
     * 
     */
    @Import(name="signatureVersion")
    private @Nullable Output<Integer> signatureVersion;

    /**
     * @return If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
     * 
     */
    public Optional<Output<Integer>> signatureVersion() {
        return Optional.ofNullable(this.signatureVersion);
    }

    /**
     * IAM role for failure feedback
     * 
     */
    @Import(name="sqsFailureFeedbackRoleArn")
    private @Nullable Output<String> sqsFailureFeedbackRoleArn;

    /**
     * @return IAM role for failure feedback
     * 
     */
    public Optional<Output<String>> sqsFailureFeedbackRoleArn() {
        return Optional.ofNullable(this.sqsFailureFeedbackRoleArn);
    }

    /**
     * The IAM role permitted to receive success feedback for this topic
     * 
     */
    @Import(name="sqsSuccessFeedbackRoleArn")
    private @Nullable Output<String> sqsSuccessFeedbackRoleArn;

    /**
     * @return The IAM role permitted to receive success feedback for this topic
     * 
     */
    public Optional<Output<String>> sqsSuccessFeedbackRoleArn() {
        return Optional.ofNullable(this.sqsSuccessFeedbackRoleArn);
    }

    /**
     * Percentage of success to sample
     * 
     */
    @Import(name="sqsSuccessFeedbackSampleRate")
    private @Nullable Output<Integer> sqsSuccessFeedbackSampleRate;

    /**
     * @return Percentage of success to sample
     * 
     */
    public Optional<Output<Integer>> sqsSuccessFeedbackSampleRate() {
        return Optional.ofNullable(this.sqsSuccessFeedbackSampleRate);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Tracing mode of an Amazon SNS topic. Valid values: `&#34;PassThrough&#34;`, `&#34;Active&#34;`.
     * 
     */
    @Import(name="tracingConfig")
    private @Nullable Output<String> tracingConfig;

    /**
     * @return Tracing mode of an Amazon SNS topic. Valid values: `&#34;PassThrough&#34;`, `&#34;Active&#34;`.
     * 
     */
    public Optional<Output<String>> tracingConfig() {
        return Optional.ofNullable(this.tracingConfig);
    }

    private TopicState() {}

    private TopicState(TopicState $) {
        this.applicationFailureFeedbackRoleArn = $.applicationFailureFeedbackRoleArn;
        this.applicationSuccessFeedbackRoleArn = $.applicationSuccessFeedbackRoleArn;
        this.applicationSuccessFeedbackSampleRate = $.applicationSuccessFeedbackSampleRate;
        this.arn = $.arn;
        this.contentBasedDeduplication = $.contentBasedDeduplication;
        this.deliveryPolicy = $.deliveryPolicy;
        this.displayName = $.displayName;
        this.fifoTopic = $.fifoTopic;
        this.firehoseFailureFeedbackRoleArn = $.firehoseFailureFeedbackRoleArn;
        this.firehoseSuccessFeedbackRoleArn = $.firehoseSuccessFeedbackRoleArn;
        this.firehoseSuccessFeedbackSampleRate = $.firehoseSuccessFeedbackSampleRate;
        this.httpFailureFeedbackRoleArn = $.httpFailureFeedbackRoleArn;
        this.httpSuccessFeedbackRoleArn = $.httpSuccessFeedbackRoleArn;
        this.httpSuccessFeedbackSampleRate = $.httpSuccessFeedbackSampleRate;
        this.kmsMasterKeyId = $.kmsMasterKeyId;
        this.lambdaFailureFeedbackRoleArn = $.lambdaFailureFeedbackRoleArn;
        this.lambdaSuccessFeedbackRoleArn = $.lambdaSuccessFeedbackRoleArn;
        this.lambdaSuccessFeedbackSampleRate = $.lambdaSuccessFeedbackSampleRate;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.owner = $.owner;
        this.policy = $.policy;
        this.signatureVersion = $.signatureVersion;
        this.sqsFailureFeedbackRoleArn = $.sqsFailureFeedbackRoleArn;
        this.sqsSuccessFeedbackRoleArn = $.sqsSuccessFeedbackRoleArn;
        this.sqsSuccessFeedbackSampleRate = $.sqsSuccessFeedbackSampleRate;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.tracingConfig = $.tracingConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TopicState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TopicState $;

        public Builder() {
            $ = new TopicState();
        }

        public Builder(TopicState defaults) {
            $ = new TopicState(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder applicationFailureFeedbackRoleArn(@Nullable Output<String> applicationFailureFeedbackRoleArn) {
            $.applicationFailureFeedbackRoleArn = applicationFailureFeedbackRoleArn;
            return this;
        }

        /**
         * @param applicationFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder applicationFailureFeedbackRoleArn(String applicationFailureFeedbackRoleArn) {
            return applicationFailureFeedbackRoleArn(Output.of(applicationFailureFeedbackRoleArn));
        }

        /**
         * @param applicationSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder applicationSuccessFeedbackRoleArn(@Nullable Output<String> applicationSuccessFeedbackRoleArn) {
            $.applicationSuccessFeedbackRoleArn = applicationSuccessFeedbackRoleArn;
            return this;
        }

        /**
         * @param applicationSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder applicationSuccessFeedbackRoleArn(String applicationSuccessFeedbackRoleArn) {
            return applicationSuccessFeedbackRoleArn(Output.of(applicationSuccessFeedbackRoleArn));
        }

        /**
         * @param applicationSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder applicationSuccessFeedbackSampleRate(@Nullable Output<Integer> applicationSuccessFeedbackSampleRate) {
            $.applicationSuccessFeedbackSampleRate = applicationSuccessFeedbackSampleRate;
            return this;
        }

        /**
         * @param applicationSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder applicationSuccessFeedbackSampleRate(Integer applicationSuccessFeedbackSampleRate) {
            return applicationSuccessFeedbackSampleRate(Output.of(applicationSuccessFeedbackSampleRate));
        }

        /**
         * @param arn The ARN of the SNS topic, as a more obvious property (clone of id)
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the SNS topic, as a more obvious property (clone of id)
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param contentBasedDeduplication Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
         * 
         * @return builder
         * 
         */
        public Builder contentBasedDeduplication(@Nullable Output<Boolean> contentBasedDeduplication) {
            $.contentBasedDeduplication = contentBasedDeduplication;
            return this;
        }

        /**
         * @param contentBasedDeduplication Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
         * 
         * @return builder
         * 
         */
        public Builder contentBasedDeduplication(Boolean contentBasedDeduplication) {
            return contentBasedDeduplication(Output.of(contentBasedDeduplication));
        }

        /**
         * @param deliveryPolicy The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
         * 
         * @return builder
         * 
         */
        public Builder deliveryPolicy(@Nullable Output<String> deliveryPolicy) {
            $.deliveryPolicy = deliveryPolicy;
            return this;
        }

        /**
         * @param deliveryPolicy The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
         * 
         * @return builder
         * 
         */
        public Builder deliveryPolicy(String deliveryPolicy) {
            return deliveryPolicy(Output.of(deliveryPolicy));
        }

        /**
         * @param displayName The display name for the topic
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name for the topic
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param fifoTopic Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
         * 
         * @return builder
         * 
         */
        public Builder fifoTopic(@Nullable Output<Boolean> fifoTopic) {
            $.fifoTopic = fifoTopic;
            return this;
        }

        /**
         * @param fifoTopic Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
         * 
         * @return builder
         * 
         */
        public Builder fifoTopic(Boolean fifoTopic) {
            return fifoTopic(Output.of(fifoTopic));
        }

        /**
         * @param firehoseFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder firehoseFailureFeedbackRoleArn(@Nullable Output<String> firehoseFailureFeedbackRoleArn) {
            $.firehoseFailureFeedbackRoleArn = firehoseFailureFeedbackRoleArn;
            return this;
        }

        /**
         * @param firehoseFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder firehoseFailureFeedbackRoleArn(String firehoseFailureFeedbackRoleArn) {
            return firehoseFailureFeedbackRoleArn(Output.of(firehoseFailureFeedbackRoleArn));
        }

        /**
         * @param firehoseSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder firehoseSuccessFeedbackRoleArn(@Nullable Output<String> firehoseSuccessFeedbackRoleArn) {
            $.firehoseSuccessFeedbackRoleArn = firehoseSuccessFeedbackRoleArn;
            return this;
        }

        /**
         * @param firehoseSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder firehoseSuccessFeedbackRoleArn(String firehoseSuccessFeedbackRoleArn) {
            return firehoseSuccessFeedbackRoleArn(Output.of(firehoseSuccessFeedbackRoleArn));
        }

        /**
         * @param firehoseSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder firehoseSuccessFeedbackSampleRate(@Nullable Output<Integer> firehoseSuccessFeedbackSampleRate) {
            $.firehoseSuccessFeedbackSampleRate = firehoseSuccessFeedbackSampleRate;
            return this;
        }

        /**
         * @param firehoseSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder firehoseSuccessFeedbackSampleRate(Integer firehoseSuccessFeedbackSampleRate) {
            return firehoseSuccessFeedbackSampleRate(Output.of(firehoseSuccessFeedbackSampleRate));
        }

        /**
         * @param httpFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder httpFailureFeedbackRoleArn(@Nullable Output<String> httpFailureFeedbackRoleArn) {
            $.httpFailureFeedbackRoleArn = httpFailureFeedbackRoleArn;
            return this;
        }

        /**
         * @param httpFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder httpFailureFeedbackRoleArn(String httpFailureFeedbackRoleArn) {
            return httpFailureFeedbackRoleArn(Output.of(httpFailureFeedbackRoleArn));
        }

        /**
         * @param httpSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder httpSuccessFeedbackRoleArn(@Nullable Output<String> httpSuccessFeedbackRoleArn) {
            $.httpSuccessFeedbackRoleArn = httpSuccessFeedbackRoleArn;
            return this;
        }

        /**
         * @param httpSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder httpSuccessFeedbackRoleArn(String httpSuccessFeedbackRoleArn) {
            return httpSuccessFeedbackRoleArn(Output.of(httpSuccessFeedbackRoleArn));
        }

        /**
         * @param httpSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder httpSuccessFeedbackSampleRate(@Nullable Output<Integer> httpSuccessFeedbackSampleRate) {
            $.httpSuccessFeedbackSampleRate = httpSuccessFeedbackSampleRate;
            return this;
        }

        /**
         * @param httpSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder httpSuccessFeedbackSampleRate(Integer httpSuccessFeedbackSampleRate) {
            return httpSuccessFeedbackSampleRate(Output.of(httpSuccessFeedbackSampleRate));
        }

        /**
         * @param kmsMasterKeyId The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
         * 
         * @return builder
         * 
         */
        public Builder kmsMasterKeyId(@Nullable Output<String> kmsMasterKeyId) {
            $.kmsMasterKeyId = kmsMasterKeyId;
            return this;
        }

        /**
         * @param kmsMasterKeyId The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
         * 
         * @return builder
         * 
         */
        public Builder kmsMasterKeyId(String kmsMasterKeyId) {
            return kmsMasterKeyId(Output.of(kmsMasterKeyId));
        }

        /**
         * @param lambdaFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder lambdaFailureFeedbackRoleArn(@Nullable Output<String> lambdaFailureFeedbackRoleArn) {
            $.lambdaFailureFeedbackRoleArn = lambdaFailureFeedbackRoleArn;
            return this;
        }

        /**
         * @param lambdaFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder lambdaFailureFeedbackRoleArn(String lambdaFailureFeedbackRoleArn) {
            return lambdaFailureFeedbackRoleArn(Output.of(lambdaFailureFeedbackRoleArn));
        }

        /**
         * @param lambdaSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder lambdaSuccessFeedbackRoleArn(@Nullable Output<String> lambdaSuccessFeedbackRoleArn) {
            $.lambdaSuccessFeedbackRoleArn = lambdaSuccessFeedbackRoleArn;
            return this;
        }

        /**
         * @param lambdaSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder lambdaSuccessFeedbackRoleArn(String lambdaSuccessFeedbackRoleArn) {
            return lambdaSuccessFeedbackRoleArn(Output.of(lambdaSuccessFeedbackRoleArn));
        }

        /**
         * @param lambdaSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder lambdaSuccessFeedbackSampleRate(@Nullable Output<Integer> lambdaSuccessFeedbackSampleRate) {
            $.lambdaSuccessFeedbackSampleRate = lambdaSuccessFeedbackSampleRate;
            return this;
        }

        /**
         * @param lambdaSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder lambdaSuccessFeedbackSampleRate(Integer lambdaSuccessFeedbackSampleRate) {
            return lambdaSuccessFeedbackSampleRate(Output.of(lambdaSuccessFeedbackSampleRate));
        }

        /**
         * @param name The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param owner The AWS Account ID of the SNS topic owner
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The AWS Account ID of the SNS topic owner
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param policy The fully-formed AWS policy as JSON.
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The fully-formed AWS policy as JSON.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param signatureVersion If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
         * 
         * @return builder
         * 
         */
        public Builder signatureVersion(@Nullable Output<Integer> signatureVersion) {
            $.signatureVersion = signatureVersion;
            return this;
        }

        /**
         * @param signatureVersion If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
         * 
         * @return builder
         * 
         */
        public Builder signatureVersion(Integer signatureVersion) {
            return signatureVersion(Output.of(signatureVersion));
        }

        /**
         * @param sqsFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder sqsFailureFeedbackRoleArn(@Nullable Output<String> sqsFailureFeedbackRoleArn) {
            $.sqsFailureFeedbackRoleArn = sqsFailureFeedbackRoleArn;
            return this;
        }

        /**
         * @param sqsFailureFeedbackRoleArn IAM role for failure feedback
         * 
         * @return builder
         * 
         */
        public Builder sqsFailureFeedbackRoleArn(String sqsFailureFeedbackRoleArn) {
            return sqsFailureFeedbackRoleArn(Output.of(sqsFailureFeedbackRoleArn));
        }

        /**
         * @param sqsSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder sqsSuccessFeedbackRoleArn(@Nullable Output<String> sqsSuccessFeedbackRoleArn) {
            $.sqsSuccessFeedbackRoleArn = sqsSuccessFeedbackRoleArn;
            return this;
        }

        /**
         * @param sqsSuccessFeedbackRoleArn The IAM role permitted to receive success feedback for this topic
         * 
         * @return builder
         * 
         */
        public Builder sqsSuccessFeedbackRoleArn(String sqsSuccessFeedbackRoleArn) {
            return sqsSuccessFeedbackRoleArn(Output.of(sqsSuccessFeedbackRoleArn));
        }

        /**
         * @param sqsSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder sqsSuccessFeedbackSampleRate(@Nullable Output<Integer> sqsSuccessFeedbackSampleRate) {
            $.sqsSuccessFeedbackSampleRate = sqsSuccessFeedbackSampleRate;
            return this;
        }

        /**
         * @param sqsSuccessFeedbackSampleRate Percentage of success to sample
         * 
         * @return builder
         * 
         */
        public Builder sqsSuccessFeedbackSampleRate(Integer sqsSuccessFeedbackSampleRate) {
            return sqsSuccessFeedbackSampleRate(Output.of(sqsSuccessFeedbackSampleRate));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param tracingConfig Tracing mode of an Amazon SNS topic. Valid values: `&#34;PassThrough&#34;`, `&#34;Active&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder tracingConfig(@Nullable Output<String> tracingConfig) {
            $.tracingConfig = tracingConfig;
            return this;
        }

        /**
         * @param tracingConfig Tracing mode of an Amazon SNS topic. Valid values: `&#34;PassThrough&#34;`, `&#34;Active&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder tracingConfig(String tracingConfig) {
            return tracingConfig(Output.of(tracingConfig));
        }

        public TopicState build() {
            return $;
        }
    }

}
