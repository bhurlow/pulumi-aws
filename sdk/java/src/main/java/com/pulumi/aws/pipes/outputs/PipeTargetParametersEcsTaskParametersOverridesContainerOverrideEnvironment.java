// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment {
    /**
     * @return Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    private @Nullable String name;
    /**
     * @return The value of the key-value pair. For environment variables, this is the value of the environment variable.
     * 
     */
    private @Nullable String value;

    private PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment() {}
    /**
     * @return Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The value of the key-value pair. For environment variables, this is the value of the environment variable.
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable String value;
        public Builder() {}
        public Builder(PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        public PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment build() {
            final var o = new PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment();
            o.name = name;
            o.value = value;
            return o;
        }
    }
}
