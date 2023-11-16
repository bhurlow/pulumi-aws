// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lex.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class V2modelsBotLocaleVoiceSettings {
    private @Nullable String engine;
    private String voiceId;

    private V2modelsBotLocaleVoiceSettings() {}
    public Optional<String> engine() {
        return Optional.ofNullable(this.engine);
    }
    public String voiceId() {
        return this.voiceId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(V2modelsBotLocaleVoiceSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String engine;
        private String voiceId;
        public Builder() {}
        public Builder(V2modelsBotLocaleVoiceSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.engine = defaults.engine;
    	      this.voiceId = defaults.voiceId;
        }

        @CustomType.Setter
        public Builder engine(@Nullable String engine) {
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder voiceId(String voiceId) {
            this.voiceId = Objects.requireNonNull(voiceId);
            return this;
        }
        public V2modelsBotLocaleVoiceSettings build() {
            final var o = new V2modelsBotLocaleVoiceSettings();
            o.engine = engine;
            o.voiceId = voiceId;
            return o;
        }
    }
}
