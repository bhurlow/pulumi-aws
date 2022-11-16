// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsMediaStoreSettings {
    /**
     * @return Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
     * 
     */
    private @Nullable Integer connectionRetryInterval;
    private @Nullable Integer filecacheDuration;
    private @Nullable String mediaStoreStorageClass;
    /**
     * @return Number of retry attempts.
     * 
     */
    private @Nullable Integer numRetries;
    /**
     * @return Number of seconds to wait until a restart is initiated.
     * 
     */
    private @Nullable Integer restartDelay;

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsMediaStoreSettings() {}
    /**
     * @return Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
     * 
     */
    public Optional<Integer> connectionRetryInterval() {
        return Optional.ofNullable(this.connectionRetryInterval);
    }
    public Optional<Integer> filecacheDuration() {
        return Optional.ofNullable(this.filecacheDuration);
    }
    public Optional<String> mediaStoreStorageClass() {
        return Optional.ofNullable(this.mediaStoreStorageClass);
    }
    /**
     * @return Number of retry attempts.
     * 
     */
    public Optional<Integer> numRetries() {
        return Optional.ofNullable(this.numRetries);
    }
    /**
     * @return Number of seconds to wait until a restart is initiated.
     * 
     */
    public Optional<Integer> restartDelay() {
        return Optional.ofNullable(this.restartDelay);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsMediaStoreSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer connectionRetryInterval;
        private @Nullable Integer filecacheDuration;
        private @Nullable String mediaStoreStorageClass;
        private @Nullable Integer numRetries;
        private @Nullable Integer restartDelay;
        public Builder() {}
        public Builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsMediaStoreSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionRetryInterval = defaults.connectionRetryInterval;
    	      this.filecacheDuration = defaults.filecacheDuration;
    	      this.mediaStoreStorageClass = defaults.mediaStoreStorageClass;
    	      this.numRetries = defaults.numRetries;
    	      this.restartDelay = defaults.restartDelay;
        }

        @CustomType.Setter
        public Builder connectionRetryInterval(@Nullable Integer connectionRetryInterval) {
            this.connectionRetryInterval = connectionRetryInterval;
            return this;
        }
        @CustomType.Setter
        public Builder filecacheDuration(@Nullable Integer filecacheDuration) {
            this.filecacheDuration = filecacheDuration;
            return this;
        }
        @CustomType.Setter
        public Builder mediaStoreStorageClass(@Nullable String mediaStoreStorageClass) {
            this.mediaStoreStorageClass = mediaStoreStorageClass;
            return this;
        }
        @CustomType.Setter
        public Builder numRetries(@Nullable Integer numRetries) {
            this.numRetries = numRetries;
            return this;
        }
        @CustomType.Setter
        public Builder restartDelay(@Nullable Integer restartDelay) {
            this.restartDelay = restartDelay;
            return this;
        }
        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsMediaStoreSettings build() {
            final var o = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsMediaStoreSettings();
            o.connectionRetryInterval = connectionRetryInterval;
            o.filecacheDuration = filecacheDuration;
            o.mediaStoreStorageClass = mediaStoreStorageClass;
            o.numRetries = numRetries;
            o.restartDelay = restartDelay;
            return o;
        }
    }
}
