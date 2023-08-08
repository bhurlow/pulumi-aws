// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs Empty = new MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs();

    /**
     * Maximum bitrate.
     * 
     */
    @Import(name="maximumBitrate")
    private @Nullable Output<Integer> maximumBitrate;

    /**
     * @return Maximum bitrate.
     * 
     */
    public Optional<Output<Integer>> maximumBitrate() {
        return Optional.ofNullable(this.maximumBitrate);
    }

    /**
     * Minimum bitrate.
     * 
     */
    @Import(name="minimumBitrate")
    private @Nullable Output<Integer> minimumBitrate;

    /**
     * @return Minimum bitrate.
     * 
     */
    public Optional<Output<Integer>> minimumBitrate() {
        return Optional.ofNullable(this.minimumBitrate);
    }

    /**
     * Priority value.
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return Priority value.
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    private MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs() {}

    private MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs(MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs $) {
        this.maximumBitrate = $.maximumBitrate;
        this.minimumBitrate = $.minimumBitrate;
        this.priority = $.priority;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs $;

        public Builder() {
            $ = new MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs();
        }

        public Builder(MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs defaults) {
            $ = new MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param maximumBitrate Maximum bitrate.
         * 
         * @return builder
         * 
         */
        public Builder maximumBitrate(@Nullable Output<Integer> maximumBitrate) {
            $.maximumBitrate = maximumBitrate;
            return this;
        }

        /**
         * @param maximumBitrate Maximum bitrate.
         * 
         * @return builder
         * 
         */
        public Builder maximumBitrate(Integer maximumBitrate) {
            return maximumBitrate(Output.of(maximumBitrate));
        }

        /**
         * @param minimumBitrate Minimum bitrate.
         * 
         * @return builder
         * 
         */
        public Builder minimumBitrate(@Nullable Output<Integer> minimumBitrate) {
            $.minimumBitrate = minimumBitrate;
            return this;
        }

        /**
         * @param minimumBitrate Minimum bitrate.
         * 
         * @return builder
         * 
         */
        public Builder minimumBitrate(Integer minimumBitrate) {
            return minimumBitrate(Output.of(minimumBitrate));
        }

        /**
         * @param priority Priority value.
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority Priority value.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        public MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettingsArgs build() {
            return $;
        }
    }

}
