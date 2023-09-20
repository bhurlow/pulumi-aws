// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appconfig.inputs;

import com.pulumi.aws.appconfig.inputs.ExtensionActionPointArgs;
import com.pulumi.aws.appconfig.inputs.ExtensionParameterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExtensionState extends com.pulumi.resources.ResourceArgs {

    public static final ExtensionState Empty = new ExtensionState();

    /**
     * The action points defined in the extension. Detailed below.
     * 
     */
    @Import(name="actionPoints")
    private @Nullable Output<List<ExtensionActionPointArgs>> actionPoints;

    /**
     * @return The action points defined in the extension. Detailed below.
     * 
     */
    public Optional<Output<List<ExtensionActionPointArgs>>> actionPoints() {
        return Optional.ofNullable(this.actionPoints);
    }

    /**
     * ARN of the AppConfig Extension.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the AppConfig Extension.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Information about the extension.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Information about the extension.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<List<ExtensionParameterArgs>> parameters;

    /**
     * @return The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
     * 
     */
    public Optional<Output<List<ExtensionParameterArgs>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The version number for the extension.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Integer> version;

    /**
     * @return The version number for the extension.
     * 
     */
    public Optional<Output<Integer>> version() {
        return Optional.ofNullable(this.version);
    }

    private ExtensionState() {}

    private ExtensionState(ExtensionState $) {
        this.actionPoints = $.actionPoints;
        this.arn = $.arn;
        this.description = $.description;
        this.name = $.name;
        this.parameters = $.parameters;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExtensionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExtensionState $;

        public Builder() {
            $ = new ExtensionState();
        }

        public Builder(ExtensionState defaults) {
            $ = new ExtensionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param actionPoints The action points defined in the extension. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder actionPoints(@Nullable Output<List<ExtensionActionPointArgs>> actionPoints) {
            $.actionPoints = actionPoints;
            return this;
        }

        /**
         * @param actionPoints The action points defined in the extension. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder actionPoints(List<ExtensionActionPointArgs> actionPoints) {
            return actionPoints(Output.of(actionPoints));
        }

        /**
         * @param actionPoints The action points defined in the extension. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder actionPoints(ExtensionActionPointArgs... actionPoints) {
            return actionPoints(List.of(actionPoints));
        }

        /**
         * @param arn ARN of the AppConfig Extension.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the AppConfig Extension.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description Information about the extension.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Information about the extension.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parameters The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<List<ExtensionParameterArgs>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder parameters(List<ExtensionParameterArgs> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param parameters The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder parameters(ExtensionParameterArgs... parameters) {
            return parameters(List.of(parameters));
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
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
         * @param version The version number for the extension.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version number for the extension.
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        public ExtensionState build() {
            return $;
        }
    }

}
