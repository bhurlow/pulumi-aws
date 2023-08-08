// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NamespaceTimeoutsArgs extends com.pulumi.resources.ResourceArgs {

    public static final NamespaceTimeoutsArgs Empty = new NamespaceTimeoutsArgs();

    @Import(name="create")
    private @Nullable Output<String> create;

    public Optional<Output<String>> create() {
        return Optional.ofNullable(this.create);
    }

    @Import(name="delete")
    private @Nullable Output<String> delete;

    public Optional<Output<String>> delete() {
        return Optional.ofNullable(this.delete);
    }

    private NamespaceTimeoutsArgs() {}

    private NamespaceTimeoutsArgs(NamespaceTimeoutsArgs $) {
        this.create = $.create;
        this.delete = $.delete;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NamespaceTimeoutsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NamespaceTimeoutsArgs $;

        public Builder() {
            $ = new NamespaceTimeoutsArgs();
        }

        public Builder(NamespaceTimeoutsArgs defaults) {
            $ = new NamespaceTimeoutsArgs(Objects.requireNonNull(defaults));
        }

        public Builder create(@Nullable Output<String> create) {
            $.create = create;
            return this;
        }

        public Builder create(String create) {
            return create(Output.of(create));
        }

        public Builder delete(@Nullable Output<String> delete) {
            $.delete = delete;
            return this;
        }

        public Builder delete(String delete) {
            return delete(Output.of(delete));
        }

        public NamespaceTimeoutsArgs build() {
            return $;
        }
    }

}
