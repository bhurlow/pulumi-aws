// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs Empty = new ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs();

    /**
     * This field is not currently supported and will not affect the output styling. Leave the default value.
     * 
     */
    @Import(name="styleControl", required=true)
    private Output<String> styleControl;

    /**
     * @return This field is not currently supported and will not affect the output styling. Leave the default value.
     * 
     */
    public Output<String> styleControl() {
        return this.styleControl;
    }

    private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs() {}

    private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs $) {
        this.styleControl = $.styleControl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs $;

        public Builder() {
            $ = new ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs();
        }

        public Builder(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs defaults) {
            $ = new ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param styleControl This field is not currently supported and will not affect the output styling. Leave the default value.
         * 
         * @return builder
         * 
         */
        public Builder styleControl(Output<String> styleControl) {
            $.styleControl = styleControl;
            return this;
        }

        /**
         * @param styleControl This field is not currently supported and will not affect the output styling. Leave the default value.
         * 
         * @return builder
         * 
         */
        public Builder styleControl(String styleControl) {
            return styleControl(Output.of(styleControl));
        }

        public ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettingsArgs build() {
            $.styleControl = Objects.requireNonNull($.styleControl, "expected parameter 'styleControl' to be non-null");
            return $;
        }
    }

}
