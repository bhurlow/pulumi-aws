// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs Empty = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs();

    /**
     * Reference ID for the destination.
     * 
     */
    @Import(name="destinationRefId", required=true)
    private Output<String> destinationRefId;

    /**
     * @return Reference ID for the destination.
     * 
     */
    public Output<String> destinationRefId() {
        return this.destinationRefId;
    }

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs() {}

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs(ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs $) {
        this.destinationRefId = $.destinationRefId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs $;

        public Builder() {
            $ = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs();
        }

        public Builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs defaults) {
            $ = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param destinationRefId Reference ID for the destination.
         * 
         * @return builder
         * 
         */
        public Builder destinationRefId(Output<String> destinationRefId) {
            $.destinationRefId = destinationRefId;
            return this;
        }

        /**
         * @param destinationRefId Reference ID for the destination.
         * 
         * @return builder
         * 
         */
        public Builder destinationRefId(String destinationRefId) {
            return destinationRefId(Output.of(destinationRefId));
        }

        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestinationArgs build() {
            $.destinationRefId = Objects.requireNonNull($.destinationRefId, "expected parameter 'destinationRefId' to be non-null");
            return $;
        }
    }

}
