// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.aws.medialive.inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsDestinationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;


public final class ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs Empty = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs();

    /**
     * Destination address and port number for RTP or UDP packets. See Destination for more details.
     * 
     */
    @Import(name="destination", required=true)
    private Output<ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsDestinationArgs> destination;

    /**
     * @return Destination address and port number for RTP or UDP packets. See Destination for more details.
     * 
     */
    public Output<ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsDestinationArgs> destination() {
        return this.destination;
    }

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs() {}

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs(ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs $) {
        this.destination = $.destination;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs $;

        public Builder() {
            $ = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs();
        }

        public Builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs defaults) {
            $ = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param destination Destination address and port number for RTP or UDP packets. See Destination for more details.
         * 
         * @return builder
         * 
         */
        public Builder destination(Output<ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsDestinationArgs> destination) {
            $.destination = destination;
            return this;
        }

        /**
         * @param destination Destination address and port number for RTP or UDP packets. See Destination for more details.
         * 
         * @return builder
         * 
         */
        public Builder destination(ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsDestinationArgs destination) {
            return destination(Output.of(destination));
        }

        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettingsArgs build() {
            $.destination = Objects.requireNonNull($.destination, "expected parameter 'destination' to be non-null");
            return $;
        }
    }

}
