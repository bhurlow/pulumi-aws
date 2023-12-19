// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.aws.medialive.inputs.ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ChannelInputAttachmentAutomaticInputFailoverSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelInputAttachmentAutomaticInputFailoverSettingsArgs Empty = new ChannelInputAttachmentAutomaticInputFailoverSettingsArgs();

    /**
     * This clear time defines the requirement a recovered input must meet to be considered healthy. The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input\_preference for the failover pair is set to PRIMARY\_INPUT\_PREFERRED, because after this time, MediaLive will switch back to the primary input.
     * 
     */
    @Import(name="errorClearTimeMsec")
    private @Nullable Output<Integer> errorClearTimeMsec;

    /**
     * @return This clear time defines the requirement a recovered input must meet to be considered healthy. The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input\_preference for the failover pair is set to PRIMARY\_INPUT\_PREFERRED, because after this time, MediaLive will switch back to the primary input.
     * 
     */
    public Optional<Output<Integer>> errorClearTimeMsec() {
        return Optional.ofNullable(this.errorClearTimeMsec);
    }

    /**
     * A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input. See Failover Condition Block for more details.
     * 
     */
    @Import(name="failoverConditions")
    private @Nullable Output<List<ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionArgs>> failoverConditions;

    /**
     * @return A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input. See Failover Condition Block for more details.
     * 
     */
    public Optional<Output<List<ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionArgs>>> failoverConditions() {
        return Optional.ofNullable(this.failoverConditions);
    }

    /**
     * Input preference when deciding which input to make active when a previously failed input has recovered.
     * 
     */
    @Import(name="inputPreference")
    private @Nullable Output<String> inputPreference;

    /**
     * @return Input preference when deciding which input to make active when a previously failed input has recovered.
     * 
     */
    public Optional<Output<String>> inputPreference() {
        return Optional.ofNullable(this.inputPreference);
    }

    /**
     * The input ID of the secondary input in the automatic input failover pair.
     * 
     */
    @Import(name="secondaryInputId", required=true)
    private Output<String> secondaryInputId;

    /**
     * @return The input ID of the secondary input in the automatic input failover pair.
     * 
     */
    public Output<String> secondaryInputId() {
        return this.secondaryInputId;
    }

    private ChannelInputAttachmentAutomaticInputFailoverSettingsArgs() {}

    private ChannelInputAttachmentAutomaticInputFailoverSettingsArgs(ChannelInputAttachmentAutomaticInputFailoverSettingsArgs $) {
        this.errorClearTimeMsec = $.errorClearTimeMsec;
        this.failoverConditions = $.failoverConditions;
        this.inputPreference = $.inputPreference;
        this.secondaryInputId = $.secondaryInputId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelInputAttachmentAutomaticInputFailoverSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelInputAttachmentAutomaticInputFailoverSettingsArgs $;

        public Builder() {
            $ = new ChannelInputAttachmentAutomaticInputFailoverSettingsArgs();
        }

        public Builder(ChannelInputAttachmentAutomaticInputFailoverSettingsArgs defaults) {
            $ = new ChannelInputAttachmentAutomaticInputFailoverSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param errorClearTimeMsec This clear time defines the requirement a recovered input must meet to be considered healthy. The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input\_preference for the failover pair is set to PRIMARY\_INPUT\_PREFERRED, because after this time, MediaLive will switch back to the primary input.
         * 
         * @return builder
         * 
         */
        public Builder errorClearTimeMsec(@Nullable Output<Integer> errorClearTimeMsec) {
            $.errorClearTimeMsec = errorClearTimeMsec;
            return this;
        }

        /**
         * @param errorClearTimeMsec This clear time defines the requirement a recovered input must meet to be considered healthy. The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input\_preference for the failover pair is set to PRIMARY\_INPUT\_PREFERRED, because after this time, MediaLive will switch back to the primary input.
         * 
         * @return builder
         * 
         */
        public Builder errorClearTimeMsec(Integer errorClearTimeMsec) {
            return errorClearTimeMsec(Output.of(errorClearTimeMsec));
        }

        /**
         * @param failoverConditions A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input. See Failover Condition Block for more details.
         * 
         * @return builder
         * 
         */
        public Builder failoverConditions(@Nullable Output<List<ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionArgs>> failoverConditions) {
            $.failoverConditions = failoverConditions;
            return this;
        }

        /**
         * @param failoverConditions A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input. See Failover Condition Block for more details.
         * 
         * @return builder
         * 
         */
        public Builder failoverConditions(List<ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionArgs> failoverConditions) {
            return failoverConditions(Output.of(failoverConditions));
        }

        /**
         * @param failoverConditions A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input. See Failover Condition Block for more details.
         * 
         * @return builder
         * 
         */
        public Builder failoverConditions(ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionArgs... failoverConditions) {
            return failoverConditions(List.of(failoverConditions));
        }

        /**
         * @param inputPreference Input preference when deciding which input to make active when a previously failed input has recovered.
         * 
         * @return builder
         * 
         */
        public Builder inputPreference(@Nullable Output<String> inputPreference) {
            $.inputPreference = inputPreference;
            return this;
        }

        /**
         * @param inputPreference Input preference when deciding which input to make active when a previously failed input has recovered.
         * 
         * @return builder
         * 
         */
        public Builder inputPreference(String inputPreference) {
            return inputPreference(Output.of(inputPreference));
        }

        /**
         * @param secondaryInputId The input ID of the secondary input in the automatic input failover pair.
         * 
         * @return builder
         * 
         */
        public Builder secondaryInputId(Output<String> secondaryInputId) {
            $.secondaryInputId = secondaryInputId;
            return this;
        }

        /**
         * @param secondaryInputId The input ID of the secondary input in the automatic input failover pair.
         * 
         * @return builder
         * 
         */
        public Builder secondaryInputId(String secondaryInputId) {
            return secondaryInputId(Output.of(secondaryInputId));
        }

        public ChannelInputAttachmentAutomaticInputFailoverSettingsArgs build() {
            $.secondaryInputId = Objects.requireNonNull($.secondaryInputId, "expected parameter 'secondaryInputId' to be non-null");
            return $;
        }
    }

}
