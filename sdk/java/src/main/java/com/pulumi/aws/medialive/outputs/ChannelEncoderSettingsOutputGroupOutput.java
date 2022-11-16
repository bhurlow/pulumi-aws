// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.aws.medialive.outputs.ChannelEncoderSettingsOutputGroupOutputOutputSettings;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ChannelEncoderSettingsOutputGroupOutput {
    /**
     * @return The names of the audio descriptions used as audio sources for the output.
     * 
     */
    private @Nullable List<String> audioDescriptionNames;
    /**
     * @return The names of the caption descriptions used as audio sources for the output.
     * 
     */
    private @Nullable List<String> captionDescriptionNames;
    /**
     * @return The name used to identify an output.
     * 
     */
    private @Nullable String outputName;
    /**
     * @return Settings for output. See Output Settings for more details.
     * 
     */
    private ChannelEncoderSettingsOutputGroupOutputOutputSettings outputSettings;
    /**
     * @return The name of the video description used as audio sources for the output.
     * 
     */
    private @Nullable String videoDescriptionName;

    private ChannelEncoderSettingsOutputGroupOutput() {}
    /**
     * @return The names of the audio descriptions used as audio sources for the output.
     * 
     */
    public List<String> audioDescriptionNames() {
        return this.audioDescriptionNames == null ? List.of() : this.audioDescriptionNames;
    }
    /**
     * @return The names of the caption descriptions used as audio sources for the output.
     * 
     */
    public List<String> captionDescriptionNames() {
        return this.captionDescriptionNames == null ? List.of() : this.captionDescriptionNames;
    }
    /**
     * @return The name used to identify an output.
     * 
     */
    public Optional<String> outputName() {
        return Optional.ofNullable(this.outputName);
    }
    /**
     * @return Settings for output. See Output Settings for more details.
     * 
     */
    public ChannelEncoderSettingsOutputGroupOutputOutputSettings outputSettings() {
        return this.outputSettings;
    }
    /**
     * @return The name of the video description used as audio sources for the output.
     * 
     */
    public Optional<String> videoDescriptionName() {
        return Optional.ofNullable(this.videoDescriptionName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelEncoderSettingsOutputGroupOutput defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> audioDescriptionNames;
        private @Nullable List<String> captionDescriptionNames;
        private @Nullable String outputName;
        private ChannelEncoderSettingsOutputGroupOutputOutputSettings outputSettings;
        private @Nullable String videoDescriptionName;
        public Builder() {}
        public Builder(ChannelEncoderSettingsOutputGroupOutput defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.audioDescriptionNames = defaults.audioDescriptionNames;
    	      this.captionDescriptionNames = defaults.captionDescriptionNames;
    	      this.outputName = defaults.outputName;
    	      this.outputSettings = defaults.outputSettings;
    	      this.videoDescriptionName = defaults.videoDescriptionName;
        }

        @CustomType.Setter
        public Builder audioDescriptionNames(@Nullable List<String> audioDescriptionNames) {
            this.audioDescriptionNames = audioDescriptionNames;
            return this;
        }
        public Builder audioDescriptionNames(String... audioDescriptionNames) {
            return audioDescriptionNames(List.of(audioDescriptionNames));
        }
        @CustomType.Setter
        public Builder captionDescriptionNames(@Nullable List<String> captionDescriptionNames) {
            this.captionDescriptionNames = captionDescriptionNames;
            return this;
        }
        public Builder captionDescriptionNames(String... captionDescriptionNames) {
            return captionDescriptionNames(List.of(captionDescriptionNames));
        }
        @CustomType.Setter
        public Builder outputName(@Nullable String outputName) {
            this.outputName = outputName;
            return this;
        }
        @CustomType.Setter
        public Builder outputSettings(ChannelEncoderSettingsOutputGroupOutputOutputSettings outputSettings) {
            this.outputSettings = Objects.requireNonNull(outputSettings);
            return this;
        }
        @CustomType.Setter
        public Builder videoDescriptionName(@Nullable String videoDescriptionName) {
            this.videoDescriptionName = videoDescriptionName;
            return this;
        }
        public ChannelEncoderSettingsOutputGroupOutput build() {
            final var o = new ChannelEncoderSettingsOutputGroupOutput();
            o.audioDescriptionNames = audioDescriptionNames;
            o.captionDescriptionNames = captionDescriptionNames;
            o.outputName = outputName;
            o.outputSettings = outputSettings;
            o.videoDescriptionName = videoDescriptionName;
            return o;
        }
    }
}
