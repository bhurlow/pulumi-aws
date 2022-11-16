// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.aws.medialive.outputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettings;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings {
    private @Nullable String h265PackagingType;
    private ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettings hlsSettings;
    /**
     * @return String concatenated to the end of the destination filename. Required for multiple outputs of the same type.
     * 
     */
    private @Nullable String nameModifier;
    private @Nullable String segmentModifier;

    private ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings() {}
    public Optional<String> h265PackagingType() {
        return Optional.ofNullable(this.h265PackagingType);
    }
    public ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettings hlsSettings() {
        return this.hlsSettings;
    }
    /**
     * @return String concatenated to the end of the destination filename. Required for multiple outputs of the same type.
     * 
     */
    public Optional<String> nameModifier() {
        return Optional.ofNullable(this.nameModifier);
    }
    public Optional<String> segmentModifier() {
        return Optional.ofNullable(this.segmentModifier);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String h265PackagingType;
        private ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettings hlsSettings;
        private @Nullable String nameModifier;
        private @Nullable String segmentModifier;
        public Builder() {}
        public Builder(ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.h265PackagingType = defaults.h265PackagingType;
    	      this.hlsSettings = defaults.hlsSettings;
    	      this.nameModifier = defaults.nameModifier;
    	      this.segmentModifier = defaults.segmentModifier;
        }

        @CustomType.Setter
        public Builder h265PackagingType(@Nullable String h265PackagingType) {
            this.h265PackagingType = h265PackagingType;
            return this;
        }
        @CustomType.Setter
        public Builder hlsSettings(ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettings hlsSettings) {
            this.hlsSettings = Objects.requireNonNull(hlsSettings);
            return this;
        }
        @CustomType.Setter
        public Builder nameModifier(@Nullable String nameModifier) {
            this.nameModifier = nameModifier;
            return this;
        }
        @CustomType.Setter
        public Builder segmentModifier(@Nullable String segmentModifier) {
            this.segmentModifier = segmentModifier;
            return this;
        }
        public ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings build() {
            final var o = new ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings();
            o.h265PackagingType = h265PackagingType;
            o.hlsSettings = hlsSettings;
            o.nameModifier = nameModifier;
            o.segmentModifier = segmentModifier;
            return o;
        }
    }
}
