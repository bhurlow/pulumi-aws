// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticbeanstalk;

import com.pulumi.aws.elasticbeanstalk.inputs.ConfigurationTemplateSettingArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConfigurationTemplateArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConfigurationTemplateArgs Empty = new ConfigurationTemplateArgs();

    /**
     * name of the application to associate with this configuration template
     * 
     */
    @Import(name="application", required=true)
    private Output<String> application;

    /**
     * @return name of the application to associate with this configuration template
     * 
     */
    public Output<String> application() {
        return this.application;
    }

    /**
     * Short description of the Template
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Short description of the Template
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the environment used with this configuration template
     * 
     */
    @Import(name="environmentId")
    private @Nullable Output<String> environmentId;

    /**
     * @return The ID of the environment used with this configuration template
     * 
     */
    public Optional<Output<String>> environmentId() {
        return Optional.ofNullable(this.environmentId);
    }

    /**
     * A unique name for this Template.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for this Template.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Option settings to configure the new Environment. These
     * override specific values that are set as defaults. The format is detailed
     * below in Option Settings
     * 
     */
    @Import(name="settings")
    private @Nullable Output<List<ConfigurationTemplateSettingArgs>> settings;

    /**
     * @return Option settings to configure the new Environment. These
     * override specific values that are set as defaults. The format is detailed
     * below in Option Settings
     * 
     */
    public Optional<Output<List<ConfigurationTemplateSettingArgs>>> settings() {
        return Optional.ofNullable(this.settings);
    }

    /**
     * A solution stack to base your Template
     * off of. Example stacks can be found in the [Amazon API documentation][1]
     * 
     */
    @Import(name="solutionStackName")
    private @Nullable Output<String> solutionStackName;

    /**
     * @return A solution stack to base your Template
     * off of. Example stacks can be found in the [Amazon API documentation][1]
     * 
     */
    public Optional<Output<String>> solutionStackName() {
        return Optional.ofNullable(this.solutionStackName);
    }

    private ConfigurationTemplateArgs() {}

    private ConfigurationTemplateArgs(ConfigurationTemplateArgs $) {
        this.application = $.application;
        this.description = $.description;
        this.environmentId = $.environmentId;
        this.name = $.name;
        this.settings = $.settings;
        this.solutionStackName = $.solutionStackName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConfigurationTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConfigurationTemplateArgs $;

        public Builder() {
            $ = new ConfigurationTemplateArgs();
        }

        public Builder(ConfigurationTemplateArgs defaults) {
            $ = new ConfigurationTemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param application name of the application to associate with this configuration template
         * 
         * @return builder
         * 
         */
        public Builder application(Output<String> application) {
            $.application = application;
            return this;
        }

        /**
         * @param application name of the application to associate with this configuration template
         * 
         * @return builder
         * 
         */
        public Builder application(String application) {
            return application(Output.of(application));
        }

        /**
         * @param description Short description of the Template
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Short description of the Template
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param environmentId The ID of the environment used with this configuration template
         * 
         * @return builder
         * 
         */
        public Builder environmentId(@Nullable Output<String> environmentId) {
            $.environmentId = environmentId;
            return this;
        }

        /**
         * @param environmentId The ID of the environment used with this configuration template
         * 
         * @return builder
         * 
         */
        public Builder environmentId(String environmentId) {
            return environmentId(Output.of(environmentId));
        }

        /**
         * @param name A unique name for this Template.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for this Template.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param settings Option settings to configure the new Environment. These
         * override specific values that are set as defaults. The format is detailed
         * below in Option Settings
         * 
         * @return builder
         * 
         */
        public Builder settings(@Nullable Output<List<ConfigurationTemplateSettingArgs>> settings) {
            $.settings = settings;
            return this;
        }

        /**
         * @param settings Option settings to configure the new Environment. These
         * override specific values that are set as defaults. The format is detailed
         * below in Option Settings
         * 
         * @return builder
         * 
         */
        public Builder settings(List<ConfigurationTemplateSettingArgs> settings) {
            return settings(Output.of(settings));
        }

        /**
         * @param settings Option settings to configure the new Environment. These
         * override specific values that are set as defaults. The format is detailed
         * below in Option Settings
         * 
         * @return builder
         * 
         */
        public Builder settings(ConfigurationTemplateSettingArgs... settings) {
            return settings(List.of(settings));
        }

        /**
         * @param solutionStackName A solution stack to base your Template
         * off of. Example stacks can be found in the [Amazon API documentation][1]
         * 
         * @return builder
         * 
         */
        public Builder solutionStackName(@Nullable Output<String> solutionStackName) {
            $.solutionStackName = solutionStackName;
            return this;
        }

        /**
         * @param solutionStackName A solution stack to base your Template
         * off of. Example stacks can be found in the [Amazon API documentation][1]
         * 
         * @return builder
         * 
         */
        public Builder solutionStackName(String solutionStackName) {
            return solutionStackName(Output.of(solutionStackName));
        }

        public ConfigurationTemplateArgs build() {
            $.application = Objects.requireNonNull($.application, "expected parameter 'application' to be non-null");
            return $;
        }
    }

}
