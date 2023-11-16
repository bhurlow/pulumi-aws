// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class V2modelsBotLocaleVoiceSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final V2modelsBotLocaleVoiceSettingsArgs Empty = new V2modelsBotLocaleVoiceSettingsArgs();

    @Import(name="engine")
    private @Nullable Output<String> engine;

    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    @Import(name="voiceId", required=true)
    private Output<String> voiceId;

    public Output<String> voiceId() {
        return this.voiceId;
    }

    private V2modelsBotLocaleVoiceSettingsArgs() {}

    private V2modelsBotLocaleVoiceSettingsArgs(V2modelsBotLocaleVoiceSettingsArgs $) {
        this.engine = $.engine;
        this.voiceId = $.voiceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(V2modelsBotLocaleVoiceSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private V2modelsBotLocaleVoiceSettingsArgs $;

        public Builder() {
            $ = new V2modelsBotLocaleVoiceSettingsArgs();
        }

        public Builder(V2modelsBotLocaleVoiceSettingsArgs defaults) {
            $ = new V2modelsBotLocaleVoiceSettingsArgs(Objects.requireNonNull(defaults));
        }

        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        public Builder voiceId(Output<String> voiceId) {
            $.voiceId = voiceId;
            return this;
        }

        public Builder voiceId(String voiceId) {
            return voiceId(Output.of(voiceId));
        }

        public V2modelsBotLocaleVoiceSettingsArgs build() {
            $.voiceId = Objects.requireNonNull($.voiceId, "expected parameter 'voiceId' to be non-null");
            return $;
        }
    }

}
