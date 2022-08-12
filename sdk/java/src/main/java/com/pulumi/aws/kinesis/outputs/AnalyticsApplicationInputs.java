// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesis.outputs;

import com.pulumi.aws.kinesis.outputs.AnalyticsApplicationInputsKinesisFirehose;
import com.pulumi.aws.kinesis.outputs.AnalyticsApplicationInputsKinesisStream;
import com.pulumi.aws.kinesis.outputs.AnalyticsApplicationInputsParallelism;
import com.pulumi.aws.kinesis.outputs.AnalyticsApplicationInputsProcessingConfiguration;
import com.pulumi.aws.kinesis.outputs.AnalyticsApplicationInputsSchema;
import com.pulumi.aws.kinesis.outputs.AnalyticsApplicationInputsStartingPositionConfiguration;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AnalyticsApplicationInputs {
    /**
     * @return The ARN of the Kinesis Analytics Application.
     * 
     */
    private final @Nullable String id;
    /**
     * @return The Kinesis Firehose configuration for the streaming source. Conflicts with `kinesis_stream`.
     * See Kinesis Firehose below for more details.
     * 
     */
    private final @Nullable AnalyticsApplicationInputsKinesisFirehose kinesisFirehose;
    /**
     * @return The Kinesis Stream configuration for the streaming source. Conflicts with `kinesis_firehose`.
     * See Kinesis Stream below for more details.
     * 
     */
    private final @Nullable AnalyticsApplicationInputsKinesisStream kinesisStream;
    /**
     * @return The Name Prefix to use when creating an in-application stream.
     * 
     */
    private final String namePrefix;
    /**
     * @return The number of Parallel in-application streams to create.
     * See Parallelism below for more details.
     * 
     */
    private final @Nullable AnalyticsApplicationInputsParallelism parallelism;
    /**
     * @return The Processing Configuration to transform records as they are received from the stream.
     * See Processing Configuration below for more details.
     * 
     */
    private final @Nullable AnalyticsApplicationInputsProcessingConfiguration processingConfiguration;
    /**
     * @return The Schema format of the data in the streaming source. See Source Schema below for more details.
     * 
     */
    private final AnalyticsApplicationInputsSchema schema;
    /**
     * @return The point at which the application starts processing records from the streaming source.
     * See Starting Position Configuration below for more details.
     * 
     */
    private final @Nullable List<AnalyticsApplicationInputsStartingPositionConfiguration> startingPositionConfigurations;
    private final @Nullable List<String> streamNames;

    @CustomType.Constructor
    private AnalyticsApplicationInputs(
        @CustomType.Parameter("id") @Nullable String id,
        @CustomType.Parameter("kinesisFirehose") @Nullable AnalyticsApplicationInputsKinesisFirehose kinesisFirehose,
        @CustomType.Parameter("kinesisStream") @Nullable AnalyticsApplicationInputsKinesisStream kinesisStream,
        @CustomType.Parameter("namePrefix") String namePrefix,
        @CustomType.Parameter("parallelism") @Nullable AnalyticsApplicationInputsParallelism parallelism,
        @CustomType.Parameter("processingConfiguration") @Nullable AnalyticsApplicationInputsProcessingConfiguration processingConfiguration,
        @CustomType.Parameter("schema") AnalyticsApplicationInputsSchema schema,
        @CustomType.Parameter("startingPositionConfigurations") @Nullable List<AnalyticsApplicationInputsStartingPositionConfiguration> startingPositionConfigurations,
        @CustomType.Parameter("streamNames") @Nullable List<String> streamNames) {
        this.id = id;
        this.kinesisFirehose = kinesisFirehose;
        this.kinesisStream = kinesisStream;
        this.namePrefix = namePrefix;
        this.parallelism = parallelism;
        this.processingConfiguration = processingConfiguration;
        this.schema = schema;
        this.startingPositionConfigurations = startingPositionConfigurations;
        this.streamNames = streamNames;
    }

    /**
     * @return The ARN of the Kinesis Analytics Application.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The Kinesis Firehose configuration for the streaming source. Conflicts with `kinesis_stream`.
     * See Kinesis Firehose below for more details.
     * 
     */
    public Optional<AnalyticsApplicationInputsKinesisFirehose> kinesisFirehose() {
        return Optional.ofNullable(this.kinesisFirehose);
    }
    /**
     * @return The Kinesis Stream configuration for the streaming source. Conflicts with `kinesis_firehose`.
     * See Kinesis Stream below for more details.
     * 
     */
    public Optional<AnalyticsApplicationInputsKinesisStream> kinesisStream() {
        return Optional.ofNullable(this.kinesisStream);
    }
    /**
     * @return The Name Prefix to use when creating an in-application stream.
     * 
     */
    public String namePrefix() {
        return this.namePrefix;
    }
    /**
     * @return The number of Parallel in-application streams to create.
     * See Parallelism below for more details.
     * 
     */
    public Optional<AnalyticsApplicationInputsParallelism> parallelism() {
        return Optional.ofNullable(this.parallelism);
    }
    /**
     * @return The Processing Configuration to transform records as they are received from the stream.
     * See Processing Configuration below for more details.
     * 
     */
    public Optional<AnalyticsApplicationInputsProcessingConfiguration> processingConfiguration() {
        return Optional.ofNullable(this.processingConfiguration);
    }
    /**
     * @return The Schema format of the data in the streaming source. See Source Schema below for more details.
     * 
     */
    public AnalyticsApplicationInputsSchema schema() {
        return this.schema;
    }
    /**
     * @return The point at which the application starts processing records from the streaming source.
     * See Starting Position Configuration below for more details.
     * 
     */
    public List<AnalyticsApplicationInputsStartingPositionConfiguration> startingPositionConfigurations() {
        return this.startingPositionConfigurations == null ? List.of() : this.startingPositionConfigurations;
    }
    public List<String> streamNames() {
        return this.streamNames == null ? List.of() : this.streamNames;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AnalyticsApplicationInputs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String id;
        private @Nullable AnalyticsApplicationInputsKinesisFirehose kinesisFirehose;
        private @Nullable AnalyticsApplicationInputsKinesisStream kinesisStream;
        private String namePrefix;
        private @Nullable AnalyticsApplicationInputsParallelism parallelism;
        private @Nullable AnalyticsApplicationInputsProcessingConfiguration processingConfiguration;
        private AnalyticsApplicationInputsSchema schema;
        private @Nullable List<AnalyticsApplicationInputsStartingPositionConfiguration> startingPositionConfigurations;
        private @Nullable List<String> streamNames;

        public Builder() {
    	      // Empty
        }

        public Builder(AnalyticsApplicationInputs defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.kinesisFirehose = defaults.kinesisFirehose;
    	      this.kinesisStream = defaults.kinesisStream;
    	      this.namePrefix = defaults.namePrefix;
    	      this.parallelism = defaults.parallelism;
    	      this.processingConfiguration = defaults.processingConfiguration;
    	      this.schema = defaults.schema;
    	      this.startingPositionConfigurations = defaults.startingPositionConfigurations;
    	      this.streamNames = defaults.streamNames;
        }

        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        public Builder kinesisFirehose(@Nullable AnalyticsApplicationInputsKinesisFirehose kinesisFirehose) {
            this.kinesisFirehose = kinesisFirehose;
            return this;
        }
        public Builder kinesisStream(@Nullable AnalyticsApplicationInputsKinesisStream kinesisStream) {
            this.kinesisStream = kinesisStream;
            return this;
        }
        public Builder namePrefix(String namePrefix) {
            this.namePrefix = Objects.requireNonNull(namePrefix);
            return this;
        }
        public Builder parallelism(@Nullable AnalyticsApplicationInputsParallelism parallelism) {
            this.parallelism = parallelism;
            return this;
        }
        public Builder processingConfiguration(@Nullable AnalyticsApplicationInputsProcessingConfiguration processingConfiguration) {
            this.processingConfiguration = processingConfiguration;
            return this;
        }
        public Builder schema(AnalyticsApplicationInputsSchema schema) {
            this.schema = Objects.requireNonNull(schema);
            return this;
        }
        public Builder startingPositionConfigurations(@Nullable List<AnalyticsApplicationInputsStartingPositionConfiguration> startingPositionConfigurations) {
            this.startingPositionConfigurations = startingPositionConfigurations;
            return this;
        }
        public Builder startingPositionConfigurations(AnalyticsApplicationInputsStartingPositionConfiguration... startingPositionConfigurations) {
            return startingPositionConfigurations(List.of(startingPositionConfigurations));
        }
        public Builder streamNames(@Nullable List<String> streamNames) {
            this.streamNames = streamNames;
            return this;
        }
        public Builder streamNames(String... streamNames) {
            return streamNames(List.of(streamNames));
        }        public AnalyticsApplicationInputs build() {
            return new AnalyticsApplicationInputs(id, kinesisFirehose, kinesisStream, namePrefix, parallelism, processingConfiguration, schema, startingPositionConfigurations, streamNames);
        }
    }
}