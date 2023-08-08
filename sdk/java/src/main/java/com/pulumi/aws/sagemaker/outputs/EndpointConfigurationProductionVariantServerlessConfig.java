// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EndpointConfigurationProductionVariantServerlessConfig {
    /**
     * @return The maximum number of concurrent invocations your serverless endpoint can process. Valid values are between `1` and `200`.
     * 
     */
    private Integer maxConcurrency;
    /**
     * @return The memory size of your serverless endpoint. Valid values are in 1 GB increments: `1024` MB, `2048` MB, `3072` MB, `4096` MB, `5120` MB, or `6144` MB.
     * 
     */
    private Integer memorySizeInMb;
    /**
     * @return The amount of provisioned concurrency to allocate for the serverless endpoint. Should be less than or equal to `max_concurrency`. Valid values are between `1` and `200`.
     * 
     */
    private @Nullable Integer provisionedConcurrency;

    private EndpointConfigurationProductionVariantServerlessConfig() {}
    /**
     * @return The maximum number of concurrent invocations your serverless endpoint can process. Valid values are between `1` and `200`.
     * 
     */
    public Integer maxConcurrency() {
        return this.maxConcurrency;
    }
    /**
     * @return The memory size of your serverless endpoint. Valid values are in 1 GB increments: `1024` MB, `2048` MB, `3072` MB, `4096` MB, `5120` MB, or `6144` MB.
     * 
     */
    public Integer memorySizeInMb() {
        return this.memorySizeInMb;
    }
    /**
     * @return The amount of provisioned concurrency to allocate for the serverless endpoint. Should be less than or equal to `max_concurrency`. Valid values are between `1` and `200`.
     * 
     */
    public Optional<Integer> provisionedConcurrency() {
        return Optional.ofNullable(this.provisionedConcurrency);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EndpointConfigurationProductionVariantServerlessConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer maxConcurrency;
        private Integer memorySizeInMb;
        private @Nullable Integer provisionedConcurrency;
        public Builder() {}
        public Builder(EndpointConfigurationProductionVariantServerlessConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxConcurrency = defaults.maxConcurrency;
    	      this.memorySizeInMb = defaults.memorySizeInMb;
    	      this.provisionedConcurrency = defaults.provisionedConcurrency;
        }

        @CustomType.Setter
        public Builder maxConcurrency(Integer maxConcurrency) {
            this.maxConcurrency = Objects.requireNonNull(maxConcurrency);
            return this;
        }
        @CustomType.Setter
        public Builder memorySizeInMb(Integer memorySizeInMb) {
            this.memorySizeInMb = Objects.requireNonNull(memorySizeInMb);
            return this;
        }
        @CustomType.Setter
        public Builder provisionedConcurrency(@Nullable Integer provisionedConcurrency) {
            this.provisionedConcurrency = provisionedConcurrency;
            return this;
        }
        public EndpointConfigurationProductionVariantServerlessConfig build() {
            final var o = new EndpointConfigurationProductionVariantServerlessConfig();
            o.maxConcurrency = maxConcurrency;
            o.memorySizeInMb = memorySizeInMb;
            o.provisionedConcurrency = provisionedConcurrency;
            return o;
        }
    }
}
