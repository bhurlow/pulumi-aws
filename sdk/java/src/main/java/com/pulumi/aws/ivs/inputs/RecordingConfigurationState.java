// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ivs.inputs;

import com.pulumi.aws.ivs.inputs.RecordingConfigurationDestinationConfigurationArgs;
import com.pulumi.aws.ivs.inputs.RecordingConfigurationThumbnailConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RecordingConfigurationState extends com.pulumi.resources.ResourceArgs {

    public static final RecordingConfigurationState Empty = new RecordingConfigurationState();

    /**
     * ARN of the Recording Configuration.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the Recording Configuration.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Object containing destination configuration for where recorded video will be stored.
     * 
     */
    @Import(name="destinationConfiguration")
    private @Nullable Output<RecordingConfigurationDestinationConfigurationArgs> destinationConfiguration;

    /**
     * @return Object containing destination configuration for where recorded video will be stored.
     * 
     */
    public Optional<Output<RecordingConfigurationDestinationConfigurationArgs>> destinationConfiguration() {
        return Optional.ofNullable(this.destinationConfiguration);
    }

    /**
     * Recording Configuration name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Recording Configuration name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
     * 
     */
    @Import(name="recordingReconnectWindowSeconds")
    private @Nullable Output<Integer> recordingReconnectWindowSeconds;

    /**
     * @return If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
     * 
     */
    public Optional<Output<Integer>> recordingReconnectWindowSeconds() {
        return Optional.ofNullable(this.recordingReconnectWindowSeconds);
    }

    /**
     * The current state of the Recording Configuration.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return The current state of the Recording Configuration.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
     * 
     */
    @Import(name="thumbnailConfiguration")
    private @Nullable Output<RecordingConfigurationThumbnailConfigurationArgs> thumbnailConfiguration;

    /**
     * @return Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
     * 
     */
    public Optional<Output<RecordingConfigurationThumbnailConfigurationArgs>> thumbnailConfiguration() {
        return Optional.ofNullable(this.thumbnailConfiguration);
    }

    private RecordingConfigurationState() {}

    private RecordingConfigurationState(RecordingConfigurationState $) {
        this.arn = $.arn;
        this.destinationConfiguration = $.destinationConfiguration;
        this.name = $.name;
        this.recordingReconnectWindowSeconds = $.recordingReconnectWindowSeconds;
        this.state = $.state;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.thumbnailConfiguration = $.thumbnailConfiguration;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RecordingConfigurationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RecordingConfigurationState $;

        public Builder() {
            $ = new RecordingConfigurationState();
        }

        public Builder(RecordingConfigurationState defaults) {
            $ = new RecordingConfigurationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the Recording Configuration.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the Recording Configuration.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param destinationConfiguration Object containing destination configuration for where recorded video will be stored.
         * 
         * @return builder
         * 
         */
        public Builder destinationConfiguration(@Nullable Output<RecordingConfigurationDestinationConfigurationArgs> destinationConfiguration) {
            $.destinationConfiguration = destinationConfiguration;
            return this;
        }

        /**
         * @param destinationConfiguration Object containing destination configuration for where recorded video will be stored.
         * 
         * @return builder
         * 
         */
        public Builder destinationConfiguration(RecordingConfigurationDestinationConfigurationArgs destinationConfiguration) {
            return destinationConfiguration(Output.of(destinationConfiguration));
        }

        /**
         * @param name Recording Configuration name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Recording Configuration name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param recordingReconnectWindowSeconds If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
         * 
         * @return builder
         * 
         */
        public Builder recordingReconnectWindowSeconds(@Nullable Output<Integer> recordingReconnectWindowSeconds) {
            $.recordingReconnectWindowSeconds = recordingReconnectWindowSeconds;
            return this;
        }

        /**
         * @param recordingReconnectWindowSeconds If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
         * 
         * @return builder
         * 
         */
        public Builder recordingReconnectWindowSeconds(Integer recordingReconnectWindowSeconds) {
            return recordingReconnectWindowSeconds(Output.of(recordingReconnectWindowSeconds));
        }

        /**
         * @param state The current state of the Recording Configuration.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state The current state of the Recording Configuration.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param thumbnailConfiguration Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
         * 
         * @return builder
         * 
         */
        public Builder thumbnailConfiguration(@Nullable Output<RecordingConfigurationThumbnailConfigurationArgs> thumbnailConfiguration) {
            $.thumbnailConfiguration = thumbnailConfiguration;
            return this;
        }

        /**
         * @param thumbnailConfiguration Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
         * 
         * @return builder
         * 
         */
        public Builder thumbnailConfiguration(RecordingConfigurationThumbnailConfigurationArgs thumbnailConfiguration) {
            return thumbnailConfiguration(Output.of(thumbnailConfiguration));
        }

        public RecordingConfigurationState build() {
            return $;
        }
    }

}
