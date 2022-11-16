// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestination {
    /**
     * @return Reference ID for the destination.
     * 
     */
    private String destinationRefId;

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestination() {}
    /**
     * @return Reference ID for the destination.
     * 
     */
    public String destinationRefId() {
        return this.destinationRefId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestination defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String destinationRefId;
        public Builder() {}
        public Builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestination defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.destinationRefId = defaults.destinationRefId;
        }

        @CustomType.Setter
        public Builder destinationRefId(String destinationRefId) {
            this.destinationRefId = Objects.requireNonNull(destinationRefId);
            return this;
        }
        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestination build() {
            final var o = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestination();
            o.destinationRefId = destinationRefId;
            return o;
        }
    }
}
