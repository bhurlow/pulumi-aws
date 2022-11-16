// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.aws.medialive.inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArchiveS3SettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs Empty = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs();

    /**
     * Archive S3 Settings. See Archive S3 Settings for more details.
     * 
     */
    @Import(name="archiveS3Settings")
    private @Nullable Output<ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArchiveS3SettingsArgs> archiveS3Settings;

    /**
     * @return Archive S3 Settings. See Archive S3 Settings for more details.
     * 
     */
    public Optional<Output<ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArchiveS3SettingsArgs>> archiveS3Settings() {
        return Optional.ofNullable(this.archiveS3Settings);
    }

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs() {}

    private ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs $) {
        this.archiveS3Settings = $.archiveS3Settings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs $;

        public Builder() {
            $ = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs();
        }

        public Builder(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs defaults) {
            $ = new ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param archiveS3Settings Archive S3 Settings. See Archive S3 Settings for more details.
         * 
         * @return builder
         * 
         */
        public Builder archiveS3Settings(@Nullable Output<ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArchiveS3SettingsArgs> archiveS3Settings) {
            $.archiveS3Settings = archiveS3Settings;
            return this;
        }

        /**
         * @param archiveS3Settings Archive S3 Settings. See Archive S3 Settings for more details.
         * 
         * @return builder
         * 
         */
        public Builder archiveS3Settings(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArchiveS3SettingsArgs archiveS3Settings) {
            return archiveS3Settings(Output.of(archiveS3Settings));
        }

        public ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettingsArgs build() {
            return $;
        }
    }

}
