// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesisanalyticsv2.inputs;

import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationRunConfigurationFlinkRunConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationApplicationConfigurationRunConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationApplicationConfigurationRunConfigurationArgs Empty = new ApplicationApplicationConfigurationRunConfigurationArgs();

    /**
     * The restore behavior of a restarting application.
     * 
     */
    @Import(name="applicationRestoreConfiguration")
    private @Nullable Output<ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfigurationArgs> applicationRestoreConfiguration;

    /**
     * @return The restore behavior of a restarting application.
     * 
     */
    public Optional<Output<ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfigurationArgs>> applicationRestoreConfiguration() {
        return Optional.ofNullable(this.applicationRestoreConfiguration);
    }

    /**
     * The starting parameters for a Flink-based Kinesis Data Analytics application.
     * 
     */
    @Import(name="flinkRunConfiguration")
    private @Nullable Output<ApplicationApplicationConfigurationRunConfigurationFlinkRunConfigurationArgs> flinkRunConfiguration;

    /**
     * @return The starting parameters for a Flink-based Kinesis Data Analytics application.
     * 
     */
    public Optional<Output<ApplicationApplicationConfigurationRunConfigurationFlinkRunConfigurationArgs>> flinkRunConfiguration() {
        return Optional.ofNullable(this.flinkRunConfiguration);
    }

    private ApplicationApplicationConfigurationRunConfigurationArgs() {}

    private ApplicationApplicationConfigurationRunConfigurationArgs(ApplicationApplicationConfigurationRunConfigurationArgs $) {
        this.applicationRestoreConfiguration = $.applicationRestoreConfiguration;
        this.flinkRunConfiguration = $.flinkRunConfiguration;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationApplicationConfigurationRunConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationApplicationConfigurationRunConfigurationArgs $;

        public Builder() {
            $ = new ApplicationApplicationConfigurationRunConfigurationArgs();
        }

        public Builder(ApplicationApplicationConfigurationRunConfigurationArgs defaults) {
            $ = new ApplicationApplicationConfigurationRunConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationRestoreConfiguration The restore behavior of a restarting application.
         * 
         * @return builder
         * 
         */
        public Builder applicationRestoreConfiguration(@Nullable Output<ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfigurationArgs> applicationRestoreConfiguration) {
            $.applicationRestoreConfiguration = applicationRestoreConfiguration;
            return this;
        }

        /**
         * @param applicationRestoreConfiguration The restore behavior of a restarting application.
         * 
         * @return builder
         * 
         */
        public Builder applicationRestoreConfiguration(ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfigurationArgs applicationRestoreConfiguration) {
            return applicationRestoreConfiguration(Output.of(applicationRestoreConfiguration));
        }

        /**
         * @param flinkRunConfiguration The starting parameters for a Flink-based Kinesis Data Analytics application.
         * 
         * @return builder
         * 
         */
        public Builder flinkRunConfiguration(@Nullable Output<ApplicationApplicationConfigurationRunConfigurationFlinkRunConfigurationArgs> flinkRunConfiguration) {
            $.flinkRunConfiguration = flinkRunConfiguration;
            return this;
        }

        /**
         * @param flinkRunConfiguration The starting parameters for a Flink-based Kinesis Data Analytics application.
         * 
         * @return builder
         * 
         */
        public Builder flinkRunConfiguration(ApplicationApplicationConfigurationRunConfigurationFlinkRunConfigurationArgs flinkRunConfiguration) {
            return flinkRunConfiguration(Output.of(flinkRunConfiguration));
        }

        public ApplicationApplicationConfigurationRunConfigurationArgs build() {
            return $;
        }
    }

}
