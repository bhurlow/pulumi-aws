// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticbeanstalk;

import com.pulumi.aws.elasticbeanstalk.inputs.ApplicationAppversionLifecycleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationArgs Empty = new ApplicationArgs();

    @Import(name="appversionLifecycle")
    private @Nullable Output<ApplicationAppversionLifecycleArgs> appversionLifecycle;

    public Optional<Output<ApplicationAppversionLifecycleArgs>> appversionLifecycle() {
        return Optional.ofNullable(this.appversionLifecycle);
    }

    /**
     * Short description of the application
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Short description of the application
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the application, must be unique within your account
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the application, must be unique within your account
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Key-value map of tags for the Elastic Beanstalk Application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of tags for the Elastic Beanstalk Application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ApplicationArgs() {}

    private ApplicationArgs(ApplicationArgs $) {
        this.appversionLifecycle = $.appversionLifecycle;
        this.description = $.description;
        this.name = $.name;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationArgs $;

        public Builder() {
            $ = new ApplicationArgs();
        }

        public Builder(ApplicationArgs defaults) {
            $ = new ApplicationArgs(Objects.requireNonNull(defaults));
        }

        public Builder appversionLifecycle(@Nullable Output<ApplicationAppversionLifecycleArgs> appversionLifecycle) {
            $.appversionLifecycle = appversionLifecycle;
            return this;
        }

        public Builder appversionLifecycle(ApplicationAppversionLifecycleArgs appversionLifecycle) {
            return appversionLifecycle(Output.of(appversionLifecycle));
        }

        /**
         * @param description Short description of the application
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Short description of the application
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the application, must be unique within your account
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the application, must be unique within your account
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags Key-value map of tags for the Elastic Beanstalk Application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of tags for the Elastic Beanstalk Application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public ApplicationArgs build() {
            return $;
        }
    }

}