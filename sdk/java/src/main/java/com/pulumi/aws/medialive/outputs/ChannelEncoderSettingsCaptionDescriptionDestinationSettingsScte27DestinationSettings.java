// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.core.annotations.CustomType;
import java.util.Objects;

@CustomType
public final class ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte27DestinationSettings {
    private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte27DestinationSettings() {}

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte27DestinationSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        public Builder() {}
        public Builder(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte27DestinationSettings defaults) {
    	      Objects.requireNonNull(defaults);
        }

        public ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte27DestinationSettings build() {
            final var o = new ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte27DestinationSettings();
            return o;
        }
    }
}
