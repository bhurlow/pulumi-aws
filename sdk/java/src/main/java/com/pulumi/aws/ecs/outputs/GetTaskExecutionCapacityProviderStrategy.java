// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTaskExecutionCapacityProviderStrategy {
    /**
     * @return The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to `0`.
     * 
     */
    private @Nullable Integer base;
    /**
     * @return Name of the capacity provider.
     * 
     */
    private String capacityProvider;
    /**
     * @return The relative percentage of the total number of launched tasks that should use the specified capacity provider. The `weight` value is taken into consideration after the `base` count of tasks has been satisfied. Defaults to `0`.
     * 
     */
    private @Nullable Integer weight;

    private GetTaskExecutionCapacityProviderStrategy() {}
    /**
     * @return The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to `0`.
     * 
     */
    public Optional<Integer> base() {
        return Optional.ofNullable(this.base);
    }
    /**
     * @return Name of the capacity provider.
     * 
     */
    public String capacityProvider() {
        return this.capacityProvider;
    }
    /**
     * @return The relative percentage of the total number of launched tasks that should use the specified capacity provider. The `weight` value is taken into consideration after the `base` count of tasks has been satisfied. Defaults to `0`.
     * 
     */
    public Optional<Integer> weight() {
        return Optional.ofNullable(this.weight);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTaskExecutionCapacityProviderStrategy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer base;
        private String capacityProvider;
        private @Nullable Integer weight;
        public Builder() {}
        public Builder(GetTaskExecutionCapacityProviderStrategy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.base = defaults.base;
    	      this.capacityProvider = defaults.capacityProvider;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder base(@Nullable Integer base) {
            this.base = base;
            return this;
        }
        @CustomType.Setter
        public Builder capacityProvider(String capacityProvider) {
            this.capacityProvider = Objects.requireNonNull(capacityProvider);
            return this;
        }
        @CustomType.Setter
        public Builder weight(@Nullable Integer weight) {
            this.weight = weight;
            return this;
        }
        public GetTaskExecutionCapacityProviderStrategy build() {
            final var o = new GetTaskExecutionCapacityProviderStrategy();
            o.base = base;
            o.capacityProvider = capacityProvider;
            o.weight = weight;
            return o;
        }
    }
}
