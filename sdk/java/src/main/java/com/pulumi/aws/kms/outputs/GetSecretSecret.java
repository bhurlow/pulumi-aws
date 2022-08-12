// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetSecretSecret {
    private final @Nullable Map<String,String> context;
    private final @Nullable List<String> grantTokens;
    private final String name;
    private final String payload;

    @CustomType.Constructor
    private GetSecretSecret(
        @CustomType.Parameter("context") @Nullable Map<String,String> context,
        @CustomType.Parameter("grantTokens") @Nullable List<String> grantTokens,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("payload") String payload) {
        this.context = context;
        this.grantTokens = grantTokens;
        this.name = name;
        this.payload = payload;
    }

    public Map<String,String> context() {
        return this.context == null ? Map.of() : this.context;
    }
    public List<String> grantTokens() {
        return this.grantTokens == null ? List.of() : this.grantTokens;
    }
    public String name() {
        return this.name;
    }
    public String payload() {
        return this.payload;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretSecret defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Map<String,String> context;
        private @Nullable List<String> grantTokens;
        private String name;
        private String payload;

        public Builder() {
    	      // Empty
        }

        public Builder(GetSecretSecret defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.context = defaults.context;
    	      this.grantTokens = defaults.grantTokens;
    	      this.name = defaults.name;
    	      this.payload = defaults.payload;
        }

        public Builder context(@Nullable Map<String,String> context) {
            this.context = context;
            return this;
        }
        public Builder grantTokens(@Nullable List<String> grantTokens) {
            this.grantTokens = grantTokens;
            return this;
        }
        public Builder grantTokens(String... grantTokens) {
            return grantTokens(List.of(grantTokens));
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder payload(String payload) {
            this.payload = Objects.requireNonNull(payload);
            return this;
        }        public GetSecretSecret build() {
            return new GetSecretSecret(context, grantTokens, name, payload);
        }
    }
}