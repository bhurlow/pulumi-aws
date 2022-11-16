// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class ChannelEncoderSettingsAudioDescriptionRemixSettingsChannelMappingInputChannelLevel {
    private Integer gain;
    private Integer inputChannel;

    private ChannelEncoderSettingsAudioDescriptionRemixSettingsChannelMappingInputChannelLevel() {}
    public Integer gain() {
        return this.gain;
    }
    public Integer inputChannel() {
        return this.inputChannel;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelEncoderSettingsAudioDescriptionRemixSettingsChannelMappingInputChannelLevel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer gain;
        private Integer inputChannel;
        public Builder() {}
        public Builder(ChannelEncoderSettingsAudioDescriptionRemixSettingsChannelMappingInputChannelLevel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.gain = defaults.gain;
    	      this.inputChannel = defaults.inputChannel;
        }

        @CustomType.Setter
        public Builder gain(Integer gain) {
            this.gain = Objects.requireNonNull(gain);
            return this;
        }
        @CustomType.Setter
        public Builder inputChannel(Integer inputChannel) {
            this.inputChannel = Objects.requireNonNull(inputChannel);
            return this;
        }
        public ChannelEncoderSettingsAudioDescriptionRemixSettingsChannelMappingInputChannelLevel build() {
            final var o = new ChannelEncoderSettingsAudioDescriptionRemixSettingsChannelMappingInputChannelLevel();
            o.gain = gain;
            o.inputChannel = inputChannel;
            return o;
        }
    }
}
