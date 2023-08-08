// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.outputs;

import com.pulumi.aws.ec2.outputs.InstanceInstanceMarketOptionsSpotOptions;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceInstanceMarketOptions {
    /**
     * @return Type of market for the instance. Valid value is `spot`. Defaults to `spot`.
     * 
     */
    private @Nullable String marketType;
    /**
     * @return Block to configure the options for Spot Instances. See Spot Options below for details on attributes.
     * 
     */
    private @Nullable InstanceInstanceMarketOptionsSpotOptions spotOptions;

    private InstanceInstanceMarketOptions() {}
    /**
     * @return Type of market for the instance. Valid value is `spot`. Defaults to `spot`.
     * 
     */
    public Optional<String> marketType() {
        return Optional.ofNullable(this.marketType);
    }
    /**
     * @return Block to configure the options for Spot Instances. See Spot Options below for details on attributes.
     * 
     */
    public Optional<InstanceInstanceMarketOptionsSpotOptions> spotOptions() {
        return Optional.ofNullable(this.spotOptions);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceInstanceMarketOptions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String marketType;
        private @Nullable InstanceInstanceMarketOptionsSpotOptions spotOptions;
        public Builder() {}
        public Builder(InstanceInstanceMarketOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.marketType = defaults.marketType;
    	      this.spotOptions = defaults.spotOptions;
        }

        @CustomType.Setter
        public Builder marketType(@Nullable String marketType) {
            this.marketType = marketType;
            return this;
        }
        @CustomType.Setter
        public Builder spotOptions(@Nullable InstanceInstanceMarketOptionsSpotOptions spotOptions) {
            this.spotOptions = spotOptions;
            return this;
        }
        public InstanceInstanceMarketOptions build() {
            final var o = new InstanceInstanceMarketOptions();
            o.marketType = marketType;
            o.spotOptions = spotOptions;
            return o;
        }
    }
}
