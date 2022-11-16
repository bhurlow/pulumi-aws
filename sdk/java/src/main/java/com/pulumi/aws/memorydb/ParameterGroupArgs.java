// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.memorydb;

import com.pulumi.aws.memorydb.inputs.ParameterGroupParameterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ParameterGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ParameterGroupArgs Empty = new ParameterGroupArgs();

    /**
     * Description for the parameter group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the parameter group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The engine version that the parameter group can be used with.
     * 
     */
    @Import(name="family", required=true)
    private Output<String> family;

    /**
     * @return The engine version that the parameter group can be used with.
     * 
     */
    public Output<String> family() {
        return this.family;
    }

    /**
     * The name of the parameter.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the parameter.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<List<ParameterGroupParameterArgs>> parameters;

    /**
     * @return Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
     * 
     */
    public Optional<Output<List<ParameterGroupParameterArgs>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ParameterGroupArgs() {}

    private ParameterGroupArgs(ParameterGroupArgs $) {
        this.description = $.description;
        this.family = $.family;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.parameters = $.parameters;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ParameterGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ParameterGroupArgs $;

        public Builder() {
            $ = new ParameterGroupArgs();
        }

        public Builder(ParameterGroupArgs defaults) {
            $ = new ParameterGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description for the parameter group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the parameter group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param family The engine version that the parameter group can be used with.
         * 
         * @return builder
         * 
         */
        public Builder family(Output<String> family) {
            $.family = family;
            return this;
        }

        /**
         * @param family The engine version that the parameter group can be used with.
         * 
         * @return builder
         * 
         */
        public Builder family(String family) {
            return family(Output.of(family));
        }

        /**
         * @param name The name of the parameter.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the parameter.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param parameters Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<List<ParameterGroupParameterArgs>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder parameters(List<ParameterGroupParameterArgs> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param parameters Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder parameters(ParameterGroupParameterArgs... parameters) {
            return parameters(List.of(parameters));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public ParameterGroupArgs build() {
            $.family = Objects.requireNonNull($.family, "expected parameter 'family' to be non-null");
            return $;
        }
    }

}
