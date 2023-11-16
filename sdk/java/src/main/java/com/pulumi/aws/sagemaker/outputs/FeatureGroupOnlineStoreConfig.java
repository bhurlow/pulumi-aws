// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.aws.sagemaker.outputs.FeatureGroupOnlineStoreConfigSecurityConfig;
import com.pulumi.aws.sagemaker.outputs.FeatureGroupOnlineStoreConfigTtlDuration;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FeatureGroupOnlineStoreConfig {
    /**
     * @return Set to `true` to disable the automatic creation of an AWS Glue table when configuring an OfflineStore.
     * 
     */
    private @Nullable Boolean enableOnlineStore;
    /**
     * @return Security config for at-rest encryption of your OnlineStore. See Security Config Below.
     * 
     */
    private @Nullable FeatureGroupOnlineStoreConfigSecurityConfig securityConfig;
    /**
     * @return Option for different tiers of low latency storage for real-time data retrieval. Valid values are `Standard`, or `InMemory`.
     * 
     */
    private @Nullable String storageType;
    /**
     * @return Time to live duration, where the record is hard deleted after the expiration time is reached; ExpiresAt = EventTime + TtlDuration.. See TTl Duration Below.
     * 
     */
    private @Nullable FeatureGroupOnlineStoreConfigTtlDuration ttlDuration;

    private FeatureGroupOnlineStoreConfig() {}
    /**
     * @return Set to `true` to disable the automatic creation of an AWS Glue table when configuring an OfflineStore.
     * 
     */
    public Optional<Boolean> enableOnlineStore() {
        return Optional.ofNullable(this.enableOnlineStore);
    }
    /**
     * @return Security config for at-rest encryption of your OnlineStore. See Security Config Below.
     * 
     */
    public Optional<FeatureGroupOnlineStoreConfigSecurityConfig> securityConfig() {
        return Optional.ofNullable(this.securityConfig);
    }
    /**
     * @return Option for different tiers of low latency storage for real-time data retrieval. Valid values are `Standard`, or `InMemory`.
     * 
     */
    public Optional<String> storageType() {
        return Optional.ofNullable(this.storageType);
    }
    /**
     * @return Time to live duration, where the record is hard deleted after the expiration time is reached; ExpiresAt = EventTime + TtlDuration.. See TTl Duration Below.
     * 
     */
    public Optional<FeatureGroupOnlineStoreConfigTtlDuration> ttlDuration() {
        return Optional.ofNullable(this.ttlDuration);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FeatureGroupOnlineStoreConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableOnlineStore;
        private @Nullable FeatureGroupOnlineStoreConfigSecurityConfig securityConfig;
        private @Nullable String storageType;
        private @Nullable FeatureGroupOnlineStoreConfigTtlDuration ttlDuration;
        public Builder() {}
        public Builder(FeatureGroupOnlineStoreConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableOnlineStore = defaults.enableOnlineStore;
    	      this.securityConfig = defaults.securityConfig;
    	      this.storageType = defaults.storageType;
    	      this.ttlDuration = defaults.ttlDuration;
        }

        @CustomType.Setter
        public Builder enableOnlineStore(@Nullable Boolean enableOnlineStore) {
            this.enableOnlineStore = enableOnlineStore;
            return this;
        }
        @CustomType.Setter
        public Builder securityConfig(@Nullable FeatureGroupOnlineStoreConfigSecurityConfig securityConfig) {
            this.securityConfig = securityConfig;
            return this;
        }
        @CustomType.Setter
        public Builder storageType(@Nullable String storageType) {
            this.storageType = storageType;
            return this;
        }
        @CustomType.Setter
        public Builder ttlDuration(@Nullable FeatureGroupOnlineStoreConfigTtlDuration ttlDuration) {
            this.ttlDuration = ttlDuration;
            return this;
        }
        public FeatureGroupOnlineStoreConfig build() {
            final var o = new FeatureGroupOnlineStoreConfig();
            o.enableOnlineStore = enableOnlineStore;
            o.securityConfig = securityConfig;
            o.storageType = storageType;
            o.ttlDuration = ttlDuration;
            return o;
        }
    }
}
