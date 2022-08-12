// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.mskconnect.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkerConfigurationState extends com.pulumi.resources.ResourceArgs {

    public static final WorkerConfigurationState Empty = new WorkerConfigurationState();

    /**
     * the Amazon Resource Name (ARN) of the worker configuration.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return the Amazon Resource Name (ARN) of the worker configuration.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * A summary description of the worker configuration.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A summary description of the worker configuration.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * an ID of the latest successfully created revision of the worker configuration.
     * 
     */
    @Import(name="latestRevision")
    private @Nullable Output<Integer> latestRevision;

    /**
     * @return an ID of the latest successfully created revision of the worker configuration.
     * 
     */
    public Optional<Output<Integer>> latestRevision() {
        return Optional.ofNullable(this.latestRevision);
    }

    /**
     * The name of the worker configuration.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the worker configuration.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
     * 
     */
    @Import(name="propertiesFileContent")
    private @Nullable Output<String> propertiesFileContent;

    /**
     * @return Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
     * 
     */
    public Optional<Output<String>> propertiesFileContent() {
        return Optional.ofNullable(this.propertiesFileContent);
    }

    private WorkerConfigurationState() {}

    private WorkerConfigurationState(WorkerConfigurationState $) {
        this.arn = $.arn;
        this.description = $.description;
        this.latestRevision = $.latestRevision;
        this.name = $.name;
        this.propertiesFileContent = $.propertiesFileContent;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkerConfigurationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkerConfigurationState $;

        public Builder() {
            $ = new WorkerConfigurationState();
        }

        public Builder(WorkerConfigurationState defaults) {
            $ = new WorkerConfigurationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn the Amazon Resource Name (ARN) of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn the Amazon Resource Name (ARN) of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description A summary description of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A summary description of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param latestRevision an ID of the latest successfully created revision of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder latestRevision(@Nullable Output<Integer> latestRevision) {
            $.latestRevision = latestRevision;
            return this;
        }

        /**
         * @param latestRevision an ID of the latest successfully created revision of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder latestRevision(Integer latestRevision) {
            return latestRevision(Output.of(latestRevision));
        }

        /**
         * @param name The name of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the worker configuration.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param propertiesFileContent Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
         * 
         * @return builder
         * 
         */
        public Builder propertiesFileContent(@Nullable Output<String> propertiesFileContent) {
            $.propertiesFileContent = propertiesFileContent;
            return this;
        }

        /**
         * @param propertiesFileContent Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
         * 
         * @return builder
         * 
         */
        public Builder propertiesFileContent(String propertiesFileContent) {
            return propertiesFileContent(Output.of(propertiesFileContent));
        }

        public WorkerConfigurationState build() {
            return $;
        }
    }

}