// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.aws.medialive.inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ChannelEncoderSettingsOutputGroupOutputArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelEncoderSettingsOutputGroupOutputArgs Empty = new ChannelEncoderSettingsOutputGroupOutputArgs();

    /**
     * The names of the audio descriptions used as audio sources for the output.
     * 
     */
    @Import(name="audioDescriptionNames")
    private @Nullable Output<List<String>> audioDescriptionNames;

    /**
     * @return The names of the audio descriptions used as audio sources for the output.
     * 
     */
    public Optional<Output<List<String>>> audioDescriptionNames() {
        return Optional.ofNullable(this.audioDescriptionNames);
    }

    /**
     * The names of the caption descriptions used as audio sources for the output.
     * 
     */
    @Import(name="captionDescriptionNames")
    private @Nullable Output<List<String>> captionDescriptionNames;

    /**
     * @return The names of the caption descriptions used as audio sources for the output.
     * 
     */
    public Optional<Output<List<String>>> captionDescriptionNames() {
        return Optional.ofNullable(this.captionDescriptionNames);
    }

    /**
     * The name used to identify an output.
     * 
     */
    @Import(name="outputName")
    private @Nullable Output<String> outputName;

    /**
     * @return The name used to identify an output.
     * 
     */
    public Optional<Output<String>> outputName() {
        return Optional.ofNullable(this.outputName);
    }

    /**
     * Settings for output. See Output Settings for more details.
     * 
     */
    @Import(name="outputSettings", required=true)
    private Output<ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs> outputSettings;

    /**
     * @return Settings for output. See Output Settings for more details.
     * 
     */
    public Output<ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs> outputSettings() {
        return this.outputSettings;
    }

    /**
     * The name of the video description used as audio sources for the output.
     * 
     */
    @Import(name="videoDescriptionName")
    private @Nullable Output<String> videoDescriptionName;

    /**
     * @return The name of the video description used as audio sources for the output.
     * 
     */
    public Optional<Output<String>> videoDescriptionName() {
        return Optional.ofNullable(this.videoDescriptionName);
    }

    private ChannelEncoderSettingsOutputGroupOutputArgs() {}

    private ChannelEncoderSettingsOutputGroupOutputArgs(ChannelEncoderSettingsOutputGroupOutputArgs $) {
        this.audioDescriptionNames = $.audioDescriptionNames;
        this.captionDescriptionNames = $.captionDescriptionNames;
        this.outputName = $.outputName;
        this.outputSettings = $.outputSettings;
        this.videoDescriptionName = $.videoDescriptionName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelEncoderSettingsOutputGroupOutputArgs $;

        public Builder() {
            $ = new ChannelEncoderSettingsOutputGroupOutputArgs();
        }

        public Builder(ChannelEncoderSettingsOutputGroupOutputArgs defaults) {
            $ = new ChannelEncoderSettingsOutputGroupOutputArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param audioDescriptionNames The names of the audio descriptions used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder audioDescriptionNames(@Nullable Output<List<String>> audioDescriptionNames) {
            $.audioDescriptionNames = audioDescriptionNames;
            return this;
        }

        /**
         * @param audioDescriptionNames The names of the audio descriptions used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder audioDescriptionNames(List<String> audioDescriptionNames) {
            return audioDescriptionNames(Output.of(audioDescriptionNames));
        }

        /**
         * @param audioDescriptionNames The names of the audio descriptions used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder audioDescriptionNames(String... audioDescriptionNames) {
            return audioDescriptionNames(List.of(audioDescriptionNames));
        }

        /**
         * @param captionDescriptionNames The names of the caption descriptions used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder captionDescriptionNames(@Nullable Output<List<String>> captionDescriptionNames) {
            $.captionDescriptionNames = captionDescriptionNames;
            return this;
        }

        /**
         * @param captionDescriptionNames The names of the caption descriptions used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder captionDescriptionNames(List<String> captionDescriptionNames) {
            return captionDescriptionNames(Output.of(captionDescriptionNames));
        }

        /**
         * @param captionDescriptionNames The names of the caption descriptions used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder captionDescriptionNames(String... captionDescriptionNames) {
            return captionDescriptionNames(List.of(captionDescriptionNames));
        }

        /**
         * @param outputName The name used to identify an output.
         * 
         * @return builder
         * 
         */
        public Builder outputName(@Nullable Output<String> outputName) {
            $.outputName = outputName;
            return this;
        }

        /**
         * @param outputName The name used to identify an output.
         * 
         * @return builder
         * 
         */
        public Builder outputName(String outputName) {
            return outputName(Output.of(outputName));
        }

        /**
         * @param outputSettings Settings for output. See Output Settings for more details.
         * 
         * @return builder
         * 
         */
        public Builder outputSettings(Output<ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs> outputSettings) {
            $.outputSettings = outputSettings;
            return this;
        }

        /**
         * @param outputSettings Settings for output. See Output Settings for more details.
         * 
         * @return builder
         * 
         */
        public Builder outputSettings(ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs outputSettings) {
            return outputSettings(Output.of(outputSettings));
        }

        /**
         * @param videoDescriptionName The name of the video description used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder videoDescriptionName(@Nullable Output<String> videoDescriptionName) {
            $.videoDescriptionName = videoDescriptionName;
            return this;
        }

        /**
         * @param videoDescriptionName The name of the video description used as audio sources for the output.
         * 
         * @return builder
         * 
         */
        public Builder videoDescriptionName(String videoDescriptionName) {
            return videoDescriptionName(Output.of(videoDescriptionName));
        }

        public ChannelEncoderSettingsOutputGroupOutputArgs build() {
            $.outputSettings = Objects.requireNonNull($.outputSettings, "expected parameter 'outputSettings' to be non-null");
            return $;
        }
    }

}
