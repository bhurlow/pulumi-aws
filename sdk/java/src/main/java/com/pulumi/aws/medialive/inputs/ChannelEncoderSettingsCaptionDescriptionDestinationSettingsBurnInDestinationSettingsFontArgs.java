// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs Empty = new ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs();

    /**
     * Key used to extract the password from EC2 Parameter store.
     * 
     */
    @Import(name="passwordParam")
    private @Nullable Output<String> passwordParam;

    /**
     * @return Key used to extract the password from EC2 Parameter store.
     * 
     */
    public Optional<Output<String>> passwordParam() {
        return Optional.ofNullable(this.passwordParam);
    }

    /**
     * Path to a file accessible to the live stream.
     * 
     */
    @Import(name="uri", required=true)
    private Output<String> uri;

    /**
     * @return Path to a file accessible to the live stream.
     * 
     */
    public Output<String> uri() {
        return this.uri;
    }

    /**
     * Username to be used.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Username to be used.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs() {}

    private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs $) {
        this.passwordParam = $.passwordParam;
        this.uri = $.uri;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs $;

        public Builder() {
            $ = new ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs();
        }

        public Builder(ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs defaults) {
            $ = new ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param passwordParam Key used to extract the password from EC2 Parameter store.
         * 
         * @return builder
         * 
         */
        public Builder passwordParam(@Nullable Output<String> passwordParam) {
            $.passwordParam = passwordParam;
            return this;
        }

        /**
         * @param passwordParam Key used to extract the password from EC2 Parameter store.
         * 
         * @return builder
         * 
         */
        public Builder passwordParam(String passwordParam) {
            return passwordParam(Output.of(passwordParam));
        }

        /**
         * @param uri Path to a file accessible to the live stream.
         * 
         * @return builder
         * 
         */
        public Builder uri(Output<String> uri) {
            $.uri = uri;
            return this;
        }

        /**
         * @param uri Path to a file accessible to the live stream.
         * 
         * @return builder
         * 
         */
        public Builder uri(String uri) {
            return uri(Output.of(uri));
        }

        /**
         * @param username Username to be used.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username to be used.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettingsFontArgs build() {
            $.uri = Objects.requireNonNull($.uri, "expected parameter 'uri' to be non-null");
            return $;
        }
    }

}
