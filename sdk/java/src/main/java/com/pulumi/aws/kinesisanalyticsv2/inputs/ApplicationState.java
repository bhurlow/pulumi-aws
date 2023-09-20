// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesisanalyticsv2.inputs;

import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationCloudwatchLoggingOptionsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationState extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationState Empty = new ApplicationState();

    /**
     * The application&#39;s configuration
     * 
     */
    @Import(name="applicationConfiguration")
    private @Nullable Output<ApplicationApplicationConfigurationArgs> applicationConfiguration;

    /**
     * @return The application&#39;s configuration
     * 
     */
    public Optional<Output<ApplicationApplicationConfigurationArgs>> applicationConfiguration() {
        return Optional.ofNullable(this.applicationConfiguration);
    }

    /**
     * The ARN of the application.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the application.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * A CloudWatch log stream to monitor application configuration errors.
     * 
     */
    @Import(name="cloudwatchLoggingOptions")
    private @Nullable Output<ApplicationCloudwatchLoggingOptionsArgs> cloudwatchLoggingOptions;

    /**
     * @return A CloudWatch log stream to monitor application configuration errors.
     * 
     */
    public Optional<Output<ApplicationCloudwatchLoggingOptionsArgs>> cloudwatchLoggingOptions() {
        return Optional.ofNullable(this.cloudwatchLoggingOptions);
    }

    /**
     * The current timestamp when the application was created.
     * 
     */
    @Import(name="createTimestamp")
    private @Nullable Output<String> createTimestamp;

    /**
     * @return The current timestamp when the application was created.
     * 
     */
    public Optional<Output<String>> createTimestamp() {
        return Optional.ofNullable(this.createTimestamp);
    }

    /**
     * A summary description of the application.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A summary description of the application.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether to force stop an unresponsive Flink-based application.
     * 
     */
    @Import(name="forceStop")
    private @Nullable Output<Boolean> forceStop;

    /**
     * @return Whether to force stop an unresponsive Flink-based application.
     * 
     */
    public Optional<Output<Boolean>> forceStop() {
        return Optional.ofNullable(this.forceStop);
    }

    /**
     * The current timestamp when the application was last updated.
     * 
     */
    @Import(name="lastUpdateTimestamp")
    private @Nullable Output<String> lastUpdateTimestamp;

    /**
     * @return The current timestamp when the application was last updated.
     * 
     */
    public Optional<Output<String>> lastUpdateTimestamp() {
        return Optional.ofNullable(this.lastUpdateTimestamp);
    }

    /**
     * The name of the application.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the application.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`, `FLINK-1_11`, `FLINK-1_13`, `FLINK-1_15`.
     * 
     */
    @Import(name="runtimeEnvironment")
    private @Nullable Output<String> runtimeEnvironment;

    /**
     * @return The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`, `FLINK-1_11`, `FLINK-1_13`, `FLINK-1_15`.
     * 
     */
    public Optional<Output<String>> runtimeEnvironment() {
        return Optional.ofNullable(this.runtimeEnvironment);
    }

    /**
     * The ARN of the IAM role used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
     * 
     */
    @Import(name="serviceExecutionRole")
    private @Nullable Output<String> serviceExecutionRole;

    /**
     * @return The ARN of the IAM role used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
     * 
     */
    public Optional<Output<String>> serviceExecutionRole() {
        return Optional.ofNullable(this.serviceExecutionRole);
    }

    /**
     * Whether to start or stop the application.
     * 
     */
    @Import(name="startApplication")
    private @Nullable Output<Boolean> startApplication;

    /**
     * @return Whether to start or stop the application.
     * 
     */
    public Optional<Output<Boolean>> startApplication() {
        return Optional.ofNullable(this.startApplication);
    }

    /**
     * The status of the application.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the application.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A map of tags to assign to the application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
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
     * The current application version. Kinesis Data Analytics updates the `version_id` each time the application is updated.
     * 
     */
    @Import(name="versionId")
    private @Nullable Output<Integer> versionId;

    /**
     * @return The current application version. Kinesis Data Analytics updates the `version_id` each time the application is updated.
     * 
     */
    public Optional<Output<Integer>> versionId() {
        return Optional.ofNullable(this.versionId);
    }

    private ApplicationState() {}

    private ApplicationState(ApplicationState $) {
        this.applicationConfiguration = $.applicationConfiguration;
        this.arn = $.arn;
        this.cloudwatchLoggingOptions = $.cloudwatchLoggingOptions;
        this.createTimestamp = $.createTimestamp;
        this.description = $.description;
        this.forceStop = $.forceStop;
        this.lastUpdateTimestamp = $.lastUpdateTimestamp;
        this.name = $.name;
        this.runtimeEnvironment = $.runtimeEnvironment;
        this.serviceExecutionRole = $.serviceExecutionRole;
        this.startApplication = $.startApplication;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.versionId = $.versionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationState $;

        public Builder() {
            $ = new ApplicationState();
        }

        public Builder(ApplicationState defaults) {
            $ = new ApplicationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationConfiguration The application&#39;s configuration
         * 
         * @return builder
         * 
         */
        public Builder applicationConfiguration(@Nullable Output<ApplicationApplicationConfigurationArgs> applicationConfiguration) {
            $.applicationConfiguration = applicationConfiguration;
            return this;
        }

        /**
         * @param applicationConfiguration The application&#39;s configuration
         * 
         * @return builder
         * 
         */
        public Builder applicationConfiguration(ApplicationApplicationConfigurationArgs applicationConfiguration) {
            return applicationConfiguration(Output.of(applicationConfiguration));
        }

        /**
         * @param arn The ARN of the application.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the application.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param cloudwatchLoggingOptions A CloudWatch log stream to monitor application configuration errors.
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchLoggingOptions(@Nullable Output<ApplicationCloudwatchLoggingOptionsArgs> cloudwatchLoggingOptions) {
            $.cloudwatchLoggingOptions = cloudwatchLoggingOptions;
            return this;
        }

        /**
         * @param cloudwatchLoggingOptions A CloudWatch log stream to monitor application configuration errors.
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchLoggingOptions(ApplicationCloudwatchLoggingOptionsArgs cloudwatchLoggingOptions) {
            return cloudwatchLoggingOptions(Output.of(cloudwatchLoggingOptions));
        }

        /**
         * @param createTimestamp The current timestamp when the application was created.
         * 
         * @return builder
         * 
         */
        public Builder createTimestamp(@Nullable Output<String> createTimestamp) {
            $.createTimestamp = createTimestamp;
            return this;
        }

        /**
         * @param createTimestamp The current timestamp when the application was created.
         * 
         * @return builder
         * 
         */
        public Builder createTimestamp(String createTimestamp) {
            return createTimestamp(Output.of(createTimestamp));
        }

        /**
         * @param description A summary description of the application.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A summary description of the application.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param forceStop Whether to force stop an unresponsive Flink-based application.
         * 
         * @return builder
         * 
         */
        public Builder forceStop(@Nullable Output<Boolean> forceStop) {
            $.forceStop = forceStop;
            return this;
        }

        /**
         * @param forceStop Whether to force stop an unresponsive Flink-based application.
         * 
         * @return builder
         * 
         */
        public Builder forceStop(Boolean forceStop) {
            return forceStop(Output.of(forceStop));
        }

        /**
         * @param lastUpdateTimestamp The current timestamp when the application was last updated.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdateTimestamp(@Nullable Output<String> lastUpdateTimestamp) {
            $.lastUpdateTimestamp = lastUpdateTimestamp;
            return this;
        }

        /**
         * @param lastUpdateTimestamp The current timestamp when the application was last updated.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdateTimestamp(String lastUpdateTimestamp) {
            return lastUpdateTimestamp(Output.of(lastUpdateTimestamp));
        }

        /**
         * @param name The name of the application.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the application.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param runtimeEnvironment The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`, `FLINK-1_11`, `FLINK-1_13`, `FLINK-1_15`.
         * 
         * @return builder
         * 
         */
        public Builder runtimeEnvironment(@Nullable Output<String> runtimeEnvironment) {
            $.runtimeEnvironment = runtimeEnvironment;
            return this;
        }

        /**
         * @param runtimeEnvironment The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`, `FLINK-1_11`, `FLINK-1_13`, `FLINK-1_15`.
         * 
         * @return builder
         * 
         */
        public Builder runtimeEnvironment(String runtimeEnvironment) {
            return runtimeEnvironment(Output.of(runtimeEnvironment));
        }

        /**
         * @param serviceExecutionRole The ARN of the IAM role used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
         * 
         * @return builder
         * 
         */
        public Builder serviceExecutionRole(@Nullable Output<String> serviceExecutionRole) {
            $.serviceExecutionRole = serviceExecutionRole;
            return this;
        }

        /**
         * @param serviceExecutionRole The ARN of the IAM role used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
         * 
         * @return builder
         * 
         */
        public Builder serviceExecutionRole(String serviceExecutionRole) {
            return serviceExecutionRole(Output.of(serviceExecutionRole));
        }

        /**
         * @param startApplication Whether to start or stop the application.
         * 
         * @return builder
         * 
         */
        public Builder startApplication(@Nullable Output<Boolean> startApplication) {
            $.startApplication = startApplication;
            return this;
        }

        /**
         * @param startApplication Whether to start or stop the application.
         * 
         * @return builder
         * 
         */
        public Builder startApplication(Boolean startApplication) {
            return startApplication(Output.of(startApplication));
        }

        /**
         * @param status The status of the application.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the application.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A map of tags to assign to the application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
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
         * @param versionId The current application version. Kinesis Data Analytics updates the `version_id` each time the application is updated.
         * 
         * @return builder
         * 
         */
        public Builder versionId(@Nullable Output<Integer> versionId) {
            $.versionId = versionId;
            return this;
        }

        /**
         * @param versionId The current application version. Kinesis Data Analytics updates the `version_id` each time the application is updated.
         * 
         * @return builder
         * 
         */
        public Builder versionId(Integer versionId) {
            return versionId(Output.of(versionId));
        }

        public ApplicationState build() {
            return $;
        }
    }

}
