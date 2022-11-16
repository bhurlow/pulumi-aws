// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.aws.medialive.outputs.ChannelInputAttachmentAutomaticInputFailoverSettings;
import com.pulumi.aws.medialive.outputs.ChannelInputAttachmentInputSettings;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ChannelInputAttachment {
    private @Nullable ChannelInputAttachmentAutomaticInputFailoverSettings automaticInputFailoverSettings;
    /**
     * @return User-specified name for the attachment.
     * 
     */
    private String inputAttachmentName;
    /**
     * @return The ID of the input.
     * 
     */
    private String inputId;
    /**
     * @return Settings of an input. See Input Settings for more details
     * 
     */
    private @Nullable ChannelInputAttachmentInputSettings inputSettings;

    private ChannelInputAttachment() {}
    public Optional<ChannelInputAttachmentAutomaticInputFailoverSettings> automaticInputFailoverSettings() {
        return Optional.ofNullable(this.automaticInputFailoverSettings);
    }
    /**
     * @return User-specified name for the attachment.
     * 
     */
    public String inputAttachmentName() {
        return this.inputAttachmentName;
    }
    /**
     * @return The ID of the input.
     * 
     */
    public String inputId() {
        return this.inputId;
    }
    /**
     * @return Settings of an input. See Input Settings for more details
     * 
     */
    public Optional<ChannelInputAttachmentInputSettings> inputSettings() {
        return Optional.ofNullable(this.inputSettings);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelInputAttachment defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ChannelInputAttachmentAutomaticInputFailoverSettings automaticInputFailoverSettings;
        private String inputAttachmentName;
        private String inputId;
        private @Nullable ChannelInputAttachmentInputSettings inputSettings;
        public Builder() {}
        public Builder(ChannelInputAttachment defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.automaticInputFailoverSettings = defaults.automaticInputFailoverSettings;
    	      this.inputAttachmentName = defaults.inputAttachmentName;
    	      this.inputId = defaults.inputId;
    	      this.inputSettings = defaults.inputSettings;
        }

        @CustomType.Setter
        public Builder automaticInputFailoverSettings(@Nullable ChannelInputAttachmentAutomaticInputFailoverSettings automaticInputFailoverSettings) {
            this.automaticInputFailoverSettings = automaticInputFailoverSettings;
            return this;
        }
        @CustomType.Setter
        public Builder inputAttachmentName(String inputAttachmentName) {
            this.inputAttachmentName = Objects.requireNonNull(inputAttachmentName);
            return this;
        }
        @CustomType.Setter
        public Builder inputId(String inputId) {
            this.inputId = Objects.requireNonNull(inputId);
            return this;
        }
        @CustomType.Setter
        public Builder inputSettings(@Nullable ChannelInputAttachmentInputSettings inputSettings) {
            this.inputSettings = inputSettings;
            return this;
        }
        public ChannelInputAttachment build() {
            final var o = new ChannelInputAttachment();
            o.automaticInputFailoverSettings = automaticInputFailoverSettings;
            o.inputAttachmentName = inputAttachmentName;
            o.inputId = inputId;
            o.inputSettings = inputSettings;
            return o;
        }
    }
}
