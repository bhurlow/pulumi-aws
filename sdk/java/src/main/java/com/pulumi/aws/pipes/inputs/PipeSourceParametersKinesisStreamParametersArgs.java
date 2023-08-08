// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.inputs;

import com.pulumi.aws.pipes.inputs.PipeSourceParametersKinesisStreamParametersDeadLetterConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipeSourceParametersKinesisStreamParametersArgs extends com.pulumi.resources.ResourceArgs {

    public static final PipeSourceParametersKinesisStreamParametersArgs Empty = new PipeSourceParametersKinesisStreamParametersArgs();

    /**
     * The maximum number of records to include in each batch. Maximum value of 10000.
     * 
     */
    @Import(name="batchSize")
    private @Nullable Output<Integer> batchSize;

    /**
     * @return The maximum number of records to include in each batch. Maximum value of 10000.
     * 
     */
    public Optional<Output<Integer>> batchSize() {
        return Optional.ofNullable(this.batchSize);
    }

    /**
     * Define the target queue to send dead-letter queue events to. Detailed below.
     * 
     */
    @Import(name="deadLetterConfig")
    private @Nullable Output<PipeSourceParametersKinesisStreamParametersDeadLetterConfigArgs> deadLetterConfig;

    /**
     * @return Define the target queue to send dead-letter queue events to. Detailed below.
     * 
     */
    public Optional<Output<PipeSourceParametersKinesisStreamParametersDeadLetterConfigArgs>> deadLetterConfig() {
        return Optional.ofNullable(this.deadLetterConfig);
    }

    /**
     * The maximum length of a time to wait for events. Maximum value of 300.
     * 
     */
    @Import(name="maximumBatchingWindowInSeconds")
    private @Nullable Output<Integer> maximumBatchingWindowInSeconds;

    /**
     * @return The maximum length of a time to wait for events. Maximum value of 300.
     * 
     */
    public Optional<Output<Integer>> maximumBatchingWindowInSeconds() {
        return Optional.ofNullable(this.maximumBatchingWindowInSeconds);
    }

    /**
     * Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, EventBridge never discards old records. Maximum value of 604,800.
     * 
     */
    @Import(name="maximumRecordAgeInSeconds")
    private @Nullable Output<Integer> maximumRecordAgeInSeconds;

    /**
     * @return Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, EventBridge never discards old records. Maximum value of 604,800.
     * 
     */
    public Optional<Output<Integer>> maximumRecordAgeInSeconds() {
        return Optional.ofNullable(this.maximumRecordAgeInSeconds);
    }

    /**
     * Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source. Maximum value of 10,000.
     * 
     */
    @Import(name="maximumRetryAttempts")
    private @Nullable Output<Integer> maximumRetryAttempts;

    /**
     * @return Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source. Maximum value of 10,000.
     * 
     */
    public Optional<Output<Integer>> maximumRetryAttempts() {
        return Optional.ofNullable(this.maximumRetryAttempts);
    }

    /**
     * Define how to handle item process failures. AUTOMATIC_BISECT halves each batch and retry each half until all the records are processed or there is one failed message left in the batch. Valid values: AUTOMATIC_BISECT.
     * 
     */
    @Import(name="onPartialBatchItemFailure")
    private @Nullable Output<String> onPartialBatchItemFailure;

    /**
     * @return Define how to handle item process failures. AUTOMATIC_BISECT halves each batch and retry each half until all the records are processed or there is one failed message left in the batch. Valid values: AUTOMATIC_BISECT.
     * 
     */
    public Optional<Output<String>> onPartialBatchItemFailure() {
        return Optional.ofNullable(this.onPartialBatchItemFailure);
    }

    /**
     * The number of batches to process concurrently from each shard. The default value is 1. Maximum value of 10.
     * 
     */
    @Import(name="parallelizationFactor")
    private @Nullable Output<Integer> parallelizationFactor;

    /**
     * @return The number of batches to process concurrently from each shard. The default value is 1. Maximum value of 10.
     * 
     */
    public Optional<Output<Integer>> parallelizationFactor() {
        return Optional.ofNullable(this.parallelizationFactor);
    }

    /**
     * The position in a stream from which to start reading. Valid values: TRIM_HORIZON, LATEST.
     * 
     */
    @Import(name="startingPosition", required=true)
    private Output<String> startingPosition;

    /**
     * @return The position in a stream from which to start reading. Valid values: TRIM_HORIZON, LATEST.
     * 
     */
    public Output<String> startingPosition() {
        return this.startingPosition;
    }

    /**
     * With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
     * 
     */
    @Import(name="startingPositionTimestamp")
    private @Nullable Output<String> startingPositionTimestamp;

    /**
     * @return With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
     * 
     */
    public Optional<Output<String>> startingPositionTimestamp() {
        return Optional.ofNullable(this.startingPositionTimestamp);
    }

    private PipeSourceParametersKinesisStreamParametersArgs() {}

    private PipeSourceParametersKinesisStreamParametersArgs(PipeSourceParametersKinesisStreamParametersArgs $) {
        this.batchSize = $.batchSize;
        this.deadLetterConfig = $.deadLetterConfig;
        this.maximumBatchingWindowInSeconds = $.maximumBatchingWindowInSeconds;
        this.maximumRecordAgeInSeconds = $.maximumRecordAgeInSeconds;
        this.maximumRetryAttempts = $.maximumRetryAttempts;
        this.onPartialBatchItemFailure = $.onPartialBatchItemFailure;
        this.parallelizationFactor = $.parallelizationFactor;
        this.startingPosition = $.startingPosition;
        this.startingPositionTimestamp = $.startingPositionTimestamp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipeSourceParametersKinesisStreamParametersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipeSourceParametersKinesisStreamParametersArgs $;

        public Builder() {
            $ = new PipeSourceParametersKinesisStreamParametersArgs();
        }

        public Builder(PipeSourceParametersKinesisStreamParametersArgs defaults) {
            $ = new PipeSourceParametersKinesisStreamParametersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param batchSize The maximum number of records to include in each batch. Maximum value of 10000.
         * 
         * @return builder
         * 
         */
        public Builder batchSize(@Nullable Output<Integer> batchSize) {
            $.batchSize = batchSize;
            return this;
        }

        /**
         * @param batchSize The maximum number of records to include in each batch. Maximum value of 10000.
         * 
         * @return builder
         * 
         */
        public Builder batchSize(Integer batchSize) {
            return batchSize(Output.of(batchSize));
        }

        /**
         * @param deadLetterConfig Define the target queue to send dead-letter queue events to. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder deadLetterConfig(@Nullable Output<PipeSourceParametersKinesisStreamParametersDeadLetterConfigArgs> deadLetterConfig) {
            $.deadLetterConfig = deadLetterConfig;
            return this;
        }

        /**
         * @param deadLetterConfig Define the target queue to send dead-letter queue events to. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder deadLetterConfig(PipeSourceParametersKinesisStreamParametersDeadLetterConfigArgs deadLetterConfig) {
            return deadLetterConfig(Output.of(deadLetterConfig));
        }

        /**
         * @param maximumBatchingWindowInSeconds The maximum length of a time to wait for events. Maximum value of 300.
         * 
         * @return builder
         * 
         */
        public Builder maximumBatchingWindowInSeconds(@Nullable Output<Integer> maximumBatchingWindowInSeconds) {
            $.maximumBatchingWindowInSeconds = maximumBatchingWindowInSeconds;
            return this;
        }

        /**
         * @param maximumBatchingWindowInSeconds The maximum length of a time to wait for events. Maximum value of 300.
         * 
         * @return builder
         * 
         */
        public Builder maximumBatchingWindowInSeconds(Integer maximumBatchingWindowInSeconds) {
            return maximumBatchingWindowInSeconds(Output.of(maximumBatchingWindowInSeconds));
        }

        /**
         * @param maximumRecordAgeInSeconds Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, EventBridge never discards old records. Maximum value of 604,800.
         * 
         * @return builder
         * 
         */
        public Builder maximumRecordAgeInSeconds(@Nullable Output<Integer> maximumRecordAgeInSeconds) {
            $.maximumRecordAgeInSeconds = maximumRecordAgeInSeconds;
            return this;
        }

        /**
         * @param maximumRecordAgeInSeconds Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, EventBridge never discards old records. Maximum value of 604,800.
         * 
         * @return builder
         * 
         */
        public Builder maximumRecordAgeInSeconds(Integer maximumRecordAgeInSeconds) {
            return maximumRecordAgeInSeconds(Output.of(maximumRecordAgeInSeconds));
        }

        /**
         * @param maximumRetryAttempts Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source. Maximum value of 10,000.
         * 
         * @return builder
         * 
         */
        public Builder maximumRetryAttempts(@Nullable Output<Integer> maximumRetryAttempts) {
            $.maximumRetryAttempts = maximumRetryAttempts;
            return this;
        }

        /**
         * @param maximumRetryAttempts Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source. Maximum value of 10,000.
         * 
         * @return builder
         * 
         */
        public Builder maximumRetryAttempts(Integer maximumRetryAttempts) {
            return maximumRetryAttempts(Output.of(maximumRetryAttempts));
        }

        /**
         * @param onPartialBatchItemFailure Define how to handle item process failures. AUTOMATIC_BISECT halves each batch and retry each half until all the records are processed or there is one failed message left in the batch. Valid values: AUTOMATIC_BISECT.
         * 
         * @return builder
         * 
         */
        public Builder onPartialBatchItemFailure(@Nullable Output<String> onPartialBatchItemFailure) {
            $.onPartialBatchItemFailure = onPartialBatchItemFailure;
            return this;
        }

        /**
         * @param onPartialBatchItemFailure Define how to handle item process failures. AUTOMATIC_BISECT halves each batch and retry each half until all the records are processed or there is one failed message left in the batch. Valid values: AUTOMATIC_BISECT.
         * 
         * @return builder
         * 
         */
        public Builder onPartialBatchItemFailure(String onPartialBatchItemFailure) {
            return onPartialBatchItemFailure(Output.of(onPartialBatchItemFailure));
        }

        /**
         * @param parallelizationFactor The number of batches to process concurrently from each shard. The default value is 1. Maximum value of 10.
         * 
         * @return builder
         * 
         */
        public Builder parallelizationFactor(@Nullable Output<Integer> parallelizationFactor) {
            $.parallelizationFactor = parallelizationFactor;
            return this;
        }

        /**
         * @param parallelizationFactor The number of batches to process concurrently from each shard. The default value is 1. Maximum value of 10.
         * 
         * @return builder
         * 
         */
        public Builder parallelizationFactor(Integer parallelizationFactor) {
            return parallelizationFactor(Output.of(parallelizationFactor));
        }

        /**
         * @param startingPosition The position in a stream from which to start reading. Valid values: TRIM_HORIZON, LATEST.
         * 
         * @return builder
         * 
         */
        public Builder startingPosition(Output<String> startingPosition) {
            $.startingPosition = startingPosition;
            return this;
        }

        /**
         * @param startingPosition The position in a stream from which to start reading. Valid values: TRIM_HORIZON, LATEST.
         * 
         * @return builder
         * 
         */
        public Builder startingPosition(String startingPosition) {
            return startingPosition(Output.of(startingPosition));
        }

        /**
         * @param startingPositionTimestamp With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
         * 
         * @return builder
         * 
         */
        public Builder startingPositionTimestamp(@Nullable Output<String> startingPositionTimestamp) {
            $.startingPositionTimestamp = startingPositionTimestamp;
            return this;
        }

        /**
         * @param startingPositionTimestamp With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
         * 
         * @return builder
         * 
         */
        public Builder startingPositionTimestamp(String startingPositionTimestamp) {
            return startingPositionTimestamp(Output.of(startingPositionTimestamp));
        }

        public PipeSourceParametersKinesisStreamParametersArgs build() {
            $.startingPosition = Objects.requireNonNull($.startingPosition, "expected parameter 'startingPosition' to be non-null");
            return $;
        }
    }

}
