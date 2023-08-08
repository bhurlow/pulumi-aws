// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServerlessCollectionTimeouts {
    private @Nullable String create;
    private @Nullable String delete;

    private ServerlessCollectionTimeouts() {}
    public Optional<String> create() {
        return Optional.ofNullable(this.create);
    }
    public Optional<String> delete() {
        return Optional.ofNullable(this.delete);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServerlessCollectionTimeouts defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String create;
        private @Nullable String delete;
        public Builder() {}
        public Builder(ServerlessCollectionTimeouts defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.create = defaults.create;
    	      this.delete = defaults.delete;
        }

        @CustomType.Setter
        public Builder create(@Nullable String create) {
            this.create = create;
            return this;
        }
        @CustomType.Setter
        public Builder delete(@Nullable String delete) {
            this.delete = delete;
            return this;
        }
        public ServerlessCollectionTimeouts build() {
            final var o = new ServerlessCollectionTimeouts();
            o.create = create;
            o.delete = delete;
            return o;
        }
    }
}
