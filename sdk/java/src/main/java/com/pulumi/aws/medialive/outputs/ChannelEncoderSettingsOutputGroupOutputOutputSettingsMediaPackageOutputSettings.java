// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.core.annotations.CustomType;
import java.util.Objects;

@CustomType
public final class ChannelEncoderSettingsOutputGroupOutputOutputSettingsMediaPackageOutputSettings {
    private ChannelEncoderSettingsOutputGroupOutputOutputSettingsMediaPackageOutputSettings() {}

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputOutputSettingsMediaPackageOutputSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        public Builder() {}
        public Builder(ChannelEncoderSettingsOutputGroupOutputOutputSettingsMediaPackageOutputSettings defaults) {
    	      Objects.requireNonNull(defaults);
        }

        public ChannelEncoderSettingsOutputGroupOutputOutputSettingsMediaPackageOutputSettings build() {
            final var o = new ChannelEncoderSettingsOutputGroupOutputOutputSettingsMediaPackageOutputSettings();
            return o;
        }
    }
}
