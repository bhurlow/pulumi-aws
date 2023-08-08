// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.outputs;

import com.pulumi.aws.pipes.outputs.PipeSourceParametersActivemqBrokerParameters;
import com.pulumi.aws.pipes.outputs.PipeSourceParametersDynamodbStreamParameters;
import com.pulumi.aws.pipes.outputs.PipeSourceParametersFilterCriteria;
import com.pulumi.aws.pipes.outputs.PipeSourceParametersKinesisStreamParameters;
import com.pulumi.aws.pipes.outputs.PipeSourceParametersManagedStreamingKafkaParameters;
import com.pulumi.aws.pipes.outputs.PipeSourceParametersRabbitmqBrokerParameters;
import com.pulumi.aws.pipes.outputs.PipeSourceParametersSelfManagedKafkaParameters;
import com.pulumi.aws.pipes.outputs.PipeSourceParametersSqsQueueParameters;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PipeSourceParameters {
    /**
     * @return The parameters for using an Active MQ broker as a source. Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersActivemqBrokerParameters activemqBrokerParameters;
    /**
     * @return The parameters for using a DynamoDB stream as a source.  Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersDynamodbStreamParameters dynamodbStreamParameters;
    /**
     * @return The collection of event patterns used to [filter events](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-filtering.html). Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersFilterCriteria filterCriteria;
    /**
     * @return The parameters for using a Kinesis stream as a source. Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersKinesisStreamParameters kinesisStreamParameters;
    /**
     * @return The parameters for using an MSK stream as a source. Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersManagedStreamingKafkaParameters managedStreamingKafkaParameters;
    /**
     * @return The parameters for using a Rabbit MQ broker as a source. Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersRabbitmqBrokerParameters rabbitmqBrokerParameters;
    /**
     * @return The parameters for using a self-managed Apache Kafka stream as a source. Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersSelfManagedKafkaParameters selfManagedKafkaParameters;
    /**
     * @return The parameters for using a Amazon SQS stream as a source. Detailed below.
     * 
     */
    private @Nullable PipeSourceParametersSqsQueueParameters sqsQueueParameters;

    private PipeSourceParameters() {}
    /**
     * @return The parameters for using an Active MQ broker as a source. Detailed below.
     * 
     */
    public Optional<PipeSourceParametersActivemqBrokerParameters> activemqBrokerParameters() {
        return Optional.ofNullable(this.activemqBrokerParameters);
    }
    /**
     * @return The parameters for using a DynamoDB stream as a source.  Detailed below.
     * 
     */
    public Optional<PipeSourceParametersDynamodbStreamParameters> dynamodbStreamParameters() {
        return Optional.ofNullable(this.dynamodbStreamParameters);
    }
    /**
     * @return The collection of event patterns used to [filter events](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-filtering.html). Detailed below.
     * 
     */
    public Optional<PipeSourceParametersFilterCriteria> filterCriteria() {
        return Optional.ofNullable(this.filterCriteria);
    }
    /**
     * @return The parameters for using a Kinesis stream as a source. Detailed below.
     * 
     */
    public Optional<PipeSourceParametersKinesisStreamParameters> kinesisStreamParameters() {
        return Optional.ofNullable(this.kinesisStreamParameters);
    }
    /**
     * @return The parameters for using an MSK stream as a source. Detailed below.
     * 
     */
    public Optional<PipeSourceParametersManagedStreamingKafkaParameters> managedStreamingKafkaParameters() {
        return Optional.ofNullable(this.managedStreamingKafkaParameters);
    }
    /**
     * @return The parameters for using a Rabbit MQ broker as a source. Detailed below.
     * 
     */
    public Optional<PipeSourceParametersRabbitmqBrokerParameters> rabbitmqBrokerParameters() {
        return Optional.ofNullable(this.rabbitmqBrokerParameters);
    }
    /**
     * @return The parameters for using a self-managed Apache Kafka stream as a source. Detailed below.
     * 
     */
    public Optional<PipeSourceParametersSelfManagedKafkaParameters> selfManagedKafkaParameters() {
        return Optional.ofNullable(this.selfManagedKafkaParameters);
    }
    /**
     * @return The parameters for using a Amazon SQS stream as a source. Detailed below.
     * 
     */
    public Optional<PipeSourceParametersSqsQueueParameters> sqsQueueParameters() {
        return Optional.ofNullable(this.sqsQueueParameters);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PipeSourceParameters defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable PipeSourceParametersActivemqBrokerParameters activemqBrokerParameters;
        private @Nullable PipeSourceParametersDynamodbStreamParameters dynamodbStreamParameters;
        private @Nullable PipeSourceParametersFilterCriteria filterCriteria;
        private @Nullable PipeSourceParametersKinesisStreamParameters kinesisStreamParameters;
        private @Nullable PipeSourceParametersManagedStreamingKafkaParameters managedStreamingKafkaParameters;
        private @Nullable PipeSourceParametersRabbitmqBrokerParameters rabbitmqBrokerParameters;
        private @Nullable PipeSourceParametersSelfManagedKafkaParameters selfManagedKafkaParameters;
        private @Nullable PipeSourceParametersSqsQueueParameters sqsQueueParameters;
        public Builder() {}
        public Builder(PipeSourceParameters defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.activemqBrokerParameters = defaults.activemqBrokerParameters;
    	      this.dynamodbStreamParameters = defaults.dynamodbStreamParameters;
    	      this.filterCriteria = defaults.filterCriteria;
    	      this.kinesisStreamParameters = defaults.kinesisStreamParameters;
    	      this.managedStreamingKafkaParameters = defaults.managedStreamingKafkaParameters;
    	      this.rabbitmqBrokerParameters = defaults.rabbitmqBrokerParameters;
    	      this.selfManagedKafkaParameters = defaults.selfManagedKafkaParameters;
    	      this.sqsQueueParameters = defaults.sqsQueueParameters;
        }

        @CustomType.Setter
        public Builder activemqBrokerParameters(@Nullable PipeSourceParametersActivemqBrokerParameters activemqBrokerParameters) {
            this.activemqBrokerParameters = activemqBrokerParameters;
            return this;
        }
        @CustomType.Setter
        public Builder dynamodbStreamParameters(@Nullable PipeSourceParametersDynamodbStreamParameters dynamodbStreamParameters) {
            this.dynamodbStreamParameters = dynamodbStreamParameters;
            return this;
        }
        @CustomType.Setter
        public Builder filterCriteria(@Nullable PipeSourceParametersFilterCriteria filterCriteria) {
            this.filterCriteria = filterCriteria;
            return this;
        }
        @CustomType.Setter
        public Builder kinesisStreamParameters(@Nullable PipeSourceParametersKinesisStreamParameters kinesisStreamParameters) {
            this.kinesisStreamParameters = kinesisStreamParameters;
            return this;
        }
        @CustomType.Setter
        public Builder managedStreamingKafkaParameters(@Nullable PipeSourceParametersManagedStreamingKafkaParameters managedStreamingKafkaParameters) {
            this.managedStreamingKafkaParameters = managedStreamingKafkaParameters;
            return this;
        }
        @CustomType.Setter
        public Builder rabbitmqBrokerParameters(@Nullable PipeSourceParametersRabbitmqBrokerParameters rabbitmqBrokerParameters) {
            this.rabbitmqBrokerParameters = rabbitmqBrokerParameters;
            return this;
        }
        @CustomType.Setter
        public Builder selfManagedKafkaParameters(@Nullable PipeSourceParametersSelfManagedKafkaParameters selfManagedKafkaParameters) {
            this.selfManagedKafkaParameters = selfManagedKafkaParameters;
            return this;
        }
        @CustomType.Setter
        public Builder sqsQueueParameters(@Nullable PipeSourceParametersSqsQueueParameters sqsQueueParameters) {
            this.sqsQueueParameters = sqsQueueParameters;
            return this;
        }
        public PipeSourceParameters build() {
            final var o = new PipeSourceParameters();
            o.activemqBrokerParameters = activemqBrokerParameters;
            o.dynamodbStreamParameters = dynamodbStreamParameters;
            o.filterCriteria = filterCriteria;
            o.kinesisStreamParameters = kinesisStreamParameters;
            o.managedStreamingKafkaParameters = managedStreamingKafkaParameters;
            o.rabbitmqBrokerParameters = rabbitmqBrokerParameters;
            o.selfManagedKafkaParameters = selfManagedKafkaParameters;
            o.sqsQueueParameters = sqsQueueParameters;
            return o;
        }
    }
}
